package main

import "fmt"

var g = "cowboy"

func main() {
	a := 10
	b := "golang"
	c := 4.15
	d := true

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
}
