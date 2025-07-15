package main

import (
    "fmt"
)

func SumSeries(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }

    return n + SumSeries(n - 2)
}


func main() {

    fmt.Println("SumSeries(0):", SumSeries(0)) // Harusnya: 0
    fmt.Println("SumSeries(1):", SumSeries(1)) // Harusnya: 1
    fmt.Println("SumSeries(2):", SumSeries(2)) // Harusnya: 2 (2 + 0)
    fmt.Println("SumSeries(3):", SumSeries(3)) // Harusnya: 4 (3 + 1)
    fmt.Println("SumSeries(5):", SumSeries(5)) // Harusnya: 9 (5 + 3 + 1)
    fmt.Println("SumSeries(6):", SumSeries(6)) // Harusnya: 12 (6 + 4 + 2 + 0)

}
