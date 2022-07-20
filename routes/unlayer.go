package routes

import (
	"fmt"
	"html/template"
	"net/http"
)

func Unlayer(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/unlayer.html")
	if err != nil {
		fmt.Fprintln(w, err)
	}
	tmpl.Execute(w, nil)
}
