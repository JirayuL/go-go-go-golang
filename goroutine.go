package main

import "fmt"
import "sync"

// import "time"

var wg sync.WaitGroup

func maiHHn() {
	wg.Add(4)
	go sayHello()
	go sayHello()
	go sayHello()
	go sayHello()
	// time.Sleep(1 * time.Second)
	wg.Wait()
	fmt.Println("World")
}

func sayHello() {
	fmt.Println("Hello")
	wg.Done()
}
