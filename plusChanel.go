package main

import "fmt"
import "time"

func sum(arr []int, c chan int) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	c <- sum
}

func mainGG() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)

	startTime := time.Now()

	go sum(arr[0:5], c)
	go sum(arr[5:10], c)
	a := <-c
	b := <-c

	fmt.Println(a + b)

	stopTime := time.Since(startTime)
	fmt.Println(stopTime)
}
