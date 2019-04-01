package main

import (
	"fmt"
	"time"
)

func workerS(sem chan int, job int) {
	time.Sleep(2 * time.Second)
	fmt.Println("work: ", job)
	<-sem
}

func mainS() {
	sem := make(chan int, 4)
	for i := 0; i < 1000; i++ {
		sem <- 1
		go workerS(sem, i)
	}
	for i := 0; i < 4; i++ {
		sem <- 1
	}
}
