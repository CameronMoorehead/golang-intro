package main

import "fmt"

func main() {
	i := 1
	for i < 10 {
		fmt.Println("test")
		i++
	}

	for j := 0; j <= 5000; j++ {
		fmt.Printf("%v - %v - %v\n", j, string(j), []byte(string(j)))
	}
}
