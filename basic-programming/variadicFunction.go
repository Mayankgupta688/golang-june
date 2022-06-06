package main

import (
	"fmt"
	"strconv"
)

func variadicFunction() {
	someFunction(1, 2, 3, 4, 5)
	someFunction(1, 2, 3, 4, 5, 6, 7, 8)
	someFunction(1, 2)
}

func someFunction(inputParameter ...int) {
	fmt.Println("Length of Array is: " + strconv.Itoa(len("Mayank Gupta")))
}
