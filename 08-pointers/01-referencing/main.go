package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	b := &a

	fmt.Println(b)

	a = 11

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(b)
}
