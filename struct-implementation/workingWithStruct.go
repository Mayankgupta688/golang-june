package main

import "fmt"

type Employee struct {
	empName        string
	empAge         byte
	empDesignation string
}

var employeeOne *Employee = new(Employee)

func workingWithStruct() {

	employeeOne.empAge = 10
	employeeOne.empDesignation = "hjdsfadha"
	employeeOne.empName = "sdgasukg"

	fmt.Println(*(&employeeOne))
	fmt.Println(employeeOne)

	//fmt.Println(employeeOne) // Give Memory Location containing Value

	(*employeeOne).empName = "Abc"
	employeeOne.empAge = 10

	fmt.Println(employeeOne.empName)

}


