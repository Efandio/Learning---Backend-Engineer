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

func AddAtEnd(head *Node, data int) *Node {
    newNode := &Node{Data: data, Next: nil}

    if head == nil {
        return newNode
    } else {
        current := head
        for current.Next != nil {
            current = current.Next
        }
        current.Next = newNode
    }

    return head
}

func main() {

    var myHead *Node = nil

    fmt.Print("List Awal (kosong): ")
    PrintList(myHead) // Output: nil

    myHead = AddAtEnd(myHead, 100) // Tambah ke list kosong
    fmt.Print("Setelah tambah 100 di akhir: ")
    PrintList(myHead) // Harusnya: 100 -> nil

    myHead = AddAtEnd(myHead, 200) // Tambah ke (100 -> nil)
    fmt.Print("Setelah tambah 200 di akhir: ")
    PrintList(myHead) // Harusnya: 100 -> 200 -> nil

    myHead = AddAtEnd(myHead, 300) // Tambah ke (100 -> 200 -> nil)
    fmt.Print("Setelah tambah 300 di akhir: ")
    PrintList(myHead) // Harusnya: 100 -> 200 -> 300 -> nil

}
