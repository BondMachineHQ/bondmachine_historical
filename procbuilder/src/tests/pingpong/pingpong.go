package main

import (
	"bondgo"
)

func pong(c chan uint8) {
	var ball uint8
	for {
		ball = <-c
		ball++
		c <- ball
	}
	bondgo.Void(ball)
}

func main() {
	var ch chan uint8
	var ball uint8

	ball = 1

	//ch = make(chan uint8)

	go pong(ch)

	for {
		ch <- ball
		ball = <-ch
	}
	bondgo.Void(ch)
	bondgo.Void(ball)
}
