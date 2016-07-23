package main

import (
	template "html/template"
	// "log"
	"net/http"
	"strings"
	// "strconv"

	"github.com/dan-frohlich/AdventurersToolBelt/adventurer"
	"github.com/dan-frohlich/AdventurersToolBelt/rules"
	"github.com/dan-frohlich/AdventurersToolBelt/storage"

	"github.com/GeertJohan/go.rice"
	"github.com/Sirupsen/logrus"
	// "github.com/Sirupsen/logrus/formatters/logstash"
)

var activeRules *rules.Rules
var theAdventurer adventurer.Adventurer

func init() {
	level := logrus.DebugLevel
	logrus.SetLevel(level)
	// logrus.SetFormatter(&logstash.LogstashFormatter{Type: "application_name"})
	logrus.WithField("level", level).Debug("set log level")
	// logrus.SetLevel(logrus.WarnLevel)

	activeRules = rules.GetAdventurersFirstEditionRules()
	activeRules.SetOptionalCanTradeCoinsForStats(true)
	// activeRules = rules.GetAdventurersRevisedRules()
	logrus.WithField("rules_edition", activeRules.RulesEdition).Info("loaded rules")

	theAdventurer = adventurer.NewAdventurer(*activeRules)
	logrus.WithField("theAdventurer", theAdventurer).Info("loaded Adventurer")
}

func main() {
	initHomeDir()
	initWebServer()
}

func initHomeDir() {
	storage.Bootstrap()
	storage.List()
}

func initWebServer() {
	logrus.Debug("initializing Web Server...")
	initTemplates()
	http.HandleFunc("/", mainHandler)

	cssBox := rice.MustFindBox("public/css/")
	cssFileServer := http.StripPrefix("/css/", http.FileServer(cssBox.HTTPBox()))
	http.Handle("/css/", cssFileServer)

	jsBox := rice.MustFindBox("public/js/")
	jsFileServer := http.StripPrefix("/js/", http.FileServer(jsBox.HTTPBox()))
	http.Handle("/js/", jsFileServer)

	imagesBox := rice.MustFindBox("public/images/")
	imagesFileServer := http.StripPrefix("/images/", http.FileServer(imagesBox.HTTPBox()))
	http.Handle("/images/", imagesFileServer)

	bind := ":8080"
	logrus.Info("listening on ", bind)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		logrus.WithField("error", err).Fatal("http.ListenAndServe failed.")
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {

	view := getView(r)
	activeMenu := getActiveMenu(view)

	model := map[string]interface{}{
		"view":       view,
		"acviveMenu": activeMenu,
	}

	addRulesToModel(model)

	addAdventurerToModel(model)

	tmplMain.Execute(w, model)

}

func addAdventurerToModel(model map[string]interface{}) {
	model["adventurer"] = theAdventurer
}

func addRulesToModel(model map[string]interface{}) {
	model["rules"] = activeRules
}

var VIEWS = []string{"home", "about", "options", "create_general", "create_stats", "create_skills", "create_gear", "create_5", "load_save_print"}

func getActiveMenu(view string) string {
	if strings.HasPrefix(view, "create") {
		return "create"
	}
	return view
}

func getView(r *http.Request) string {
	path := r.URL.Path
	pathElements := strings.Split(path, "/")
	view := "home"
	if len(pathElements) > 1 && pathElements[1] != "" {
		for _, v := range VIEWS {
			if v == pathElements[1] {
				view = pathElements[1]
			}
		}
	}
	// logrus.WithFields(logrus.Fields{"view": view, "path": path}).Debug("found view")
	return view
}

var templateBox *rice.Box

var tmplMain *template.Template

// var tmplH1 *template.Template
var err error

func initTemplates() {
	// find a rice.Box
	templateBox, err = rice.FindBox("templates")
	if err != nil {
		logrus.Fatal(err)
	}
	tmplMain = initTemplate("main", "main.tmpl")

	// tmplH1 = initTemplate("header_1", "header_1.tmpl")

}

func initTemplate(name, filename string) *template.Template {
	logrus.WithFields(logrus.Fields{
		"name":     name,
		"template": filename,
	}).Info("initializing template...")
	// get file contents as string
	text, err := templateBox.String(filename)
	if err != nil {
		logrus.Fatal(err)
	}

	tmpl, err := template.New(name).Parse(text)
	if err != nil {
		logrus.Fatal(err)
	}
	return tmpl
}
