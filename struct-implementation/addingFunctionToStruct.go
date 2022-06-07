package main

import "fmt"

type EmpStructure struct {
	empName   string
	empSalary int16
	empBonus  int16
}

func (emp *EmpStructure) updateName() {
	(*emp).empName = "Random Name"
	emp.empName = "Other Random"
}

func (emp EmpStructure) getDetails() {
	fmt.Printf("\n\nEmployee %v has Effective salary of %v Rs", emp.empName, emp.empBonus+emp.empSalary)
}

func (emp EmpStructure) calculateTotalSalary() int16 {
	return emp.empSalary + emp.empBonus
}

func (emp EmpStructure) getName() {
	fmt.Printf("\n\nEmployee %v", emp.empName)
}

func (emp EmpStructure) getSalary() {
	fmt.Printf("\n\nEmployee has Effective salary of %v Rs", emp.empBonus+emp.empSalary)
}

func addingFunctionToStruct() {
	var empOne EmpStructure
	empOne.empName = "Mayank"
	empOne.empBonus = 10
	empOne.empSalary = 100
	empOne.updateName()
	fmt.Println(empOne.empName)
	fmt.Println(empOne.calculateTotalSalary())
}

func returnData(emp EmpStructure) EmpStructure {
	fmt.Printf("\n\nEmployee %v has Effective salary of %v Rs", emp.empName, emp.empBonus+emp.empSalary)
	return emp
}
