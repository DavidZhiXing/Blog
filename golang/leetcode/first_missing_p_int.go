package leetcode

// https://leetcode.com/problems/first-missing-positive/
// 
// Given an unsorted integer array, find the first missing positive integer.
//
// For example,
// Given [1,2,0] return 3,
// and [3,4,-1,1] return 2.
//
// Your algorithm should run in O(n) time and uses constant space.

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

// test
func TestFirstMissingPositive() {
	println(firstMissingPositive([]int{1, 2, 0}))
	println(firstMissingPositive([]int{3, 4, -1, 1}))
}

func main() {
	TestFirstMissingPositive()
}