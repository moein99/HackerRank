package main

import (
	"fmt"
	"strings"
)

func main() {
	var length int
	fmt.Scanf("%d", &length)

	for i := 0; i < length; i++ {
		spaces := strings.Repeat(" ", length - 1 - i)
		sharps := strings.Repeat("#", i + 1)
		fmt.Println(spaces + sharps)
	}
}
