package main

import (
	"html/template"
	"net/http"
	"path"
)

func renderWithData() {
	http.HandleFunc("/", ShowMainPage)
	http.ListenAndServe(":8000", nil)
}

func ShowMainPage(writer http.ResponseWriter, request *http.Request) {
	employeeOne := Employee{"Mayank", "Developer"}
	fp := path.Join("template", "employeeDetails.html")
	employeeTemplate, _ := template.ParseFiles(fp)
	employeeTemplate.Execute(writer, employeeOne)
}
