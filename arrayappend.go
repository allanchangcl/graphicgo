package main

func myArrayAppend() []int {
	var scores = []int{90, 70, 50, 80, 60, 85}

	var length = len(scores)
	var tempArray = make([]int, length+1)

	for i := 0; i < length; i++ {
		tempArray[i] = scores[i]
	}
	tempArray[length] = 75
	scores = tempArray

	return tempArray

}
