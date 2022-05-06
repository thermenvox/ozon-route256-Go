// https://cups.online/ru/tasks/1218

package main

import (
	"fmt"
)

func main() {
	var masksNeed int
	fmt.Scan(&masksNeed)

	const pack = 24
	const box = 16 * pack

	boxes := masksNeed / box
	ost1 := masksNeed - boxes*box
	packs := ost1 / pack
	masks := ost1 - packs*pack

	if masks >= 20 {
		packs += 1
		masks = 0
	}

	if packs > 14 {
		boxes += 1
		packs = 0
	}

	fmt.Println(masks, packs, boxes)
}
