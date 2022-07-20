package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/LinuxSploit/emailsoft/engine"
)

func Setting(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("./templates/setting.html")
		if err != nil {
			fmt.Fprintln(w, err)
		}
		tmpl.Execute(w, nil)
	case "POST":
		r.ParseForm()

		host := r.FormValue("host")
		port := r.FormValue("port")

		youremail := r.FormValue("youremail")
		passcode := r.FormValue("passcode")

		if host != "" && port != "" && youremail != "" && passcode != "" {
			err := engine.AppSetting.Write(host, port, youremail, passcode)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Fprintln(w, "Please fill all the fields of forms")
			return
		}

		//

		tmpl, err := template.ParseFiles("./templates/setting.html")
		if err != nil {
			fmt.Fprintln(w, err)
		}
		tmpl.Execute(w, nil)

	}
}
