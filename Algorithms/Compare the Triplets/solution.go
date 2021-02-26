package main

import (
	"fmt"
)

func readOneLineOfNums(length int) []int {
	nums := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &nums[i])
	}

	return nums
}

func main() {
	aliceNums := readOneLineOfNums(3)
	bobNums := readOneLineOfNums(3)
	aliceScore, bobScore := 0, 0

	for i := 0; i < len(aliceNums); i++ {
		if aliceNums[i] == bobNums[i] {
			continue
		} else if aliceNums[i] < bobNums[i] {
			bobScore += 1
		} else {
			aliceScore += 1
		}
	}

	fmt.Println(aliceScore, bobScore)
}
