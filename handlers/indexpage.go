package handlers

import (
	"html/template"
	"log"
	"net/http"
	"sensormonitor/model"
	"sensormonitor/syscmd"
)

//IndexPage handeler
type IndexPage struct {
	logger *log.Logger
}

//IndexPage function returning reference to IndexPage handle
func NewIndexPage(logger *log.Logger) *IndexPage {
	return &IndexPage{logger}
}

///@todo this function crash if data.message is nil
func (indexPage *IndexPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	indexPage.logger.Println("(OPEN): " + model.HttpDir + "templates/statusPage.html")
	statusPage := model.StatusPage{}
	osinfo, err := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	if err != nil {
		indexPage.logger.Println(err)
	}

	statusPage.FooterData.OsString = osinfo["NAME"]
	statusPage.FooterData.OsVersion = osinfo["VERSION_ID"]

	statusPage.NavPages = GetMenu(r.URL.Path)

	indexPage.logger.Println("Reading status")
	data, err := syscmd.ReadStatus()
	if err != nil {
		indexPage.logger.Println(err)
	}

	statusPage.SystemdStatus = *data
	var temp = model.HttpDir + "templates/error.html"

	if data.Active {
		indexPage.logger.Println("Systemd service works")
		temp = model.HttpDir + "templates/message.html"
	}

	navbar := model.HttpDir + "templates/navbar.html"
	footer := model.HttpDir + "templates/footer.html"

	t, err := template.ParseFiles(model.HttpDir+"templates/statusPage.html", temp, navbar, footer)

	if err != nil {
		indexPage.logger.Println("Error ", err)
	}

	exec_err := t.Execute(w, statusPage)

	if exec_err != nil {
		indexPage.logger.Println(exec_err.Error())
	}
}
