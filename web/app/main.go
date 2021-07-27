package main

import (
	"log"
	"net/http"
	"os"

	"github.com/brianshef/knavigator/web/handlers"
)

const defaultPort = "3000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/character", handlers.CharacterHandler)
	log.Fatal(http.ListenAndServe(":"+port, mux))

}
