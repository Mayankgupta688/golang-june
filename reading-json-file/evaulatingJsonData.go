package main

import "fmt"

type StockData struct {
	code               string
	message            string
	company            string
	pricechange        float64
	pricepercentchange float64
}

// JSON Marshaling / UnMarshling

func main() {
	stockDetailsSBI := `{
		"code": "200",
		"message": "Success",
		"company": "SBI",
		"pricechange": "-1.0500",
		"pricepercentchange": "-0.2265"
	}`

	stockDetailsNalco := `{
		"code": "200",
		"message": "Success",
		"company": "NALCO",
		"pricechange": "-1.0500",
		"pricepercentchange": "-0.2265"
	}`

	// Display Change value from the above Data

	fmt.Println(stockDetailsSBI)
	fmt.Println(stockDetailsNalco)
}
