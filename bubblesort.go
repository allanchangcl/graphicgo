package main

func myBubbleSort(myArrays []int, length int) {
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if myArrays[j] > myArrays[j+1] {
				var flag = myArrays[j]
				myArrays[j] = myArrays[j+1]
				myArrays[j+1] = flag
			}
		}
	}
}
