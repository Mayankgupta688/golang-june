package main

import "fmt"
import "io/ioutil"

func multipleValuesReturnFromFunction() {
	a, _, c, d := getDataFromFunction()

	fmt.Printf("Value of A is %v \n", a)
	fmt.Printf("Value of C is %v \n", c)
	fmt.Printf("Value of C is %v \n", d)
}

func getDataFromFunction() (byte, string, bool, float32) {
	return 1, "string Data", false, 1.5
}
