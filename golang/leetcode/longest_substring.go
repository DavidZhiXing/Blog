// Given a string, find the length of the longest substring without repeating characters.
// For example, the longest substring without repeating letters for "abcabcbb" is "abc", which the length is 3.
// For "bbbbb" the longest substring is "b", with the length of 1.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var max int
	var start int
	var end int
	var charMap = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := charMap[s[i]]; ok {
			if charMap[s[i]] >= start {
				start = charMap[s[i]] + 1
			}
		}
		charMap[s[i]] = i
		if i-start+1 > max {
			max = i - start + 1
		}
	}
	return max
}

func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	var max int
	var charMap = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		charMap[s[i]] = i
		if len(charMap) > max {
			max = len(charMap)
		}
	}
	return max
}