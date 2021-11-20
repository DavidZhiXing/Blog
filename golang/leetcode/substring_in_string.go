package leetcode

// you are given a string, s, and a list of words, words, that are all of the same length. Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.
//
// For example, given:
// s: "barfoothefoobarman"
// words: ["foo", "bar"]
//
// You should return the indices: [0,9].
// (order does not matter).
//
// Subscribe to see which companies asked this question
//

func substringInString(s string, words []string) []int {
	var res []int
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(words); j++ {
			if i+len(words[j]) <= len(s) && s[i:i+len(words[j])] == words[j] {
				res = append(res, i)
			}
		}
	}
	return res
}

// test
func TestSubstringInString(t *testing.T) {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	res := substringInString(s, words)
	if len(res) != 2 || res[0] != 0 || res[1] != 9 {
		t.Error("error")
	}
}