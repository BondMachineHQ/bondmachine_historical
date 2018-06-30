package main

import ()

func test01(a uint8) uint8 {
	return 5 + a
}

func main() {
	var reg_prova1 uint8
	reg_prova1 = test01(4)
}
