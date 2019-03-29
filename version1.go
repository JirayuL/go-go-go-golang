package main

import "fmt"

func mainn() {
	var store string
	var temp string
	var command string
	for {
		fmt.Print("> ")
		fmt.Scanf("%v %v", &command, &temp)

		if command == "GET" {
			fmt.Printf("%v\n", store)
		} else if command == "SET" {
			store = temp
			fmt.Printf("OK\n")
		}
	}
}
