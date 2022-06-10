package main

import (
	"html/template"
	"os"
)

type Employee struct {
	Name        string
	Designation string
}

func (emp Employee) ShowMessage() string {
	return "Employee Name is " + emp.Name
}

func basicTemplate() {
	newEmployee := Employee{"Mayank", "Developer"}
	templateString := `
		Hello {{ .Name }} 
		Designated as {{ .Designation }}, 
		Welcome to Golang Templating Session 
		{{ .ShowMessage }}
	`

	newTemplate, _ := template.New("customTemplates").Parse(templateString)
	newTemplate.Execute(os.Stdout, newEmployee)
}
