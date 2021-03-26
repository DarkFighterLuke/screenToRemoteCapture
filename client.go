package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/kbinani/screenshot"
	hook "github.com/robotn/gohook"
	"gopkg.in/yaml.v2"
	"image"
	"image/png"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

type clientConfigurationData struct {
	Apps map[string]appsData `yaml:"capture,omitempty"`
}

type captureArea struct {
	X      int
	Y      int
	Width  int
	Height int
}

type serverData struct {
	Ip   string
	Port int
}

type appsData struct {
	Server serverData
	Area   captureArea
}

func clientMode() {
	log.Println("--- Running in Client Mode ---")
	var clientConf clientConfigurationData

	content, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatalln("An error occurred when opening configuration file.\nCheck that the file exists.\n", err)
	}

	err = yaml.Unmarshal(content, &clientConf)
	if err != nil {
		log.Fatalln("An error occurred when opening configuration file.\nCheck that the file exists.\n", err)
	}

	appName := inputAppName(clientConf)
	fmt.Println("App configuration found. Using \"" + appName + "\"")

	servData := clientConf.Apps[appName].Server
	capData := clientConf.Apps[appName].Area

	// Launch goroutine to handle user keystrokes
	capture := make(chan bool)
	go handleUserKeystroke(capture)

	for {
		select {
		case toCapture := <-capture:
			if toCapture {
				fmt.Println("Taking screenshot...")
				screenCap, err := takeScreenshot(capData)
				if err == nil {
					fmt.Println("Screenshot taken. Now sending...")
					go sendToServer(servData, screenCap)
				}
			}
		}
	}
}

func inputAppName(clientConf clientConfigurationData) string {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Input the name of the application you want to capture: ")
		appNameBytes, _, _ := reader.ReadLine()
		appName := string(appNameBytes)

		if _, ok := clientConf.Apps[appName]; ok {
			return appName
		} else {
			fmt.Println("No application with this name was found in configuration.yaml.")
		}
	}

}

func takeScreenshot(capData captureArea) (image.Image, error) {
	img, err := screenshot.Capture(capData.X, capData.Y, capData.Width, capData.Height)
	if err != nil {
		log.Println("Error while taking screenshot: ", err)
	}

	return img, err
}

func sendToServer(servData serverData, img image.Image) {
	c := &http.Client{}
	url := "http://" + servData.Ip + ":" + strconv.Itoa(servData.Port)

	var b bytes.Buffer
	mw := multipart.NewWriter(&b)
	var fw io.Writer

	fw, err := mw.CreateFormFile("image", "image")
	if err != nil {
		log.Println("Error while creating FormFile: ", err)
		return
	}

	err = png.Encode(fw, img)
	if err != nil {
		log.Println("Error while encoding image to PNG: ", err)
		return
	}

	err = mw.Close()
	if err != nil {
		log.Println("Error while closing multipart: ", err)
		return
	}

	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		log.Println("Error while creating HTTP request: ", err)
		return
	}
	req.Header.Set("Content-Type", mw.FormDataContentType())
	res, err := c.Do(req)
	if err != nil {
		log.Println("Error while setting HTTP request content-type header: ", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		log.Println("Status Code: ", res.StatusCode)
		return
	}

}

func handleUserKeystroke(capture chan bool) {
	hook.Register(hook.KeyDown, []string{"ctrl", "k"}, func(e hook.Event) {
		capture <- true
	})

	s := hook.Start()
	<-hook.Process(s)
}
