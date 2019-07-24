package main

import (
	"fmt"
	"sort"
)

func OnlyNum(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 0 {
		return -1
	}

	sort.Ints(nums)

	for i := 0; i < len(nums); i += 2 {

		if len(nums)-i == 1 {
			return nums[i]
		}
		// fmt.Println(nums[i], nums[i+1])
		if nums[i] != nums[i+1] {

			return nums[i]
		}

		if nums[i] == nums[i+1] {
			continue
		}
	}
	// for i :=0; i<len(nums); i++{

	// 	for j :=0; j<len(nums); j++ {

	// 		// 同一个位置忽略
	// 		if i==j{
	// 			continue
	// 		}

	// 		if nums[i]==nums[j]{
	// 			break
	// 		}

	// 	}
	// }

	return -1
}

func main() {
	nums := []int{4, 1, 2, 1, 2}

	result := OnlyNum(nums)

	fmt.Println(result)
}
