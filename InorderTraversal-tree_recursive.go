package main

import (
	"fmt"
)

type Node struct {
	Data int
	Left *Node
	Right *Node
}

func InorderTraversal(node *Node) {
	// basis case
	if node == nil {
		return 
	}

	// recursive dari node kiri
	InorderTraversal(node.Left)

	fmt.Printf("%d", node.Data)

	InorderTraversal(node.Right)


}

func main() {

	root := &Node{Data: 1}
	node2 := &Node{Data: 2}
	node3 := &Node{Data: 3}
	node4 := &Node{Data: 4}
	node5 := &Node{Data: 5}

	root.Left = node2
	root.Right = node3

	node2.Left = node4
	node2.Right = node5

	fmt.Println("\n Hasil Traversal: ")
	InorderTraversal(root)
}
