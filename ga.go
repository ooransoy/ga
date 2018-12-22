package ga

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Glossary:
	genotype -> a defining string, or a DNA with a custom encoding

	offspring -> child genotype of two genotypes created by crossover or
		crossover + mutation

	genetic operator -> an operator or function that contributes to the
		computation of a genetic algorithm; one of crossover, mutation,
		selection

	tournament selection -> a method of selection TODO explain
*/

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// Crossover applies the crossover genetic operator to two given genotypes, and
// returns an offspring.
// This function will be overridable, the current value is just a placeholder.
var Crossover = func(a []interface{}, b []interface{}) ([]interface{}, error) {
	if len(a) != len(b) {
		return []interface{}{}, fmt.Errorf("Lengths don't match, %v and %v", a, b)
	}

	sp := r.Intn(len(a)) // Split point
	if r.Int()%2 == 0 {
		// There is a 50% chance that the part before the split point is
		// taken from the genotype b instead of a.
		return append(b[:sp], a[sp:]...), nil
	}
	return append(a[:sp], b[sp:]...), nil
}

// Mutate applies the mutation genetic operator to a given genotype and returns
// a mutated genotype.
// This function will be overridable, the current value is just a placeholder.
var Mutate = func(g []interface{}, palette []interface{}) ([]interface{}, error) {
	if len(palette) == 0 {
		return []interface{}{}, fmt.Errorf("Palette is empty")
	}

	g[r.Intn(len(g))] = palette[r.Intn(len(palette))]
	return g, nil
}

// RandomGenotype creates and returns a random genotype with a given length and
// and using a given charset.
func RandomGenotype(palette []interface{}, l int) []interface{} {
	b := make([]interface{}, l, l)
	for i, _ := range b {
		b[i] = palette[r.Intn(len(palette))]
	}
	return b
}

// RandomPop creates and returns a random population of a given size that
// contains random genotypes with a given length using a given charset.
func RandomPop(palette []interface{}, gl, ps int) [][]interface{} {
	pop := make([][]interface{}, ps, ps)
	for i, _ := range pop {
		pop[i] = RandomGenotype(palette, gl)
	}
	return pop
}

// EvolvePop TODO write doc
func EvolvePop(pop [][]interface{}, ff func([]interface{}) float64, k int, palette []interface{}, mr float64) ([][]interface{}, error) {
	for _, g := range pop {
		if r.Float64() < mr {
			var err error
			g, err = Mutate(g, palette)
			if err != nil {
				return pop, err
			}
		}
	}
	o, err := Offspring(pop, ff, k, palette)
	if err != nil {
		return pop, err
	}

	wi := WorstIndex(pop, ff) // Worst Genotype Index

	return append(append(pop[:wi], pop[wi+1:]...), o), nil
}

func WorstIndex(pop [][]interface{}, ff func([]interface{}) float64) int {
	var wf = ff(pop[0]) // Worst Fitness
	var wi int          // Worst Genotype Index

	for i := 1; i < len(pop); i++ {
		g := pop[i]
		f := ff(g)
		if f < wf {
			wf = f
			wi = i
		}
	}

	return wi
}

func Offspring(pop [][]interface{}, ff func([]interface{}) float64, k int, palette []interface{}) ([]interface{}, error) {
	a := Select(SelectSet(pop, k), ff)

	set := SelectSet(pop, k)
	for i, g := range set {
		if same(g, a) {
			set = append(set[:i], set[i+1:]...)
			break
		}
	}
	b := Select(set, ff)

	o, err := Crossover(a, b)

	if err != nil {
		return []interface{}{}, err
	}
	return o, nil
}

// SelectSet selects a completely random subset of size k from a given
// population.
func SelectSet(pop [][]interface{}, k int) [][]interface{} {
	s := make([][]interface{}, k, k) // Selection
	p := rand.Perm(len(pop))[:k]     // k numbers picked from [0,len(pop))

	for i := range s {
		s[i] = pop[p[i]]
	}

	return s
}

// Select finds and returns the fittest genotype from a given set of genotypes
// using a given fitness function.
func Select(set [][]interface{}, ff func([]interface{}) float64) []interface{} {
	var s []interface{} // Selection
	var bf float64      // Best Fitness
	for _, g := range set {
		f := ff(g)
		if f >= bf {
			bf = f
			s = g
		}
	}

	return s
}

func same(a, b []interface{}) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
