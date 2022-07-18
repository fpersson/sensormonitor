package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sensormonitor/model"
)

//UpdateSettings Handeler
type UpdateSettings struct {
	logger *log.Logger
}

//UpdateSettings function returning reference to UpdateSettings handle
func NewUpdateSettings(logger *log.Logger) *UpdateSettings {
	return &UpdateSettings{logger}
}

func (updateSettings *UpdateSettings) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	updateSettings.logger.Println("Update settings.")
	settings := &model.Settings{}

	if r.Method == "GET" {
		updateSettings.logger.Println("TODO: implement GET for UpdateSettings")
		http.Redirect(w, r, "/", 301)
	} else {
		settings.Sensor = r.FormValue("sensor")
		settings.Name = r.FormValue("name")
		settings.Influx.Host = r.FormValue("host")
		settings.Influx.Token = r.FormValue("token")
		settings.Influx.Apiorg = r.FormValue("api")
		settings.Influx.Bucket = r.FormValue("bucket")
		settings.Influx.Interval = r.FormValue("interval")
		settings.Grafana.Host = r.FormValue("grafana-host")

		data, err := json.MarshalIndent(&settings, "", "	")

		if err != nil {
			updateSettings.logger.Println("Error ", err)
		}

		err = ioutil.WriteFile(model.SettingsPath, data, 0666)

		if err != nil {
			updateSettings.logger.Println("Error ", err)
		}

		http.Redirect(w, r, "/", 301)
	}
}
