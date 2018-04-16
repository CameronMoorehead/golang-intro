package main

import "fmt"

func main() {
	naturals()
}

func naturals() {
	var num int
	i := 999
	for 0 < i {
		if i%3 == 0 {
			num += i
		} else if i%5 == 0 {
			num += i
		}
		i--
	}
	fmt.Println(num)
}
