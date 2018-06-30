package main

import ()

func test02(a uint8) uint8 {
	var b uint8
	b = a + 1
}

func main() {
	var reg_prova1 uint8
	var reg_prova2 uint8
	reg_prova1 = 4
	go test02(reg_prova1)
}
