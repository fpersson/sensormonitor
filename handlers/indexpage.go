package handlers

import (
	"example/user/webserver/logreader"
	"example/user/webserver/model"
	"example/user/webserver/syscmd"
	"html/template"
	"log"
	"net/http"
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

	indexPage.logger.Println("Reading status")
	data, err := logreader.ReadStatus()
	if err != nil {
		indexPage.logger.Println(err)
	}

	statusPage.SystemdStatus = *data
	var temp = model.HttpDir + "templates/error.html"

	if data.Active == true {
		indexPage.logger.Println("Systemd service works")
		temp = model.HttpDir + "templates/message.html"
	}

	footer := model.HttpDir + "templates/footer.html"

	t, err := template.ParseFiles(model.HttpDir+"templates/statusPage.html", temp, footer)

	if err != nil {
		indexPage.logger.Println("Error ", err)
	}

	t.Execute(w, statusPage)
}
