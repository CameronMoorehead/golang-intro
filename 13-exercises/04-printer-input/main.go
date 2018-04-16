package main

import "fmt"

func main() {
	fmt.Println("Please enter a small and large number:")
	var smallNumber int
	var largeNumber int

	fmt.Scan(&smallNumber, &largeNumber)
	fmt.Println(largeNumber % smallNumber)
}
