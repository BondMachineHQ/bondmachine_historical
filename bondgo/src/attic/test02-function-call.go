package main

import ()

func test01(a uint8) uint8 {
	return 5 + a
}

func test02(a uint8) uint8 {
	return 6 + test01(a)
}

func main() {
	var reg_prova1 uint8
	var reg_prova2 uint8
	reg_prova1 = test02(4)
}
