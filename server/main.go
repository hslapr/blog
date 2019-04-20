package main

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// templateDir is the directory of templates
var templateDir = "../templates/"

// init does initialization work
func init() {

}

// templatePath returns the path of the template file
func templatePath(filename string) string {
	return filepath.Join(templateDir, filename+".html")
}

// parseTemplates parses given templates
func parseTemplates(filenames ...string) (*template.Template, error) {
	paths := make([]string, len(filenames))
	for i, filename := range filenames {
		paths[i] = templatePath(filename)
	}
	return template.ParseFiles(paths...)
}

// indexHandler handles the index page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := parseTemplates("layouts/default", "post")
	if err != nil {
		log.Fatal(err)
	}
	templ.ExecuteTemplate(w, "layout", nil)
	// fmt.Fprintf(w, "Hello, world!\n")
}

// main starts the program
func main() {
	// data.Download()
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":2333", nil))
}
