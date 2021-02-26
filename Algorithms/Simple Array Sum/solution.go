package main

import (
    "fmt"
)

func sumIntSlice(nums []int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    
    return total
}

func readOneLineOfNums(length int) []int {
    nums := make([]int, length)
    for i := 0; i < length; i++ {
        fmt.Scanf("%d", &nums[i])
    }

    return nums
}

func main() {
    var i int
    fmt.Scanf("%d", &i)
    nums := readOneLineOfNums(i)
    total := sumIntSlice(nums)
    fmt.Print(total)
}
