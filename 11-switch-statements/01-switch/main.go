package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Danny")
	case "Cameron":
		fmt.Println("Cam")
	case "Medhi":
		fmt.Println("Who?")
		fallthrough
	default:
		fmt.Println("Nope")
	}
}
