package engine

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"time"
)

var (
	TotalCompagins int
	DomainName     string = "http://localhost:4556"
)

type Message struct {
	Subject string
	Body    string
}
type SendTo struct {
	Subscribe    bool
	Unsubscribe  bool
	Notsubscribe bool
}
type Trackmail struct {
	Compagin string
	Email    []string
}

type Compagin struct {
	UID         int
	Title       string
	ToAddresses SendTo
	Themessage  Message
	RunTime     time.Time
	Pending     bool
	Running     bool
	Done        bool
	Attachment  string

	// tracking email member
	ViewdEmail []string
	SentList   []string
	BounceList []string
}

func (a *Compagin) ViewCount() int {
	return len(a.ViewdEmail)
}
func (a *Compagin) TotalEmailSent() int {
	fmt.Println("total-email-sent", a.SentList)
	return len(a.SentList)
}
func (a *Compagin) BounceCount() int {
	return len(a.BounceList)
}
func (a *Compagin) AddViewedEmail(email string) {
	a.ViewdEmail = append(a.ViewdEmail, email)
}

func (a *Compagin) Run() Compagin {
	var tmpcomp Compagin = *a
	// reading settings from config needed to email Auth
	err := AppSetting.Read()
	if err != nil {
		log.Fatal("cannot Read Config File, Check Permissions")
	}
	//

	fmt.Println("running")

	/////////detete empty records\\\\\\\\\\
	To := delete_empty(a.SentList)
	////////////////////////
	auth := smtp.PlainAuth("", AppSetting.YourEmail, AppSetting.passcode, AppSetting.Host)

	///////////////////////////////////////////////////////////////////

	for i, q := range To {
		// varifying email \\
		// ... varify syntax of email
		// ... varify valid domain of email address
		//                  \\
		thebody := BuildMail(a.Title, AppSetting.YourEmail, q, a.Themessage.Subject, a.Themessage.Body, a.Attachment)

		err = smtp.SendMail(AppSetting.Host+":"+AppSetting.port, auth, AppSetting.YourEmail, []string{q}, thebody)

		if err != nil {
			tmpcomp.BounceList = append(a.BounceList, q)
			tmpcomp.SentList = RemoveIndex(tmpcomp.SentList, i)
		} else {
			fmt.Println("Sent to :", q)
		}
	}
	return tmpcomp
}

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

/////////

func BuildMail(comp string, from string, to string, subject string, body string, attachment string) []byte {
	fmt.Println("building mail:", to)
	// generating view tracker
	viewTrackerID := base64.StdEncoding.EncodeToString([]byte(comp + ":" + to))
	viewTrackertag := "<img src=\"" + DomainName + "/pixel?id=" + viewTrackerID + "\" width=\"1px\" height=\"1px\">"

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("From: %s\r\n", from))
	buf.WriteString(fmt.Sprintf("To: %s\r\n", to))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))

	boundary := "my-boundary-779"
	buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))

	buf.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	buf.WriteString("Content-Type: text/html; charset=\"utf-8\"\r\n\r\n")
	buf.WriteString(body + viewTrackertag)

	buf.WriteString(fmt.Sprintf("\r\n--%s", boundary))

	if attachment != "" {
		buf.WriteString("\r\nContent-Type: text/plain; charset=\"utf-8\"\r\n")
		buf.WriteString("Content-Transfer-Encoding: base64\r\n")
		buf.WriteString("Content-Disposition: attachment; filename=" + attachment + "\r\n")
		buf.WriteString("Content-ID: <" + attachment + ">\r\n\r\n")

		data := readFile("./uploads/" + attachment)

		b := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
		base64.StdEncoding.Encode(b, data)
		buf.Write(b)
		buf.WriteString(fmt.Sprintf("\r\n--%s", boundary))
	}
	buf.WriteString("--")

	return buf.Bytes()
}

func readFile(fileName string) []byte {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	return data
}
