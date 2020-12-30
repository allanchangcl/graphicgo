package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinValue(t *testing.T) {
	var scores = []int{60, 80, 95, 50, 70, 23}
	var length = len(scores)
	var minValue = mymin(scores, length)
	assert.Equal(t, minValue, 23, "The values should be the same")
}
