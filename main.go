package main

import (
	"errors"
	"example/user/webserver/handlers"
	"example/user/webserver/model"
	"fmt"
	"net/http"
	"os"
	"os/user"
	"strings"
)

const configfile = "tempsensor/settings.json"
const http_path = "/usr/share/demo/"

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

///TODO flytta och testa denna funktionen
func findConfigFile(configpaths string) (confile string) {
	v := strings.Split(configpaths, ":")

	for _, element := range v {
		file := element + "/" + configfile
		fmt.Println(file)
		if exists(file) {
			return file
		}
	}

	return ""
}

func main() {
	port := os.Getenv("PORT")
	configdir := os.Getenv("CONFIG")
	model.HttpDir = os.Getenv("HTML")

	if port == "" {
		port = ":8080"
	}

	if configdir == "" {
		configdir = os.Getenv("XDG_DATA_DIRS")
	}
	model.SettingsPath = findConfigFile(configdir)

	if model.SettingsPath == "" {
		fmt.Println("ERROR: " + configfile + " not found in " + configdir)
	}

	if model.HttpDir == "" {
		model.HttpDir = http_path
	}

	user, err := user.Current()
	if err != nil {
		fmt.Printf("Error:" + err.Error())
	}

	fmt.Printf("Username: %s - Uid; %s - Gid %s\n", user.Username, user.Uid, user.Gid)

	fmt.Println("Starting at localhost" + port)
	fmt.Println("Http dir: " + model.HttpDir)
	http.HandleFunc("/", handlers.StatusFunc)
	http.HandleFunc("/update", handlers.UpdateSettings)
	http.HandleFunc("/settings.html", handlers.SettingsFunc)
	http.HandleFunc("/restart.html", handlers.RebootFunc)
	http.HandleFunc("/status.html", handlers.StatusFunc)
	http.HandleFunc("/log.html", handlers.LogFunc)
	http.HandleFunc("/restart_webui", handlers.RestartWebui)
	http.HandleFunc("/restart_sensor", handlers.RestartSensor)
	http.ListenAndServe(port, nil)
}
