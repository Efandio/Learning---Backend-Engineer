package main

import (
    "fmt"
)

func ProductIndex(arr []int, target int) int {
    low := 0
    high := len(arr) - 1

    for low <= high {
        mid := low + (high - low) / 2

        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return -1
}


func main() {

    idProduct := []int{121, 122, 123, 124, 125}
    target := 124
    getProductIndex := ProductIndex(idProduct, target)

    fmt.Printf("Index produk %d adalah %d", target, getProductIndex)

}
