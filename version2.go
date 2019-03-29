package main

import "fmt"

func main() {
	store := map[string]string{}
	var key string
	var temp string
	var command string
	for {
		fmt.Print("> ")
		fmt.Scanf("%v %v %v", &command, &key, &temp)

		if command == "GET" {
			value, ok := store[key]
			if ok {
				fmt.Printf("%v\n", value)
			} else {
				fmt.Println("Not Found")
			}
		} else if command == "SET" {
			store[key] = temp
			fmt.Printf("OK\n")
		} else {
			fmt.Println("Please input the right command!")
		}
	}
}
