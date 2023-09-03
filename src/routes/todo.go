package routes

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/onlyno2/go_htmx/src/db"
	"github.com/onlyno2/go_htmx/src/models"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	db.Db.Find(&todos)
	template := template.Must(template.ParseFiles("src/views/todos/index.html"))
	fmt.Println(len(todos))

	template.Execute(w, todos)
}
