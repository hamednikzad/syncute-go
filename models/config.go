package models

type Config struct {
	Server struct {
		RemoteAddress string `yaml:"remote-address"`
		AccessToken   string `yaml:"access-token"`
	} `yaml:"server"`
	Client struct {
		RepositoryPath string `yaml:"repository-path"`
	} `yaml:"client"`
}
