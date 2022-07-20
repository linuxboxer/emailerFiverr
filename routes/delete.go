package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/LinuxSploit/emailsoft/globle"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uid := r.FormValue("uid")
	uidint, err := strconv.Atoi(uid)
	if err != nil {
		fmt.Fprintln(w, "invalid input!")
		return
	}
	for i, v := range globle.Compagins {
		if uidint == v.UID {
			globle.Compagins = append(globle.Compagins[:i], globle.Compagins[i+1:]...)

		}
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
