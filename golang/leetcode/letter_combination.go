package leetcode 

// Given a digit string, return all possible letter combinations that the number could represent.
//
// A mapping of digit to letters (just like on the telephone buttons) is given below.
//
//
//
// Input:Digit string "23"
// Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//
//
// Note:
// Although the above answer is in lexicographical order, your answer could be in any order you want.

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var res []string
	var m = map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
	var dfs func(string, string)
	dfs = func(digits, cur string) {
		if len(digits) == 0 {
			res = append(res, cur)
			return
		}
		for _, v := range m[digits[0]] {
			dfs(digits[1:], cur + v)
		}
	}
	dfs(digits, "")
	return res
}