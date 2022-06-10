package main

import (
	"fmt"
	"sync"
)

func main111() {

	var executionNotifier sync.WaitGroup

	executionNotifier.Add(2)

	go func() {
		var dataToRender DataRendering
		dataToRender.ShowBoldName = false
		dataToRender.EmployeeList = getEmployees()
		executionNotifier.Done()
	}()

	go func() {
		GetSBIStockValue()
		executionNotifier.Done()
	}()

	fmt.Println("Done")

	fmt.Println("Done Again")

	executionNotifier.Wait()
}
