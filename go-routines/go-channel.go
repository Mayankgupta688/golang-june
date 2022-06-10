package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	accountChannel := make(chan string)
	var syncData sync.WaitGroup

	syncData.Add(1)

	go func() {

		// Receiver Function

		time.Sleep(14 * time.Second)

		fmt.Println(<-accountChannel)
		syncData.Done()
	}()

	go func() {

		// Sender Function

		time.Sleep(15 * time.Second)

		accountChannel <- "Sample Data"
		syncData.Done()
	}()

	fmt.Println("Account Complete...")

	syncData.Wait()
	
	fmt.Println("One of the Process Complete")
}
