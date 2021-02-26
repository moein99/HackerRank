package main

import (
	"fmt"
	"math"
)

func main() {
	nums := make([]int, 5)
	var sum, min, max = 0, int(math.Pow10(10)), -1
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &nums[i])
		sum += nums[i]
	}

	for _, val := range nums {
		if sum - val > max {
			max = sum - val
		}
		if sum - val < min {
			min = sum - val
		}
	}
	fmt.Printf("%d %d", min, max)
}
