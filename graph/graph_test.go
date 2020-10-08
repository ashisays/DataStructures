package graph

import (
    "fmt"
    "testing"
)

func TestGraph(t *testing.T) {
    g := NewGraph()

    nA := Node{value:"A"}
    nB := Node{value:"B"}
    nC := Node{value:"C"}
    nD := Node{value:"D"}
    nE := Node{value:"E"}
    nF := Node{value:"F"}
    g.AddNode(&nA)
    g.AddNode(&nB)
    g.AddNode(&nC)
    g.AddNode(&nD)
    g.AddNode(&nE)
    g.AddNode(&nF)

    g.AddEdge(&nA, &nB)
    g.AddEdge(&nA, &nC)
    g.AddEdge(&nB, &nE)
    g.AddEdge(&nC, &nE)
    g.AddEdge(&nE, &nF)
    g.AddEdge(&nD, &nA)

    if g.Display() == "A--> B; C; D;\nB--> A; E;\nC--> A; E;\nD--> A;\nE--> B; C; F;\nF--> E;\n" {
      fmt.Println(g.Display())
    }
}