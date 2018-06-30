package main

import (
	"fmt"
)

func test() (a uint8, b uint8) {
	return 5, 6
}

func main() {
	a, b := test()
}
