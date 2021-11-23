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

func searchIndex(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[right] {
			if nums[mid] > target && nums[left] <= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return left
}