package handlers

import (
	"example/user/webserver/logreader"
	"example/user/webserver/model"
	"example/user/webserver/syscmd"
	"html/template"
	"log"
	"net/http"
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

	logHandle.logger.Println("Reading logs")
	data, err := logreader.ReadLog()
	if err != nil {
		logHandle.logger.Println(err)
	}

	logPage.AllMessages = *data

	footer := model.HttpDir + "templates/footer.html"

	t, err := template.ParseFiles(model.HttpDir+"templates/logPage.html", footer)
	if err != nil {
		logHandle.logger.Println(err)
	}

	t.Execute(w, logPage)
}
