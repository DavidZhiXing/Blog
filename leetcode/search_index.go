package leetcode 

// given a sorted array and a target value, return the index if the target is found. if not, return the index where it would be if it were inserted in order.
//
// example:
//
// input: [1,3,5,6], 5
// output: 2
//
// input: [1,3,5,6], 2
// output: 1
//
// input: [1,3,5,6], 7
// output: 4
//
// input: [1,3,5,6], 0
// output: 0
//
// input: [1,3,5,6], 4
// output: 3
//
// input: [1,3,5,6], 6
// output: 4
//

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] >= target {
			high = mid - 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] >= target) {
				return mid + 1
			}
			low = mid + 1
		}
	}
	return 0
}