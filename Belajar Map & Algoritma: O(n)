package main

import (
	"fmt"
)


func countChar(input string) map[rune]int {
    counts := make(map[rune]int)

    for _, char := range input {
        counts[char]++
    }
    return counts
}


func main() {

    tes1 := "Hello, World"

    mapResult := countChar(tes1)

    for charCode, count := range mapResult {
        fmt.Printf("%c: %d\n", charCode, count)
    }

}
