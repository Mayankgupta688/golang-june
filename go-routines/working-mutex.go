package main

import (
	"fmt"
	"sync"
)

var balance = 5000
var mu = &sync.Mutex{}
var waitGrp sync.WaitGroup

func main1221() {
	waitGrp.Add(5)
	go ReadWriteSynchronousData()
	go ReadWriteSynchronousData()
	go ReadWriteSynchronousData()
	go ReadWriteSynchronousData()
	go ReadWriteSynchronousData()
	waitGrp.Wait()
}

func ReadWriteSynchronousData() {
	mu.Lock()

	balance = balance + 500
	fmt.Println(balance)

	mu.Unlock()
	waitGrp.Done()
}
