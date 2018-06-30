package main

import (
	"bondgo"
)

func main() {
	var reg_prova1 uint8
	var reg_prova2 uint8
	reg_prova1 = 1
	reg_prova2 = 2
	reg_prova1, reg_prova2 = 2, 1
	bondgo.Void(reg_prova1)
	bondgo.Void(reg_prova2)
}
