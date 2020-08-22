package main

import (
	"log"
	"net/http"
	"./routes"
	"./comhtml"
)

func main() {

	http.HandleFunc("/", routehello.HandlerHello)
	http.Handle("/comhtml", comhtml.HandlerComHTML)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
