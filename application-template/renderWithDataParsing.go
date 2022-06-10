package main

import (
	"net/http"
	"text/template"
)

var templateContainer *template.Template

func renderWithDataParsing() {

	templateContainer = template.New("templateContainer")
	template.Must(templateContainer.ParseGlob("template/*.html"))

	http.HandleFunc("/", ShowMainPageData)
	http.HandleFunc("/help", ShowHelpPageData)
	http.ListenAndServe(":8000", nil)
}

func ShowMainPageData(writer http.ResponseWriter, request *http.Request) {
	employeeOne := Employee{"Anshul", "Developer"}
	employeeTemplate := templateContainer.Lookup("employeeDetails.html")
	employeeTemplate.Execute(writer, employeeOne)
}

func ShowHelpPageData(writer http.ResponseWriter, request *http.Request) {
	employeeTemplate := templateContainer.Lookup("about.html")
	employeeTemplate.Execute(writer, nil)
}
