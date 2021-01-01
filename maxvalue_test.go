package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxValue(t *testing.T) {
	var scores = []int{60, 80, 95, 50, 70, 23}
	var length = len(scores)
	var maxValue = myMax(scores, length)
	assert.Equal(t, maxValue, 23, "The values should be the same")
}
