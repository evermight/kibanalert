package notify

import (
	"kibanalert/alerts"
	"os"
	"strings"
)

type sendfunc func(alerts.Source) error

func Notify(source alerts.Source) []error {
	var errorList []error
	notifyMethods := strings.Split(os.Getenv("NOTIFY_METHODS"), ",")

	adapters := map[string]sendfunc{
		"sendgrid": SendGrid,
		"smtp":     SMTP,
	}

	for _, meth := range notifyMethods {
		err := adapters[strings.TrimSpace(meth)](source)
		if err != nil {
			if errorList == nil {
				errorList = []error{}
			}
			errorList = append(errorList, err)
		}
	}
	return errorList
}
