/*
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (ie, buy one and sell one share of the stock), design an algorithm to find the maximum profit.
*/

package main

import "fmt"

var testData = []struct {
	prices   []int
	expected int
}{
	{[]int{}, 0},
	{[]int{7, 1, 5, 3, 6, 4}, 5},
	{[]int{7, 6, 4, 3, 1}, 0},
	{[]int{1, 100, 2, 3, 4, 5, 99}, 99},
	{[]int{1, 4, 5, 2, 7, 8, 0, 1, 9}, 9},
	{[]int{1, 4, 5, 2, 7, 8, -3, 1, 9}, 12},
	{[]int{-1, -4, -5, -2, -7, -8, -3, -1, -9}, 7},
	{[]int{-1, -4, -5, -2, 3, -8, -3, -1, -9}, 8},
}

func main() {

	for _, tt := range testData {
		profit := maxProfit(tt.prices)
		if profit != tt.expected {
			fmt.Printf("Error! Prices: %v Max Profit: %d, Expected: %d\n", tt.prices, profit, tt.expected)
		} else {
			fmt.Printf("Prices: %v Max Profit: %d, Expected: %d\n", tt.prices, profit, tt.expected)
		}
	}
}

func maxProfit(prices []int) int {
	maxProfit := 0 // If there is no profit to be made, return 0
	if len(prices) == 0 {
		return maxProfit
	}
	minPrice := prices[0] // Start at the beginning and set the lowest price
	for _, curPrice := range prices {
		if curPrice > minPrice { // There is some sort of profit
			curProfit := curPrice - minPrice
			if maxProfit < curProfit { // If the profit is higher than our current profit, we have a winner
				maxProfit = curPrice - minPrice
			}
		} else {
			minPrice = curPrice // The current stock is cheaper so it is a better deal
		}
	}
	return maxProfit
}
