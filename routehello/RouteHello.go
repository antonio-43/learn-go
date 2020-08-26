package routehello

import (
	"fmt"
	"net/http"
)

func HandlerHello(w http.ResponseWriter, r *http.Request) {

	// handle in route /
	/*
	*
	 */

	if r.URL.Path != "/" {
		http.Error(w, "404 NOT FOUND TROLLEIIIIIIIIII! HEHE BOY", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "NETHOD ERROR", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello World!")
}
