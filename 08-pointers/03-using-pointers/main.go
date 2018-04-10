package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 42

	c := b

	fmt.Println(a)
	fmt.Println(*c)
}
