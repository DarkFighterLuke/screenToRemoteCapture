package main

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type configurationData struct {
	Mode string `yaml:"mode,omitempty"`
}

var configFilePath = "configuration.yaml"

func main() {
	var confData configurationData

	if len(os.Args) > 1 {
		configFilePath = os.Args[1]
	}
	content, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatalln("An error occurred when opening configuration file.\nCheck that the file exists.\n", err)
	}

	err = yaml.Unmarshal(content, &confData)
	if err != nil {
		log.Fatalln("An error occurred when reading configuration file.")
	}

	if confData.Mode == "server" {
		serverMode()
	} else if confData.Mode == "client" {
		clientMode()
	} else {
		log.Fatalln("Invalid value for parameter mode in configuration.yaml.")
	}

}
