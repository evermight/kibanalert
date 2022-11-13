package notify

import (
	"errors"
	"fmt"
	"kibanalert/alerts"
	"net/smtp"
	"os"
	"strings"
)

func SMTP(source alerts.Source) error {

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
		msg := []byte(fmt.Sprintf("From: %v <%v>\r\nTo: %v\r\nSubject: %v\r\n\r\n%v\r\n",
			os.Getenv("SMTP_FROM_NAME"),
			os.Getenv("SMTP_FROM_EMAIL"),
			to,
			subject,
			plainTextContent,
		))

		auth := smtp.PlainAuth("", user, password, host)

		err := smtp.SendMail(addr, auth, from, []string{to}, msg)
		if err != nil {
			return errors.New(fmt.Sprintf("SMTP: %v %v", host, err))
		}
	}
	return nil
}
