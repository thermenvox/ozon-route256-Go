// https://cups.online/ru/tasks/1219

package main

import (
	"fmt"
)

func main() {
	var days int
	var price int

	var prices []int

	fmt.Scan(&days)

	for i := 0; i < days; i++ {
		fmt.Scan(&price)
		prices = append(prices, price)
	}
	sell(prices, 0, 0, 0)

}

func sell(array []int, maxPrice int, dayOfMaxPrice int, revenue int) {
	revenue += maxPrice * dayOfMaxPrice

	if len(array) > 0 {
		array = array[dayOfMaxPrice:]
		maxPrice, dayOfMaxPrice = findMax(array)
		sell(array, maxPrice, dayOfMaxPrice, revenue)
	} else {
		fmt.Println(revenue)
	}
}

func findMax(array []int) (int, int) {
	var maxPrice int = 0
	var dayOfMaxPrice int = 0
	for i := 0; i < len(array); i++ {
		if array[i] > maxPrice {
			maxPrice = array[i]
			dayOfMaxPrice = i + 1
		}
	}
	return maxPrice, dayOfMaxPrice
}
