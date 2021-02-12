package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApparyAppend(t *testing.T) {
	var scores = []int{90, 70, 50, 80, 60, 85, 75}
	var myArrayAppend = myArrayAppend()
	assert.Equal(t, scores, myArrayAppend, "The values should be the same")
}
