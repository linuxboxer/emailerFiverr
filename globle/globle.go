package globle

import (
	"text/template"
	"time"

	"github.com/LinuxSploit/emailsoft/engine"
)

var (
	Compagins        []engine.Compagin = []engine.Compagin{}
	Duration         time.Duration     = time.Duration(1) * time.Second
	TK               time.Ticker       = *time.NewTicker(Duration)
	Allfile          []string
	AllContact       Contacts
	Singlelistemails []string
	CSVemailcol      int
	Contactfuncmap   template.FuncMap = template.FuncMap{
		"SubCount":    AllContact.CountSub(),
		"UnsubCount":  AllContact.CountUnsub(),
		"NotsubCount": AllContact.CountNotsub(),
	}
)

type Contacts struct {
	Subscribe    [][]string
	Subzcount    int
	Unsubscribe  [][]string
	Unsubzcount  int
	Notsubscribe [][]string
	Notsubzcount int
}

func Counter(zzz [][]string) (i int) {
	for i = range zzz {
	}
	return i
}
func (z *Contacts) CountSub() (v int) {

	return Counter(AllContact.Subscribe)
}
func (z *Contacts) CountUnsub() (v int) {

	return Counter(AllContact.Unsubscribe)
}
func (z *Contacts) CountNotsub() (v int) {

	return Counter(AllContact.Notsubscribe)
}
