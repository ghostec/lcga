package ga

import "fmt"

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
}

type Edge struct {
	From Node
	To   Node
}

type CommonGraph struct {
	incomingEdges map[string]*Set
	nodes         map[string]Node
	outgoingEdges map[string]*Set
}

func NewCommonGraph() *CommonGraph {
	return &CommonGraph{
		incomingEdges: map[string]*Set{},
		nodes:         map[string]Node{},
		outgoingEdges: map[string]*Set{},
	}
}

func (g CommonGraph) Clone() Graph {
	gg := NewCommonGraph()
	for k, v := range g.nodes {
		vv := v.Clone()
		vv.SetGraph(gg)
		gg.nodes[k] = vv
	}
	for k, v := range g.incomingEdges {
		gg.incomingEdges[k] = v.Clone()
	}
	for k, v := range g.outgoingEdges {
		gg.outgoingEdges[k] = v.Clone()
	}
	return gg
}

func (g *CommonGraph) NewNode(value interface{}) Node {
	node := NewCommonNode(value)
	g.AddNode(node)
	return node
}

func (g *CommonGraph) NewNodeWithKey(key string, value interface{}) Node {
	node := NewCommonNode(value)
	node.key = key
	g.AddNode(node)
	return node
}

func (g *CommonGraph) AddNode(node Node) {
	node.SetGraph(g)
	g.incomingEdges[node.Key()] = NewSet()
	g.nodes[node.Key()] = node
	g.outgoingEdges[node.Key()] = NewSet()
}

func (g *CommonGraph) AddEdge(from, to Node) {
	g.incomingEdges[to.Key()].Add(from.Key())
	g.outgoingEdges[from.Key()].Add(to.Key())
}

func (g *CommonGraph) RemoveEdge(from, to Node) {
	g.incomingEdges[to.Key()].Remove(from.Key())
	g.outgoingEdges[from.Key()].Remove(to.Key())
}

func (g CommonGraph) Degree(node Node) int {
	return g.incomingEdges[node.Key()].Size()
}

func (g CommonGraph) IncomingEdges(node Node) []Node {
	in := g.incomingEdges[node.Key()].Slice()
	return g.getNodesByKeys(toStringSlice(in))
}

func (g CommonGraph) OutgoingEdges(node Node) []Node {
	out := g.outgoingEdges[node.Key()].Slice()
	return g.getNodesByKeys(toStringSlice(out))
}

func (g *CommonGraph) getNodesByKeys(keys []string) []Node {
	ns := make([]Node, 0, len(keys))
	for _, key := range keys {
		ns = append(ns, g.nodes[key])
	}
	return ns
}

func (g CommonGraph) Nodes() []Node {
	list := make([]Node, 0, len(g.nodes))
	for _, node := range g.nodes {
		list = append(list, node)
	}
	return list
}

func (g CommonGraph) Node(key string) Node {
	return g.nodes[key]
}

func (g CommonGraph) Edges() []Edge {
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
