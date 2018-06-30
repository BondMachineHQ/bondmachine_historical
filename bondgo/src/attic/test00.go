package main

import (
	"bondgo"
)

func test00(c chan uint8) {
	var b uint8
	for {
		b = <-c
	}
	bondgo.Void(b)
}

func main() {
	var ch chan uint8
	var a uint8

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
