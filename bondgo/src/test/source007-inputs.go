package main

import ()

func main() {
	var prova1 uint8
	var in_2 bondgo.Input
	in_2 = bondgo.Make(bondgo.Input, 1)
	prova1 = bondgo.IORead(in_2)
}
