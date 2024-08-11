package main

import (
	"log"
	"net/http"

	"github.com/Ayomided/ItsAfullTimeJob.git/components"
	"github.com/a-h/templ"
)

// //go:embed static/*
// var static embed.FS

func main() {
	homepage := components.Index()
	formPage := components.Form()
	counter := components.Counter()

	router := http.NewServeMux()

	router.Handle("/", templ.Handler(homepage))
	router.Handle("/form", templ.Handler(formPage))
	router.Handle("/counter", templ.Handler(counter))

	log.Fatalln(http.ListenAndServe(":3000", router))
}
