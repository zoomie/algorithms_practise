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

func addNode(parentNode, newNode *Node) {
	fmt.Println("hit")
	fmt.Println(parentNode.key)
	fmt.Println(newNode.key)
	if parentNode.key < newNode.key {
		if parentNode.right == nil {
			parentNode.right = newNode
		} else {
			addNode(parentNode.right, newNode)
		}
	} else if parentNode.key > newNode.key {
		if parentNode.left == nil {
			parentNode.left = newNode
		} else {
			addNode(parentNode.right, newNode)
		}
	} else {
		parentNode = newNode
	}
}
func check(node *Node) bool {
	if node.left == nil && node.right == nil {
		return true
	} else if node.left != nil && node.right != nil {
		if node.left.key > node.right.key {
			rightCheck := check(node.left)
			leftCheck := check(node.right)
			if rightCheck && leftCheck {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else if node.left != nil && node.right == nil {
		if node.key > node.left.key {
			rightCheck := check(node.left)
			if rightCheck {
				return true
			} else {
				return false
			}

		} else {
			return false
		}
	} else if node.left == nil && node.right != nil {
		if node.key < node.right.key {
			leftCheck := check(node.right)
			if leftCheck {
				return true
			} else {
				return false
			}
		} else {
			return false
		}

	}
	return false
}

func (b *BST) checkValid() bool {
	if b.root == nil {
		return true
	} else {
		return check(b.root)
	}
}
func main() {
	b := BST{}
	b.addData(1, "one")
	b.addData(2, "two")
	b.addData(3, "three")
	result := b.checkValid()
	fmt.Println(result)
}
