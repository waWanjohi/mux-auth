package handlers

import (
	"net/http"
	"text/template"
)

var IndexHandler = func(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("static/index.html"))

	if r.Method != http.MethodGet {
		t.Execute(w, "<h1>Method Not Allowed!</h1>")
	}
	t.Execute(w, nil)
}
