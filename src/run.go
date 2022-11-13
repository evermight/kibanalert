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
	debug := os.Getenv("DEBUG") == "1"

	for true {
		currentRules := rules.Get(
			os.Getenv("KIBANA_URL"),
			os.Getenv("ELASTIC_API_KEY"),
		)
		for _, rule := range currentRules.Rules {
			if rule.ExecutionStatus.Status == "active" {
				ruleId := rule.RuleId
				if _, ok := previousHitId[ruleId]; !ok {
					previousHitId[ruleId] = ""
				}
				currentAlert := alerts.Get(
					ruleId,
					os.Getenv("CONNECTOR_INDEX_NAME"),
					os.Getenv("ELASTIC_URL"),
					os.Getenv("ELASTIC_API_KEY"),
				)
				if len(currentAlert.Hits.Hits) > 0 {
					hit := currentAlert.Hits.Hits[0]
					if previousHitId[ruleId] != hit.HitId {
						if debug {
							fmt.Println(fmt.Sprintf("Notifying %v hit.HitId: %v", rule.Name, hit.HitId))
						}
						if errs := notify.Notify(hit.Source); errs != nil {
							fmt.Println("Notification Failures: ", errs)
						}
						previousHitId[ruleId] = hit.HitId
					} else {
						if debug {
							fmt.Println(fmt.Sprintf("Skipping %v hit.HitId: %v", rule.Name, hit.HitId))
						}
					}
				}
			}
		}
		time.Sleep(time.Duration(scanInterval) * time.Second)
	}
}
