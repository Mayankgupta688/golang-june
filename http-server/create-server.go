package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createserver() {
	http.HandleFunc("/simpleprint", PrintSimpleString)
	http.HandleFunc("/showjsondata", PrintSimplejson)
	http.HandleFunc("/middleware", logging(PrintSimplejson))
	http.ListenAndServe(":4000", nil)
}

type Employee struct {
	Name string
	Age  int16
}

func PrintSimpleString(w http.ResponseWriter, r *http.Request) {
	var userName string = "Mayank Gupta"
	fmt.Fprintf(w, "Hello "+userName)
}

func PrintSimplejson(w http.ResponseWriter, r *http.Request) {
	emp := Employee{"Mayank", 20}
	dataOutput, _ := json.Marshal(emp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataOutput)
}

func logging(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Loging Function is Called as middleware")
		handler(w, r)
	}
}
