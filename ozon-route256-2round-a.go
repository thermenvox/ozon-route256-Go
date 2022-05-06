// https://cups.online/ru/tasks/1222

package main

import (
	"fmt"
)

func main() {
	var friends int
	fmt.Scan(&friends)

	for i := 0; i < friends+1; i++ {
		var districts uint64
		var mod uint64

		fmt.Scan(&districts, &mod)

		if districts >= mod {
			fmt.Println(0)
		} else {
			variants := factorial(districts)
			variant := variants % mod
			fmt.Println(variant)
		}
	}
}

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
