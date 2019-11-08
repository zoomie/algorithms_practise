package main

import "fmt"

type Node struct {
	key   int
	value string
	left  *Node
	right *Node
}
type BST struct {
	root *Node
}

func (b *BST) addData(key int, value string) {
	newNode := &Node{key: key, value: value}
	if b.root == nil {
		b.root = newNode
	} else {
		addNode(b.root, newNode)
	}
}

func (b *BST) getData(key int) string {
	if b.root == nil {
		return "none"
	} else {
		return getNodeValue(b.root, key)
	}
}
func getNodeValue(node *Node, key int) string {
	if node.key == key {
		return node.value
	} else if node.key > key {
		if node.right == nil {
			return "No value"
		} else {
			return getNodeValue(node.right, key)
		}
	} else {
		if node.left == nil {
			return "No value"
		} else {
			return getNodeValue(node.left, key)
		}
	}
}

func addNode(parentNode, newNode *Node) {
	if parentNode.key > newNode.key {
		if parentNode.right == nil {
			parentNode.right = newNode
		} else {
			addNode(parentNode.right, newNode)
		}
	} else if parentNode.key < newNode.key {
		if parentNode.left == nil {
			parentNode.left = newNode
		} else {
			addNode(parentNode.right, newNode)
		}
	} else {
		parentNode = newNode
	}
}

func main() {
	b := BST{}
	b.addData(1, "value")
	b.addData(3, "Test input string")
	fmt.Println(b.getData(4))
}
