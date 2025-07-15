package main

import (
    "fmt"
)

type Node struct {
    Data string
    Left *Node
    Right *Node
}

func PrintTree(node *Node) string {

    if node == nil || node.Data == "" {
        return ""
    }
    
    return node.Data + PrintTree(node.Left) + PrintTree(node.Right)
}

func main() {

    root := &Node{Data: "A"}
    root.Left = &Node{Data: "B"}
    root.Right = &Node{Data: "C"}

    root.Left.Left = &Node{Data: "D"}
    root.Left.Right = &Node{Data: "E"}
    root.Right.Right = &Node{Data: "F"}

    fmt.Println("Root:", root.Data)
	fmt.Println("Child Kiri Root:", root.Left.Data)
	fmt.Println("Child Kanan Root:", root.Right.Data)
	fmt.Println("Leaf Node D:", root.Left.Left.Data)

    fmt.Println("Tree: ", PrintTree(root))

}
