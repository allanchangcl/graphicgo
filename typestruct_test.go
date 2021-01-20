package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeStruct(t *testing.T) {
	pet := myStructtype().name
	assert.Equal(t, pet, "Ricky", "The values should be the same")

}
