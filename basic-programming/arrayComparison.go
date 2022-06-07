package main

import "fmt"

func arrayComparison() {
	sampleArray := [3]int{1, 2, 3}
	sampleArrayOther := [3]int{1, 2, 3}
	otherArray := [...]int{1, 2, 3}

	// This is "Deep" copy
	copyArray := sampleArray

	// Array of 5 Elements, assign value to 35 an 57th element = 20, 30

	myArray := [70]int{35: 20, 57: 30}

	fmt.Println(sampleArray == sampleArrayOther)
	fmt.Println(sampleArray == otherArray)

	fmt.Println(myArray)
	fmt.Println(copyArray)
	fmt.Println(sampleArray)

	fmt.Println(&copyArray == &sampleArray)
}
