package leetcode

// https://leetcode.com/problems/trapping-rain-water/
// 
// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.
//
// For example,
// Given [0,1,0,2,1,0,1,3,2,1,2,1], return 6.
//
// The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!
//
//
//
// Credits:Special thanks to @memoryless for adding this problem and creating all test cases.
//
// Subscribe to see which companies asked this question
//
// Show Tags
// 
// Array
// Two Pointers
//
// Show Similar Problems
// 
// (M) Container With Most Water
// (M) Maximal Rectangle
// (M) Trap Rain Water

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	res := 0
	for left < right {
		if leftMax < rightMax {
			left++ 
			if height[left] > leftMax {
				leftMax = height[left]
			}
			res += leftMax - height[left]
		} else {
			right--
			if height[right] > rightMax {
				rightMax = height[right]
			}
			res += rightMax - height[right]
		}
	}
	return res
}

type height struct {
	h int
	i int
}

