package handlers

import (
	"log"
	"net/http"
	"sensormonitor/syscmd"
)

type RestartWebui struct {
	logger *log.Logger
}

func NewRestartWebui(logger *log.Logger) *RestartWebui {
	return &RestartWebui{logger}
}

func (restartWebui *RestartWebui) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	restartWebui.logger.Println("Restart webui.")
	if r.Method == "GET" {
		restartWebui.logger.Println("TODO: implement GET for RestartWebui")
		http.Redirect(w, r, "/", 301)
	} else {
		restartWebui.logger.Println("Restarting webui.")
		err := syscmd.CmdRestart(syscmd.Webservice)
		if err != nil {
			restartWebui.logger.Println(err)
		}
		http.Redirect(w, r, "/", 301)
	}
}
