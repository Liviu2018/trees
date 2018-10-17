package main_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trees/interpolation_search/main/main"
)

// TestInterpolationSearch does a simple input search
func TestInterpolationSearch(t *testing.T) {
	input := generateInput(500)
	resultIndex, found := main.InterpolationSearch(input, input[5])

	fmt.Println("aaaaaaaaaaaa")

	assert.Equal(t, 5, resultIndex, "should be 5, true")
	assert.Equal(t, true, found, "should be 5, true")
}

func generateInput(n int) []int {
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = 2 * i
	}

	return result
}
