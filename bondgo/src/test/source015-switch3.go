package main

import (
	"bondgo"
)

func main() {
	var a uint8
	switch {
	case true:
		a = 11
	case false:
		a = 12
	}

	bondgo.Void(a)
	bondgo.Void(b)
}
