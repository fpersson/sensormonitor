package main

import (
	"context"
	"errors"
	"example/user/webserver/handlers"
	"example/user/webserver/model"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
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
	logger := log.New(os.Stdout, "sensor-monitor: ", log.LstdFlags)
	port := os.Getenv("PORT")
	configdir := os.Getenv("CONFIG")
	model.HttpDir = os.Getenv("HTML")

	serveMux := http.NewServeMux()

	if port == "" {
		port = ":8080"
	}

	if configdir == "" {
		configdir = os.Getenv("XDG_DATA_DIRS")
	}
	logger.Println(configdir)
	model.SettingsPath = findConfigFile(configdir)

	if model.SettingsPath == "" {
		logger.Println("ERROR: " + configfile + " not found in " + configdir)
	}

	if model.HttpDir == "" {
		model.HttpDir = http_path
	}

	logger.Println("Starting at localhost" + port)
	logger.Println("Http dir: " + model.HttpDir)

	indexPageHandler := handlers.NewIndexPage(logger)
	logHandler := handlers.NewLogHandle(logger)
	restartHandler := handlers.NewReboot(logger)
	restartSensorHandler := handlers.NewRestartSensor(logger)
	restartWebuiHandler := handlers.NewRestartWebui(logger)
	settingsHandler := handlers.NewSettingsHandler(logger)
	updateSettingsHandler := handlers.NewUpdateSettings(logger)

	serveMux.Handle("/", indexPageHandler)
	serveMux.Handle("/log.html", logHandler)
	serveMux.Handle("/restart.html", restartHandler)
	serveMux.Handle("/restart_sensor", restartSensorHandler)
	serveMux.Handle("/restart_webui", restartWebuiHandler)
	serveMux.Handle("/settings.html", settingsHandler)
	serveMux.Handle("/status.html", indexPageHandler)
	serveMux.Handle("/update", updateSettingsHandler)

	server := &http.Server{
		Addr:         port,
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Sigterm: ", sig)
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.Shutdown(tc)

	logger.Println("** EXIT **")
}
