package circuits

import (
	"math/rand"

	"github.com/ghostec/lcga/ds"
)

type CircuitIndividual struct {
	circuit *Circuit
	fitness float64
}

func RandomCircuitIndividual(numInputs, numOutputs int) *CircuitIndividual {
	ci := NewCircuitIndividual()
	factories := OperatorsFactories()
	validInputs := make([]ds.Node, 0)
	outputs := make([]ds.Node, 0)
	for i := 0; i < numInputs; i++ {
		input := NewBit(0)
		validInputs = append(validInputs, input)
		ci.circuit.AddInput(input)
	}
	for i := 0; i < numOutputs; i++ {
		output := NewBit(0)
		outputs = append(outputs, output)
		ci.circuit.AddOutput(output)
	}
	for rand.Float64() < 0.45 {
		factory := factories[rand.Intn(len(factories))]
		op := factory()
		ci.circuit.AddOperator(op)
		numValidInputs := len(validInputs)
		for op.Degree() < op.RequiredInputs() {
			j := rand.Intn(numValidInputs)
			ci.circuit.AddEdge(validInputs[j], op)
		}
		validInputs = append(validInputs, op)
	}
	numValidInputs := len(validInputs)
	for i := 0; i < numOutputs; i++ {
		j := rand.Intn(numValidInputs)
		ci.circuit.AddEdge(validInputs[j], outputs[i])
	}
	// validInputs is already topsort
	return ci
}

func NewCircuitIndividual() *CircuitIndividual {
	return &CircuitIndividual{circuit: NewCircuit()}
}

func (c CircuitIndividual) Execute(input interface{}) interface{} {
	output, _ := c.circuit.Execute(input.([]int))
	return output
}

func (c CircuitIndividual) Fitness() float64 {
	return c.fitness
}

func (c *CircuitIndividual) CalculateFitness(inputs, outputs []interface{}) {
	f := float64(0)
	for i := range inputs {
		correct := float64(0)
		output := c.Execute(inputs[i]).([]int)
		expected := outputs[i].([]int)
		for j := range output {
			if output[j] == expected[j] {
				correct += 1
			}
		}
		f += correct / float64(len(output))
	}
	c.fitness = f / float64(len(outputs))
}
