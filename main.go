package main

import (
	"log"
	"net/http"
	"split/views"

	"github.com/a-h/templ"
)

func init() {
	MakeMigrations()
}

func main() {

	http.Handle("/", templ.Handler(views.Index()))

	log.Println("🚀 Starting up on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
