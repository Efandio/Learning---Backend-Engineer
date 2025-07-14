package main

import (
    "fmt"
)


func TargetSum (arr []int, sumTarget int) bool {
    seenNums := make(map[int]bool)

    for i := 0; i < len(arr); i++ {
        complemet := sumTarget - arr[i]

        if seenNums[complemet] {
            return true
        }
        seenNums[arr[i]] = true
    }
    return false
}


func main() {

    tes1 := []int{1,2,3,4,5}
    tesTarget := 10

    fmt.Println(TargetSum(tes1, tesTarget))

}
