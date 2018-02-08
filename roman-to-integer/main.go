/*
Given a roman numeral, convert it to an integer.

Input is guaranteed to be within the range from 1 to 3999.
*/

package main

import "fmt"

var testData = []struct {
	data     string
	expected int
}{
	{"I", 1},
	{"V", 5},
	{"X", 10},
	{"L", 50},
	{"C", 100},
	{"D", 500},
	{"M", 1000},
	{"IV", 4},
	{"IX", 9},
	{"XX", 20},
	{"XXXIX", 39},
	{"XLI", 41},
	{"XCIX", 99},
	{"DCXXI", 621},
	{"MCMXCVI", 1996},
	{"MCDLXXVI", 1476},
	{"MMCDXLIX", 2449},
	{"B", 0}, // B is an invalid roman numeral and should be ignored
	{"BIV", 4},
	{"BIBV", 4},
	{"MBMCBDXLBIX", 2449},
}

func main() {

	for _, tt := range testData {
		result := romanToInt(tt.data)
		if result != tt.expected {
			fmt.Printf("Error! Data: %v Result: %d, Expected: %d\n", tt.data, result, tt.expected)
		} else {
			fmt.Printf("Data: %v Result: %d, Expected: %d\n", tt.data, result, tt.expected)
		}
	}
}

func romanToInt(s string) int {
	total := 0
	current := 0
	previous := 0
	size := len(s)
	for i := 0; i < size; i++ {
		invalid := false
		switch string(s[i]) {
		case "I":
			current = 1
		case "V":
			current = 5
		case "X":
			current = 10
		case "L":
			current = 50
		case "C":
			current = 100
		case "D":
			current = 500
		case "M":
			current = 1000
		default:
			invalid = true
		}
		if !invalid {
			if current > previous {
				total = total - previous - previous + current
			} else {
				total += current
			}
			previous = current
		}
	}
	return total
}
