package main

import (
	"fmt"
	"time"
)

func worker1(c chan int) {
	for {
		time.Sleep(1 * time.Second)
		c <- 10
	}
}

func worker2(c chan int) {
	for {
		time.Sleep(4 * time.Second)
		c <- 20
	}
}

func worker3(c chan int) {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println(<-c)
	}
}

func timer(stop chan int) {
	time.Sleep(5 * time.Second)
}

func mainFF() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	go worker1(c1)
	go worker2(c2)

	for {
		select {
		case n := <-c1:
			fmt.Println(n)
		case n := <-c2:
			fmt.Println(n)
		case c3 <- 10:
			fmt.Println("Send to c3")
		default:
			time.Sleep(1 * time.Second)
			fmt.Println(".")
		}
	}
}
