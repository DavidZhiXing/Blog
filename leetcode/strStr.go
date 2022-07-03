package leetcode

// Implement strStr().
// Returns the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
// Subscribe to see which companies asked this question
//
// Tags
// Two Pointers String
//
// Related Topics
// Two Pointers

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}

	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			if haystack[i+j] != needle[j] {
				break
			}
		}
	}
}

// test
func TestStrStr(t *testing.T) {
	cases := []struct {
		haystack string
		needle   string
		expect   int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"aaaaa", "", 0},
		{"", "", 0},
		{"", "a", -1},
	}

	for _, c := range cases {
		if actual := strStr(c.haystack, c.needle); actual != c.expect {
			t.Fatalf("%s, %s, expect %d, actual %d\n", c.haystack, c.needle, c.expect, actual)
		}
	}
}