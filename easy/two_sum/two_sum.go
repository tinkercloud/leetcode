package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	// rs := []int{}

	for index, value := range nums {
		comp := target - value

		if _, ok := m[comp]; ok {
			return []int{m[comp], index}
		} else {
			m[value] = index
		}

	}

	return []int{}

	// for i := 0; i < len(nums); i++ {
	// 	comp := target - nums[i]

	// 	if m[comp] != 0 {
	// 		fmt.Println(m[comp])
	// 	}

	// 	m[nums[i]] = i
	// }
	// for j := 0; j < len(nums); j++ {

	// 	// 获取 target - 索引值
	// 	comp := target - nums[j]
	// 	// fmt.Println(comp)
	// 	// fmt.Println(m[comp])
	// 	if m[comp] != 0 {
	// 		rs = append(rs, nums[j])
	// 	}
	// }

}

func main() {

	nums := []int{2, 7, 11, 15}
	target := 9

	result := TwoSum(nums, target)

	fmt.Println(result)
	http.ListenAndServe("0.0.0.0:8080", nil)

}
