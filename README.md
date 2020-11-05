# go-goldapi
### Go package for consuming the Goldapi.io API

Get your API key from https://www.goldapi.io/ and see full API documentation.

This API and package can be used to query current and historical prices of gold, silver, platinum and palladium in the following currencies:

* USD - United States dollar
* AUD - Australian dollar
* GBP - British pound
* EUR - European Euro
* CHF - Swiss franc
* CAD - Canadian dollar
* JPY - Japanese yen
* INR - Indian rupee
* SGD - Singapore Dollar
* BTC - Bitcoin
* CZK - Czech Krona
* RUB - Russian Ruble
* PLN - Polish Złoty
* MYR - Malaysian Ringgit
* XAG - Gold/Silver Ratio

To use download the package with `go get github.com/joel-g/go-goldapi`

Import the package into your project.

Create an API client with `api := goldapi.NewAPIClient("your-API-key")`

Then call the GetPrice method `api.GetPrice(goldapi.Metal, string, string)`

`GetPrice()` takes 3 arguments:
* metal type (use one of the 4 constants from the package `goldapi.Gold`, `goldapi.Silver`, `goldapi.Platinum` or `goldapi.Palladium`)
* currency code as a string (use one of the 3 digits codes from the list above)
* date (optional, pass empty string for today's prices or pass YYYYMMDD formatted string)


### Example:

```go
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

```


### Sample `MetalReport`:

```go
{Timestamp:1604535350 Metal:XAG Currency:USD Exchange:FOREXCOM Symbol:FOREXCOM:XAGUSD PrevClosePrice:23.895 OpenPrice:23.895 LowPrice:23.895 HighPrice:24.042 OpenTime:1604527200 Price:24.04 Ch:0.145 Chp:0.61 Ask:24.059 Bid:24.011}
```
