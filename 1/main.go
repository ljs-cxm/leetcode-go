package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	ret := twoSum(nums, target)
	fmt.Printf("answer: %v\n", ret)
}

func twoSum(nums []int, target int) []int {
	n := len(nums)
	ret := []int{}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				ret = append(ret, i, j)
				return ret
			}
		}
	}
	return ret
}

func twoSum(nums []int, target int) []int {
	tables := make(map[int]int)
	for index, value := range nums {
		tmp := target - value
		if i, ok := tables[tmp]; ok {
			return []int{i, index}
		}
		tables[value] = index
	}
	return []int{}
}
