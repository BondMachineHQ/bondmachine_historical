package main

import (
	"bondgo"
	"fmt"
	"time"
)

func two() {
	var out_0 bondgo.Output
	var val uint8
	out_0.Make(1)
	bondgo.AllWait()
	val = uint8(5)
	out_0.Write(val)
}

func three() {
	var in_0 bondgo.Input
	in_0.Make(1, uint8(0))
	bondgo.AllWait()
	val := in_0.Read().(uint8)
	fmt.Println(val)
}

func main() {
	bondgo.AllInit(3)
	go two()
	go three()
	var in_0 bondgo.Input
	in_0.Make(1, uint8(0))
	bondgo.AllWait()
	for i := 0; i < 10; i++ {
		val := in_0.Read().(uint8)
		fmt.Println(val)
		time.Sleep(1 * time.Millisecond)
	}
}
