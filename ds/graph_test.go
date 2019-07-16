package ds_test

import (
	"testing"

	"github.com/ghostec/lcga/ds"
)

func TestSimpleGraphWithEdge(t *testing.T) {
	g := ds.NewSimpleGraph()
	a := g.NewNode(nil)
	b := g.NewNode(nil)
	g.AddEdge(a, b)
	if a.Degree() != 0 {
		t.Errorf("Expected a to have degree = 0. Got: %d", a.Degree())
	}
	if b.Degree() != 1 {
		t.Errorf("Expected b to have degree = 1. Got: %d", b.Degree())
	}
	aie := g.IncomingEdges(a)
	if len(aie) != 0 {
		t.Fatalf("Expected a to have 0 incomings edges. Got: %d", len(aie))
	}
	bie := g.IncomingEdges(b)
	if len(bie) != 1 {
		t.Fatalf("Expected a to have 1 incoming edges. Got: %d", len(bie))
	}
	if bie[0] != a {
		t.Error("Expected b incoming edge to be from a")
	}
	aoe := g.OutgoingEdges(a)
	if len(aoe) != 1 {
		t.Fatalf("Expected a to have 0 outgoing edges. Got: %d", len(aoe))
	}
	if aoe[0] != b {
		t.Error("Expected a outgoing edge to be to b")
	}
	boe := g.OutgoingEdges(b)
	if len(boe) != 0 {
		t.Fatalf("Expected a to have 0 outgoing edges. Got: %d", len(boe))
	}
}

func TestTopsortStraightLine(t *testing.T) {
	g := ds.NewSimpleGraph()
	a := g.NewNodeWithKey("a", nil)
	b := g.NewNodeWithKey("b", nil)
	c := g.NewNodeWithKey("c", nil)
	d := g.NewNodeWithKey("d", nil)
	e := g.NewNodeWithKey("e", nil)
	g.AddEdge(a, b)
	g.AddEdge(b, c)
	g.AddEdge(c, d)
	g.AddEdge(d, e)
	topsort, err := ds.TopSort(g)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(topsort) != 5 {
		t.Fatalf("Expected topsort to have 5 nodes. Has: %d", len(topsort))
	}
	expected := []ds.Node{a, b, c, d, e}
	for i := range expected {
		if expected[i] != topsort[i] {
			got := make([]string, 0, len(topsort))
			for _, n := range topsort {
				got = append(got, n.Key())
			}
			t.Fatalf("Expected topsort to be [a, b, c, d, e]. Got: %#v", got)
		}
	}
}

func TestTopsortMoreEdges(t *testing.T) {
	g := ds.NewSimpleGraph()
	a := ds.NewSimpleNodeWithKey("a", nil)
	b := ds.NewSimpleNodeWithKey("b", nil)
	c := ds.NewSimpleNodeWithKey("c", nil)
	d := ds.NewSimpleNodeWithKey("d", nil)
	e := ds.NewSimpleNodeWithKey("e", nil)
	g.AddNode(a)
	g.AddNode(b)
	g.AddNode(c)
	g.AddNode(d)
	g.AddNode(e)
	g.AddEdge(a, c)
	g.AddEdge(a, e)
	g.AddEdge(b, a)
	g.AddEdge(b, c)
	g.AddEdge(b, d)
	g.AddEdge(c, e)
	g.AddEdge(d, a)
	g.AddEdge(d, c)
	topsort, err := ds.TopSort(g)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(topsort) != 5 {
		t.Fatalf("Expected topsort to have 5 nodes. Has: %d", len(topsort))
	}
	expected := []ds.Node{b, d, a, c, e}
	for i := range expected {
		if expected[i] != topsort[i] {
			got := make([]string, 0, len(topsort))
			for _, n := range topsort {
				got = append(got, n.Key())
			}
			t.Fatalf("Expected topsort to be [a, b, c, d, e]. Got: %#v", got)
		}
	}
}
