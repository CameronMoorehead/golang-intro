package main

import "fmt"

func zero(y *int) {
	*y = 0
}

func main() {
	x := 5
	y := &x

	fmt.Println(*y)
	fmt.Println(x)

	zero(y)

	fmt.Println(*y)
	fmt.Println(x)
}
