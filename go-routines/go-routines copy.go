package main

import (
	"fmt"
	"time"
)

func main11() {
	func() {
		var dataToRender DataRendering
		dataToRender.ShowBoldName = false
		dataToRender.EmployeeList = getEmployees()
	}()

	func() {
		time.Sleep(10 * time.Second)
		fmt.Println("Function Two Done")
	}()

	fmt.Println("Done")

	time.Sleep(5 * time.Second)

	fmt.Println("Done Again")
}
