package main

import (
	"fmt"
	"ga"
	"math/rand"
	"strconv"
	"time"
)

// RNG seeder, calls an anonymous function that seeds the RNG and returns an
// empty struct. There is a return value and that return value is assigned to _
// so that the function runs before anything else. The return value is an empty
// struct because it does not carry any data, and we do not need any data.
var _ = func() struct{} { rand.Seed(time.Now().UnixNano()); return struct{}{} }()

func main() {
	palette := make([]interface{}, 16)
	for i := range palette {
		palette[i] = uint8(i)
	}

	ps := 50  // Population Size
	gl := 2   // Genotype Length
	mr := 0.1 // Mutation Rate
	ec := 200 // Evolve count
	k := 3    // k used in selection

	ff := func(s []interface{}) float64 {
		var sum float64
		for _, b := range s {
			sum += float64(b.(uint8))
		}

		return sum
	}

	pop := ga.RandomPop(palette, gl, ps)
	for i := 0; i < ec; i++ {
		fmt.Println("Step", i+1)
		pop, _ = ga.EvolvePop(pop, ff, k, palette, float64(mr))
		fmt.Println(pop)
	}

	best := ga.Select(pop, ff)
	fmt.Println(
		"Fitness of best genotype in population after evolving",
		ec, "times in a population of", ps, "individiuals:",
		strconv.FormatFloat(ff(best), 'f', -1, 64),
	)
}
