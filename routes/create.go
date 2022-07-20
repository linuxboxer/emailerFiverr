package routes

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/LinuxSploit/emailsoft/engine"
	"github.com/LinuxSploit/emailsoft/globle"
	"github.com/vanng822/go-premailer/premailer"
)

var (
	RFC3339local string = "2006-01-02T15:04Z"
)

func Create(w http.ResponseWriter, r *http.Request) {
	funcmap := template.FuncMap{
		"SubCount":    globle.AllContact.CountSub(),
		"UnsubCount":  globle.AllContact.CountUnsub(),
		"NotsubCount": globle.AllContact.CountNotsub(),
	}
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("./templates/create.html")
		if err != nil {
			fmt.Fprintln(w, "Error loading template!")
		}
		tmpl.Execute(w, funcmap)
	case "POST":
		r.ParseForm()
		title := r.FormValue("title")
		subject := r.FormValue("subject")
		exportedtmpl := r.FormValue("exportedtmpl")
		inliner, err := premailer.NewPremailerFromString(exportedtmpl, premailer.NewOptions())
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		exportedtmpl, err = inliner.Transform()
		if err != nil {
			fmt.Fprintln(w, err)
		}
		//send list option
		tosub := r.FormValue("tosubscribe")
		tononsub := r.FormValue("tonotsubscribe")
		tounsub := r.FormValue("tounsubscribe")
		if tosub != "send" && tounsub != "send" && tononsub != "send" {
			fmt.Fprintln(w, "Please select contact from form!")
		}
		//
		// optional
		theattachment := r.FormValue("attachmentoye")

		//
		runtime := r.FormValue("runtime")

		loc, err := time.LoadLocation("America/New_York")
		if err != nil {
			fmt.Fprintln(w, "Time Location Error!")
			return
		}
		t1, err := time.ParseInLocation(RFC3339local, runtime+"Z", loc)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		var tmpsendto engine.SendTo
		if tosub == "send" {
			tmpsendto.Subscribe = true
		}
		if tounsub == "send" {
			tmpsendto.Unsubscribe = true
		}
		if tononsub == "send" {
			tmpsendto.Notsubscribe = true
		}

		if title != "" && subject != "" && runtime != "" && exportedtmpl != "" {
			// check duplicate time
			for _, v := range globle.Compagins {
				if v.RunTime.Format("2006-01-02T15:04Z") == t1.Format("2006-01-02T15:04Z") {
					t1 = t1.Add(time.Minute * 2)
				}
			}
			// reading email list and adding it struct

			To := []string{}

			//////////Adding Email contacts///////////////
			if tmpsendto.Subscribe {
				tmpfile, err := os.Open("./CSV/subscribe.csv")
				if err != nil {
					log.Fatal(err)
				}
				defer tmpfile.Close()
				tmpcsv := csv.NewReader(tmpfile)
				_, err = tmpcsv.Read()
				if err != nil {
					fmt.Println("Not records in subscribe csv")
				}

				// reading all remaining record without heading(first line of csv)
				allrecords, err := tmpcsv.ReadAll()
				if err == nil {
					for _, record := range allrecords {
						To = append(To, record[2])
					}
				}
			}
			if tmpsendto.Unsubscribe {
				tmpfile, err := os.Open("./CSV/unsubscribe.csv")
				if err != nil {
					log.Fatal(err)
				}
				defer tmpfile.Close()
				tmpcsv := csv.NewReader(tmpfile)
				_, err = tmpcsv.Read()
				if err != nil {
					fmt.Println("Not records in unsubscribe csv")
				}

				// reading all remaining record without heading(first line of csv)
				allrecords, err := tmpcsv.ReadAll()
				if err == nil {
					for _, record := range allrecords {
						To = append(To, record[2])
					}
				}
			}
			if tmpsendto.Notsubscribe {
				tmpfile, err := os.Open("./CSV/notsubscribe.csv")
				if err != nil {
					log.Fatal(err)
				}
				defer tmpfile.Close()
				tmpcsv := csv.NewReader(tmpfile)
				_, err = tmpcsv.Read()
				if err != nil {
					fmt.Println("Not records in not-subscribe csv")
				}

				// reading all remaining record without heading(first line of csv)
				allrecords, err := tmpcsv.ReadAll()
				if err == nil {
					for _, record := range allrecords {
						To = append(To, record[2])
					}
				}
			}

			// create new campagin
			tmp := engine.Compagin{
				UID: engine.TotalCompagins, Title: title, Themessage: engine.Message{Subject: subject, Body: exportedtmpl}, RunTime: t1, Pending: true, Attachment: theattachment, ToAddresses: tmpsendto, SentList: To,
			}
			globle.Compagins = append(globle.Compagins, tmp)
			engine.TotalCompagins++

			// debug line
			fmt.Println(globle.Compagins)

			http.Redirect(w, r, "/", http.StatusPermanentRedirect)

		} else {
			fmt.Fprintln(w, "missing any non-optional form value!")
		}
	}
}
