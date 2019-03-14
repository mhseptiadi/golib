package sendgmail

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

func Send(recipient string, subject string, body string, sender string, password string) {

	auth := smtp.PlainAuth("", sender, password, "smtp.gmail.com")

	to := strings.Split(recipient, ",")

	msg := []byte("To: " + recipient + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail("smtp.gmail.com:587", auth, sender, to, msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("email sent")
}
