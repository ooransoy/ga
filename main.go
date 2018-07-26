package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/*
	Shorthands:
	cset -> charset
	charset -> character set
	pop -> population
	genotype -> geno


	Glossary:
	genotype -> a defining string, or a DNA with a custom charset

	offspring -> child genotype of two genotypes created by crossover or
		crossover + mutation

	genetic operator -> an operator or function that contributes to the
		compution of a genetic algorithm; one of crossover, mutation,
		selection

	tournament selection -> a method of selection TODO explain
*/

// RNG seeder, calls an anonymous function that seeds the RNG and returns an
// empty struct. There is a return value and that return value is assigned to _
// so that the function runs before anything else. The return value is an empty
// struct because it does not carry any data, and we do not need any data.
var _ = func() struct{} { rand.Seed(time.Now().UnixNano()); return struct{}{} }()

// Crossover applies the crossover genetic operator to two given strings, and
// returns an offspring.
// The implementation is fairly simple, a split point is picked, TODO finish doc
func Crossover(a string, b string) (string, error) {
	if len(a) != len(b) {
		return "", errors.New(fmt.Sprintf("Lengths don't match, %s and %s", a, b))
	}

	splitPoint := rand.Intn(len(a))
	if rand.Int()%2 == 0 {
		// The reason behind this clause is to maximise stochasticity.
		// There is a 50% chance that the part before the split point is
		// taken from the genotype b instead of a.
		return b[:splitPoint] + a[splitPoint:], nil
	}
	return a[:splitPoint] + b[splitPoint:], nil
}

// Mutate applies the mutation genetic operator to a given string and returns
// a mutated string.
func Mutate(s string, cset string) string {
	mutationPoint := rand.Intn(len(s))
	bytes := []byte(s)
	bytes[mutationPoint] = sampleByte(cset)
	return string(bytes)
}

// RandomGenotype creates and returns a random genotype with a given length and
// and using a given charset.
func RandomGenotype(cset string, l int) string {
	b := make([]byte, l, l)
	for i, _ := range b {
		b[i] = sampleByte(cset)
	}
	return string(b)
}

// RandomPopulation creates and returns a random population of a given size that
// contains random genotypes with a given length using a given charset.
func RandomPopulation(cset string, genotypeLen, popSize int) []string {
	pop := make([]string, popSize, popSize)
	for i, _ := range pop {
		pop[i] = RandomGenotype(cset, genotypeLen)
	}
	return pop
}

// EvolvePopulation TODO write doc
func EvolvePopulation(pop []string, fitnessFunction func(string) float64, k int, cset string, mutationRate float64) ([]string, error) {
	offspring, err := CreateOffspring(pop, fitnessFunction, k, cset, mutationRate)
	if err != nil {
		return pop, err
	}

	var worstFitness = fitnessFunction(pop[0])
	var worstGIndex int
	for i := 1; i < len(pop); i++ {
		genotype := pop[i]
		f := fitnessFunction(genotype)
		if f < worstFitness {
			worstFitness = f
			worstGIndex = i
		}
	}
	return append(append(pop[:worstGIndex], pop[worstGIndex+1:]...), offspring), nil
}

func CreateOffspring(pop []string, fitnessFunction func(string) float64, k int, cset string, mutationRate float64) (string, error) {
	a := SelectFromSubset(SelectSubset(pop, k), fitnessFunction)
declareB:
	b := SelectFromSubset(SelectSubset(pop, k), fitnessFunction)

	if a == b {
		goto declareB
	}
	offspring, err := Crossover(a, b)

	if err != nil {
		return "", err
	}
	if rand.Float64() * 100 < mutationRate {
		return Mutate(offspring, cset), nil
	}
	return offspring, nil
}

// SelectSubset selects a completely random subset of size k from a given
// population.
func SelectSubset(pop []string, k int) []string {
	selected := make([]string, k, k)
	for i, _ := range selected {
	back:
		selected[i] = sample(pop)
		if i == 0 {
			continue
		}

		for j, _ := range selected {
			// If the newly selected genotype is already in the
			// selected subset,
			if selected[i] == selected[j] && j != i {
				goto back // Try again.
			}
		}
	}
	return selected
}

// SelectFromSubset finds and returns the fittest genotype from a given subset
// using a given fitness function.
func SelectFromSubset(subset []string, fitnessFunction func(string) float64) string {
	var selected string
	var bestFitness float64
	for _, genotype := range subset {
		fitness := fitnessFunction(genotype)
		if fitness > bestFitness {
			bestFitness = fitness
			selected = genotype
		}
	}
	return selected
}

func main() {
	cset := "0123456789"
	popSize := 10000
	genotypeLen := 5
	mutationRate := 80
	k := 500
	fitnessFunction := func(s string) float64 {
		n, _ := strconv.Atoi(s)
		return float64(n)
	}

	pop := RandomPopulation(cset, genotypeLen, popSize)
	for i := 0; i < 500; i++ {
		pop, _ = EvolvePopulation(pop, fitnessFunction, k, cset, float64(mutationRate))
	}
	fmt.Println("Fittest:", SelectFromSubset(pop, fitnessFunction))
}

func sampleByte(s string) byte {
	return s[rand.Intn(len(s))]
}

func sample(s []string) string {
	return s[rand.Intn(len(s))]
}
