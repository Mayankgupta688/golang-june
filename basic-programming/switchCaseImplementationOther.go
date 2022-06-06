package main

import (
	"fmt"
)

func switchCaseImplementationOther() {
	userdata := "Employee"

	switch userdata {
	case "Director":
		fmt.Println("User was a Director")
		fallthrough
	case "VP":
		fmt.Println("User was a VP")
		fallthrough
	case "Manager":
		fmt.Println("User was a Manager")
		fallthrough
	case "Employee":
		fmt.Println("User was a Employee")
	}

}
