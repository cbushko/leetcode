package main

import "fmt"
import "strconv"

var testData = []struct {
	number   int
	expected bool
}{
	{1221, true},
	{121, true},
	{345, false},
	{1234455556, false},
	{123456789987654321, true},
	{12345678987654321, true},
}

func main() {
	for _, tt := range testData {
		result := isPalindrome(tt.number)
		if result == tt.expected {
			fmt.Printf("+ Have: %t, Want: %t\n", result, tt.expected)
		} else {
			fmt.Printf("- Have: %t, Want: %t\n", result, tt.expected)
		}
	}
}

func isPalindrome(x int) bool {
	palString := strconv.Itoa(x)

	start := 0
	end := len(palString) - 1
	for ; start < end; start++ {
		if palString[start] != palString[end] {
			return false
		}
		end--
	}
	return true
}
