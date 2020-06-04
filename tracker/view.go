package tracker

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	Render(w, "index.html", r)
}

func Search(w http.ResponseWriter, r *http.Request) {
	Render(w, "search.html", r)
}

func Find(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	userName := r.FormValue("findUser")
	userInfo := GitHubInfo{}
	userInfo.Name = userName
	fmt.Println("Your username is", userName)
	RenderWithView(w, "results.html", r, userInfo)
}


func Render(w http.ResponseWriter, tmpl string, r *http.Request) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Fatalln(err)
	}

	err = t.Execute(w, nil)
}

func RenderWithView(w http.ResponseWriter, tmpl string, r *http.Request, data interface{})  {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Fatalln(err)
	}

	err = t.Execute(w, data)
}