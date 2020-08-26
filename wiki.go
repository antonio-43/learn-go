package main

import (
	"log"
	"net/http"

	"./comhtml"
	"./routehello"
)

func main() {

	http.HandleFunc("/", routehello.HandlerHello)
	http.HandleFunc("/comhtml", comhtml.HandlerComload)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
