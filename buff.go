package main

import "fmt"

func workerB(buff chan int) {
	fmt.Println(<-buff)
	fmt.Println(<-buff)
	fmt.Println(<-buff)
}

func mainBBB() {
	bufferChannel := make(chan int, 3)
	go workerB(bufferChannel)

	bufferChannel <- 10
	bufferChannel <- 11
	bufferChannel <- 12
}
