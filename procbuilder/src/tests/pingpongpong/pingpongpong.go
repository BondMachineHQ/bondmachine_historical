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
	var ch1 chan uint8
	var ch2 chan uint8
	var ball1 uint8
	var ball2 uint8
	var sel uint8

	ball1 = 1
	ball2 = 1
	sel = 0

	//ch1 = make(chan uint8)
	//ch2 = make(chan uint8)

	go pong(ch1)
	go pong(ch2)

	for {
		select {
		case ch1 <- ball1:
			sel = 1
		case ball1 := <-ch1:
			sel = 2
		case ch2 <- ball2:
			sel = 3
		case ball2 = <-ch2:
			sel = 4
		default:
			sel = 5
		}
	}
	bondgo.Void(ch1)
	bondgo.Void(ch2)
	bondgo.Void(ball1)
	bondgo.Void(ball2)
	bondgo.Void(sel)
}
