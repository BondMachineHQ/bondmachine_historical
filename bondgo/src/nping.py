#!/usr/bin/env python2

n=128

print """
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
"""
for i in range(n):
    print "\tvar ch"+str(i)+" chan uint8"

for i in range(n):
    print "\tvar ball"+str(i)+" uint8"

for i in range(n):
    print "\tball"+str(i)+"=1"

print """
	var sel uint8
"""

print """
	sel = 0

	//ch1 = make(chan uint8)
	//ch2 = make(chan uint8)


"""

for i in range(n):
    print "\tgo pong(ch"+str(i)+")"

print """
	for {
		select {
"""
for i in range(n):
    print "\t\tcase ch"+str(i)+" <- ball"+str(i)+":"
    print "\t\tsel = "+str(i*2+1)
    print "\t\tcase ball"+str(i)+" = <-ch"+str(i)+":"
    print "\t\tsel = "+str(i*2+2)

print """
		default:
			sel = 10000
		}
	}

"""

for i in range(n):
    print "\tbondgo.Void(ch"+str(i)+")"

for i in range(n):
    print "\tbondgo.Void(ball"+str(i)+")"

print """
    bondgo.Void(sel)
}
"""
