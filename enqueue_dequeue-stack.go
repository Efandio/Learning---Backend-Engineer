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
		return 0, fmt.Errorf("Item queue kosong")
	}

	firstItem := q.items[0]
	q.items = q.items[ 1: ]

	return firstItem, nil
}

func main() {

		fmt.Println("--- Implementasi Queue (Enqueue & Dequeue) ---")

	myQueue := Queue{} // Inisialisasi Queue baru

	fmt.Println("\n--- Operasi Enqueue ---")
	myQueue.Enqueue(100) // Output: Enqueued 100. Queue: [100]
	myQueue.Enqueue(200) // Output: Enqueued 200. Queue: [100 200]
	myQueue.Enqueue(300) // Output: Enqueued 300. Queue: [100 200 300]

	fmt.Println("\n--- Operasi Dequeue ---")

	// Test Dequeue
	item, err := myQueue.Dequeue() // Output: Dequeued 100. Queue: [200 300]
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Item Dequeued:", item) // Output: Item Dequeued: 100
	}

	item, err = myQueue.Dequeue() // Output: Dequeued 200. Queue: [300]
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Item Dequeued:", item) // Output: Item Dequeued: 200
	}

	item, err = myQueue.Dequeue() // Output: Dequeued 300. Queue: []
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Item Dequeued:", item) // Output: Item Dequeued: 300
	}

	fmt.Println("\n--- Dequeue dari Queue Kosong ---")
	item, err = myQueue.Dequeue() // Output: Error
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Item Dequeued:", item)
	}

	fmt.Println("\nFinal Queue:", myQueue.items) // Final Queue: []

}
