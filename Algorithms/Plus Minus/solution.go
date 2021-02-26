package main

import "fmt"

func main() {
	var length int
	fmt.Scanf("%d", &length)

	var num int
	var pos, neg, zero float64 = 0, 0, 0
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &num)
		if num == 0 {
			zero += 1
		} else if num > 0 {
			pos += 1
		} else {
			neg += 1
		}
	}

	fmt.Printf("%f\n%f\n%f\n", pos / float64(length), neg / float64(length), zero / float64(length))
}
