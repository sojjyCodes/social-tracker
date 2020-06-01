package tracker

import (
	"html/template"
	"log"
	"net/http"
)

func render(w http.ResponseWriter, tmpl string, r *http.Request) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Fatalln(err)
	}

	err = t.Execute(w, nil)
}