package main

import (
	"fmt"
	"github.com/gabielfemi/my-version-social-tracker/tracker"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	router.HandleFunc("/search", Search).Methods("GET")
	router.HandleFunc("/search", Find).Methods("POST")
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))


	server := &http.Server{
		Addr: "127.0.0.1:8000",
		Handler: router,

	}
	log.Println("Listening on 127.0.0.1:8000")
	log.Fatalln(server.ListenAndServe())

}

func Index(w http.ResponseWriter, r *http.Request) {
	tracker.Render(w, "index.html", r)
}

func Search(w http.ResponseWriter, r *http.Request) {
	tracker.Render(w, "search.html", r)
}

func Find(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	userName := r.FormValue("findUser")
	fmt.Println("Your username is", userName)
	w.WriteHeader(200)
}