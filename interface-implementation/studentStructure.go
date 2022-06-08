package main

import "fmt"

type Student struct {
	studName string
	studAge  int
}

func (emp Student) getDetails() {
	fmt.Println(emp.studName)
	fmt.Println(emp.studAge)
}

func (emp Student) getName() {
	fmt.Println(emp.studName)
}

func (emp Student) getAge() {
	fmt.Println(emp.studAge)
}
