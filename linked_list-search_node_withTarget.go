package main

import (
	"fmt"
)

type Node struct {
    Data int
    Next *Node
}

func PrintList(head *Node) {
    current := head

    for current != nil {
        fmt.Printf("%d -> ", current.Data)
        current = current.Next
    }
    fmt.Println("nil")
}

func SearchNode(head *Node, targetData int) bool {
    current := head
    for current != nil {
        if current.Data == targetData {
            return true
        } else {
            current = current.Next
        }
    }
    return false
}

func main() {

	head := &Node{Data: 10}
	head.Next = &Node{Data: 20}
	head.Next.Next = &Node{Data: 30}
	head.Next.Next.Next = &Node{Data: 40}

	fmt.Print("List: ")
	PrintList(head) 

	fmt.Println("Mencari 30:", SearchNode(head, 30)) // Harusnya: true
	fmt.Println("Mencari 10:", SearchNode(head, 10)) // Harusnya: true
	fmt.Println("Mencari 50:", SearchNode(head, 50)) // Harusnya: false
	fmt.Println("Mencari 40:", SearchNode(head, 40)) // Harusnya: true
	fmt.Println("Mencari 0 di list kosong:", SearchNode(nil, 0)) // Harusnya: false (test dengan list kosong)
}
