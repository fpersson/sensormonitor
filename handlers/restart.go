package handlers

import (
	"html/template"
	"log"
	"net/http"
	"sensormonitor/model"
	"sensormonitor/syscmd"
)

type Reboot struct {
	logger *log.Logger
}

func NewReboot(logger *log.Logger) *Reboot {
	return &Reboot{logger}
}

func (reboot *Reboot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reboot.logger.Println("(OPEN): " + model.HttpDir + "templates/restartPage.html")
	rebootPage := model.RebootPage{}
	osinfo, err := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	if err != nil {
		reboot.logger.Println(err)
	}

	rebootPage.FooterData.OsString = osinfo["NAME"]
	rebootPage.FooterData.OsVersion = osinfo["VERSION_ID"]

	rebootPage.NavPages = GetMenu(r.URL.Path)

	navbar := model.HttpDir + "templates/navbar.html"
	footer := model.HttpDir + "templates/footer.html"
	t, err := template.ParseFiles(model.HttpDir+"templates/restartPage.html", navbar, footer)

	if err != nil {
		reboot.logger.Println(err)
	}

	t.Execute(w, rebootPage)
}
