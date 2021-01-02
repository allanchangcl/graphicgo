package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	var scores = []int{60, 80, 95, 50, 70, 23}
	var length = len(scores)
	myBubbleSort(scores, length)
	var sorted = []int{23, 50, 60, 70, 80, 95}

	for i := 0; i < length; i++ {
		assert.Equal(t, scores[i], sorted[i], "The values should be the same")
	}
}
