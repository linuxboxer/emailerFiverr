package routes

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Getfilenames(w http.ResponseWriter, r *http.Request) {
	finfo, err := ioutil.ReadDir("./uploads/")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range finfo {
		if !v.IsDir() {

			w.Write([]byte(v.Name() + "/"))

		}
	}

}
