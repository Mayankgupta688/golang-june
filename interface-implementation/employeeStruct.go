package main

import "fmt"

type Employee struct {
	empName string
	empAge  int
}

func (emp Employee) getDetails() {
	fmt.Println(emp.empName)
	fmt.Println(emp.empAge)
}

func (emp Employee) getName() {
	fmt.Println(emp.empName)
}

func (emp Employee) getAge() {
	fmt.Println(emp.empAge)
}
