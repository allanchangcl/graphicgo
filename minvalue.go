package main

func mymin(arrays []int, length int) int {
	var minIndex = 0
	for j := 1; j < length; j++ {
		if arrays[minIndex] > arrays[j] {
			minIndex = j
		}
	}
	return arrays[minIndex]
}
