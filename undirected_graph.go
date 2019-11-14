package main

import "fmt"

type node struct {
	value int
}
type graph struct {
	nodes       []*node
	connections map[node][]*node
}

func (g *graph) addNode(n *node) {
	g.nodes = append(g.nodes, n)
}

func (g *graph) connect(n1, n2 *node) {
	if g.connections == nil {
		g.connections = make(map[node][]*node)
	}
	g.connections[*n1] = append(g.connections[*n1], n2)
	g.connections[*n2] = append(g.connections[*n2], n1)
}

func (g *graph) isConnected(n1, n2 *node) bool {
	visited := make(map[int]bool)
	findValue := n2.value
	visit(g.connections, visited)
}

func visit(conns []*node, visited map[int]bool) {
	for i := 0; i < len(conns); i++ {
		visited[conns[i].value] = true
		_, ok := g.connections[conns[i]]
		if ok {
			visit()
		}
	}
}

func main() {
	g := graph{}
	n1 := &node{value: 1}
	n2 := &node{value: 2}
	g.addNode(n1)
	g.addNode(n2)
	g.connect(n1, n2)
	fmt.Println(g.isConnected(n1, n1))
}
