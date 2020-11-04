package goldapi

import (
	"encoding/json"
	"fmt"
	"strings"
)

type MetalReport struct {
	Timestamp      int     `json:"timestamp"`
	Metal          string  `json:"metal"`
	Currency       string  `json:"currency"`
	Exchange       string  `json:"exchange"`
	Symbol         string  `json:"symbol"`
	PrevClosePrice float64 `json:"prev_close_price"`
	OpenPrice      float64 `json:"open_price"`
	LowPrice       float64 `json:"low_price"`
	HighPrice      float64 `json:"high_price"`
	OpenTime       int     `json:"open_time"`
	Price          float64 `json:"price"`
	Ch             float64 `json:"ch"`
	Chp            float64 `json:"chp"`
	Ask            float64 `json:"ask"`
	Bid            float64 `json:"bid"`
}

type Metal string

const (
	Gold Metal = "XAU"
	Silver Metal = "XAG"
	Platinum Metal= "XPT"
	Palladium  Metal = "XPD"
)

func (a *API) GetPrice(metal Metal, currency string, date string) (*MetalReport, error){
	if len(currency) != 3 {
		return nil, fmt.Errorf("Currency must be in ISO 4217 format. Got %s", currency)
	}
	if date != "" && len(date) != 8 {
		return nil, fmt.Errorf("Date must be either blank or in YYYYMMDD format. Got %s", date)
	}

	if date != "" {
		date = "/" + date
	}

	path := fmt.Sprintf("%s/%s%s", metal, currency, date)

	status, byt, err := a.get(path)
	if err != nil {
		return nil, err
	}
	if status != 200 || strings.Contains(string(byt), "error") {
		return nil, fmt.Errorf("API error: %s", string(byt))
	}

	var m MetalReport
	err = json.Unmarshal(byt, &m)

	return &m, err
}
