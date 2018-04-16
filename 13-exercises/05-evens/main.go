package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if isEven(i) {
			fmt.Println(i)
		}
	}
}

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}
