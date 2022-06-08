package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Only the fields in Capital are the one that will be Unmarshaled

type EmployeeStruct struct {
	ID        string
	CreatedAt string
	Name      string
	Avatar    string
}

func main() {
	var employeeNameArray []string

	readFile("employeeList.json")
	employeeDataByte := readFile("employeeList.json")

	if len(employeeDataByte) > 0 {
		var employeeList []EmployeeStruct
		unmarshalError := unMarshalData(employeeDataByte, &employeeList)

		if unmarshalError != nil {
			fmt.Println("Error Parsing Data")
		} else {

			// Try Writing Function for Iteration as well

			for _, emp := range employeeList {
				employeeNameArray = append(employeeNameArray, emp.Name)
			}

			fmt.Println(employeeNameArray)
		}
	}
}

func readFile(fileName string) []byte {
	employeeDataByte, err := ioutil.ReadFile("employeeList.json")
	if err != nil {
		return []byte("")
	} else {
		return employeeDataByte
	}
}

func unMarshalData(byteData []byte, dataPointer *[]EmployeeStruct) error {
	return json.Unmarshal(byteData, &dataPointer)
}
