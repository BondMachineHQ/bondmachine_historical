package main

import (
	"bondgo"
)

func main() {
	var a uint8
	var b uint8
	a = 1
	switch {
	case 1 == 1, 6 == 5:
		b = 11
	default:
		b = 4
	}
	bondgo.Void(a)
	bondgo.Void(b)
}
