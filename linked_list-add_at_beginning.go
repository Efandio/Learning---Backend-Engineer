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

func AddAtBeginning(head *Node, data int) *Node {
    
    newNode := &Node{Data: data, Next: nil}
    newNode.Next = head

    return newNode
}

func main() {

    var myHead *Node = nil


    fmt.Print("List Awal: ")
    PrintList(myHead)

    myHead = AddAtBeginning(myHead, 10)
    fmt.Print("Setelah tambah 10: ")
    PrintList(myHead)

    myHead = AddAtBeginning(myHead, 5)
    fmt.Print("Setelah tambah 5: ")
    PrintList(myHead)

    myHead = AddAtBeginning(myHead, 1)
    fmt.Print("Setelah tambah 1: ")
    PrintList(myHead)
}
