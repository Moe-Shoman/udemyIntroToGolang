package main

import (
	"fmt"
	"sync"
)

//have multiple threads running

func oneWay() {
	fmt.Println("************* oneWay Started ****************")
	for i := 0; i < 1000; i++ {
		fmt.Println("oneWay: ", i)
	}
	fmt.Println("************* oneWay Finished ****************")
	yas.Done()
}
func twoWay() {
	fmt.Println("************* oneWay Started ****************")
	for i := 0; i < 1000; i++ {
		fmt.Println("twoWay: ", i)
	}
	fmt.Println("************* twoWay Finished ****************")
	yas.Done()
}

var yas sync.WaitGroup

func main() {
	yas.Add(2)
	go oneWay()
	go twoWay()
	yas.Wait()
}
