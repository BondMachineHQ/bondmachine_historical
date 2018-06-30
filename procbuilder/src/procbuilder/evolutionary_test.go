package procbuilder

import (
	"fmt"
	"math/rand"
	"mel"
	"testing"
	"time"
)

func TestEvolutionaryMachine(t *testing.T) {
	rand.Seed(int64(time.Now().Unix()))
	ep := new(mel.Evolution_parameters)
	ep.Pars = make(map[string]string)
	ep.Pars["procbuilder:opcodes"] = "nop,rset"
	fmt.Println(Machine_Program_Generate(ep))
}
