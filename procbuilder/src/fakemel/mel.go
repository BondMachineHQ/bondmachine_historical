package mel

import (
	"math/rand"
	"time"
)

// The main interface, it states: It is a mel object
type Me3li interface {
	Mel_init(*Evolution_parameters)
	Mel_copy() Me3li
}

type Me3li_string_import interface {
	Mel_string_import(string)
}

type Me3li_dump interface {
	Mel_dump()
}

func Init() {
	rand.Seed(int64(time.Now().Unix()))
}
