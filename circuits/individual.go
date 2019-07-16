package circuits

import (
	"math/rand"

	"github.com/ghostec/lcga/ds"
	"github.com/ghostec/lcga/ga"
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
	for rand.Float64() < 0.8 {
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

func (c CircuitIndividual) Clone() ga.Individual {
	cc := &CircuitIndividual{circuit: c.circuit.Clone().(*Circuit), fitness: 0}
	return cc
}

func (c CircuitIndividual) Execute(input interface{}) interface{} {
	output, _ := c.circuit.Execute(input.([]int))
	return output
}

func (c CircuitIndividual) Fitness() float64 {
	return c.fitness
}

func (c CircuitIndividual) Circuit() *Circuit {
	return c.circuit
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

func (c CircuitIndividual) Mate(other ga.Individual) ga.Individual {
	a := c.Clone().(*CircuitIndividual)
	b := other.Clone().(*CircuitIndividual)
	if a.circuit.NumOperators() == 0 {
		return a
	}
	if b.circuit.NumOperators() == 0 {
		return b
	}
	numInputs := a.circuit.NumInputs()
	// keeping start of a
	aTopSort, _ := a.circuit.TopSort()
	aEnd := numInputs + rand.Intn(a.circuit.NumOperators())
	for i := aEnd; i < len(aTopSort); i++ {
		a.circuit.RemoveNode(aTopSort[i])
	}
	aTopSort = aTopSort[:aEnd]
	// keeping end of b
	// placing outputs at the end of bTopSort
	bTopSort, _ := b.circuit.TopSort()
	bStart := numInputs + rand.Intn(b.circuit.NumOperators())
	for i := 0; i < bStart; i++ {
		b.circuit.RemoveNode(bTopSort[i])
	}
	bTopSort = bTopSort[bStart:]
	// adding a nodes+edges to b
	for i := 0; i < numInputs; i++ {
		b.circuit.AddInput(aTopSort[i].(*Bit))
	}
	for i := numInputs; i < len(aTopSort); i++ {
		b.circuit.AddOperator(aTopSort[i].(Operator))
	}
	for _, edge := range a.circuit.Edges() {
		b.circuit.AddEdge(edge.From, edge.To)
	}
	// filling missing inputs in b elements with a+b
	randomNodeBeforeIndex := func(bi int) ds.Node {
		j := rand.Intn(len(aTopSort) + bi)
		if j < len(aTopSort) {
			return aTopSort[j]
		}
		return bTopSort[j-len(aTopSort)]
	}
	for bi, node := range bTopSort {
		// either Operator or *Bit (output)
		switch node.(type) {
		case Operator:
			for node.Degree() < node.(Operator).RequiredInputs() {
				from := randomNodeBeforeIndex(bi)
				b.circuit.AddEdge(from, node)
			}
		case *Bit:
			for node.Degree() == 0 {
				from := randomNodeBeforeIndex(bi)
				b.circuit.AddEdge(from, node)
			}
		}
	}
	return b
}
