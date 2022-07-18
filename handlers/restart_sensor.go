package handlers

import (
	"log"
	"net/http"
	"sensormonitor/syscmd"
)

type RestartSensor struct {
	logger *log.Logger
}

func NewRestartSensor(logger *log.Logger) *RestartSensor {
	return &RestartSensor{logger}
}

func (restartSensor *RestartSensor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	restartSensor.logger.Println("Restart sensor.")
	if r.Method == "GET" {
		restartSensor.logger.Println("TODO: implement GET for RestartSensor")
		http.Redirect(w, r, "/", 301)
	} else {
		restartSensor.logger.Println("Restarting sensor.")
		err := syscmd.CmdRestart(syscmd.Tempsensorservice)
		if err != nil {
			restartSensor.logger.Println(err)
		}
		http.Redirect(w, r, "/", 301)
	}
}
