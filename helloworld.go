package main

import "fmt"

func mainNN() {
	// var helloWorld string
	// helloWorld = "Hello World!"
	// fmt.Println(helloWorld)

	// var hello, world = "Hello", "World"
	// fmt.Println(hello, world)

	// helloo, worldd := "Hello", "World"
	// fmt.Println(helloo, worldd)

	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println("Even")
	// 	}
	// 	fmt.Println(i)
	// }

	// // FizzBuzz
	// for i := 0; i <= 100; i++ {
	// 	if i%3 == 0 && i%5 == 0 {
	// 		fmt.Println("FizzBuzz")
	// 	} else if i%3 == 0 {
	// 		fmt.Println("Fizz")
	// 	} else if i%5 == 0 {
	// 		fmt.Println("Buzz")
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }

	// var x int
	// fmt.Printf("Please input the number: ")
	// fmt.Scanf("%d", &x)
	// fmt.Scanln("%d", &x)
	// fmt.Printf("%d\n", x)
	// %v for all format
	// fmt.Println("%v", x)

	// Array
	// var a [3]int
	// a[0] = 1
	// a[1] = 2
	// a[2] = 69
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(a[i])
	// }

	// Slices
	// a := []int{1, 42, 99}
	// fmt.Println(len(a))

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(a[i])
	// }

	// a = append(a, 250)
	// fmt.Println(len(a))

	// fmt.Println(a[1:3])

	// Map
	// m := make(map[string]string)
	// m["Hello"] = "World"
	// m["Hi"] = "yoyo"

	// value, ok := m["Hello"]

	// fmt.Println(value, ok)

	// x := []string{"a", "b", "c", "d"}
	// for _, value := range x {
	// 	fmt.Println(value)
	// }

	x := map[string]string{"A": "a", "B": "b", "C": "c", "D": "d"}
	for key, value := range x {
		fmt.Println(key, value)
	}
}
