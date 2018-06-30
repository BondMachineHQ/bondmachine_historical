package main

import (
	"fmt"
	"time"
)

func main() {

	testchan := make(chan uint8)

	go func(c chan uint8) {
		time.Sleep(1 * time.Second)
		select {
		case c <- 0:
		}
	}(testchan)

	go func(c chan uint8) {
		time.Sleep(2 * time.Second)
		select {
		case c <- 1:
		default:
		}
	}(testchan)

	time.Sleep(4 * time.Second)

	select {
	case dato := <-testchan:
		fmt.Println(dato)
	}
}
