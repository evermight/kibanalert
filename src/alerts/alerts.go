package alerts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(ruleId, indexName, apiUrl, apiKey string) Alerts {
	var jsonData = []byte(`{
  "query": {
   "term": {
     "rule_id": "` + ruleId + `"
   }
  },
  "size": 1,
  "sort": [{"date":{"order":"desc"}}]}`)
	req, err := http.NewRequest(http.MethodGet, apiUrl+"/"+indexName+"/_search", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error", err)
		return Alerts{}
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "ApiKey "+apiKey)
	// req.SetBasicAuth(username, password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error", err)
		return Alerts{}
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error", err)
		return Alerts{}
	}
	return Parse(respBody)
}
func Parse(content []uint8) Alerts {
	var dat Alerts
	if err := json.Unmarshal(content, &dat); err != nil {
		return Alerts{}
	}
	return dat
}

type Source struct {
	AlertId     string `json:"alert_id"`
	RuleId      string `json:"rule_id"`
	Reason      string `json:"reason"`
	ServiceName string `json:"service_name"`
	Date        string `json:"date"`
}
type Hit struct {
	HitId  string `json:"_id"`
	Source Source `json:"_source"`
}
type Hits struct {
	Hits []Hit `json:"hits"`
}
type Alerts struct {
	Hits Hits `json:"hits"`
}
