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

func DeleteNode(head *Node, targetData int) *Node {
    if head == nil {
        return nil
    }

    if head.Data == targetData {
        return head.Next
    }

    previous := head
    current := head.Next
    for current != nil {
        if current.Data == targetData {
            previous.Next = current.Next
            return head
        } 
        previous = current
        current = current.Next
    }
    return head
}

func main() {

	head := &Node{Data: 10}
	head.Next = &Node{Data: 20}
	head.Next.Next = &Node{Data: 30}
	head.Next.Next.Next = &Node{Data: 40}

	fmt.Print("List: ")
	PrintList(head) 

    // Hapus 30 (tengah)
	head = DeleteNode(head, 30)
	fmt.Print("Setelah hapus 30: ")
	PrintList(head) // Harusnya: 10 -> 20 -> 40 -> nil

	// Hapus 10 (head)
	head = DeleteNode(head, 10)
	fmt.Print("Setelah hapus 10: ")
	PrintList(head) // Harusnya: 20 -> 40 -> nil

	// Hapus 40 (akhir)
	head = DeleteNode(head, 40)
	fmt.Print("Setelah hapus 40: ")
	PrintList(head) // Harusnya: 20 -> nil

	// Hapus 20 (Node terakhir)
	head = DeleteNode(head, 20)
	fmt.Print("Setelah hapus 20: ")
	PrintList(head) // Harusnya: nil

	// Coba hapus dari list kosong
	head = DeleteNode(head, 99)
	fmt.Print("Setelah hapus 99 dari kosong: ")
	PrintList(head) // Harusnya: nil

	// Buat list baru untuk testing tidak ditemukan
	head2 := &Node{Data: 1}
	head2.Next = &Node{Data: 2}
	head2.Next.Next = &Node{Data: 3}
	fmt.Print("List baru: ")
	PrintList(head2) // Output: 1 -> 2 -> 3 -> nil

	// Coba hapus yang tidak ada
	head2 = DeleteNode(head2, 5)
	fmt.Print("Setelah hapus 5 (tidak ada): ")
	PrintList(head2) // Harusnya: 1 -> 2 -> 3 -> nil (tidak berubah)
}
