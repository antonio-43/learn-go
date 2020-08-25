package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"

	//"./comhtml"
	"./routes"
)

var templates *template.Template

func main() {
	// http.HandleFunc("/mclaren", fs)
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/", routehello.HandlerHello)
	r.HandleFunc("/comhtml", ServeTemplate)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func ServeTemplate(w http.ResponseWriter, r *http.Request) {
	//templates := template.Must(template.ParseGlob("templates/*.html"))

	templates.ExecuteTemplate(w, "jj.html", nil)
}
