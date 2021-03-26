package main

import (
	"github.com/ZacJoffe/clipboard/xclip"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type serverConfigurationData struct {
	ServerInfo serverInfo `yaml:"server,omitempty"`
}

type serverInfo struct {
	ListenIp   string
	ListenPort int
}

func serverMode() {
	log.Println("--- Running in Server Mode ---")
	var servConf serverConfigurationData

	content, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatalln("An error occurred when opening configuration file.\nCheck that the file exists.\n", err)
	}

	err = yaml.Unmarshal(content, &servConf)
	if err != nil {
		log.Fatalln("An error occurred when reading configuration file: ", err)
	}

	listenAddress := servConf.ServerInfo.ListenIp + ":" + strconv.Itoa(servConf.ServerInfo.ListenPort)
	log.Println("Listening on ", listenAddress, "...")
	http.HandleFunc("/", handleIncomingImage)
	err = http.ListenAndServe(listenAddress, nil)
	if err != nil {
		log.Println("Can't start server with given listen address: ", err)
	}
}

func handleIncomingImage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(524288000) //500MB
	if err != nil {
		log.Println("Error while parsing form data: ", err)
	}

	image, _, err := r.FormFile("image")
	if err != nil {
		log.Println("Error while parsing received image: ", err)
	}

	imageBytes, err := ioutil.ReadAll(image)
	if err != nil {
		log.Println(err)
	}

	imageFile, err := ioutil.TempFile("./", ".tempimg*.png")
	if err != nil {
		log.Println("Error while creating temp image file: ", err)
		return
	}

	_, err = imageFile.Write(imageBytes)
	if err != nil {
		log.Println("Error while writing to temp image file: ", err)
		return
	}

	// TODO: Distinguish between Linux and Windows
	err = xclip.WriteImage(imageFile)
	if err != nil {
		log.Println("Error while writing image to clipboard: ", err)
	}

	log.Println("Received screenshot")
	w.WriteHeader(http.StatusOK)
}
