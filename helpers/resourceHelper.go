package helpers

import (
	"log"
	"os"
	"syncute-go/messages/resources"
)

var repoPath = "C:/SyncRepo"

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetAllFilesWithChecksum() []resources.Resource {
	return nil
}

func CheckRepository() {
	ex, err := exists(repoPath)
	if err != nil {
		panic(err)
	}
	if !ex {
		err := os.Mkdir(repoPath, os.ModePerm)
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}
}
