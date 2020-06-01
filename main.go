package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello world")
	http.HandleFunc("'/", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {

}
