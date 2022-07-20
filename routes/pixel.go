package routes

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/LinuxSploit/emailsoft/globle"
)

func Pixel(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		r.ParseForm()
		view := r.FormValue("id")
		viewdata, err := base64.StdEncoding.DecodeString(view)
		if err != nil {
			return
		}
		viewdataS := strings.Split(string(viewdata), ":")
		fmt.Println(viewdataS)

		for i, v := range globle.Compagins {
			if v.Title == viewdataS[0] {
				for _, d := range v.ViewdEmail {
					if d == viewdataS[1] {
						return
					}
				}
				v.AddViewedEmail(viewdataS[1])

				globle.Compagins = append(globle.Compagins[:i], globle.Compagins[i+1:]...)

				globle.Compagins = append(globle.Compagins, v)
			}
		}
		imgg, err := ioutil.ReadFile("./lee.jpg")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(w, string(imgg))
	}
}
