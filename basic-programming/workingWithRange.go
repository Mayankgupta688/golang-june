package main

import "fmt"

func workingWithRange() {
	myUserAge := []int{10, 20, 30, 40, 50, 60}

	for i := 0; i < len(myUserAge); i++ {
		fmt.Println(myUserAge[i])
	}

	for index, value := range myUserAge {
		fmt.Println(index)
		fmt.Println(value)
	}
}
