/*
 * Copyright (c) 2023 Marco Massenzio. All rights reserved.
 */
package graph_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/massenz/algos/graph"
)

var _ = Describe("Graph", func() {
	var (
		g *graph.Graph
	)

	BeforeEach(func() {
		g = &graph.Graph{
			Vertices: map[string]*graph.Node{},
			Type:     graph.Directed,
		}
	})

	Describe("DFS", func() {
		It("traverses the graph using Depth-First Search", func() {
			g.Vertices["node1"] = &graph.Node{Name: "node1", Children: map[string]interface{}{"node2": nil}}
			g.Vertices["node2"] = &graph.Node{Name: "node2", Children: map[string]interface{}{}}
			var visited []string
			g.DFS("node1", func(n *graph.Node) {
				visited = append(visited, n.Name)
			})
			Expect(visited).To(Equal([]string{"node2", "node1"}))
		})
	})

	Describe("BFS", func() {
		It("traverses the graph using Breadth-First Search", func() {
			g, _, err := graph.LoadGraphs("../testdata/animals.yaml")
			Expect(err).NotTo(HaveOccurred())
			Expect(g).To(HaveLen(1))
			var visited []string
			g[0].BFS("animals", func(n *graph.Node) {
				visited = append(visited, n.Name)
			})
			Expect(visited).To(HaveLen(14))
			// The actual order is not guaranteed, but we know that the root node is visited first
			// and the last two nodes are the leaf nodes.
			Expect(visited[0]).To(Equal("animals"))
			Expect(visited[10:]).To(ContainElements("spaniels", "poodles", "labradors"))
		})
	})

	Describe("HasCycles", func() {
		It("returns true if the undirected graph has cycles", func() {
			g, _, _ := graph.LoadGraphs("../testdata/cyclic.yaml")
			Expect(g).To(HaveLen(1))
			g[0].DFS("foo", func(n *graph.Node) {})
			Expect(g[0].HasCycles()).To(BeTrue())
		})
		It("returns true if the directed graph has cycles", func() {
			g, _, _ := graph.LoadGraphs("../testdata/delivery.yaml")
			Expect(g).To(HaveLen(1))
			g[0].DFS("sent", func(n *graph.Node) {})
			Expect(g[0].HasCycles()).To(BeTrue())
		})
		It("returns false if the graph does not have cycles", func() {
			g, roots, _ := graph.LoadGraphs("../testdata/animals.yaml")
			Expect(g).To(HaveLen(1))
			g[0].DFS(roots[0], func(n *graph.Node) {})
			Expect(g[0].HasCycles()).To(BeFalse())
		})
	})

	Describe("FromAdjacencyList", func() {
		It("loads a graph from an adjacency list", func() {
			adjList := graph.AdjacencyList{
				Vertices: map[string][]string{"node1": {"node2"}},
				Kind:     graph.Directed,
			}
			g, err := graph.FromAdjacencyList(adjList)
			Expect(err).NotTo(HaveOccurred())
			Expect(g.Vertices["node1"].Children).To(HaveKey("node2"))
		})
	})

	Describe("LoadGraphs", func() {
		It("loads multiple graphs from a YAML file containing multiple adjacency lists", func() {
			graphs, roots, err := graph.LoadGraphs("../testdata/multidocs.yaml")
			Expect(err).NotTo(HaveOccurred())
			Expect(graphs).To(HaveLen(2))
			Expect(roots).To(Equal([]string{"main", ""}))
		})
	})
})
