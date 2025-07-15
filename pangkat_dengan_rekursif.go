package main

import (
    "fmt"
)

func HitungPangkat(base, exponent int) int {
    if exponent == 0 {
        return 1
    }

    if exponent == 1 {
        return base
    }

    return base * HitungPangkat(base, exponent - 1)

}

func main() {

    fmt.Println("2 pangkat 3:", HitungPangkat(2, 3)) // Harusnya: 8
    fmt.Println("5 pangkat 2:", HitungPangkat(5, 2)) // Harusnya: 25
    fmt.Println("7 pangkat 1:", HitungPangkat(7, 1)) // Harusnya: 7
    fmt.Println("10 pangkat 0:", HitungPangkat(10, 0)) // Harusnya: 1
    fmt.Println("3 pangkat 4:", HitungPangkat(3, 4)) // Harusnya: 81


}
