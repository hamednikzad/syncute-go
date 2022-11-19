package connections

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"syncute-go/helpers"
	"syncute-go/messages"
)

type Client struct {
	Token         string
	RemoteAddress string
	connection    *websocket.Conn
}

func (client *Client) Start() {
	go client.configRepositoryWatcher()

	client.connect()

	client.consume()
}

var watcher *fsnotify.Watcher

func (client *Client) configRepositoryWatcher() {
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	// starting at the root of the project, walk each file/directory searching for directories
	if err := filepath.Walk(helpers.RepoPath, watchDir); err != nil {
		log.Println("ERROR in watcher Walk", err)
	}

	done := make(chan bool)

	func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Create) {
					log.Printf("Change detect: %s name: %s\n", "Create", event.Name)

					if messages.IsResourceExistAlready(event.Name) {
						log.Printf("Watcher: file %s already synced\n", event.Name)
					} else {
						newResource := helpers.GetResourceWithoutChecksum(event.Name, "")
						//messages.UploadResource(newResource, client.sendBinaryMessage)
						data := messages.SerializeResource(newResource)
						client.sendBinaryMessage(data)
					}
				}
			case err := <-watcher.Errors:
				log.Println("ERROR in watcher", err)
			}
		}
	}()

	<-done
}

func watchDir(path string, fi os.FileInfo, err error) error {
	// since fsnotify can watch all the files in a directory, watchers only need
	// to be added to each nested directory
	if fi.Mode().IsDir() {
		return watcher.Add(path)
	}

	return nil
}

func (client *Client) connect() {
	log.Printf("Connecting to %s...\n", client.RemoteAddress)
	var err error
	client.connection, _, err = websocket.DefaultDialer.Dial(client.RemoteAddress, http.Header{
		"access_token": []string{client.Token},
	})
	if err != nil {
		log.Println(err)
		return
	}
}

func (client *Client) consume() {
	messages.CurrentResources = helpers.GetAllFilesWithChecksum()

	log.Println("Waiting for new message...")

	for {
		messageType, message, err := client.connection.ReadMessage()

		if err != nil {
			return
		}

		switch messageType {
		case websocket.TextMessage:
			log.Println("server: text message received>>")
			go messages.ProcessTextMessage(client.sendTextMessage, client.sendBinaryMessage, message)
			break
		case websocket.BinaryMessage:
			go messages.ProcessBinaryMessage(message)
			break
		default:
			log.Println("Unknown")
		}
	}
}

func (client *Client) sendTextMessage(message []byte) {
	log.Println("sending: ", string(message))
	client.connection.WriteMessage(websocket.TextMessage, message)
}

func (client *Client) sendBinaryMessage(message []byte) {
	log.Printf("sending binary message with size: %d\n", len(message))
	client.connection.WriteMessage(websocket.BinaryMessage, message)
}
