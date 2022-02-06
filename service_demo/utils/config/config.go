package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Database Database `yaml:"database"`
}

type Database struct {
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	DatabaseType   string `yaml:"databasetype"`
	ConnectionType string `yaml:"connectiontype"`
	Schema         string `yaml:"schema"`
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	Charset        string `yaml:"charset"`
}

func ReadConfig() (config Configuration) {

	yfile, err := ioutil.ReadFile("items.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yfile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return config
}
