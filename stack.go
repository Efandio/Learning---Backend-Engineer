package main

import (
	"fmt"
)

type Stack struct {
	Items []int
}


func (s *Stack) Push(item int) {
	s.Items = append(s.Items, item)
	fmt.Printf("Pushed %d. Stack %v\n", item, s.Items)
}

func (s *Stack) Pop() (int, error) {
	if len(s.Items) == 0 {
		return 0, fmt.Errorf("Item kosong, tidak bisa pop")
	}

	lastItem := len(s.Items) - 1
	item := s.Items[lastItem]
	s.Items = s.Items[ :lastItem]

	return item, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.Items) == 0 {
		return 0, fmt.Errorf("Stack kosong")
	}

	lastItem := len(s.Items) - 1
	item := s.Items[lastItem]

	return item, nil
}

func (s *Stack) IsEmpty() bool {
	if len(s.Items) == 0 {
		return true
	}
	return false
}

func main() {

	fmt.Println("--- Implementasi Stack Lengkap ---")

	myStack := Stack{} 

	fmt.Println("\n--- Operasi Push ---")
	myStack.Push(10) 
	myStack.Push(20) 
	myStack.Push(30) 

	fmt.Println("\n--- Operasi Peek ---")
	item, err := myStack.Peek()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Peeked item:", item) // Harusnya: 30
	}
	fmt.Printf("Stack after Peek: %v\n", myStack.Items) // Harusnya: [10 20 30] (tidak berubah)

	fmt.Println("\n--- Operasi IsEmpty ---")
	fmt.Printf("Is stack empty? %t\n", myStack.IsEmpty()) // Harusnya: false

	fmt.Println("\n--- Operasi Pop ---")
	myStack.Pop() 
	myStack.Pop() 
	myStack.Pop() 

	fmt.Println("\n--- Operasi IsEmpty setelah Pop semua ---")
	fmt.Printf("Is stack empty? %t\n", myStack.IsEmpty()) // Harusnya: true

	fmt.Println("\n--- Operasi Peek dari Stack Kosong ---")
	item, err = myStack.Peek()
	if err != nil {
		fmt.Println("Error:", err) // Harusnya: Error
	} else {
		fmt.Println("Peeked item:", item)
	}

}
