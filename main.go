/*
 * Copyright (c) 2023 Marco Massenzio. All rights reserved.
 */

package main

import (
	. "algorithms/graph"
	"flag"
	"fmt"
	"os"
)

func main() {
	nodes := flag.String("graph", "", "YAML file containing graph adjacency list")
	start := flag.String("start", "", "An optional start node for the graph traversal")
	flag.Parse()

	if *nodes == "" {
		fmt.Println("Must specify a YAML file containing a graph adjacency list")
		os.Exit(1)
	}
	gs, roots, err := LoadGraphs(*nodes)
	if err != nil {
		fmt.Printf("Error loading graphs: %v\n", err)
		os.Exit(1)
	}
	for n, g := range gs {
		fmt.Printf("Loaded graph:\n%s\n", g)
		startNode := roots[n]
		if startNode == "" {
			startNode = *start
		}
		fmt.Println("Traversing graph using DFS:")
		g.DFS(startNode, func(n *Node) {
			fmt.Println(n.Name)
		})
		var has string
		if g.HasCycles() {

			has = "has"
		} else {
			has = "does not have"
		}
		fmt.Printf("The graph %s cycles\n", has)
		fmt.Println("\n-------\nTraversing graph using BFS:")
		g.BFS(startNode, func(n *Node) {
			fmt.Println(n.Name)
		})
	}
}
