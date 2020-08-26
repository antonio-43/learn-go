package comhtml

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func HandlerComload(w http.ResponseWriter, r *http.Request) {

	// random comment /

	// cjrhviuwbviuh

	if r.URL.Path != "/comhtml" {
		fmt.Fprintf(w, "Path diferente")
		return
	}

	if r.Method != "GET" {
		fmt.Fprintf(w, "erro de metodo")
		return
	}

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, ""); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
