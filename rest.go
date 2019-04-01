package main

import "sync"

var wggg sync.WaitGroup
var mutex sync.Mutex
var x int

func add() {
	mutex.Lock()
	x++
	mutex.Unlock()
}

func worker() {
	for i := 0; i < 1000000; i++ {
		add()
	}
	wggg.Done()
}

func mainRR() {
	wggg.Add(2)
	go worker()
	go worker()
	wggg.Wait()

}
