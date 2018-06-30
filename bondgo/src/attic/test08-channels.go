package main

import (
	"fmt"
	"time"
)

func main() {

	testchan := make(chan uint8)

	go func(c chan uint8) {
		time.Sleep(2 * time.Second)
		select {
		case c <- 0:
		default:
		}
	}(testchan)

	select {
	case dato := <-testchan:
		fmt.Println(dato)
	default:
	}

}
