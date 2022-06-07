package main

import "fmt"

func workingWithArrays() {
	marks := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(len(marks))
	fmt.Println(cap(marks))

	marksData := [10]int{1, 2, 3, 4, 5}
	fmt.Println(len(marksData))
	fmt.Println(cap(marksData))
	fmt.Println(marksData[7])

	var myArray [3]string
	myArray[0] = "Hello"
	myArray[1] = " World"
	myArray[2] = " Hi"

	fmt.Println(myArray)
}
