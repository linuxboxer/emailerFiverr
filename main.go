package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/LinuxSploit/emailsoft/globle"
	"github.com/LinuxSploit/emailsoft/routes"
)

func main() {

	go tickeroye()

	http.HandleFunc("/", routes.Home)
	http.HandleFunc("/delete", routes.Delete)
	http.HandleFunc("/create", routes.Create)
	http.HandleFunc("/unlayer", routes.Unlayer)
	http.HandleFunc("/upload", routes.Upload)
	http.HandleFunc("/csv", routes.CSV)
	http.HandleFunc("/analytics", routes.Analytics)
	http.HandleFunc("/comp", routes.Comp)

	http.HandleFunc("/getfilenames", routes.Getfilenames)
	http.HandleFunc("/setting", routes.Setting)

	// tracking route
	http.HandleFunc("/pixel", routes.Pixel)
	// debug route
	http.HandleFunc("/debug", routes.Allemails)
	http.ListenAndServe(":4556", nil)
}

func tickeroye() {
	for range globle.TK.C {
		for _, v := range globle.Compagins {
			if v.RunTime.Format("2006-01-02T15:04:00Z") == time.Now().Format("2006-01-02T15:04:50Z") {
				updatedcomp := v.Run()

				for i, x := range globle.Compagins {
					if v.UID == x.UID {
						v = updatedcomp
						v.Pending = false
						v.Done = true
						fmt.Println("Delte!")
						globle.Compagins = append(globle.Compagins[:i], globle.Compagins[i+1:]...)

						globle.Compagins = append(globle.Compagins, v)
					}
				}

			}
		}
	}
}
