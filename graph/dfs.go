package graph

import (
  "fmt"
  "strings"
)
var marked map[string] bool = make(map[string]bool)
var edgeto map [string] string = make(map[string]string)

func hasPathTo(v string) bool {
  return marked[v]
}

func pathTo(v string,source string) string {
  var path []string
  if !hasPathTo(v) {
    return ""
  }

  for x := v; x != source; x =edgeto[x]{
    path = append(path,x)
  }
  return strings.Join(path, ",")
}

func dfs(g *Graph, n *Node)  {
  v := n.value
  marked[v] = true
  for _,w := range g.edges[*n] {
    if _,found := marked[w.value]; !found{
      dfs(g,w)
      edgeto[w.value] = v
    fmt.Printf("Traversing : %v\n",w.value)
    }
  } 
}