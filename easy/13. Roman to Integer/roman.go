package main

import (
	"fmt"
	"strings"
)

func main() {

	v := romanToInt("IVXX")
	fmt.Println(v)
}

func romanToInt(s string) int {

	dic := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	rv := 0
	arrs := strings.Split(s, "")

	for i := 0; i < len(arrs); i++ {

		if i+1 < len(arrs) && dic[arrs[i+1]] > dic[arrs[i]] {
			rv -= dic[arrs[i]]
		} else {
			rv += dic[arrs[i]]
		}
	}

	return rv
}
