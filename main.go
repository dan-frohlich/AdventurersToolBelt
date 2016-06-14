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

  jsBox := rice.MustFindBox("public/js/")
	jsFileServer := http.StripPrefix("/js/", http.FileServer(jsBox.HTTPBox()))
	http.Handle("/js/", jsFileServer)

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

	view := getView( r )
  activeMenu := getActiveMenu( view )
	m := map[string]string{
		"Message":      "Hello, world!",
		"Name":         "Fred",
		"TitleMessage": "This is the stuff you love!",
		"view":         view,
    "acviveMenu":   activeMenu,
	}

  // add query params for now...
  values := r.URL.Query()
  for k, z := range values {
    for _, v := range z {
      m[k] = v
    }
  }

  log.Println("m", m)
	tmplMain.Execute(w, m)

}

var VIEWS = []string{"home", "about", "options", "create_1", "create_2", "create_3", "create_4", "create_5", "load_save_print"}

func getActiveMenu( view string ) string {
  if strings.HasPrefix(view, "create") {
    return "create"
  }
  return view
}

func getView(r *http.Request ) string {
  path:= r.URL.Path
  pathElements := strings.Split(path, "/")
	view := "home"
	if len(pathElements) > 1 && pathElements[1] != "" {
    for _, v := range VIEWS {
      if v == pathElements[1] {
        view = pathElements[1]
      }
    }
	}
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
