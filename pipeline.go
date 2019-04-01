package main

import "fmt"

func generator() chan int {
	output := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			output <- i
		}
		close(output)
	}()
	return output
}

func multiply(input chan int) chan int {
	output := make(chan int)
	go func() {
		for num := range input {
			output <- num * num
		}
		close(output)
	}()
	return output
}

func print(input chan int) {
	for num := range input {
		fmt.Println(num)
	}

}

func addd(input chan int) chan int {
	output := make(chan int)
	go func() {
		for num := range input {
			output <- num + num
		}
		close(output)
	}()
	return output
}

func mainPP() {
	g := generator()
	m := multiply(g)
	a := addd(m)
	print(a)
}
