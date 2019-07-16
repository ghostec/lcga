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
	i := circuits.RandomCircuitIndividual(4, 2)
	i.CalculateFitness(inputs, outputs)
	fmt.Printf("%f\n", i.Fitness())
}
