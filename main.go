package main

import (
	"github.com/gabielfemi/my-version-social-tracker/tracker"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Println("Listening on 127.0.0.1:8000")
	_ = http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	tracker.Render(w, "index.html", r)
}
