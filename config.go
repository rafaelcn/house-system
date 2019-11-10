package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type (
	// DatabaseConfiguration represents the configuration of a database.
	DatabaseConfiguration struct {
		Host         string
		Port         string
		User         string
		Password     string
		DatabaseName string
	}

	// ServerConfiguration represents the configuration struct encoded in json.
	ServerConfiguration struct {
		Address  string
		Port     string
		Database DatabaseConfiguration
	}
)

// Parse parsee the configuration file pointed by the filename
func Parse(filename string) ServerConfiguration {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("Server configuration file not found. Error: %v", err)
	}

	var config ServerConfiguration

	json.Unmarshal(file, &config)

	return config
}
