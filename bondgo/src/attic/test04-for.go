package main

import ()

func main() {
	var reg_aa uint8
	var reg_ab uint8
	for reg_aa = 10; reg_aa; reg_aa-- {
		reg_ab = reg_aa + 1
		if reg_ab {
			continue
		}
	}
}
