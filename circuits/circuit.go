package circuits

import (
	"fmt"

	"github.com/ghostec/lcga/ds"
)

type Bit struct {
	ds.Node
}

func NewBit(value int) *Bit {
	return &Bit{Node: ds.NewSimpleNode(value)}
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
	inputs    []*Bit
	operators map[string]Operator
	outputs   []*Bit
}

func NewCircuit() *Circuit {
	return &Circuit{
		Graph:     ds.NewSimpleGraph(),
		inputs:    []*Bit{},
		operators: map[string]Operator{},
		outputs:   []*Bit{},
	}
}

func (c Circuit) Clone() ds.Graph {
	cc := &Circuit{
		Graph:     ds.NewSimpleGraph(),
		inputs:    make([]*Bit, 0, len(c.inputs)),
		operators: map[string]Operator{},
		outputs:   make([]*Bit, 0, len(c.outputs)),
	}
	for _, x := range c.inputs {
		cc.AddInput(x.Clone().(*Bit))
	}
	for _, x := range c.operators {
		cc.AddOperator(x.Clone().(Operator))
	}
	for _, x := range c.outputs {
		cc.AddOutput(x.Clone().(*Bit))
	}
	for _, e := range c.Edges() {
		cc.Graph.AddEdge(e.From, e.To)
	}
	return cc
}

func (c *Circuit) RemoveNode(node ds.Node) {
	idx := -1
	for i := range c.inputs {
		if c.inputs[i] == node {
			idx = i
			break
		}
	}
	if idx != -1 {
		c.inputs = append(c.inputs[:idx], c.inputs[idx+1:]...)
	}
	idx = -1
	for i := range c.outputs {
		if c.outputs[i] == node {
			idx = i
			break
		}
	}
	if idx != -1 {
		c.outputs = append(c.outputs[:idx], c.outputs[idx+1:]...)
	}
	delete(c.operators, node.Key())
	c.Graph.RemoveNode(node)
}

func (c Circuit) NumInputs() int {
	return len(c.inputs)
}

func (c Circuit) NumOperators() int {
	return len(c.operators)
}

func (c Circuit) NumOutputs() int {
	return len(c.outputs)
}

func (c *Circuit) AddInput(input *Bit) {
	c.inputs = append(c.inputs, input)
	c.AddNode(input)
}

func (c *Circuit) AddOutput(output *Bit) {
	c.outputs = append(c.outputs, output)
	c.AddNode(output)
}

func (c *Circuit) AddOperator(operator Operator) {
	c.operators[operator.Key()] = operator
	c.AddNode(operator)
}

func (c *Circuit) TopSort() ([]ds.Node, error) {
	topsort, err := ds.TopSort(c)
	if err != nil {
		return nil, err
	}
	// inputs are placed at the start due to ds.TopSort's implementation
	// placing outputs at the end of topsort in the same order they were added
	// to c
	oMap := map[string]*Bit{}
	for _, o := range c.outputs {
		oMap[o.Key()] = o
	}
	toRemove := make([]int, 0, len(c.outputs))
	for i, node := range topsort {
		if _, ok := oMap[node.Key()]; ok {
			toRemove = append(toRemove, i)
		}
	}
	for i, idx := range toRemove {
		topsort = append(topsort[:idx-i], topsort[idx-i+1:]...)
	}
	for _, o := range c.outputs {
		topsort = append(topsort, o)
	}
	return topsort, nil
}

func (c *Circuit) Execute(inputs []int) ([]int, error) {
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

type Operator interface {
	ds.Node
	RequiredInputs() int
}

type AndOperator struct {
	ds.Node
}

func NewAndOperator() Operator {
	return &AndOperator{Node: ds.NewSimpleNode(nil)}
}

func (a AndOperator) Clone() ds.Node {
	return &AndOperator{Node: a.Node.Clone()}
}

func (a AndOperator) Value() interface{} {
	inputs := a.Graph().IncomingEdges(a)
	return inputs[0].Value().(int) & inputs[1].Value().(int)
}

func (a AndOperator) RequiredInputs() int {
	return 2
}

type OrOperator struct {
	ds.Node
}

func NewOrOperator() Operator {
	return &OrOperator{Node: ds.NewSimpleNode(nil)}
}

func (o OrOperator) Clone() ds.Node {
	return &OrOperator{Node: o.Node.Clone()}
}

func (o OrOperator) Value() interface{} {
	inputs := o.Graph().IncomingEdges(o)
	return inputs[0].Value().(int) | inputs[1].Value().(int)
}

func (o OrOperator) RequiredInputs() int {
	return 2
}

type OperatorFactory func() Operator

func OperatorsFactories() []OperatorFactory {
	return []OperatorFactory{NewAndOperator, NewOrOperator}
}
