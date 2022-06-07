package main

import (
	"fmt"
)

func slicingArray() {
	someArray := [11]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	otherArray := someArray[:8]
	otherArray = someArray[:8]
	otherArray = someArray[2:5]
	fmt.Println(otherArray)
}
