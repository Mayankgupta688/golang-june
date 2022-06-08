package main

import "fmt"

func usingInterfaces() {
	var userType string
	fmt.Println("Enter 1 to create Employee and 2 to create Student")
	fmt.Scanln(&userType)

	if userType == "1" {
		// Create Employee

		employeeDetails := new(Employee)

		fmt.Println("Enter Employee Name")
		fmt.Scanln(&employeeDetails.empName)

		fmt.Println("Enter Age")
		fmt.Scanln(&employeeDetails.empAge)

		getData(employeeDetails)
	} else {
		// Create Student

		studentDetails := new(Student)

		fmt.Println("Enter Student Name")
		fmt.Scanln(&studentDetails.studName)

		fmt.Println("Enter Age")
		fmt.Scanln(&studentDetails.studAge)

		getData(studentDetails)
	}
}

type collegeEntry interface {
	getDetails()
	getName()
	getAge()
}

func getData(entry collegeEntry) {
	entry.getDetails()
}

func getDataForStudent(entry Student) {
	entry.getDetails()
}

func getDataForEmployee(entry Employee) {
	entry.getDetails()
}
