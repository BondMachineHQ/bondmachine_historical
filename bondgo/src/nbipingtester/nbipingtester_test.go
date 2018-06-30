package nbipingtester

import (
	"runtime"
	"testing"
)

func BenchmarkBiping(b *testing.B) {
	runtime.GOMAXPROCS(8)
	for n := 0; n < b.N; n++ {
		Bstart()
	}
}

func BenchmarkBipinge(b *testing.B) {
	runtime.GOMAXPROCS(8)
	for n := 0; n < b.N; n++ {
		Bstarte()
	}
}

func TestBiping(t *testing.T) {
	Bstart()
}
