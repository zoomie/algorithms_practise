package main

import (
	"fmt"
	"sync"
)

type Node struct {
	value int
}

type ItemGraph struct {
	nodes []*Node
	edges map[Node]*Node
	lock  sync.RWMutex
}

func (i *ItemGraph) addNode(j int) {
	n := &Node{value: j}
	i.nodes = append(i.nodes, n)
}

func (i *ItemGraph) addEdge(a, j int) {
	if i.edges == nil {
		i.edges = make(map[Node]*Node)
	}
	nodeI := &Node{value: a}
	nodeJ := &Node{value: j}
	i.edges[*nodeI] = nodeJ
}

func (i *ItemGraph) display() {
	for j := 0; j < len(i.nodes); j++ {
		fmt.Println(i.nodes[j])
	}
}

func main() {
	var graph ItemGraph
	graph.addNode(10)
	graph.addNode(1)
	graph.addEdge(1, 10)
	graph.display()
}
