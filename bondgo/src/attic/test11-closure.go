package main

import (
	"fmt"
)

func main() {
	var test func()
	test = func() {
	}
	fmt.Println(test)
}
