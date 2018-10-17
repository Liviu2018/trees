package main

import (
	"fmt"
)

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(input)
	fmt.Println(InterpolationSearch(input, 5))
}

// InterpolationSearch does a search for toBeFound in input slice
func InterpolationSearch(input []int, toBeFound int) (int, bool) {
	if toBeFound < input[0] {
		return -1, false
	}

	if toBeFound > input[len(input)-1] {
		return len(input), false
	}

	start, stop := 0, len(input)-1
	min, max := input[start], input[stop]

	for start <= stop {
		if min == max {
			return min, input[min] == toBeFound
		}

		if input[start] > toBeFound {
			return start, false
		}

		if input[stop] < toBeFound {
			return stop, false
		}

		if input[start] == toBeFound {
			return start, true
		}

		if input[stop] == toBeFound {
			return stop, true
		}

		mid := start + (stop-start)*(toBeFound-input[start])/(input[stop]-input[start])
		if input[mid] == toBeFound {
			return mid, true
		}

		if input[mid] < toBeFound {
			start = mid
		} else {
			stop = mid
		}
	}

	return -1, false
}
