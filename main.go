package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ghostec/lcga/circuits"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	inputs := []interface{}{
		[]int{0, 0, 0, 0}, // 0, 0
		[]int{0, 0, 0, 1}, // 0, 1
		[]int{0, 0, 1, 0}, // 1, 0
		[]int{0, 0, 1, 1}, // 1, 1
		[]int{0, 1, 0, 0}, // 0, 1
		[]int{0, 1, 0, 1}, // 1, 0
		[]int{0, 1, 1, 0}, // 1, 1
		[]int{0, 1, 1, 1}, // 0, 0
		[]int{1, 0, 0, 0}, // 1, 0
		[]int{1, 0, 0, 1}, // 1, 1
		[]int{1, 0, 1, 0}, // 0, 0
		[]int{1, 0, 1, 1}, // 0, 1
		[]int{1, 1, 0, 0}, // 1, 1
		[]int{1, 1, 0, 1}, // 0, 0
		[]int{1, 1, 1, 0}, // 0, 1
		[]int{1, 1, 1, 1}, // 1, 0
	}
	outputs := []interface{}{
		[]int{0, 0},
		[]int{0, 1},
		[]int{1, 0},
		[]int{1, 1},
		[]int{0, 1},
		[]int{1, 0},
		[]int{1, 1},
		[]int{0, 0},
		[]int{1, 0},
		[]int{1, 1},
		[]int{0, 0},
		[]int{0, 1},
		[]int{1, 1},
		[]int{0, 0},
		[]int{0, 1},
		[]int{1, 0},
	}
	// individualFactory := ga.IndividualFactory(func() ga.Individual {
	// 	return circuits.RandomCircuitIndividual(4, 2)
	// })
	// algo := ga.NewAlgorithm(ga.AlgorithmConfig{
	// 	IndividualFactory:   individualFactory,
	// 	PopulationSize:      10,
	// 	Epochs:              3,
	// 	GenerationsPerEpoch: 1000,
	// })
	// algo.Execute(inputs, outputs)
	i := circuits.RandomCircuitIndividual(4, 2)
	i.CalculateFitness(inputs, outputs)
	fmt.Printf("%f\n", i.Fitness())
	ii := circuits.RandomCircuitIndividual(4, 2)
	ii.CalculateFitness(inputs, outputs)
	fmt.Printf("%f\n", ii.Fitness())
	iii := i.Mate(ii)
	iii.CalculateFitness(inputs, outputs)
	fmt.Printf("%f\n", iii.Fitness())
}
