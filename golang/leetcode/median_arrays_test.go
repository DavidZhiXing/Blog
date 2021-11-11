package leetcode

import (
	"testing"
)

type TestCase struct {
	nums1   []int
	nums2   []int
	median  float64
	median2 float64
}

func TestMedianOfTwoSortedArrays(t *testing.T) {
	var tests = []TestCase{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{1, 2, 3}, []int{4, 5, 6}, 3.5},
		{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, 4.5},
		{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}, 5.5},
		{[]int{1, 2, 3, 4, 5, 6}, []int{7, 8, 9, 10, 11, 12}, 6.5},
	}

	for _, test := range tests {
		median := findMedianSortedArrays(test.nums1, test.nums2)
		if median != test.median {
			t.Errorf("findMedianSortedArrays(%v, %v) = %v, want %v", test.nums1, test.nums2, median, test.median)
		}
	}
}
