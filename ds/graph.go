package ds

import (
	"fmt"

	"github.com/ghostec/lcga/helpers"
)

type Graph interface {
	AddEdge(from, to Node)
	AddNode(Node)
	Clone() Graph
	Degree(Node) int
	Edges() []Edge
	IncomingEdges(Node) []Node
	NewNode(value interface{}) Node
	NewNodeWithKey(key string, value interface{}) Node
	Node(key string) Node
	Nodes() []Node
	OutgoingEdges(Node) []Node
	RemoveEdge(from, to Node)
	RemoveNode(Node)
}

type Edge struct {
	From Node
	To   Node
}

type SimpleGraph struct {
	incomingEdges map[string]*Set
	nodes         map[string]Node
	outgoingEdges map[string]*Set
}

func NewSimpleGraph() *SimpleGraph {
	return &SimpleGraph{
		incomingEdges: map[string]*Set{},
		nodes:         map[string]Node{},
		outgoingEdges: map[string]*Set{},
	}
}

func (g SimpleGraph) Clone() Graph {
	gg := NewSimpleGraph()
	for _, v := range g.nodes {
		gg.AddNode(v.Clone())
	}
	for _, e := range g.Edges() {
		gg.AddEdge(e.From, e.To)
	}
	return gg
}

func (g *SimpleGraph) NewNode(value interface{}) Node {
	node := NewSimpleNode(value)
	g.AddNode(node)
	return node
}

func (g *SimpleGraph) NewNodeWithKey(key string, value interface{}) Node {
	node := NewSimpleNode(value)
	node.key = key
	g.AddNode(node)
	return node
}

func (g *SimpleGraph) AddNode(node Node) {
	node.SetGraph(g)
	g.incomingEdges[node.Key()] = NewSet()
	g.nodes[node.Key()] = node
	g.outgoingEdges[node.Key()] = NewSet()
}

func (g *SimpleGraph) AddEdge(from, to Node) {
	g.incomingEdges[to.Key()].Add(from.Key())
	g.outgoingEdges[from.Key()].Add(to.Key())
}

func (g *SimpleGraph) RemoveEdge(from, to Node) {
	g.incomingEdges[to.Key()].Remove(from.Key())
	g.outgoingEdges[from.Key()].Remove(to.Key())
}

func (g *SimpleGraph) RemoveNode(node Node) {
	nodeKey := node.Key()
	delete(g.incomingEdges, nodeKey)
	delete(g.outgoingEdges, nodeKey)
	delete(g.nodes, nodeKey)
	for _, set := range g.outgoingEdges {
		set.Remove(nodeKey)
	}
	for _, set := range g.incomingEdges {
		set.Remove(nodeKey)
	}
}

func (g SimpleGraph) Degree(node Node) int {
	return g.incomingEdges[node.Key()].Size()
}

func (g SimpleGraph) IncomingEdges(node Node) []Node {
	in := g.incomingEdges[node.Key()].Slice()
	return g.getNodesByKeys(helpers.ToStringSlice(in))
}

func (g SimpleGraph) OutgoingEdges(node Node) []Node {
	out := g.outgoingEdges[node.Key()].Slice()
	return g.getNodesByKeys(helpers.ToStringSlice(out))
}

func (g *SimpleGraph) getNodesByKeys(keys []string) []Node {
	ns := make([]Node, 0, len(keys))
	for _, key := range keys {
		ns = append(ns, g.nodes[key])
	}
	return ns
}

func (g SimpleGraph) Nodes() []Node {
	list := make([]Node, 0, len(g.nodes))
	for _, node := range g.nodes {
		list = append(list, node)
	}
	return list
}

func (g SimpleGraph) Node(key string) Node {
	return g.nodes[key]
}

func (g SimpleGraph) Edges() []Edge {
	list := make([]Edge, 0)
	for fromKey, set := range g.outgoingEdges {
		toKeys := set.Slice()
		fromNode := g.nodes[fromKey]
		for _, toKey := range toKeys {
			list = append(list, Edge{From: fromNode, To: g.nodes[toKey.(string)]})
		}
	}
	return list
}

func TopSort(gg Graph) ([]Node, error) {
	g := gg.Clone()
	topsort := make([]Node, 0)
	noIncomingEdges := NewQueue()
	for _, node := range g.Nodes() {
		if node.Degree() == 0 {
			noIncomingEdges.Enqueue(node)
		}
	}
	for !noIncomingEdges.Empty() {
		n := noIncomingEdges.Dequeue().(Node)
		topsort = append(topsort, n)
		for _, m := range g.OutgoingEdges(n) {
			g.RemoveEdge(n, m)
			if m.Degree() == 0 {
				noIncomingEdges.Enqueue(m)
			}
		}
	}
	if len(g.Edges()) > 0 {
		return nil, fmt.Errorf("g is not a DAG")
	}
	for i := range topsort {
		topsort[i] = gg.Node(topsort[i].Key())
	}
	return topsort, nil
}
