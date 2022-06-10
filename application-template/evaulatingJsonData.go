package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"text/template"
)

var templateContainerEmployee *template.Template

// Only the fields in Capital are the one that will be Unmarshaled

type DataRendering struct {
	EmployeeList []EmployeeStruct
	ShowBoldName bool
}

type EmployeeStruct struct {
	ID        string
	CreatedAt string
	Name      string
	Avatar    string
}

func (emp *EmployeeStruct) IsIdEven() bool {
	convertedInteger, _ := strconv.Atoi(emp.ID)
	isEven := convertedInteger%2 == 0
	return isEven
}

func main() {
	templateContainer = template.New("templateContainer")
	template.Must(templateContainer.ParseGlob("template/*.html"))

	http.HandleFunc("/", ShowEmployeeList)
	
	http.ListenAndServe(":8000", nil)
}

func ShowEmployeeList(writer http.ResponseWriter, request *http.Request) {
	var dataToRender DataRendering
	dataToRender.ShowBoldName = false
	dataToRender.EmployeeList = getEmployees()
	employeeTemplate := templateContainer.Lookup("showEmployeeList.html")
	employeeTemplate.Execute(writer, dataToRender)
}

func getEmployees() []EmployeeStruct {

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
