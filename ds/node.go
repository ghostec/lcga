package ds

import "github.com/google/uuid"

type Node interface {
	Degree() int
	Graph() Graph
	Key() string
	Value() interface{}
	SetValue(interface{})
	SetGraph(Graph)
	Clone() Node
}

type SimpleNode struct {
	graph Graph
	key   string
	value interface{}
}

func NewSimpleNode(value interface{}) *SimpleNode {
	return &SimpleNode{
		key:   uuid.New().String(),
		value: value,
	}
}

func NewSimpleNodeWithKey(key string, value interface{}) *SimpleNode {
	return &SimpleNode{
		key:   key,
		value: value,
	}
}

func (c SimpleNode) Clone() Node {
	return &SimpleNode{
		key:   c.key,
		value: c.value,
	}
}

func (c SimpleNode) Key() string {
	return c.key
}

func (c SimpleNode) Value() interface{} {
	return c.value
}

func (c *SimpleNode) SetValue(value interface{}) {
	c.value = value
}

func (c *SimpleNode) Degree() int {
	return c.graph.Degree(c)
}

func (c SimpleNode) Graph() Graph {
	return c.graph
}

func (c *SimpleNode) SetGraph(g Graph) {
	c.graph = g
}
