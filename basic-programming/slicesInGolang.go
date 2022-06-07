package main

import "fmt"

func slicesInGolang() {
	dataSlice := []int{10, 20, 30}
	dataLocation := &dataSlice[0]

	fmt.Printf("\t Length is: %d", len(dataSlice))
	fmt.Printf("\t Capacity is: %d", cap(dataSlice))
	fmt.Printf("\t %d \n", &dataSlice[0])

	for i := 0; i < 8; i++ {
		dataSlice = append(dataSlice, i)
		fmt.Printf("\t %v", dataSlice)
		fmt.Printf("\t %d", &dataSlice[0])
		fmt.Printf("\t Length is: %d", len(dataSlice))
		fmt.Printf("\t Capacity is: %d \n", cap(dataSlice))

	}

	fmt.Println(*dataLocation)

}
