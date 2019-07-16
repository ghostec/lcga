package ga

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

type CommonNode struct {
	graph Graph
	key   string
	value interface{}
}

func NewCommonNode(value interface{}) *CommonNode {
	return &CommonNode{
		key:   uuid.New().String(),
		value: value,
	}
}

func NewCommonNodeWithKey(key string, value interface{}) *CommonNode {
	return &CommonNode{
		key:   key,
		value: value,
	}
}

func (c CommonNode) Clone() Node {
	return &CommonNode{
		key:   c.key,
		value: c.value,
	}
}

func (c CommonNode) Key() string {
	return c.key
}

func (c CommonNode) Value() interface{} {
	return c.value
}

func (c *CommonNode) SetValue(value interface{}) {
	c.value = value
}

func (c *CommonNode) Degree() int {
	return c.graph.Degree(c)
}

func (c CommonNode) Graph() Graph {
	return c.graph
}

func (c *CommonNode) SetGraph(g Graph) {
	c.graph = g
}
