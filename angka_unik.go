package main

import (
    "fmt"
)


func AngkaUnik(arr []int) int {
    count := make(map[int]int)

    for _, val := range arr {
        count[val]++
    }

    for _, val := range arr {
        if count[val] == 1 {
            return val
        }
    }
    return 0
}

func main() {

    numbers1 := []int{1, 2, 2, 3, 4, 3, 5, 1}
    fmt.Printf("angka unik pertama di %v adalah: %d\n", numbers1, AngkaUnik(numbers1))

}
