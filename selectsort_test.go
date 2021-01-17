package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectSort(t *testing.T) {
	var scores = []int{90, 70, 50, 80, 60, 85}
	var length = len(scores)
	var sorted = []int{50, 60, 70, 80, 85, 90}
	mySelectSort(scores, length)
	assert.Equal(t, scores, sorted, "The values should be the same")

}
