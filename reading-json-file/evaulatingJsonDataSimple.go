package main

import (
	"encoding/json"
	"fmt"
)

// Only the fields in Capital are the one that will be Unmarshaled

type StockData struct {
	Code               string `json:"data_id"`
	Message            string `json:"user_message"`
	Company            string `json:"company_name"`
	Pricechange        string
	Pricepercentchange string
}

// JSON Marshaling / UnMarshling

var stockDataObjectArray []StockData

func evaulatingJsonDataSimple() {

	stockDataArray := `[{
		"user_message": "Success",
		"company_name": "SBI",
		"data_id": "200",
		"pricechange": "-1.0500",
		"message": "Failure",
		"pricepercentchange": "-0.2265"
	}, {
		"user_message": "FAILURE",
		"company_name": "NALCO",
		"data_id": "200",
		"pricechange": "-1.0500",
		"message": "Failure",
		"pricepercentchange": "-0.2265"
	}]`

	err := json.Unmarshal([]byte(stockDataArray), &stockDataObjectArray)
	if len(stockDataObjectArray) > 0 {
		for index, dataValue := range stockDataObjectArray {
			fmt.Println(index)
			fmt.Println(dataValue)
		}
	}

	if err != nil {
		fmt.Println("Cannot Convert JSON to Struct")
	}
}
