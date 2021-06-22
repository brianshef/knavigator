package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

const defaultPort = "3000"

var templates = template.Must(template.ParseFiles("web/template/default.html"))

type Page struct {
	Title string
	Body  string
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World!</h1>"))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	renderTemplate(w, title, &Page{
		Title: title,
		Body:  "Character information will be displayed here",
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":"+port, mux))

}
