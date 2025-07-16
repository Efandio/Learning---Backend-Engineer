package main

import (
	"fmt"
)

type Node struct {
	Data int
	Left *Node
	Right *Node
}

func PreorderTraversal(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%d", node.Data)
	PreorderTraversal(node.Left)
	PreorderTraversal(node.Right)
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

	fmt.Println("\nHasil Preorder Traversal:")
	PreorderTraversal(root) // Harusnya: 1 2 4 5 3
	fmt.Println()

}
