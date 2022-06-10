package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
	"time"
)

var templateContainerEmployee *template.Template

// Only the fields in Capital are the one that will be Unmarshaled

type DataRendering struct {
	EmployeeList []EmployeeStruct
	ShowBoldName bool
}

type EmployeeStruct struct {
	Name string
}

func main1() {
	http.HandleFunc("/", ShowEmployeeList)
	http.ListenAndServe(":8000", nil)
}

func ShowEmployeeList(writer http.ResponseWriter, request *http.Request) {
	var dataToRender DataRendering
	dataToRender.ShowBoldName = false
	dataToRender.EmployeeList = getEmployees()

	jsonData, _ := json.Marshal(dataToRender.EmployeeList)
	fmt.Fprint(writer, string(jsonData))
}

func getEmployees() []EmployeeStruct {

	time.Sleep(10 * time.Second)

	// Reading Data from the File System
	employeeDataByte := readFile("data/employeeList.json")

	if len(employeeDataByte) > 0 {

		// Conterting the JSON text into EmployeeStruct Type
		var employeeList []EmployeeStruct
		unmarshalError := unMarshalData(employeeDataByte, &employeeList)

		if unmarshalError != nil {
			return []EmployeeStruct{}
		} else {
			return employeeList
		}
	} else {
		return []EmployeeStruct{}
	}
}

func readFile(fileName string) []byte {
	employeeDataByte, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []byte("")
	} else {
		return employeeDataByte
	}
}

func unMarshalData(byteData []byte, dataPointer *[]EmployeeStruct) error {
	return json.Unmarshal(byteData, &dataPointer)
}
