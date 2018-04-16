package main

import "fmt"

func main() {
	if true {
		fmt.Println("This ran")
	}

	if false {
		fmt.Println("This did not run")
	}

	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	// fmt.Println(food) - undefined
}
