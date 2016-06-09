package main

import (
	template "html/template"
	"log"
	"net/http"
	"strings"

	"github.com/GeertJohan/go.rice"
	// "github.com/thrisp/djinn"
)

func main() {
	log.Println("initializing...")
	initTemplates()
	http.HandleFunc("/", mainHandler)

	cssBox := rice.MustFindBox("public/css/")
	cssFileServer := http.StripPrefix("/css/", http.FileServer(cssBox.HTTPBox()))
	http.Handle("/css/", cssFileServer)

	imagesBox := rice.MustFindBox("public/images/")
	imagesFileServer := http.StripPrefix("/images/", http.FileServer(imagesBox.HTTPBox()))
	http.Handle("/images/", imagesFileServer)

	log.Println("listening...")
	http.ListenAndServe(":8080", nil)
}

// var J *djinn.Djinn
// var templateBox *rice.Box
// var err error
//
// func initTemplates() {
// 	templateBox, err = rice.FindBox("templates")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	m := map[string]string{
// 		"main":     load("main.tmpl"),
// 		"header_1": load("header_1.tmpl"),
// 	}
// 	J := djinn.New()
//
//   djinn.Loaders(&djinn.MapLoader{ TemplateMap: m})(J)
// 	// J.AddLoaders(&MapLoader{m: &m})
//
// }
//
// func load(file string) string {
// 	text, err := templateBox.String(file)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return text
// }

// func mainHandler(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// 	// parse and execute the template
// 	J.Render(w, "main", map[string]string{
// 		"Message":      "Hello, world!",
// 		"Name":         "Fred",
// 		"TitleMessage": "This is the stuff you love!",
// 	})
//
// }

func mainHandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	components := strings.Split(path, "/")
  log.Println("components", components)
  log.Println("len", len(components))
  log.Println("cap", cap(components))
	view := "home"
	if len(components) > 1 && components[1] != "" {
		view = components[1]
	}
	m := map[string]string{
		"Message":      "Hello, world!",
		"Name":         "Fred",
		"TitleMessage": "This is the stuff you love!",
		"view":         view,
	}
  log.Println("m", m)
	tmplMain.Execute(w, m)

}

var templateBox *rice.Box

var tmplMain *template.Template

// var tmplH1 *template.Template
var err error

func initTemplates() {
	// find a rice.Box
	templateBox, err = rice.FindBox("templates")
	if err != nil {
		log.Fatal(err)
	}
	tmplMain = initTemplate("main", "main.tmpl")

	// tmplH1 = initTemplate("header_1", "header_1.tmpl")

}

func initTemplate(name, filename string) *template.Template {
	// get file contents as string
	text, err := templateBox.String(filename)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New(name).Parse(text)
	if err != nil {
		log.Fatal(err)
	}
	return tmpl
}
