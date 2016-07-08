package main

import (
	template "html/template"
	// "log"
	"net/http"
	"strings"
	// "strconv"

	"adventurers_tools/adventurer"
	"adventurers_tools/rules"

	"github.com/GeertJohan/go.rice"
	"github.com/Sirupsen/logrus"
	// "github.com/Sirupsen/logrus/formatters/logstash"
)

var activeRules rules.Rules
var theAdventurer adventurer.Adventurer

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	// logrus.SetFormatter(&logstash.LogstashFormatter{Type: "application_name"})
	logrus.Debug("set log level debug")
	// logrus.SetLevel(logrus.WarnLevel)

	activeRules = rules.GetAdventurersFirstEditionRules()
	activeRules.SetOptionalCanTradeCoinsForStats(true)
	// activeRules = rules.GetAdventurersRevisedRules()
	logrus.WithField("rules_edition", activeRules.RulesEdition()).Info("loaded rules")

  theAdventurer = adventurer.NewAdventurer()
	logrus.WithField("theAdventurer", theAdventurer).Info("loaded Adventurer")
}

func main() {
	logrus.Info("initializing...")
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
	http.ListenAndServe(bind, nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {

	view := getView(r)
	activeMenu := getActiveMenu(view)

	model := map[string]interface{}{
		"view":         view,
		"acviveMenu":   activeMenu,
	}

	addRulesToModel(model)

  addAdventurerToModel(model)

	tmplMain.Execute(w, model)

}

func addAdventurerToModel(model map[string]interface{}) {
  model["char"] = theAdventurer
	// model["char_strength"] = theAdventurer.Strength()
	// model["char_agility"] = theAdventurer.Agility()
	// model["char_mind"] = theAdventurer.Mind()
	model["CreationPoints"] = theAdventurer.CreationPoints()
}

func addRulesToModel(model map[string]interface{}) {
  model["rules"] = activeRules
	model["rules_min_stat"] = activeRules.MinStat()
	model["rules_max_stat"] = activeRules.MaxStat()
	model["rules_min_end"] = activeRules.MinEnd()
	model["rules_max_end"] = activeRules.MaxEnd()
	model["rules_initial_stat_points"] = activeRules.InitialStatPoints()
	model["rules_max_stat_points"] = activeRules.MaxStatPoints()
}

var VIEWS = []string{"home", "about", "options", "create_1", "create_2", "create_3", "create_4", "create_5", "load_save_print"}

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
	logrus.WithFields(logrus.Fields{"view": view, "path": path}).Debug("found view")
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
		"name": name,
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
