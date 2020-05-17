package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	root *Node
}

func (l *LinkedList) appendValue(i int) {
	newNode := &Node{value: i}
	if l.root == nil {
		l.root = newNode
	} else {
		node := l.root
		for node.next != nil {
			node = node.next
		}
		node.next = newNode
	}
}

func (l *LinkedList) length() int {
	if l.root == nil {
		listLength := 0
	} else {
		node := l.root
		listLength := 1
		for node.next != nil {
			node = node.next
			listLength++
		}
	}
	return listLength
}

func main() {
	var l LinkedList
	l.appendValue(2)
	l.appendValue(3)
	fmt.Println(l.length())
	// l.printList()
}
