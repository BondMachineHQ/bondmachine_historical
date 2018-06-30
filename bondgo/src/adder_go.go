package main

import (
	"bondgo"
)

func test() {
	var in0 bondgo.Input
	var out1 bondgo.Output
	var out0 bondgo.Output

	in0 = bondgo.Make(bondgo.Input, 3)
	out1 = bondgo.Make(bondgo.Output, 7)
	out0 = bondgo.Make(bondgo.Output, 5)

	for {
		bondgo.IOWrite(out0, bondgo.IORead(in0)+1)
	}

	bondgo.Void(in0)
	bondgo.Void(out0)
}

func main() {
	var in0 bondgo.Input
	var in1 bondgo.Input
	var out0 bondgo.Output

	in0 = bondgo.Make(bondgo.Input, 5)
	in1 = bondgo.Make(bondgo.Input, 8)
	out0 = bondgo.Make(bondgo.Output, 3)

device_1:
	go test()

	for {
		bondgo.IOWrite(out0, bondgo.IORead(in0)+1)
	}

	bondgo.Void(in0)
	bondgo.Void(out0)
}
