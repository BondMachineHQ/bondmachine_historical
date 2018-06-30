package ping

import (
	"bondgo"
	"testing"
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

func BenchmarkPing(b *testing.B) {
	var ch chan uint8
	var ball uint8

	ball = 1

	ch = make(chan uint8)

	go pong(ch)

	for i := 0; i < b.N; i++ {
		ch <- ball
		ball = <-ch
	}
	bondgo.Void(ch)
	bondgo.Void(ball)
}
