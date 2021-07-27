package handlers

import "html/template"

const templatesDir = "web/template/"

var templates = template.Must(
	template.ParseFiles(
		templatesDir+"default.html",
		templatesDir+"character.html",
	))
