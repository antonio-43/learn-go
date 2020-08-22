package comhtml

import (
	"html/template"
	"log"
	"net/http"
)

func HandlerComHTML(w http.ResponseWriter, r *http.Request) {

	temp, err := template.ParseFiles("templates/index.html")
	if r.URL.Path != "/comhtml" {
		http.Error(w, "Error", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "ERROR METHOD", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}
