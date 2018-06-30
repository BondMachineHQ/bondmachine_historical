package main

import (
	"bondgo"
)

func main() {
	var reg_prova uint8
	reg_prova = 1
	{
		var reg_prova uint8
		reg_prova = 2
		reg_prova = 3
		bondgo.Void(reg_prova)
	}
	reg_prova = 4

	var prova uint8
	prova = 1
	{
		var prova uint8
		prova = 2
		prova = 3
		bondgo.Void(prova)
	}
	prova = 4
	bondgo.Void(reg_prova)
	bondgo.Void(prova)
}
