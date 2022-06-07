package main

import "fmt"

type EmployeeStructure struct {
	name        string
	age         int
	designation string
	bonus       int
}

func arrayOfStructures() {
	namesArray := [4]string{"A", "B", "C", "D"}
	ageArray := [4]int{10, 20, 30, 40}
	designationArray := [4]string{"X", "Y", "Z", "XY"}
	bonusArray := [4]int{2, 3, 4, 5}

	var empList []EmployeeStructure

	for i := 0; i < len(designationArray); i++ {
		empDetails := EmployeeStructure{namesArray[i], ageArray[i], designationArray[i], bonusArray[i]}
		empList = append(empList, empDetails)
	}

	fmt.Println(empList)
}
