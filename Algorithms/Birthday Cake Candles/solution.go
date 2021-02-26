package main

import "fmt"

func readOneLineOfNums(length int) []int {
	nums := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &nums[i])
	}

	return nums
}

func max(nums []int) int {
	var max = -100000000
	for _, item := range nums {
		if item > max {
			max = item
		}
	}

	return max
}

func count(val int, nums []int) int {
	c := 0
	for _, item := range nums {
		if val == item {
			c += 1
		}
	}
	return c
}

func main() {
	var length int
	fmt.Scanf("%d", &length)
	nums := readOneLineOfNums(length)
	m := max(nums)
	c := count(m, nums)
	fmt.Println(c)
}
