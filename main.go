package main

import (
	"flag"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"syncute-go/connections"
	"syncute-go/helpers"
	"syncute-go/models"
)

var (
	remoteAddress *string
	token         *string
	repoPath      *string
)

func init() {
	remoteAddress = flag.String("address", "", "Remote Address")
	token = flag.String("token", "", "Access Token")
	repoPath = flag.String("path", "", "Repository Path")
}

func main() {
	flag.Parse()

	f, err := os.Open("config.yml")
	if err != nil {

	}
	defer f.Close()

	var cfg models.Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal("Error in reading config.yml ", err)
	}

	if *remoteAddress == "" {
		*remoteAddress = cfg.Server.RemoteAddress
	}

	if *token == "" {
		*token = cfg.Server.AccessToken
	}

	if *repoPath == "" {
		*repoPath = cfg.Client.RepositoryPath
	}

	log.Printf("Config=> RemoteAddress: %s, RepositoryPath: %s\n", *remoteAddress, *repoPath)

	helpers.CheckRepo(cfg.Client.RepositoryPath)

	client := connections.Client{
		RemoteAddress: *remoteAddress,
		Token:         *token,
	}
	client.Start()
}
