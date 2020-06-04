package main

import (
	"github.com/gabielfemi/my-version-social-tracker/tracker"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", tracker.Index)
	router.HandleFunc("/search", tracker.Search).Methods("GET")
	router.HandleFunc("/search", tracker.Find).Methods("POST")
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))


	server := &http.Server{
		Addr: "127.0.0.1:8000",
		Handler: router,

	}
	log.Println("Listening on", server.Addr)
	log.Fatalln(server.ListenAndServe())

}

