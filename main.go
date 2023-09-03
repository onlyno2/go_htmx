package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/onlyno2/go_htmx/src/db"
	"github.com/onlyno2/go_htmx/src/models"
	"github.com/onlyno2/go_htmx/src/routes"
)

func main() {
	db.Open()

	db.Db.AutoMigrate(&models.Todo{})

	assetsFs := http.FileServer(http.Dir("src/public/assets"))
	http.Handle("/stylesheets/", assetsFs)

	http.HandleFunc("/todos", routes.GetTodos)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("src/views/index.html"))
		template.Execute(w, nil)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
