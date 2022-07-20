package routes

import (
	"fmt"
	"net/http"

	"github.com/LinuxSploit/emailsoft/globle"
)

func Allemails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, globle.Singlelistemails)
}
