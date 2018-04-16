package main

import "fmt"

func main() {
	var name string
	fmt.Println("Enter a name")
	fmt.Scan(&name)
	fmt.Printf("Hello, %v\n", name)
}
