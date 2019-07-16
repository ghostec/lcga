package ga

import "fmt"

type Bit struct {
	Node
}

func NewBit(value int) *Bit {
	return &Bit{Node: NewCommonNode(value)}
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

func (b Bit) Clone() Node {
	return &Bit{Node: b.Node.Clone()}
}

type Circuit struct {
	*CommonGraph
	inputs  []*Bit
	outputs []*Bit
}

func NewCircuit() *Circuit {
	return &Circuit{
		CommonGraph: NewCommonGraph(),
		inputs:      []*Bit{},
		outputs:     []*Bit{},
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

func (c *Circuit) AddOperator(operator Node) {
	c.AddNode(operator)
}

func (c Circuit) Execute(inputs []int) ([]int, error) {
	if len(inputs) != len(c.inputs) {
		return nil, fmt.Errorf("inputs should have length %d", len(c.inputs))
	}
	for i := range inputs {
		c.inputs[i].SetValue(inputs[i])
	}
	topsort, err := TopSort(c)
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
	Node
}

func NewAndOperator() *AndOperator {
	return &AndOperator{Node: NewCommonNode(nil)}
}

func (a AndOperator) Clone() Node {
	return &AndOperator{Node: a.Node.Clone()}
}

func (a AndOperator) Value() interface{} {
	inputs := a.Graph().IncomingEdges(a)
	return inputs[0].Value().(int) & inputs[1].Value().(int)
}

type OrOperator struct {
	Node
}

func NewOrOperator() *OrOperator {
	return &OrOperator{Node: NewCommonNode(nil)}
}

func (o OrOperator) Clone() Node {
	return &OrOperator{Node: o.Node.Clone()}
}

func (o OrOperator) Value() interface{} {
	inputs := o.Graph().IncomingEdges(o)
	return inputs[0].Value().(int) | inputs[1].Value().(int)
}

type CircuitIndividual struct {
	*Circuit
	fitness float64
}

func NewCircuitIndividual() *CircuitIndividual {
	return &CircuitIndividual{Circuit: NewCircuit()}
}

func (c *CircuitIndividual) Execute(input interface{}) interface{} {
	return nil
}
