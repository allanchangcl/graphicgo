package main

func myMax(myArrays []int, length int) int {
	for i := 0; i < length-1; i++ {
		if myArrays[i] < myArrays[i+1] {
			var temp = myArrays[i]
			myArrays[i] = myArrays[i+1]
			myArrays[i+1] = temp
		}
	}
	var maxValue = myArrays[length-1]
	return maxValue
}
