package routes

import (
	"net/http"
	"text/template"
)

type MainPage struct {
}

func ServeMainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/main/main.html"))
	data := MainPage{}

	tmpl.Execute(w, data)
}
