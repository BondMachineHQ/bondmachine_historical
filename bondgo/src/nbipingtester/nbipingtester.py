#!/usr/bin/env python2

import subprocess
import re

REALCORES=8
ITERATIONS=2048
CPRANGE=[2,4,8,16,32,64,128]

for CP in CPRANGE:

    result= """package nbipingtester

import (
	"bondgo"
	"sync"
)

func pinge(wg *sync.WaitGroup, wgi *sync.WaitGroup) {
	var ball uint8
	var ch chan uint8

	ch = make(chan uint8)

	go ponge(ch)

        wgi.Done()
        wgi.Wait()

	ball = 0

	wg.Done()

	bondgo.Void(ch)
	bondgo.Void(ball)
}

func ponge(c chan uint8) {
	var ball uint8

	bondgo.Void(ball)
}

func ping(wg *sync.WaitGroup, wgi *sync.WaitGroup) {
	var ball uint8
	var ch chan uint8

	ch = make(chan uint8)

	go pong(ch)
        
        wgi.Done()
        wgi.Wait()

	ball = 0

	for i := 0; i < """+str(ITERATIONS)+"""; i++ {
		ch <- ball
		ball = <-ch
	}

	wg.Done()

	bondgo.Void(ch)
	bondgo.Void(ball)
}

func pong(c chan uint8) {
	var ball uint8
	for i := 0; i < """+str(ITERATIONS)+"""; i++ {
		ball = <-c
		ball++
		c <- ball
	}
	bondgo.Void(ball)
}

func Bstarte() {
	var sel1 uint8
	var sel2 uint8
	sel1 = 0
	sel2 = sel1

	var wg sync.WaitGroup
	var wgi sync.WaitGroup
"""

    result+="      wg.Add("+str(CP)+")\n"
    result+="      wgi.Add("+str(CP)+")\n"

    for i in range(CP):
        result+="       go pinge(&wg,&wgi)\n"

    result+="""

	wg.Wait()

	bondgo.Void(sel1)
	bondgo.Void(sel2)
}

func Bstart() {
	var sel1 uint8
	var sel2 uint8
	sel1 = 0
	sel2 = sel1

	var wg sync.WaitGroup
	var wgi sync.WaitGroup
"""

    result+="      wg.Add("+str(CP)+")\n"
    result+="      wgi.Add("+str(CP)+")\n"

    for i in range(CP):
        result+="       go ping(&wg,&wgi)\n"

    result+="""

	wg.Wait()

	bondgo.Void(sel1)
	bondgo.Void(sel2)
}

"""

    libfile = open("nbipingtester.go","w")
    libfile.write(result)
    libfile.close()

    result = """package nbipingtester

import (
	"runtime"
	"testing"
)

func BenchmarkBiping(b *testing.B) {
	runtime.GOMAXPROCS("""+str(REALCORES)+""")
	for n := 0; n < b.N; n++ {
		Bstart()
	}
}

func BenchmarkBipinge(b *testing.B) {
	runtime.GOMAXPROCS("""+str(REALCORES)+""")
	for n := 0; n < b.N; n++ {
		Bstarte()
	}
}

func TestBiping(t *testing.T) {
	Bstart()
}
"""

    testfile = open("nbipingtester_test.go","w")
    testfile.write(result)
    testfile.close()

    totaltime=0
    inittime=0

    p=subprocess.Popen("go test -bench=.", shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE, close_fds=True)
    p.wait()
    if p.returncode==0:
        while True:
            o = p.stdout.readline()
            result = re.match('BenchmarkBipinge.* (\d+) ns/op.*', o)
            if result:
                inittime = int(result.group(1))
                #print inittime
            result = re.match('BenchmarkBiping[^e].* (\d+) ns/op.*', o)
            if result:
                totaltime = int(result.group(1))
                #print totaltime
            if o == '' and p.poll() != None:
                break
    nettime=(totaltime-inittime)
    timepercp=nettime/CP
    timepercpint=timepercp/ITERATIONS
    timeperint=nettime/ITERATIONS
    print CP, timeperint
