package main

import (
	"bondgo"
)

func main() {
	var prova1 uint8
	var out_3 bondgo.Output
	out_3 = bondgo.Make(bondgo.Output, 5)
	prova1 = 4
	bondgo.IOWrite(out_3, prova1)
}
