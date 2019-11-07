// package main

// import (
// 	"sync"
// )

// type Node struct {
// 	key   int
// 	value string
// 	left  *Node //left
// 	right *Node //right
// }

// type ItemBinarySearchTree struct {
// 	root *Node
// 	lock sync.RWMutex
// }

// func main() {
// 	var t ItemBinarySearchTree
// 	t.addKey(1, "string")
// 	t.addKey(2, "string")
// 	t.addKey(0, "string")
// 	t.addKey(3, "string")
// 	t.addKey(5, "string")
// 	t.addKey(7, "string")
// 	t.addKey(9, "string")
// }

// func (bst *ItemBinarySearchTree) addKey(key int, value string) {
// 	newNode := &Node{key, value, nil, nil}
// 	if bst.root == nil {
// 		bst.root = newNode
// 	} else {
// 		appendNode(bst.root, newNode)
// 	}
// }

// func appendNode(parent, child *Node) {
// 	if parent.key > child.key {
// 		if parent.left == nil {
// 			parent.left = child
// 		} else {
// 			appendNode(parent.left, child)
// 		}
// 	} else {
// 		if parent.right == nil {
// 			parent.right = child
// 		} else {
// 			appendNode(parent.right, child)
// 		}
// 	}
// }
