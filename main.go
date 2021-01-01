package main

import (
	"fmt"
)

func main() {
	var scores = []int{60, 80, 95, 50, 70, 23}
	var length = len(scores)
	var minValue = myMin(scores, length)
	fmt.Printf("Min Value = %d\n", minValue)

}
