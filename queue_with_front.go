package main

import (
	"fmt"
)

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
	fmt.Printf("Enqueue %d. Queue: %v\n", item, q.items)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, fmt.Errorf("item queue kosong")
	}

	firstItem := q.items[0]
	q.items = q.items[ 1: ]

	return firstItem, nil
}

func (q *Queue) Front() (int, error) { // Front ini sama dengan peek di stack
	if len(q.items) == 0 {
		return 0, fmt.Errorf("Queue kosong")
	}

	return q.items[0], nil
}

func (q *Queue) IsEmpty() bool {
	if len(q.items) == 0 {
		return true
	}
	return false
}


func main() {

fmt.Println("--- Implementasi Queue Lengkap ---")

	myQueue := Queue{} 

	fmt.Println("\n--- Operasi Enqueue ---")
	myQueue.Enqueue(100) 
	myQueue.Enqueue(200) 
	myQueue.Enqueue(300) 

	fmt.Println("\n--- Operasi Front ---")
	item, err := myQueue.Front()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Front item:", item) // Harusnya: 100
	}
	fmt.Printf("Queue after Front: %v\n", myQueue.items) // Harusnya: [100 200 300] (tidak berubah)

	fmt.Println("\n--- Operasi IsEmpty ---")
	fmt.Printf("Is queue empty? %t\n", myQueue.IsEmpty()) // Harusnya: false

	fmt.Println("\n--- Operasi Dequeue ---")
	myQueue.Dequeue() 
	myQueue.Dequeue() 
	myQueue.Dequeue() 

	fmt.Println("\n--- Operasi IsEmpty setelah Dequeue semua ---")
	fmt.Printf("Is queue empty? %t\n", myQueue.IsEmpty()) // Harusnya: true

	fmt.Println("\n--- Operasi Front dari Queue Kosong ---")
	item, err = myQueue.Front()
	if err != nil {
		fmt.Println("Error:", err) // Harusnya: Error: Queue kosong, tidak bisa front
	} else {
		fmt.Println("Front item:", item)
	}

}
