package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func gettingSbiStockData() {
	http.HandleFunc("/sbistockinfo", GetSBIStockValue)
	http.ListenAndServe(":4000", nil)
}

func GetSBIStockValue(writer http.ResponseWriter, request *http.Request) {

	var stockInfo SbiStockInfo
	response, _ := http.Get("https://priceapi.moneycontrol.com/pricefeed/bse/equitycash/SBI")
	data, _ := ioutil.ReadAll(response.Body)

	unMarshalSbiData(data, &stockInfo)

	jsonData, _ := json.Marshal(stockInfo)

	fmt.Println(data)
	fmt.Fprint(writer, string(jsonData))
}
