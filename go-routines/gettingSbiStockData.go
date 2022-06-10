package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type StockData struct {
	Pricechange        string
	Pricepercentchange string
	Pricecurrent       string
}

type SbiStockInfo struct {
	Data StockData
}

func unMarshalSbiData(byteData []byte, dataPointer *SbiStockInfo) error {
	return json.Unmarshal(byteData, &dataPointer)
}

func GetSBIStockValue() {

	var stockInfo SbiStockInfo
	response, _ := http.Get("https://priceapi.moneycontrol.com/pricefeed/bse/equitycash/SBI")
	data, _ := ioutil.ReadAll(response.Body)

	unMarshalSbiData(data, &stockInfo)

	jsonData, _ := json.Marshal(stockInfo)

	fmt.Println(jsonData)
	time.Sleep(5 * time.Second)
}
