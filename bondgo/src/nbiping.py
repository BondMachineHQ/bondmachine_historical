#!/usr/bin/env python2

n=16

print """
package main

import (
	"bondgo"
)

func ping() {
	var ball uint8
        var ch chan uint8

        go pong(ch)

        ball = 0

	for {
		ch <- ball
		ball = <-ch
	}

	bondgo.Void(ch)
	bondgo.Void(ball)
}

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
        var sel1 uint8
        var sel2 uint8
        sel1 = 0
        sel2 = sel1
"""

for i in range(n):
    print "\tgo ping()"

print """
        bondgo.Void(sel1)
        bondgo.Void(sel2)
}
"""
