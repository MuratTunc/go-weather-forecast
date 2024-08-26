package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const port = ":8084"
const tmpl_page = "../../templates/home_page.html"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, tmpl_page)
	})

	fmt.Printf("WEB PAGE FRONT END SERVICES IS STARTED!!! AT LOCALHOST PORT %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Panic(err)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles(tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Panic("Error parsing template html page...")
		return
	}
}
