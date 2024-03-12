package main

import (
	"math"
)

func maxProfit(prices []int) int {
	minPrice, maxProfit := math.MaxInt, 0
	for _, value := range prices {
		if value < minPrice {
			minPrice = value
		} else if (value - minPrice) > maxProfit {
			maxProfit = value - minPrice
		}
	}
	return maxProfit
}
