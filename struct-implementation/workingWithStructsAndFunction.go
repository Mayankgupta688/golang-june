package main

import "fmt"

type Student struct {
	studName        string
	studAge         byte
	studDesignation string
}

func workingWithStructsAndFunction() {
	studentOne := new(Student)
	studentOne.studName = "Mayank"
	studentOne.studAge = 10
	studentOne.studDesignation = "Developer"

	studTwo := Student{"Mayank", 20, "Director"}
	fmt.Println(studTwo)

	studThree := Student{studDesignation: "VP", studAge: 100}
	fmt.Println(studThree)

	modifyStudent(studentOne)
	fmt.Println(studentOne.studName)
}

func modifyStudent(stud *Student) {
	stud.studName = "Abc"
	(*stud).studName = "Xyz"
	fmt.Println(stud.studName)
}
