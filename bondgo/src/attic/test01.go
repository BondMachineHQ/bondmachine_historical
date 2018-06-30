package main

import (
	"bondgo"
)

func test00(c chan uint16) {
	var b uint16
	for {
		b = <-c
	}
	bondgo.Void(b)
}

func main() {
	var ch chan uint16
	var a uint16

	a = 0

	go test00(ch)

	for {
		if a == 10 {
			break
		}
		ch <- a
		a++
	}
}
