package main

import (
	"net/smtp"
	"os"
)

func sendSimpleMail() {
	msg := "Subject: TimeTable verification code\nYour Code is 1234. Kindly share this with everyone, it's just a test, doesn't matter :)"
	user := os.Getenv("user")
	appPass := os.Getenv("appPassword")
	auth := smtp.PlainAuth("", user, appPass, "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:587", auth, user, []string{"ha470474@gmail.com"}, []byte(msg))
	if err != nil {
		panic(err)
	}

}

func main() {
	sendSimpleMail()
}
