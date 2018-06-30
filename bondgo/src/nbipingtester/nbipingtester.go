package nbipingtester

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

	for i := 0; i < 2048; i++ {
		ch <- ball
		ball = <-ch
	}

	wg.Done()

	bondgo.Void(ch)
	bondgo.Void(ball)
}

func pong(c chan uint8) {
	var ball uint8
	for i := 0; i < 2048; i++ {
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
      wg.Add(128)
      wgi.Add(128)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)
       go pinge(&wg,&wgi)


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
      wg.Add(128)
      wgi.Add(128)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)
       go ping(&wg,&wgi)


	wg.Wait()

	bondgo.Void(sel1)
	bondgo.Void(sel2)
}

