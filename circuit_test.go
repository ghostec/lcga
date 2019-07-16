package ga_test

import (
	"testing"

	ga "github.com/ghostec/lcga"
)

func TestCircuitAndOperator(t *testing.T) {
	c := ga.NewCircuit()
	a := ga.NewBit(1)
	b := ga.NewBit(1)
	o := ga.NewAndOperator()
	c.AddNode(a)
	c.AddNode(b)
	c.AddNode(o)
	c.AddEdge(a, o)
	c.AddEdge(b, o)
	v := o.Value().(int)
	if v != 1 {
		t.Errorf("Expected v to be 1. Got: %d", v)
	}
}

func TestCircuitExecute(t *testing.T) {
	type testCase struct {
		inputs   []int
		expected int
	}
	testCases := []testCase{
		testCase{[]int{0, 0}, 0},
		testCase{[]int{0, 1}, 0},
		testCase{[]int{1, 0}, 0},
		testCase{[]int{1, 1}, 1},
	}
	for _, tt := range testCases {
		c := ga.NewCircuit()
		a := ga.NewBit(1)
		b := ga.NewBit(1)
		ao := ga.NewAndOperator()
		ot := ga.NewBit(1)
		c.AddInput(a)
		c.AddInput(b)
		c.AddOperator(ao)
		c.AddOutput(ot)
		c.AddEdge(a, ao)
		c.AddEdge(b, ao)
		c.AddEdge(ao, ot)
		outputs, err := c.Execute(tt.inputs)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if len(outputs) != 1 {
			t.Fatalf("Expected outputs length to be 1. Got: %d", len(outputs))
		}
		if outputs[0] != tt.expected {
			t.Fatalf("Expected outputs to be [0]. Got: %#v", outputs)
		}
	}
}
