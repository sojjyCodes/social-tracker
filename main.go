package main

import (
	"github.com/gabielfemi/my-version-social-tracker/tracker"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/search", Search)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Println("Listening on 127.0.0.1:8000")
	_ = http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	tracker.Render(w, "index.html", r)
}

func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//_, _ = fmt.Fprintf(w, "You searched for %s", r.FormValue("searchPhrase"))
		tracker.Render(w, "search.html", r)
	} else {
		http.Redirect(w, r , "/", 301)
	}

}