/*
 * Copyright (c) 2023 Marco Massenzio. All rights reserved.
 */

package graph

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/massenz/algos/common"
)

type Node struct {
	Name string
	// We use this map as a Set of children names; the Nodes are kept in the Graph.Vertices map
	Children map[string]interface{}
}

func (n *Node) String() string {
	s := fmt.Sprintf("[%s] -> {", n.Name)
	if n.Children != nil {
		for child := range n.Children {
			s = s + child + " "
		}
	}
	s = s + "}"
	return s
}

type Kind string

const (
	Directed   Kind = "directed"
	Undirected Kind = "undirected"
)

type AdjacencyList struct {
	Vertices map[string][]string `yaml:"vertices"`
	Kind     Kind                `yaml:"kind"`
	Root     string              `yaml:"root"`
}

type Graph struct {
	Vertices map[string]*Node
	Size     int
	Type     Kind

	// Internal data structures used for DFS and BFS
	discovered map[string]bool
	visited    map[string]bool
	hasCycles  bool
}

func (g *Graph) String() string {
	s := fmt.Sprintf("Graph [Size: %d, Type: %s]\n", g.Size, g.Type)
	for _, node := range g.Vertices {
		s = s + node.String() + "\n"
	}
	return s
}

// HasCycles returns true if the graph has cycles, false otherwise
// It MUST be called after DFS or BFS, otherwise it will return false regardless
// of whether the graph has cycles or not.
func (g *Graph) HasCycles() bool {
	return g.hasCycles
}

// DFS traverses the graph using Depth-First Search
func (g *Graph) DFS(start string, visit func(*Node)) {
	g.discovered = make(map[string]bool)
	g.visited = make(map[string]bool)
	if _, found := g.Vertices[start]; !found {
		panic("Start node not found")
	}
	g.discovered[start] = true
	g.dfs(g.Vertices[start], visit)
}

// dfs is the recursive internal function for DFS
// it returns true if it finds a cycle
func (g *Graph) dfs(node *Node, visit func(*Node)) {
	for child := range node.Children {
		if g.discovered[child] {
			// when we discover a node that has already been discovered, in an undirected graph,
			// it means we have a cycle
			if g.Type == Undirected {
				g.hasCycles = true
				continue
			}
			// in directed graphs, nodes that have already been discovered,
			// but not been visited yet, indicate the existence of a cycle
			if g.Type == Directed && !g.visited[child] {
				g.hasCycles = true
				continue
			}
		}
		g.discovered[child] = true
		n := g.Vertices[child]
		g.dfs(n, visit)
	}
	visit(node)
}

// BFS traverses the graph using Breadth-First Search
func (g *Graph) BFS(start string, visit func(*Node)) {
	queue := common.NewQueue(g.Size)
	g.discovered = make(map[string]bool)
	g.visited = make(map[string]bool)
	if _, found := g.Vertices[start]; !found {
		panic("Start node not found")
	}
	g.discovered[start] = true
	err := queue.Enqueue(g.Vertices[start])
	if err != nil {
		panic(err)
	}
	for !queue.IsEmpty() {
		node := queue.Dequeue().(*Node)
		if node == nil {
			panic("Unexpected nil Node from a non-empty queue")
		}
		for child := range node.Children {
			if g.discovered[child] {
				continue
			}
			g.discovered[child] = true
			err := queue.Enqueue(g.Vertices[child])
			if err != nil {
				panic("Unexpected error enqueuing child node: " + err.Error())
			}
		}
		visit(node)
		g.visited[node.Name] = true
	}
}

// LoadGraphs loads multiple graphs from a YAML file containing multiple adjacency lists.
// If your YAML file contains multiple documents, you must separate them with `---`.
// It returns a slice of Graphs and a slice of root node names; some of the root names
// may be empty, in which case the caller must specify a start node for the traversal.
func LoadGraphs(filepath string) ([]*Graph, []string, error) {
	// Read the YAML file
	var graphs []*Graph
	var roots []string
	graphDesc, err := os.ReadFile(filepath)
	if err != nil {
		return nil, nil, err
	}
	docs := strings.Split(string(graphDesc), "---")
	for _, doc := range docs {
		var adjList AdjacencyList
		err = yaml.Unmarshal([]byte(doc), &adjList)
		if err != nil {
			return nil, nil, err
		}
		g, err := FromAdjacencyList(adjList)
		if err != nil {
			return nil, nil, err
		}
		graphs = append(graphs, g)
		roots = append(roots, adjList.Root)
	}
	return graphs, roots, nil
}

// FromAdjacencyList loads a graph from a YAML file containing an adjacency list.
func FromAdjacencyList(adjList AdjacencyList) (*Graph, error) {
	var g = Graph{Vertices: map[string]*Node{}, Type: adjList.Kind, Size: 0}

	for name, adjNodeNames := range adjList.Vertices {
		vertex, found := g.Vertices[name]
		if !found {
			vertex = &Node{Name: name, Children: make(map[string]interface{})}
			g.Vertices[name] = vertex
			g.Size = g.Size + 1
		}

		for _, adjNodeName := range adjNodeNames {
			adj, found := g.Vertices[adjNodeName]
			if !found {
				adj = &Node{Name: adjNodeName, Children: map[string]interface{}{}}
				g.Vertices[adjNodeName] = adj
				g.Size = g.Size + 1
			}
			// we don't care about the value, all we need to keep track of is the
			// presence of the key in the map (Go doesn't have a Set type)
			vertex.Children[adjNodeName] = nil
			if g.Type == Undirected {
				// if the graph is undirected, we need to add the edge in the other direction
				adj.Children[name] = nil
			}
		}
	}
	return &g, nil
}
