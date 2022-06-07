package main

import "fmt"

func main() {

	// Lifo Way to execute the Defere keyword
	defer helloAllOthers()
	defer helloAll()

	employeeMap := make(map[int]int)
	employeeMap[1] = 2

	employeeMapOther := map[string]string{
		"name":        "Mayank",
		"age":         "30",
		"designation": "developer",
	}

	delete(employeeMapOther, "age")
	delete(employeeMapOther, "name")

	for key, keyValue := range employeeMapOther {
		fmt.Println(key)
		fmt.Println(keyValue)
	}
}

func helloAll() {
	fmt.Println("Hello All")
}

func helloAllOthers() {
	fmt.Println("Hello All Others")
}
