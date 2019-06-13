package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	testCases := []struct {
		inputs   []int
		expected int
	}{
		{[]int{1, 4, 6}, 6},
		{[]int{0, -1, -3}, 0},
		{[]int{1}, 1},
		{[]int{-1, -2, -3}, -1},
	}
	for _, testCase := range testCases {
    actual := max(testCase.inputs)
		if actual != testCase.expected {
			t.Errorf("invalid actual:%d,expected:%d", actual, testCase.expected)
		}
	}
}
