package main

import ()

func test() chan uint8 {
	newchan := make(chan uint8)
	return newchan
}

func main() {
	var testchan chan uint8
	testchan = test()
}
