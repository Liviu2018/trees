package main_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trees/interpolation_search/main"
)

// benchmark for interpolation search method
func BenchmarkInterpolationSearch(b *testing.B) {
	b.StopTimer()
	input := generateInput(b.N)
	b.StartTimer()

	resultIndex, found := interpolationsearch.InterpolationSearch(input, input[b.N/2])

	fmt.Println("aaaa benchmark", b.N)

	assert.Equal(b, b.N/2, resultIndex, "should be b.N/2, true")
	assert.Equal(b, true, found, "should be b.N/2, true")
}

// TestInterpolationSearch does a simple input search
func TestInterpolationSearch(t *testing.T) {
	input := generateInput(500)
	resultIndex, found := interpolationsearch.InterpolationSearch(input, input[5])

	fmt.Println("aaaaaabbaaaaaa")

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
