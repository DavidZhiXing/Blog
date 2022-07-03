package leetcode

// Problem Link: https://leetcode.com/problems/combination-sum/
// 
// Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.
//
// The same repeated number may be chosen from candidates unlimited number of times.
//
// Note:
//
// All numbers (including target) will be positive integers.
// The solution set must not contain duplicate combinations.
// Example 1:
//
// Input: candidates = [2,3,6,7], target = 7,
// A solution set is:
// [
//   [7],
//   [2,2,3]
// ]
// Example 2:
//
// Input: candidates = [2,3,5], target = 8,
// A solution set is:
// [
//   [2,2,2,2],
//   [2,3,3],
//   [3,5]
// ]

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	var dfs func(int, int)
	dfs = func(sum, index int) {
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		if sum > target || index >= len(candidates) {
			return
		}
		path = append(path, candidates[index])
		dfs(sum+candidates[index], index)
		path = path[:len(path)-1]
		dfs(sum, index+1)
	}
	dfs(0, 0)
	return res
}


// test
func TestCombinationSum() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	res := combinationSum(candidates, target)
	fmt.Println(res)
}

func TestCombinationSum2() {
	candidates := []int{2, 3, 5}
	target := 8
	res := combinationSum(candidates, target)
	fmt.Println(res)
}

func main() {
	TestCombinationSum()
	TestCombinationSum2()
}