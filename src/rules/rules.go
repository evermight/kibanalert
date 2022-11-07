package rules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(apiUrl, apiKey string) Rules {
	req, err := http.NewRequest(http.MethodGet, apiUrl+"/api/alerting/rules/_find", nil)
	if err != nil {
		fmt.Println("Error", err)
		return Rules{}
	}
	req.Header.Set("Authorization", "ApiKey "+apiKey)
	//req.SetBasicAuth(username, password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error", err)
		return Rules{}
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error", err)
		return Rules{}
	}
	return Parse(respBody)
}
func Parse(content []uint8) Rules {

	//content, _ := ioutil.ReadFile("./rules/test.json")
	var dat Rules
	if err := json.Unmarshal(content, &dat); err != nil {
		return Rules{}
	}

	return dat
}

type ExecutionStatus struct {
	Status string `json:"status"`
}
type Rule struct {
	RuleId          string          `json:"id"`
	Name            string          `json:"name"`
	ExecutionStatus ExecutionStatus `json:"execution_status"`
}
type Rules struct {
	Rules []Rule `json:"data"`
}
