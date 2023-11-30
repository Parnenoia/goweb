package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "FFC"},
				{Title: "Blade Runnder", Director: "Ridley Scott"},
			},
		}
		tmpl.Execute(w, films)
	}
	http.HandleFunc("/", h1)

	h2 := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		htmlStr := fmt.Sprintf("<h2>%s</h2><h4>%s</h4>", title, director)
		tmpl, _ := template.New("new_movie").Parse(htmlStr)
		tmpl.Execute(w, nil)
		fmt.Print("Movie with title:", title, "and director:", director, "added\n")
		log.Print("HTMX request received")
		log.Print(r.Header.Get("HX-request"))
	}
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
