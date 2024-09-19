package main

import (
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("temp/inputPage.html"))
	tmpl.Execute(w, nil)
}
