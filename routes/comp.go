package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/LinuxSploit/emailsoft/engine"
	"github.com/LinuxSploit/emailsoft/globle"
)

func Comp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		r.ParseForm()
		compid := r.FormValue("id")
		id, err := strconv.Atoi(compid)
		if err != nil {
			fmt.Fprintln(w, "Invalid ID!")
			return
		}

		for _, v := range globle.Compagins {
			if v.UID == id {
				tmpl, err := template.ParseFiles("./templates/comp.html")
				if err != nil {
					fmt.Fprintln(w, err)
				}
				tmpl.Execute(w, []engine.Compagin{v})
			}
		}

	}
}
