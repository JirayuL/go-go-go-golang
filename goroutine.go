package main

import "fmt"
import "sync"

// import "time"

var wgg sync.WaitGroup

func maiHHn() {
	wgg.Add(4)
	go sayHello()
	go sayHello()
	go sayHello()
	go sayHello()
	// time.Sleep(1 * time.Second)
	wgg.Wait()
	fmt.Println("World")
}

func sayHello() {
	fmt.Println("Hello")
	wgg.Done()
}
