package handler

import (
	"html/template"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	films := map[string][]Film{
		"Films": {
			{Title: "The Godfather", Director: "FFC"},
			{Title: "Blade Runnder", Director: "Ridley Scott"},
		},
	}
	tmpl.Execute(w, films)

	// title := r.PostFormValue("title")
	// director := r.PostFormValue("director")
	// htmlStr := fmt.Sprintf("<h2>%s</h2><h4>%s</h4>", title, director)
	// tmpl, _ := template.New("new_movie").Parse(htmlStr)
	// tmpl.Execute(w, nil)
	// fmt.Print("Movie with title:", title, "and director:", director, "added\n")
	// log.Print("HTMX request received")
	// log.Print(r.Header.Get("HX-request"))
}
func main() {
	http.HandleFunc("/", Handler)
	// http.HandleFunc("/add-film/", h2)

	// log.Fatal(http.ListenAndServe(":8000", nil))

}
