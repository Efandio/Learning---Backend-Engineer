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

func PreorderTraversal(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%d", node.Data)
	PreorderTraversal(node.Left)
	PreorderTraversal(node.Right)
}

func PostorderTraversal(node *Node) {
	if node == nil {
		return
	}

	PostorderTraversal(node.Left)
	PostorderTraversal(node.Right)
	fmt.Printf("%d", node.Data)
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
		fmt.Println("--- Traversal Binary Tree ---")

	fmt.Println("\nHasil Preorder Traversal:")
	PreorderTraversal(root) // Harusnya: 1 2 4 5 3
	fmt.Println()

	fmt.Println("\nHasil Postorder Traversal:")
	PostorderTraversal(root) // Harusnya: 4 5 2 3 1
	fmt.Println()
}
