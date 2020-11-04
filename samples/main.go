package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joel-g/go-goldapi/goldapi"
)

func main() {
	// Get your API key from https://www.goldapi.io
	api := goldapi.NewAPIClient(os.Getenv("gold_api_key"))

	// Use one of the 4 const Metals from goldapi package:
	// Must include currency in ISO 4217.
	// Date is optional but must be in YYYYMMDD format.
	// If date is left blank today's date will be used.
	silverReport, err := api.GetPrice(goldapi.Silver, "USD", "")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The price of silver is %f\n", silverReport.Price)
	fmt.Printf("%+v", silverReport)

	// A report of your API usage
	stats, err := api.GetStats()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nYou have used %d requests this month", stats.RequestsMonth)
}
