package routes

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/LinuxSploit/emailsoft/globle"
)

func CSV(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		var tmpsinglelist []string

		asubscribe, err := os.Open("./CSV/subscribe.csv")
		if err != nil {
			fmt.Fprintln(w, err)
		}
		subreader := csv.NewReader(asubscribe)
		subrecords, _ := subreader.ReadAll()
		for o, u := range subrecords {
			tmpsinglelist = append(tmpsinglelist, u[2])
			globle.AllContact.Subzcount = o
		}
		globle.AllContact.Subscribe = subrecords
		///////////////////////////////////////////////
		aunsubscribe, err := os.Open("./CSV/unsubscribe.csv")
		if err != nil {
			fmt.Fprintln(w, err)
		}
		unsubreader := csv.NewReader(aunsubscribe)
		unsubrecords, _ := unsubreader.ReadAll()
		for p, u := range unsubrecords {
			tmpsinglelist = append(tmpsinglelist, u[2])
			globle.AllContact.Unsubzcount = p
		}
		globle.AllContact.Unsubscribe = unsubrecords

		///////////////////////////////////////////////
		anotsubscribe, err := os.Open("./CSV/notsubscribe.csv")
		if err != nil {
			fmt.Fprintln(w, err)
		}
		notsubreader := csv.NewReader(anotsubscribe)
		notsubrecords, _ := notsubreader.ReadAll()
		for q, u := range notsubrecords {
			tmpsinglelist = append(tmpsinglelist, u[2])
			globle.AllContact.Notsubzcount = q
		}
		globle.AllContact.Notsubscribe = notsubrecords

		// updating all single email list
		globle.Singlelistemails = tmpsinglelist

		tmpl, err := template.ParseFiles("./templates/CSV.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, globle.AllContact)

	case "POST":
		r.ParseForm()

		var filedata bytes.Buffer

		// file reading and error handling
		file, header, err := r.FormFile("file")
		if err != nil {
			fmt.Fprintln(w, "Error uploading file!")
			return
		}
		tmpfilnameparts := strings.Split(header.Filename, ".")
		if tmpfilnameparts[len(tmpfilnameparts)-1] != "csv" {
			fmt.Fprintln(w, tmpfilnameparts[len(tmpfilnameparts)-1], " invalid file format please upload csv file format")
			return
		}

		upcategory := r.FormValue("type")
		if upcategory == "" {
			fmt.Fprintln(w, "Please select Category in the form")
			return
		}
		upaction := r.FormValue("action")
		if upaction == "" {
			fmt.Fprintln(w, "Please select replace/append in the form")
			return
		}

		io.Copy(&filedata, file)
		if len(filedata.Bytes()) > 26214400 {
			fmt.Fprintln(w, "file size must be less than 25MB!")
			return
		}

		if upcategory == "subscribe" {
			switch upaction {
			case "replace":
				upcsv := csv.NewReader(&filedata)
				csvtitle, err := upcsv.Read()
				if err == io.EOF {
					fmt.Fprintln(w, "Debug: #Empty-File#")
					return
				}
				if err != nil {
					fmt.Fprintln(w, "Debug: Error in File!")
				}
				os.Remove("./CSV/subscribe.csv")
				downfile, err := os.OpenFile("./CSV/subscribe.csv", os.O_CREATE|os.O_WRONLY, 0644)

				if err != nil {
					fmt.Fprintln(w, err)
				}
				defer downfile.Close()

				downcsv := csv.NewWriter(downfile)
				downcsv.Write(csvtitle)
				downcsv.Flush()

				var dup bool
				for {
					uprecord, err := upcsv.Read()
					if err == io.EOF {
						fmt.Println("Debug: #File-END#")
						break
					}
					if err != nil {
						fmt.Println("Debug: #File-END#")
						return
					}
					for _, v := range globle.Singlelistemails {
						if uprecord[2] == v {
							dup = true
							break
						}
					}
					if !dup {

						err := downcsv.Write(uprecord)
						downcsv.Flush()
						fmt.Println("Debug: [Write][err]", uprecord, err)

					}
					dup = false
				}
			case "append":
				upcsv := csv.NewReader(&filedata)
				_, err := upcsv.Read()
				if err == io.EOF {
					fmt.Fprintln(w, "Debug: #Empty-File#")
					return
				}
				if err != nil {
					fmt.Fprintln(w, "Debug: Error in File!")
				}
				downfile, err := os.OpenFile("./CSV/subscribe.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

				if err != nil {
					fmt.Fprintln(w, err)
				}
				defer downfile.Close()

				downcsv := csv.NewWriter(downfile)

				var dup bool
				for {
					uprecord, err := upcsv.Read()
					if err == io.EOF {
						fmt.Println("Debug: #File-END#")
						break
					}
					if err != nil {
						fmt.Println("Debug: #File-END#")
						return
					}
					for _, v := range globle.Singlelistemails {
						if uprecord[2] == v {
							dup = true
							break
						}
					}
					if !dup {

						err := downcsv.Write(uprecord)
						downcsv.Flush()
						fmt.Println("Debug: [Write][err]", uprecord, err)

					}
					dup = false
				}
			}
		} else if upcategory == "unsubscribe" {
			switch upaction {
			case "replace":
				upcsv := csv.NewReader(&filedata)
				csvtitle, err := upcsv.Read()
				if err == io.EOF {
					fmt.Fprintln(w, "Debug: #Empty-File#")
					return
				}
				if err != nil {
					fmt.Fprintln(w, "Debug: Error in File!")
				}
				os.Remove("./CSV/unsubscribe.csv")
				downfile, err := os.OpenFile("./CSV/unsubscribe.csv", os.O_CREATE|os.O_WRONLY, 0644)

				if err != nil {
					fmt.Fprintln(w, err)
				}
				defer downfile.Close()

				downcsv := csv.NewWriter(downfile)
				downcsv.Write(csvtitle)
				downcsv.Flush()

				var dup bool
				for {
					uprecord, err := upcsv.Read()
					if err == io.EOF {
						fmt.Println("Debug: #File-END#")
						break
					}
					if err != nil {
						fmt.Println("Debug: #File-END#")
						return
					}
					for _, v := range globle.Singlelistemails {
						if uprecord[2] == v {
							dup = true
							break
						}
					}
					if !dup {

						err := downcsv.Write(uprecord)
						downcsv.Flush()
						fmt.Println("Debug: [Write][err]", uprecord, err)

					}
					dup = false
				}
			case "append":
				upcsv := csv.NewReader(&filedata)
				_, err := upcsv.Read()
				if err == io.EOF {
					fmt.Fprintln(w, "Debug: #Empty-File#")
					return
				}
				if err != nil {
					fmt.Fprintln(w, "Debug: Error in File!")
				}
				downfile, err := os.OpenFile("./CSV/unsubscribe.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

				if err != nil {
					fmt.Fprintln(w, err)
				}
				defer downfile.Close()

				downcsv := csv.NewWriter(downfile)

				var dup bool
				for {
					uprecord, err := upcsv.Read()
					if err == io.EOF {
						fmt.Println("Debug: #File-END#")
						break
					}
					if err != nil {
						fmt.Println("Debug: #File-END#")
						return
					}
					for _, v := range globle.Singlelistemails {
						if uprecord[2] == v {
							dup = true
							break
						}
					}
					if !dup {

						err := downcsv.Write(uprecord)
						downcsv.Flush()
						fmt.Println("Debug: [Write][err]", uprecord, err)

					}
					dup = false
				}
			}
		} else if upcategory == "notsubscribe" {
			switch upaction {
			case "replace":
				upcsv := csv.NewReader(&filedata)
				csvtitle, err := upcsv.Read()
				if err == io.EOF {
					fmt.Fprintln(w, "Debug: #Empty-File#")
					return
				}
				if err != nil {
					fmt.Fprintln(w, "Debug: Error in File!")
				}
				os.Remove("./CSV/notsubscribe.csv")
				downfile, err := os.OpenFile("./CSV/notsubscribe.csv", os.O_CREATE|os.O_WRONLY, 0644)

				if err != nil {
					fmt.Fprintln(w, err)
				}
				defer downfile.Close()

				downcsv := csv.NewWriter(downfile)
				downcsv.Write(csvtitle)
				downcsv.Flush()

				var dup bool
				for {
					uprecord, err := upcsv.Read()
					if err == io.EOF {
						fmt.Println("Debug: #File-END#")
						break
					}
					if err != nil {
						fmt.Println("Debug: #File-END#")
						return
					}
					for _, v := range globle.Singlelistemails {
						if uprecord[2] == v {
							dup = true
							break
						}
					}
					if !dup {

						err := downcsv.Write(uprecord)
						downcsv.Flush()
						fmt.Println("Debug: [Write][err]", uprecord, err)

					}
					dup = false
				}
			case "append":
				upcsv := csv.NewReader(&filedata)
				_, err := upcsv.Read()
				if err == io.EOF {
					fmt.Fprintln(w, "Debug: #Empty-File#")
					return
				}
				if err != nil {
					fmt.Fprintln(w, "Debug: Error in File!")
				}
				downfile, err := os.OpenFile("./CSV/notsubscribe.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

				if err != nil {
					fmt.Fprintln(w, err)
				}
				defer downfile.Close()

				downcsv := csv.NewWriter(downfile)

				var dup bool
				for {
					uprecord, err := upcsv.Read()
					if err == io.EOF {
						fmt.Println("Debug: #File-END#")
						break
					}
					if err != nil {
						fmt.Println("Debug: #File-END#")
						return
					}
					for _, v := range globle.Singlelistemails {
						if uprecord[2] == v {
							dup = true
							break
						}
					}
					if !dup {

						err := downcsv.Write(uprecord)
						downcsv.Flush()
						fmt.Println("Debug: [Write][err]", uprecord, err)

					}
					dup = false
				}
			}
		}

		/// show template html
		var tmpsinglelist []string

		asubscribe, err := os.Open("./CSV/subscribe.csv")
		if err != nil {
			fmt.Fprintln(w, err)
		}
		subreader := csv.NewReader(asubscribe)
		subrecords, _ := subreader.ReadAll()
		for o, u := range subrecords {
			tmpsinglelist = append(tmpsinglelist, u[2])
			globle.AllContact.Subzcount = o
		}
		globle.AllContact.Subscribe = subrecords
		///////////////////////////////////////////////
		aunsubscribe, err := os.Open("./CSV/unsubscribe.csv")
		if err != nil {
			fmt.Fprintln(w, err)
		}
		unsubreader := csv.NewReader(aunsubscribe)
		unsubrecords, _ := unsubreader.ReadAll()
		for q, u := range unsubrecords {
			tmpsinglelist = append(tmpsinglelist, u[2])
			globle.AllContact.Unsubzcount = q
		}
		globle.AllContact.Unsubscribe = unsubrecords

		///////////////////////////////////////////////
		anotsubscribe, err := os.Open("./CSV/notsubscribe.csv")
		if err != nil {
			fmt.Fprintln(w, err)
		}
		notsubreader := csv.NewReader(anotsubscribe)
		notsubrecords, _ := notsubreader.ReadAll()
		for p, u := range notsubrecords {
			tmpsinglelist = append(tmpsinglelist, u[2])
			globle.AllContact.Notsubzcount = p
		}
		globle.AllContact.Notsubscribe = notsubrecords

		// updating all single email list
		globle.Singlelistemails = tmpsinglelist

		tmpl, err := template.ParseFiles("./templates/CSV.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, globle.AllContact)
	}
}
