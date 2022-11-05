package notify

import (
	"kibanalert/alerts"
	"os"
	"strings"
)

type sendfunc func(alerts.Source)

func Notify(source alerts.Source) {

	notifyMethods := strings.Split(os.Getenv("NOTIFY_METHODS"), ",")

	adapters := map[string]sendfunc{
		"sendgrid": SendGrid,
		"smtp":     SMTP,
	}

	for _, meth := range notifyMethods {
		adapters[strings.TrimSpace(meth)](source)
	}
}
