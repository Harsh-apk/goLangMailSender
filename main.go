package main

import (
	"encoding/json"
	"net/smtp"
	"os"
)

type data struct {
	User        string
	AppPassword string
}

func sendSimpleMail() {
	var data data
	file, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		panic(err)
	}
	msg := "Subject: Verification code\nYour code is 1232123.\n\tCheck out github.com/Harsh-apk/goLangMailSender"
	auth := smtp.PlainAuth("", data.User, data.AppPassword, "smtp.gmail.com")
	err = smtp.SendMail("smtp.gmail.com:587", auth, data.User, []string{"ha470474@gmail.com"}, []byte(msg))
	if err != nil {
		panic(err)
	}

}

func main() {
	sendSimpleMail()
}
