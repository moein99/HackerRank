package main

import "fmt"

func readOneLineOfNums(length int) []int64 {
	nums := make([]int64, length)
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &nums[i])
	}

	return nums
}

func sumIntSlice(nums []int64) int64 {
	var total int64 = 0
	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	var i int
	fmt.Scanf("%d", &i)
	nums := readOneLineOfNums(i)
	result := sumIntSlice(nums)
	fmt.Println(result)
}
