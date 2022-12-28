package handlers

import (
	"html/template"
	"log"
	"net/http"
	"sensormonitor/model"
	"sensormonitor/syscmd"
)

type LogHandle struct {
	logger *log.Logger
}

func NewLogHandle(logger *log.Logger) *LogHandle {
	return &LogHandle{logger}
}

func (logHandle *LogHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logHandle.logger.Println("(OPEN): " + model.HttpDir + "templates/logPage.html")
	logPage := model.LogPage{}
	osinfo, err := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	if err != nil {
		logHandle.logger.Println(err)
	}

	logPage.FooterData.OsString = osinfo["NAME"]
	logPage.FooterData.OsVersion = osinfo["VERSION_ID"]

	logPage.NavPages = GetMenu(r.URL.Path)

	logHandle.logger.Println("Reading logs")
	data, err := syscmd.ReadLog()
	if err != nil {
		logHandle.logger.Println(err)
	}

	logPage.AllMessages = *data

	navbar := model.HttpDir + "templates/navbar.html"
	footer := model.HttpDir + "templates/footer.html"

	t, err := template.ParseFiles(model.HttpDir+"templates/logPage.html", navbar, footer)
	if err != nil {
		logHandle.logger.Println(err)
	}

	t.Execute(w, logPage)
}
