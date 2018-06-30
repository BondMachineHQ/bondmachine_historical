package main

import (
	"bondgo"
)

func main() {
	var a uint8
	var b uint8
	a = 1
	switch a {
	case 1:
		b = 11
	case 2:
		b = 12
		fallthrough
	case 3:
		b = 13
	}

	switch {
	}

	bondgo.Void(a)
	bondgo.Void(b)
}
