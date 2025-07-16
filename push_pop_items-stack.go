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

func main() {

	fmt.Println("--- Implementasi Stack (Push & Pop) ---")

	myStack := Stack{}

	fmt.Println("\n--- Operasi Push ---")
	myStack.Push(10) 
	myStack.Push(20) 
	myStack.Push(30)
	myStack.Push(40)

	fmt.Println("\n--- Operasi Pop ---")

	// Test Pop
	item, err := myStack.Pop() 
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Item Pop:", item) 
	}

	item, err = myStack.Pop() 
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Item Pop:", item) 
	}

	item, err = myStack.Pop() 
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Item Pop:", item) 
	}

	fmt.Println("\n--- Pop dari Stack Kosong ---")
	item, err = myStack.Pop() 
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Item Pop:", item)
	}

	fmt.Println("\nFinal Stack:", myStack.Items) 

}
