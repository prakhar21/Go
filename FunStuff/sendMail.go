package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"os"
)

func check(e error) {

	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	arguments := os.Args[1:]
	email := arguments[0]
	passwd := arguments[1]
	to := arguments[2]

	mail := gomail.NewMessage()
	mail.SetHeader("From", email)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", "Test GO Mail")
	mail.SetBody("text/plain", "Hello, How are you!")
	//mail.Attach("")

	dialer := gomail.NewDialer("smtp.gmail.com", 587, email, passwd)

	check(dialer.DialAndSend(mail))
	fmt.Println("Mail Send")
}
