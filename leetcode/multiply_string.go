package leetcode

// https://leetcode.com/problems/multiply-strings/
// 
// Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.
//
// Example 1:
//
//	input: num1 = "2", num2 = "3"
//	output: "6"
//
// Example 2:
//
//	input: num1 = "123", num2 = "456"
//	output: "56088"
//
// Note:
//
//	The length of both num1 and num2 is < 110.
//	Both num1 and num2 contain only digits 0-9.
//	Both num1 and num2 do not contain any leading zero, except the number 0 itself.

func multiplyString(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 {
		return ""
	}
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if num1 == "1" {
		return num2
	}
	if num2 == "1" {
		return num1
	}
	var result []byte
	for i := len(num2) - 1; i >= 0; i-- {
		var carry byte
		var tmp []byte
		for j := len(num1) - 1; j >= 0; j-- {
			val := (num2[i] - '0') * (num1[j] - '0') + carry
			carry = val / 10
			tmp = append(tmp, byte(val%10)+'0')
		}
		if carry > 0 {
			tmp = append(tmp, carry+'0')
		}
		for k := 0; k < len(num2)-1-i; k++ {
			tmp = append(tmp, '0')
		}
		result = add(result, tmp)
	}
	return string(result)
}

func add(a []byte, b []byte) []byte {
	if len(a) < len(b) {
		a, b = b, a
	}
	var carry byte
	for i := 0; i < len(b); i++ {
		val := a[i] + b[i] + carry
		carry = val / 10
		a[i] = val % 10
	}
	for i := len(b); carry > 0; i++ {
		val := a[i] + carry
		carry = val / 10
		a[i] = val % 10
	}
	return a
}

type testCase struct {
	num1 string
	num2 string
	want string
}


func main() {
	cases := []testCase{
		{
			num1: "2",
			num2: "3",
			want: "6",
		},
		{
			num1: "123",
			num2: "456",
			want: "56088",
		},
		}
	for _, test := range testCases {
		result := multiplyString(test.num1, test.num2)
		if result != test.want {
			panic(fmt.Sprintf("%s * %s = %s, expect %s", test.num1, test.num2, result, test.want))
		}
	}
}