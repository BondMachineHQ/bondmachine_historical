package bondmachine

import (
	"fmt"
	"math/rand"
	"mel"
	"testing"
	"time"
)

func TestEvolutionaryBondmachine(t *testing.T) {
	rand.Seed(int64(time.Now().Unix()))
	ep := new(mel.Evolution_parameters)
	ep.Pars = make(map[string]string)
	fmt.Println(ep)
}

func TestEvolutionaryFitnell(t *testing.T) {
	rand.Seed(int64(time.Now().Unix()))
	ep := new(mel.Evolution_parameters)
	ep.Pars = make(map[string]string)
	fmt.Println(ep)
}
