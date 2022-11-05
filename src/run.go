package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"kibanalert/alerts"
	"kibanalert/notify"
	"kibanalert/rules"
	"os"
	"strconv"
	"time"
)

func main() {
	godotenv.Load()
	scanInterval, err := strconv.Atoi(os.Getenv("SCAN_INTERVAL"))
	if err != nil {
		scanInterval = 60
	}

	previousHitId := map[string]string{}

	for true {
		currentRules := rules.Get(
			os.Getenv("KIBANA_URL"),
			os.Getenv("KIBANA_USER"),
			os.Getenv("KIBANA_PASS"),
		)
		for _, rule := range currentRules.Rules {
			ruleId := rule.RuleId
			if _, ok := previousHitId[ruleId]; !ok {
				previousHitId[ruleId] = ""
			}
			currentAlert := alerts.Get(
				ruleId,
				os.Getenv("CONNECTOR_INDEX_NAME"),
				os.Getenv("ELASTIC_URL"),
				os.Getenv("ELASTIC_USER"),
				os.Getenv("ELASTIC_PASS"),
			)
			hit := currentAlert.Hits.Hits[0]
			if previousHitId[ruleId] != hit.HitId {
				notify.Notify(hit.Source)
				previousHitId[ruleId] = hit.HitId
				fmt.Println("Emailed " + hit.HitId)
			} else {
				fmt.Println("Skipped " + hit.HitId)
			}
		}
		time.Sleep(time.Duration(scanInterval) * time.Second)
	}
}
