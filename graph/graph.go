package graph

import (
  "fmt"
)

type Node struct{
  value string
}

func (n *Node) String() string {
  return fmt.Sprintf("%v", n.value)
}

type Graph struct {
  vertices []*Node
  edges map[Node][]*Node
} 

func NewGraph() Graph{
  return Graph{}
}

func (g *Graph) AddNode(n *Node) {
  g.vertices = append(g.vertices,n)
}

func (g *Graph) AddEdge(n1,n2 *Node) {
  if g.edges == nil {
    g.edges = make(map[Node][]*Node)
  }
  g.edges[*n1] = append(g.edges[*n1],n2)
  g.edges[*n2] = append(g.edges[*n2],n1)
}

func (g *Graph) Display() string {
  s := ""
  for _ , v := range g.vertices {
    s += v.String() + "--> "
    for _,neighbor :=  range g.edges[*v]{
      s += neighbor.String() + "; "
    }
    s += "\n"
  }
  fmt.Println(s)
  return s
}

func Print(){
  fmt.Println("Graph Algorithms !!")
}
