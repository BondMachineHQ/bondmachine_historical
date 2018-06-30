package main

import (
	"fmt"
	"time"
)

func streamprocessor(a *[]uint8, b *[]uint8, c *[]uint8, gid uint8) {
	(*c)[gid] = (*a)[gid] + (*b)[gid]
}
func main() {

	a := make([]uint8, 256)
	b := make([]uint8, 256)
	c := make([]uint8, 256)

	for i := 0; i < 256; i++ {
		a[i] = uint8(i)
		b[i] = uint8(2)
		go streamprocessor(&a, &b, &c, uint8(i))
	}
	time.Sleep(3 * time.Second)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
