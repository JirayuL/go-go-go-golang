package main

import "fmt"
import "strings"
import "sync"

var wg sync.WaitGroup

func main2() {
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
		} else if command == "MSET" {
			keys := strings.Split(key, ",")
			wg.Add(len(keys))
			for i := 0; i < len(keys); i++ {
				go func(i int) {

					value, ok := store[keys[i]]
					if ok {
						fmt.Printf("%v: %v\n", keys[i], value)
					} else {
						fmt.Printf("%v: NOT FOUND\n", keys[i])
					}
					wg.Done()
				}(i)
			}
			fmt.Printf("OK\n")
		} else {
			fmt.Println("Please input the right command!")
		}
	}
}
