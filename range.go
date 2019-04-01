package main

import "fmt"

func count(c chan int) {
	for i := 0; i < 20; i++ {
		c <- 1
	}
	close(c)
}

func mainR() {
	c := make(chan int)
	go count(c)
	for i := range c {
		fmt.Println(i)
	}
}
