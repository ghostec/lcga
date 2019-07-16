package circuits

import (
	"fmt"

	"github.com/ghostec/lcga/ds"
)

type Bit struct {
	ds.Node
}

func NewBit(value int) *Bit {
	return &Bit{Node: ds.NewCommonNode(value)}
}

func (b Bit) Value() interface{} {
	es := b.Graph().IncomingEdges(b)
	if len(es) == 0 {
		// input bit
		return b.Node.Value()
	}
	// output bit
	return es[0].Value()
}

func (b Bit) Clone() ds.Node {
	return &Bit{Node: b.Node.Clone()}
}

type Circuit struct {
	ds.Graph
	inputs  []*Bit
	outputs []*Bit
}

func NewCircuit() *Circuit {
	return &Circuit{
		Graph:   ds.NewCommonGraph(),
		inputs:  []*Bit{},
		outputs: []*Bit{},
	}
}

func (c *Circuit) AddInput(input *Bit) {
	c.inputs = append(c.inputs, input)
	c.AddNode(input)
}

func (c *Circuit) AddOutput(output *Bit) {
	c.outputs = append(c.outputs, output)
	c.AddNode(output)
}

func (c *Circuit) AddOperator(operator ds.Node) {
	c.AddNode(operator)
}

func (c Circuit) Execute(inputs []int) ([]int, error) {
	if len(inputs) != len(c.inputs) {
		return nil, fmt.Errorf("inputs should have length %d", len(c.inputs))
	}
	for i := range inputs {
		c.inputs[i].SetValue(inputs[i])
	}
	topsort, err := ds.TopSort(c)
	if err != nil {
		return nil, err
	}
	values := make([]int, 0, len(topsort))
	for _, node := range topsort {
		values = append(values, node.Value().(int))
	}
	return values[len(values)-len(c.outputs):], nil
}

type AndOperator struct {
	ds.Node
}

func NewAndOperator() *AndOperator {
	return &AndOperator{Node: ds.NewCommonNode(nil)}
}

func (a AndOperator) Clone() ds.Node {
	return &AndOperator{Node: a.Node.Clone()}
}

func (a AndOperator) Value() interface{} {
	inputs := a.Graph().IncomingEdges(a)
	return inputs[0].Value().(int) & inputs[1].Value().(int)
}

type OrOperator struct {
	ds.Node
}

func NewOrOperator() *OrOperator {
	return &OrOperator{Node: ds.NewCommonNode(nil)}
}

func (o OrOperator) Clone() ds.Node {
	return &OrOperator{Node: o.Node.Clone()}
}

func (o OrOperator) Value() interface{} {
	inputs := o.Graph().IncomingEdges(o)
	return inputs[0].Value().(int) | inputs[1].Value().(int)
}

// type CircuitIndividual struct {
// 	circuit *Circuit
// 	fitness float64
// }
//
// func NewCircuitIndividual() *CircuitIndividual {
// 	return &CircuitIndividual{circuit: NewCircuit()}
// }
//
// func (c CircuitIndividual) Execute(input interface{}) interface{} {
// 	output, _ := c.circuit.Execute(input.([]int))
// 	return output
// }
//
// func (c CircuitIndividual) Fitness() float64 {
// 	return c.fitness
// }
//
// func (c *CircuitIndividual) CalculateFitness(inputs, outputs []interface{}) {
// 	f := float64(0)
// 	for i := range inputs {
// 		correct := float64(0)
// 		output := c.Execute(inputs[i]).([]int)
// 		expected := outputs[i].([]int)
// 		for j := range output {
// 			if output[j] == expected[j] {
// 				correct += 1
// 			}
// 		}
// 		f += correct / float64(len(output))
// 	}
// 	c.fitness = f / float64(len(outputs))
// }
