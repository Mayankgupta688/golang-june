package main

import "fmt"

type EmployeeAddress struct {
	street, nearby, city, state, country string
	pinCode                              int32
	isPermanent                          bool
}

type EmployeeDetailedInfo struct {
	empName    string
	empAge     int
	empSalary  int16
	isMarried  bool
	empAddress EmployeeAddress
}

func main() {
	employeeOne := new(EmployeeDetailedInfo)
	employeeOne.empName = "Mayank"
	employeeOne.empAge = 20
	employeeOne.empSalary = 100
	employeeOne.empAddress = EmployeeAddress{street: "A-8 Milansar", city: "Delhi", pinCode: 110085}

	fmt.Println(employeeOne)
}
