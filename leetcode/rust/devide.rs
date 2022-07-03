fn devideTwoInteger(dividend: i32, divisor: i32) -> i32 {
    if dividend == 0 {
        return 0;
    }
    if dividend == divisor {
        return 1;
    }
    if divisor == 1 {
        return dividend;
    }
    if divisor == -1 {
        return -dividend;
    }
    if dividend == 1 && divisor == -1 {
        return 0;
    }
    if dividend == -1 && divisor == 1 {
        return 0;
    }
    if dividend == -1 && divisor == -1 {
        return 1;
    }
    let mut dividend = dividend;
    let mut divisor = divisor;
    let mut sign = 1;
    if dividend < 0 {
        sign = -sign;
        dividend = -dividend;
    }
    if divisor < 0 {
        sign = -sign;
        divisor = -divisor;
    }
    let mut result = 0;
    let mut i = 0;
    let mut j = 0;
    let mut nums = vec![];
    while dividend >= divisor {
        nums.push(dividend);
        dividend = dividend - divisor;
        i += 1;
    }
    nums.push(dividend);
    while j < i {
        if nums[j] != nums[j + 1] {
            j += 1;
            nums[j] = nums[j + 1];
        }
    }
    result = nums[j];
    if sign == -1 {
        result = -result;
    }
    result
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_devideTwoInteger() {
        assert_eq!(devideTwoInteger(10, 3), 3);
        assert_eq!(devideTwoInteger(7, -3), -2);
        assert_eq!(devideTwoInteger(1, 1), 1);
        assert_eq!(devideTwoInteger(1, -1), 0);
        assert_eq!(devideTwoInteger(1, -2), 0);
        assert_eq!(devideTwoInteger(1, -3), 0);
        assert_eq!(devideTwoInteger(1, -4), 0);
        assert_eq!(devideTwoInteger(1, -5), 0);
        assert_eq!(devideTwoInteger(1, -6), 0);
        assert_eq!(devideTwoInteger(1, -7), 0);
        assert_eq!(devideTwoInteger(1, -8), 0);
        assert_eq!(devideTwoInteger(1, -9), 0);
        assert_eq!(devideTwoInteger(1, -10), 0);
        assert_eq!(devideTwoInteger(1, -11), 0);
        assert_eq!(devideTwoInteger(1, -12), 0);
        assert_eq!(devideTwoInteger(1, -13), 0);
        assert_eq!(devideTwoInteger(1, -14), 0);
        assert_eq!(devideTwoInteger(1, -15), 0);
        assert_eq!(devideTwoInteger(1, -16), 0);
        assert_eq!(devideTwoInteger(1, -17), 0);
        assert_eq!(devideTwoInteger(1, -18), 0);
        assert_eq!(devideTwoInteger(1, -19), 0);
        assert_eq!(devideTwoInteger(1, -20), 0);
        assert_eq!(devideTwoInteger(1, -21), 0);
        assert_eq!(devideTwoInteger(1, -22), 0);
    }
}
