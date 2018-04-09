package main

import (
	"fmt"

	"github.com/CameronMoorehead/golang-intro/02-package/stringutils"
)

func main() {
	fmt.Println(stringutils.Reverse("Hello World"))
	fmt.Println(stringutils.MyName)
}
