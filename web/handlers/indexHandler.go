package handlers

import "net/http"

// IndexHandler is responsible for serving up an initial index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1><a href=\"/character\">Generate a character</a></h1>"))
}