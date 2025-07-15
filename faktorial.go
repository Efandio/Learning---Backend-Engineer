package main

import (
    "fmt"
)

func Faktorial(n int) int {
    if n == 0 || n == 1 {
        return 1
    }

    

    return n * Faktorial(n-1)
}

func main() {

fmt.Println("Faktorial dari 7:", Faktorial(7)) // hasilnya: 5040

}
