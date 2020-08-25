package comhtml

import (
	"html/template"
	"net/http"
)

/*
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

	p := "./html"

	http.FileServer(http.Dir(p))

}
*/

func ServeTemplate(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseGlob("templates/*.html"))

	templates.ExecuteTemplate(w, "jj.html", nil)
}
