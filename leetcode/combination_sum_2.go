package leetcode

// Problem Link: https://leetcode.com/problems/combination-sum-ii/
//
// Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.
//
// Each number in candidates may only be used once in the combination.
//
// Note:
//
// All numbers (including target) will be positive integers.
// The solution set must not contain duplicate combinations.
// Example 1:
//
// Input: candidates = [10,1,2,7,6,1,5], target = 8,
// A solution set is:
// [
//   [1, 7],
//   [1, 2, 5],
//   [2, 6],
//   [1, 1, 6]
// ]
// Example 2:
//
// Input: candidates = [2,5,2,1,2], target = 5,
// A solution set is:
// [
//   [1,2,2],
//   [5]
// ]
//
// Solution:
//
// 1. Sort the candidates array
// 2. Use backtracking to find all combinations
// 3. If the sum of the combination is equal to the target, add the combination to the result
// 4. If the sum of the combination is greater than the target, skip the combination
// 5. If the sum of the combination is less than the target, add the combination to the result and recursively call the function with the remaining candidates
//
// Time Complexity: O(n^2)
// Space Complexity: O(n)
func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(candidates)
	backtrackCombinationSum2(candidates, target, 0, []int{}, &result)
	return result
}

func backtrackCombinationSum2(candidates []int, target int, start int, combination []int, result *[][]int) {
	if target == 0 {
		*result = append(*result, append([]int{}, combination...))
		return
	}
	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		if candidates[i] > target {
			break
		}
		backtrackCombinationSum2(candidates, target-candidates[i], i+1, append(combination, candidates[i]), result)
	}
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	result := combinationSum2(candidates, target)
	fmt.Println(result)
}