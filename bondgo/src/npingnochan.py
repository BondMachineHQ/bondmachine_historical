#!/usr/bin/env python2

n=2048

print """
package main

import (
	"bondgo"
)

func pong() {
	var ball uint8
	var in_ball uint8
        var out_ball uint8
	for {
                ball=in_ball
		ball++
                out_ball=ball
	}
	bondgo.Void(ball)
	bondgo.Void(in_ball)
	bondgo.Void(out_ball)
}

func main() {
	var ball uint8
	var in_ball uint8
        var out_ball uint8
"""

for i in range(n):
        print "\tgo pong()"

print """
	for {
                ball=in_ball
		ball++
                out_ball=ball
	}
	bondgo.Void(ball)
	bondgo.Void(in_ball)
	bondgo.Void(out_ball)
"""

print """
}
"""
