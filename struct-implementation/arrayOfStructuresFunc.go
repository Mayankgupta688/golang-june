package main

import "fmt"

func arrayOfStructuresFunc() {
	namesArray := [4]string{"A", "B", "C", "D"}
	ageArray := [4]int{10, 20, 30, 40}
	designationArray := [4]string{"X", "Y", "Z", "XY"}
	bonusArray := [4]int{2, 3, 4, 5}

	var empList []EmployeeStructure

	for i := 0; i < len(designationArray); i++ {
		returnedEmployee := createEmployee(ageArray[i], bonusArray[i], namesArray[i], designationArray[i])
		empList = append(empList, returnedEmployee)
	}

	fmt.Println(empList)
}

func createEmployee(age, bonus int, name, designation string) EmployeeStructure {
	return EmployeeStructure{name, age, designation, bonus}
}
