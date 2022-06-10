package main

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
)

func renderSimpleHtmlPage() {
	http.HandleFunc("/", ShowHtmlTemplate)
	http.HandleFunc("/about", ShowAboutTemplate)
	http.Handle("/css/", http.FileServer(http.Dir("style")))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("appScripts"))))
	http.ListenAndServe(":8000", nil)
}

// localhost:8000/css/index.css

func ShowHtmlTemplate(response http.ResponseWriter, request *http.Request) {
	filePath := path.Join("template", "index.html")
	fmt.Println(filePath)
	applicationTemplate, _ := template.ParseFiles(filePath)
	applicationTemplate.Execute(response, nil)
}

func ShowAboutTemplate(response http.ResponseWriter, request *http.Request) {
	filePath := path.Join("template", "about.html")
	fmt.Println(filePath)
	applicationTemplate, _ := template.ParseFiles(filePath)
	applicationTemplate.Execute(response, nil)
}
