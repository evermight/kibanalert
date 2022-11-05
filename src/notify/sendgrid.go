package notify

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"kibanalert/alerts"
	"os"
	"strings"
)

func SendGrid(source alerts.Source) {
	apiKey := os.Getenv("SENDGRID_KEY")

	from := mail.NewEmail(os.Getenv("SENDGRID_FROM_NAME"), os.Getenv("SENDGRID_FROM_EMAIL"))

	subject := fmt.Sprintf("Alert: %v %v", source.ServiceName, source.Reason)
	plainTextContent := subject
	htmlContent := subject

	recipients := strings.Split(os.Getenv("SENDGRID_TO_EMAIL"), ",")
	for i, recipient := range recipients {
		recipients[i] = strings.TrimSpace(recipient)
	}

	for _, recipient := range recipients {
		to := mail.NewEmail(recipient, recipient)
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		client := sendgrid.NewSendClient(apiKey)
		//response, err := client.Send(message)
		_, err := client.Send(message)
		if err != nil {
			fmt.Println(err)
		} else {
			/*
				fmt.Println(response.StatusCode)
				fmt.Println(response.Body)
				fmt.Println(response.Headers)
			*/
		}
	}
}
