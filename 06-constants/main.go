package main

import "fmt"

const p string = "death & taxes"

const (
	_  = iota
	kb = 1<<(iota*10) - 1
	mb = 1<<(iota*10) - 1
)

func main() {
	// Invalid reassignment
	// p = "whisky & bbq"

	fmt.Println(p)
	fmt.Printf("%b\n", kb)
	fmt.Printf("%b\n", mb)
}
