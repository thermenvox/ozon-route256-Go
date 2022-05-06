// https://cups.online/ru/tasks/1223

package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)

	for i := 0; i < number; i++ {
		var points uint64
		var stepSize uint64
		fmt.Scan(&points, &stepSize)

		if stepSize == 0 {
			fmt.Println(0)
		} else {
			maxSteps := points / 2
			var closestNeighbours uint64
			if maxSteps%stepSize == 0 {
				closestNeighbours = maxSteps / stepSize
			} else {
				closestNeighbours = maxSteps/stepSize + 1
			}
			fmt.Println(closestNeighbours)
		}
	}
}
