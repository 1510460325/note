package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(constructRectangle(8))
}

func constructRectangle(area int) (ans []int) {
	a := int(math.Sqrt(float64(area)))
	for {
		if area % a == 0 && a >= area / a {
			return []int{a, area / a}
		}
		a++
	}
}