package main

import (
    "fmt"
)

func Fibonacci(n int) int {
    // basis
    if n == 0 {
        return 0
    }

    if n == 1 {
        return 1
    }

    return Fibonacci(n - 1) + Fibonacci(n - 2)
}

func main() {

    fmt.Println("Fibonacci ke-0:", Fibonacci(0)) // Harusnya: 0
    fmt.Println("Fibonacci ke-2:", Fibonacci(2)) // Harusnya: 1
    fmt.Println("Fibonacci ke-3:", Fibonacci(3)) // Harusnya: 2

}
