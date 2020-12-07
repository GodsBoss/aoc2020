package main

import (
	"fmt"
)

func findCombination(lines []int, sum int, count int) ([]int, error) {
	if count <= 0 {
		return nil, fmt.Errorf("invalid count")
	}
	maxIndex := len(lines) - 1
	indices := make([]int, count)
	for {
		if allIndicesAreUnequal(indices) && matchSum(lines, sum, indices) {
			result := make([]int, count)
			for i := range indices {
				result[i] = lines[indices[i]]
			}
			return result, nil
		}
		if !increaseIndices(indices, maxIndex) {
			return nil, fmt.Errorf("no matching combination found")
		}
	}
}

func allIndicesAreUnequal(indices []int) bool {
	for i := range indices {
		for j := range indices {
			if i != j && indices[i] == indices[j] {
				return false
			}
		}
	}
	return true
}

func matchSum(lines []int, sum int, indices []int) bool {
	actual := 0
	for i := range indices {
		actual += lines[indices[i]]
	}
	return actual == sum
}

// increaseIndices returns wether the index combination is still viable.
func increaseIndices(indices []int, maxIndex int) bool {
	indices[len(indices)-1]++
	for i := len(indices) - 1; i > 0; i-- {
		if indices[i] > maxIndex {
			indices[i] = 0
			indices[i-1]++
		}
	}
	return indices[0] <= maxIndex
}
