package main

import "fmt"
import "strconv"
import "math"

/*
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output:  321

Example 2:

Input: -123
Output: -321

Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only hold integers within the 32-bit signed integer range. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

var testData = []struct {
	input    int
	expected int
}{
	{123, 321},
	{-123, -321},
	{120, 21},
	{1534236469, 0},
	{-2147483648, 0},
}

func main() {
	for _, tt := range testData {
		result := reverse(tt.input)
		if result != tt.expected {
			fmt.Printf("Error! input: %d, result: %d, should be: %d\n", tt.input, result, tt.expected)
		} else {
			fmt.Printf("Good! input: %d, result: %d\n", tt.input, result)
		}
	}
}

func reverse(x int) int {

	negative := false
	value := x
	if x < 0 {
		value = x * -1
		negative = true
	}

	intString := strconv.Itoa(value)
	newString := ""

	i := len(intString)
	if intString[i-1] == '0' {
		i--
	}

	for ; i > 0; i-- {
		newString += string(intString[i-1])
	}

	newInt, err := strconv.Atoi(newString)
	if err != nil {
		return 0
	}

	if negative == true {
		newInt = newInt * -1
	}
	if newInt > math.MaxInt32 || newInt < math.MinInt32 {
		return 0
	}
	return newInt
}
