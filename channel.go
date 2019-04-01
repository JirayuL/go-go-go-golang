package main

import "fmt"

func workerr(c chan int) {
	c <- 10
}

func mainJJ() {
	c := make(chan int)
	go workerr(c)
	num := <-c
	fmt.Println(num)
}
