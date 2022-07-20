package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/LinuxSploit/emailsoft/globle"
)

func Analytics(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/analytics.html")
	if err != nil {
		fmt.Fprintln(w, err)
	}
	tmpl.Execute(w, globle.Compagins)

}
