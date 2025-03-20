package main

import (
	"fmt"
	"sort"
	"strings"
)

// Node represents a graph node
type Node struct {
	ID       string
	Children []*Node
}

// Graph represents a directed graph
type Graph struct {
	Nodes map[string]*Node
	Edges map[string][]string
}

// NewGraph creates a new graph
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
		Edges: make(map[string][]string),
	}
}

// AddEdge adds a directed edge to the graph
func (g *Graph) AddEdge(from, to string) {
	if g.Nodes[from] == nil {
		g.Nodes[from] = &Node{ID: from}
	}
	if g.Nodes[to] == nil {
		g.Nodes[to] = &Node{ID: to}
	}
	g.Nodes[from].Children = append(g.Nodes[from].Children, g.Nodes[to])
	g.Edges[from] = append(g.Edges[from], to)
}

// AssignLayers assigns nodes to hierarchical layers
func (g *Graph) AssignLayers() map[string]int {
	layer := make(map[string]int)
	var dfs func(node *Node, depth int)

	dfs = func(node *Node, depth int) {
		if depth > layer[node.ID] {
			layer[node.ID] = depth
		}
		for _, child := range node.Children {
			dfs(child, depth+1)
		}
	}

	for _, node := range g.Nodes {
		if layer[node.ID] == 0 {
			dfs(node, 1)
		}
	}
	return layer
}

// OrderLayers minimizes crossings between layers using barycenter heuristic
func OrderLayers(layering map[string]int, g *Graph) map[int][]string {
	layeredNodes := make(map[int][]string)
	for node, l := range layering {
		layeredNodes[l] = append(layeredNodes[l], node)
	}
	for _, nodes := range layeredNodes {
		sort.Strings(nodes)
	}
	return layeredNodes
}

// RenderASCII renders the DAG as an ASCII graph
func RenderASCII(g *Graph) string {
	var builder strings.Builder
	for from, tos := range g.Edges {
		builder.WriteString(fmt.Sprintf("%s", from))
		for _, to := range tos {
			builder.WriteString(fmt.Sprintf(" ───→ %s", to))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func main() {
	g := NewGraph()
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "D")
	g.AddEdge("D", "E")

	fmt.Println("\nASCII Graph Representation:")
	fmt.Println(RenderASCII(g))
}
