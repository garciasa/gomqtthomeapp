package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

/**
	Modified from https://gist.github.com/jpillora/cb46d183eca0710d909a
	Thank you very much.
**/

// SMTP_SERVER
const (
	SMTPServer = "smtp.gmail.com"
)

// Sender - Structure for Sender
type Sender struct {
	User     string
	Password string
}

// NewSender - function for creating new sender
func NewSender(Username, Password string) Sender {

	return Sender{Username, Password}
}

// SendMail - Sending Email
func (sender Sender) SendMail(Dest []string, Subject, bodyMessage string) {

	msg := "From: " + sender.User + "\n" +
		"To: " + strings.Join(Dest, ",") + "\n" +
		"Subject: " + Subject + "\n" + bodyMessage

	err := smtp.SendMail(SMTPServer+":587",
		smtp.PlainAuth("", sender.User, sender.Password, SMTPServer),
		sender.User, Dest, []byte(msg))

	if err != nil {

		fmt.Printf("smtp error: %s", err)
		return
	}

	fmt.Println("Mail sent successfully!")
}
