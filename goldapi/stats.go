package goldapi

import (
	"encoding/json"
	"fmt"
	"strings"
)

type StatsReport struct {
	RequestsToday     int `json:"requests_today"`
	RequestsYesterday int `json:"requests_yesterday"`
	RequestsMonth     int `json:"requests_month"`
	RequestsLastMonth int `json:"requests_last_month"`
}

func (a *API) GetStats() (*StatsReport, error){

	status, byt, err := a.get("stat")
	if err != nil {
		return nil, err
	}
	if status != 200 || strings.Contains(string(byt), "error") {
		return nil, fmt.Errorf("API error: %s", string(byt))
	}

	var s StatsReport
	err = json.Unmarshal(byt, &s)

	return &s, err
}
