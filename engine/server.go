package engine

import (
	"encoding/csv"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

var (
	AppSetting Setting
)

type Setting struct {
	Host      string
	port      string
	YourEmail string
	passcode  string
}

func (a *Setting) Read() error {
	// reading settings from config
	afile, err := os.Open("./config")
	if err != nil {
		log.Fatal(err)
		return errors.New("cannot open config file!")
	}
	defer afile.Close()
	setting := csv.NewReader(afile)
	allsetting, err := setting.Read()
	if err != nil {
		log.Fatal(err)
		return errors.New("cannot read opened config file!")
	}
	a.Host = allsetting[0]
	a.port = allsetting[1]
	a.YourEmail = allsetting[2]
	a.passcode = allsetting[3]

	return nil
}

func (a *Setting) Write(ahost string, aport string, ayouremail string, apasscode string) error {
	err := ioutil.WriteFile("./config", []byte(ahost+","+aport+","+ayouremail+","+apasscode), 0644)
	if err != nil {
		return err
	}
	err = a.Read()
	if err != nil {
		return err

	}
	return nil
}
