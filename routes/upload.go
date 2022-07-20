package routes

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/LinuxSploit/emailsoft/globle"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	var allfile []string

	switch r.Method {
	case "GET":
		r.ParseForm()
		delfile := r.FormValue("del")
		if delfile != "" {
			err := os.Remove("./uploads/" + delfile)
			if err != nil {
				fmt.Fprintln(w, "delete failed!\n filename:", delfile)
			}
		}

		finfo, err := ioutil.ReadDir("./uploads/")
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range finfo {
			if !v.IsDir() {
				allfile = append(allfile, v.Name())
			}
		}
		fmt.Println(allfile)
		globle.Allfile = allfile

		tmpl, err := template.ParseFiles("./templates/upload.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, allfile)
	case "POST":
		r.ParseForm()

		var filedata bytes.Buffer
		file, header, err := r.FormFile("file")
		if err != nil {
			fmt.Fprintln(w, "Error uploading file!")
			return
		}
		io.Copy(&filedata, file)
		if len(filedata.Bytes()) > 26214400 {
			fmt.Fprintln(w, "file size must be less than 25MB!")
			return
		}
		err = ioutil.WriteFile("./uploads/"+header.Filename, filedata.Bytes(), 0644)
		if err != nil {
			fmt.Fprintln(w, "error saving file!")
		}

		finfo, err := ioutil.ReadDir("./uploads/")
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range finfo {
			if !v.IsDir() {
				allfile = append(allfile, v.Name())
			}
		}

		tmpl, err := template.ParseFiles("./templates/upload.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, allfile)
	}
}
