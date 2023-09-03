package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	assetsFs := http.FileServer(http.Dir("src/public/assets"))
	http.Handle("/stylesheets/", assetsFs)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("src/views/index.html"))
		template.Execute(w, nil)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
