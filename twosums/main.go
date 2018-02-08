// Given an array of integers, return indices of the two numbers such that they add up to a specific target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example:
// Given nums = [2, 7, 11, 15], target = 9,

// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

package main

import (
	"fmt"
)

var testData = []struct {
	nums     []int
	target   int
	expected []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{2, 7, 11, 15}, 17, []int{0, 3}},
}

func main() {
	for _, tt := range testData {
		result := twoSum(tt.nums, tt.target)
		fmt.Printf("Result: %d, should be %d\n", result, tt.expected)
	}
}

func twoSum(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			val := nums[i] + nums[j]
			if val == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
