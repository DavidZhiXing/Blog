// the string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"
// Write the code that will take a string and make this conversion given a number of rows:
// string convert(string text, int nRows);
// convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".
//
// Tags: string
// Difficulty: Medium

// exampel:
// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"

// exampel2:
// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I

// exampel3:
// Input: s = "A", numRows = 1
// Output: "A"

// constraint:
// 1 <= s.length <= 1000
// s consists of English letters (lower-case and upper-case), ' ', and '\n'.
// 1 <= numRows <= 1000

func zigzag(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var res string
	var row = 0
	var down = true
	for i := 0; i < len(s); i++ {
		if row == 0 || row == numRows-1 {
			down = !down
		}
		res += string(s[i])
		if down {
			row++
		} else {
			row--
		}
	}
	return res
}