class LeetCode:
    def divide(self, dividend, divisor):
        if divisor == 0:
            return None
        if dividend == 0:
            return 0
        if dividend == -2147483648 and divisor == -1:
            return 2147483647
        sign = 1
        if (dividend > 0 and divisor < 0) or (dividend < 0 and divisor > 0):
            sign = -1
        dividend = abs(dividend)
        divisor = abs(divisor)
        res = 0
        while dividend >= divisor:
            temp = divisor
            i = 1
            while dividend >= temp:
                dividend -= temp
                res += i
                i <<= 1
                temp <<= 1
        return res * sign

# Test
if __name__ == '__main__':
    leetcode = LeetCode()
    print(leetcode.divide(10, 3))
    print(leetcode.divide(7, -3))
    print(leetcode.divide(1, 1))
    print(leetcode.divide(1, -1))
    print(leetcode.divide(1, -2147483648))
    print(leetcode.divide(-2147483648, -1))
    print(leetcode.divide(-2147483648, 1))
    print(leetcode.divide(-2147483648, 2))
    print(leetcode.divide(-2147483648, -2))
    print(leetcode.divide(-2147483648, -2147483648))