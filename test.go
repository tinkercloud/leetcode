package main

import "fmt"

func main() {
	arr := []int{1, 2, 8, 2, 1}

	var result = 0

	for i := 0; i < len(arr); i++ {
		result = result ^ arr[i]
		fmt.Println(result)
	}

	fmt.Println(result)
}
