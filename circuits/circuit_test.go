package circuits_test

import (
	"testing"

	"github.com/ghostec/lcga/circuits"
)

func TestCircuitAndOperator(t *testing.T) {
	c := circuits.NewCircuit()
	a := circuits.NewBit(1)
	b := circuits.NewBit(1)
	o := circuits.NewAndOperator()
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

func TestCircuitExecuteWithAnd(t *testing.T) {
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
		c := circuits.NewCircuit()
		a := circuits.NewBit(1)
		b := circuits.NewBit(1)
		ao := circuits.NewAndOperator()
		ot := circuits.NewBit(1)
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

func TestCircuitExecuteWithOr(t *testing.T) {
	type testCase struct {
		inputs   []int
		expected int
	}
	testCases := []testCase{
		testCase{[]int{0, 0}, 0},
		testCase{[]int{0, 1}, 1},
		testCase{[]int{1, 0}, 1},
		testCase{[]int{1, 1}, 1},
	}
	for _, tt := range testCases {
		c := circuits.NewCircuit()
		a := circuits.NewBit(1)
		b := circuits.NewBit(1)
		ao := circuits.NewOrOperator()
		ot := circuits.NewBit(1)
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
