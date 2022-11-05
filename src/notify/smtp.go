package notify

import (
	"fmt"
	"kibanalert/alerts"
	"net/smtp"
	"os"
	"strings"
)

func SMTP(source alerts.Source) {

	subject := fmt.Sprintf("Alert: %v %v", source.ServiceName, source.Reason)
	plainTextContent := subject

	from := os.Getenv("SMTP_FROM_EMAIL")

	user := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASS")
	addr := os.Getenv("SMTP_ADDR")
	host := os.Getenv("SMTP_HOST")

	recipients := strings.Split(os.Getenv("SMTP_TO_EMAIL"), ",")
	for _, recipient := range recipients {
		to := strings.TrimSpace(recipient)
		msg := []byte("From: " + os.Getenv("SMTP_FROM_NAME") + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n\r\n" +
			plainTextContent + "\r\n")

		auth := smtp.PlainAuth("", user, password, host)

		err := smtp.SendMail(addr, auth, from, []string{to}, msg)

		if err != nil {
			fmt.Println(err)
		}

	}
	// fmt.Println("Email sent successfully")

}
