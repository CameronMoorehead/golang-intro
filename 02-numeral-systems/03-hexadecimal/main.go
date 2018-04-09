package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("%d - %#x \n", 42, 42)
	fmt.Printf("%d \t %#X \n", 42, 42)
	fmt.Println(rand.Intn(100))
}
