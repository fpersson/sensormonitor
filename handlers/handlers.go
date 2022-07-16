package handlers

import (
	"encoding/json"
	"example/user/webserver/logreader"
	"example/user/webserver/model"
	"example/user/webserver/syscmd"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateSettings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update settings.")
	settings := &model.Settings{}
	if r.Method == "GET" {
		fmt.Println("TODO: implement GET for UpdateSettings")
		http.Redirect(w, r, "/", 301)
	} else {
		///TODO validate input

		settings.Sensor = r.FormValue("sensor")
		settings.Name = r.FormValue("name")
		settings.Influx.Host = r.FormValue("host")
		settings.Influx.Token = r.FormValue("token")
		settings.Influx.Apiorg = r.FormValue("api")
		settings.Influx.Bucket = r.FormValue("bucket")
		settings.Influx.Interval = r.FormValue("interval")
		settings.Grafana.Host = r.FormValue("grafana-host")

		data, err := json.MarshalIndent(&settings, "", "	")
		checkError(err)
		err = ioutil.WriteFile(model.SettingsPath, data, 0666)

		if err != nil {
			fmt.Println("Error ", err)
		}

		http.Redirect(w, r, "/", 301)
	}
}

/// @todo swap settings and status as indexpage
func IndexFunc(w http.ResponseWriter, r *http.Request) {
	idxPage := model.IndexPage{}
	osinfo := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	idxPage.FooterData.OsString = osinfo["NAME"]
	idxPage.FooterData.OsVersion = osinfo["VERSION_ID"]

	idxPage.Settings = model.ListAllSettings()
	fmt.Println("(OPEN): " + model.HttpDir + "templates/indexPage.html")
	footer := model.HttpDir + "templates/footer.html"
	t, err := template.ParseFiles(model.HttpDir+"templates/indexPage.html", footer)
	checkError(err)
	t.Execute(w, idxPage)
}

func SettingsFunc(w http.ResponseWriter, r *http.Request) {
	idxPage := model.IndexPage{}
	osinfo := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	idxPage.FooterData.OsString = osinfo["NAME"]
	idxPage.FooterData.OsVersion = osinfo["VERSION_ID"]

	idxPage.Settings = model.ListAllSettings()
	fmt.Println("(OPEN): " + model.HttpDir + "templates/settings.html")
	footer := model.HttpDir + "templates/footer.html"
	t, err := template.ParseFiles(model.HttpDir+"templates/settings.html", footer)
	checkError(err)
	t.Execute(w, idxPage)
}

func LogFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("(OPEN): " + model.HttpDir + "templates/logPage.html")
	logPage := model.LogPage{}
	osinfo := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	logPage.FooterData.OsString = osinfo["NAME"]
	logPage.FooterData.OsVersion = osinfo["VERSION_ID"]

	data, err := logreader.ReadLog()
	checkError(err)
	logPage.AllMessages = *data

	footer := model.HttpDir + "templates/footer.html"
	t, err := template.ParseFiles(model.HttpDir+"templates/logPage.html", footer)
	checkError(err)
	t.Execute(w, logPage)
}

func RebootFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("(OPEN): " + model.HttpDir + "templates/rebootPage.html")
	rebootPage := model.RebootPage{}
	osinfo := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	rebootPage.FooterData.OsString = osinfo["NAME"]
	rebootPage.FooterData.OsVersion = osinfo["VERSION_ID"]

	footer := model.HttpDir + "templates/footer.html"
	t, err := template.ParseFiles(model.HttpDir+"templates/rebootPage.html", footer)
	checkError(err)
	t.Execute(w, rebootPage)
}

///@todo denna buggar om data.message nill
func StatusFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("(OPEN): " + model.HttpDir + "templates/statusPage.html")
	statusPage := model.StatusPage{}
	osinfo := syscmd.ParseOsRelease(syscmd.OsReleasePath)

	statusPage.FooterData.OsString = osinfo["NAME"]
	statusPage.FooterData.OsVersion = osinfo["VERSION_ID"]

	data := logreader.ReadStatus()
	statusPage.SystemdStatus = *data
	var temp = model.HttpDir + "templates/error.html"

	if data.Active == true {
		fmt.Println("Systemd service works")
		temp = model.HttpDir + "templates/message.html"
	}

	footer := model.HttpDir + "templates/footer.html"

	t, err := template.ParseFiles(model.HttpDir+"templates/statusPage.html", temp, footer)
	checkError(err)
	t.Execute(w, statusPage)
}

func RestartWebui(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Restart webui.")
	if r.Method == "GET" {
		fmt.Println("TODO: implement GET for RestartWebui")
		http.Redirect(w, r, "/", 301)
	} else {
		fmt.Println("TODO: implement POST for RestartWebui")
		syscmd.CmdRestart(syscmd.Webservice)
		http.Redirect(w, r, "/", 301)
	}
}

func RestartSensor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Restart sensor.")
	if r.Method == "GET" {
		fmt.Println("TODO: implement GET for RestartSensor")
		http.Redirect(w, r, "/", 301)
	} else {
		syscmd.CmdRestart(syscmd.Tempsensorservice)
		///@todo redirect to restarting.html
		http.Redirect(w, r, "/", 301)
	}
}
