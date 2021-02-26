package main

import "fmt"

func convertToStr(val int) string {
	if val < 10 {
		return "0" + fmt.Sprint(val)
	}
	return fmt.Sprint(val)
}

func main() {
	var h, m, s int
	var str string
	fmt.Scanf("%d:%d:%d%s", &h, &m, &s, &str)
	if h == 12 && str == "AM" {
		h = 0
	}
	if h < 12 && h > 0 && str == "PM" {
		h += 12
	}

	fmt.Printf("%s:%s:%s", convertToStr(h), convertToStr(m), convertToStr(s))
}
