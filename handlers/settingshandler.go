package handlers

import (
	"html/template"
	"log"
	"net/http"
	"sensormonitor/model"
	"sensormonitor/syscmd"
)

type SettingsHandler struct {
	logger *log.Logger
}

func NewSettingsHandler(logger *log.Logger) *SettingsHandler {
	return &SettingsHandler{logger}
}

func (settingsHandler *SettingsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	idxPage := model.IndexPage{}
	osinfo, err := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	if err != nil {
		settingsHandler.logger.Println(err)
	}

	idxPage.FooterData.OsString = osinfo["NAME"]
	idxPage.FooterData.OsVersion = osinfo["VERSION_ID"]

	idxPage.NavPages = GetMenu(r.URL.Path)

	data, err := model.ListAllSettings()
	if err != nil {
		settingsHandler.logger.Println(err)
	}

	idxPage.Settings = data
	settingsHandler.logger.Println("(OPEN): " + model.HttpDir + "templates/settings.html")

	navbar := model.HttpDir + "templates/navbar.html"
	footer := model.HttpDir + "templates/footer.html"
	t, err := template.ParseFiles(model.HttpDir+"templates/settings.html", navbar, footer)

	if err != nil {
		settingsHandler.logger.Println(err)
	}

	t.Execute(w, idxPage)
}
