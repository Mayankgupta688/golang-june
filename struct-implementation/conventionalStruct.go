package main

import "fmt"

type EmployeeStruct struct {
	empName        string
	empAge         byte
	empDesignation string
}

type EmployeeData struct {
	empName        string
	empAge         byte
	empDesignation string
}

func conventionalStruct() {
	employeeTwo := EmployeeStruct{
		empName:        "Mayank",
		empAge:         30,
		empDesignation: "VP",
	}

	employeeOne := new(EmployeeData)
	(*employeeOne).empAge = 10
	employeeOne.empDesignation = "hjdsfadha"
	employeeOne.empName = "sdgasukg"

	fmt.Println(&employeeTwo)
	fmt.Println(employeeTwo.empName)

	fmt.Println(employeeOne)
	fmt.Println(employeeOne.empName)
}
