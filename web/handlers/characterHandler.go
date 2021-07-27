package handlers

import (
	"net/http"
	

	"github.com/brianshef/knavigator/api"
)

func renderCharacter(w http.ResponseWriter, tmpl string, c *api.Character) {
	err := templates.ExecuteTemplate(w, tmpl+".html", c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// CharacterHandler is responsible for serving up
// a rendered character template
func CharacterHandler(w http.ResponseWriter, r *http.Request) {
	c, err := api.GenerateCharacter()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	renderCharacter(w, "character", c)
}
