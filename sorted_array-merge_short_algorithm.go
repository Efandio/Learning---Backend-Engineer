package main

import (
    "fmt"
)

func Merge(arr1, arr2 []int) []int {
    result := []int{}
    i := 0
    j := 0

    // Urutin arr
    for i < len(arr1) && j < len(arr2) {
        if arr1[i] < arr2[j] {
            result = append(result, arr1[i])
            i++
        } else {
            result = append(result, arr2[j])
            j++
        }
    }

    // sisa arr1
    for i < len(arr1) {
        result = append(result, arr1[i])
        i++
    }

    // sisa arr2
    for j < len(arr2) {
        result = append(result, arr2[j])
        j++
    }

    return result
}

func MergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    // ambil nilai midnya
    mid := len(arr) / 2

    // slice awal ke mid
    halfLeft := arr[ :mid ]
    // slice mid ke akhir
    halfRight := arr[ mid: ]

    
    leftSort := MergeSort(halfLeft)
    rightSort := MergeSort(halfRight)

    // urutkan
    return Merge(leftSort, rightSort)
}

func main() {

  unsortedArr := []int{7, 40, 2, 9, 1, 5, 8, 3, 6}
	fmt.Println("Unsorted Array:", unsortedArr)

	sortedArr := MergeSort(unsortedArr)
	fmt.Println("Sorted Array:", sortedArr)

}
