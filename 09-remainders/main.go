package main

import "fmt"

func main() {
	x := 13 / 5
	fmt.Println(x)

	for i := 0; i < 100; i++ {
		if i%2 == 1 {
			fmt.Println("odd")
		} else {
			fmt.Println("even")
		}
	}
}
