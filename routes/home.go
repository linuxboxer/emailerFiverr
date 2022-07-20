package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/LinuxSploit/emailsoft/globle"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cache-Control", "no-cache")
	tmpl, err := template.ParseFiles("./templates/home.html")
	if err != nil {
		fmt.Fprintln(w, "Error loading template!")
	}

	tmpl.Execute(w, globle.Compagins)
}
