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

	var ch0 chan uint8
	var ch1 chan uint8
	var ch2 chan uint8
	var ch3 chan uint8
	var ch4 chan uint8
	var ch5 chan uint8
	var ch6 chan uint8
	var ch7 chan uint8
	var ball0 uint8
	var ball1 uint8
	var ball2 uint8
	var ball3 uint8
	var ball4 uint8
	var ball5 uint8
	var ball6 uint8
	var ball7 uint8
	ball0 = 1
	ball1 = 1
	ball2 = 1
	ball3 = 1
	ball4 = 1
	ball5 = 1
	ball6 = 1
	ball7 = 1

	var sel uint8

	sel = 0

	//ch1 = make(chan uint8)
	//ch2 = make(chan uint8)

	go pong(ch0)
	go pong(ch1)
	go pong(ch2)
	go pong(ch3)
	go pong(ch4)
	go pong(ch5)
	go pong(ch6)
	go pong(ch7)

	for {
		select {

		case ch0 <- ball0:
			sel = 1
		case ball0 = <-ch0:
			sel = 2
		case ch1 <- ball1:
			sel = 3
		case ball1 = <-ch1:
			sel = 4
		case ch2 <- ball2:
			sel = 5
		case ball2 = <-ch2:
			sel = 6
		case ch3 <- ball3:
			sel = 7
		case ball3 = <-ch3:
			sel = 8
		case ch4 <- ball4:
			sel = 9
		case ball4 = <-ch4:
			sel = 10
		case ch5 <- ball5:
			sel = 11
		case ball5 = <-ch5:
			sel = 12
		case ch6 <- ball6:
			sel = 13
		case ball6 = <-ch6:
			sel = 14
		case ch7 <- ball7:
			sel = 15
		case ball7 = <-ch7:
			sel = 16

		default:
			sel = 100
		}
	}

	bondgo.Void(ch0)
	bondgo.Void(ch1)
	bondgo.Void(ch2)
	bondgo.Void(ch3)
	bondgo.Void(ch4)
	bondgo.Void(ch5)
	bondgo.Void(ch6)
	bondgo.Void(ch7)
	bondgo.Void(ball0)
	bondgo.Void(ball1)
	bondgo.Void(ball2)
	bondgo.Void(ball3)
	bondgo.Void(ball4)
	bondgo.Void(ball5)
	bondgo.Void(ball6)
	bondgo.Void(ball7)

	bondgo.Void(sel)
}
