package leetcode

import (
	"testing"
)


type TestCase struct {
	input  []int
	output int
	}

func TestMaxArea(t *testing.T) {
	var TestCases = []TestCase
		{
		 []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		 49,
		}
	for _, testCase := range TestCases {
		result := maxArea(testCase.input)
		if result != testCase.output {
			t.Errorf("Test Failed! input: %v, expected: %v, received: %v", testCase.input, testCase.output, result)
		}
	}
}

