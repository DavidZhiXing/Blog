const devide = (dividend, divisor) => {
    if (dividend === 0) return 0
    if (divisor === 1) return dividend
    if (divisor === -1) return -dividend
    if (divisor === -2147483648) {
        if (dividend === -2147483648) return 2147483647
        if (dividend === 2147483647) return -2147483648
    }
    let sign = 1
    if (dividend < 0) {
        sign = -sign
        dividend = -dividend
    }
    if (divisor < 0) {
        sign = -sign
        divisor = -divisor
    }
    let result = 0
    while (dividend >= divisor) {
        let temp = divisor
        let count = 1
        while (dividend >= temp) {
            dividend -= temp
            result += count
            count *= 2
            temp *= 2
        }
    }
    return sign * result
}

// test
console.log(devide(10, 3));
console.log(devide(7, -3));
console.log(devide(7, -2));
console.log(devide(7, -1));
console.log(devide(7, -2147483648));
console.log(devide(2147483647, -2147483648));
console.log(devide(2147483647, -1));
console.log(devide(2147483647, -2));
console.log(devide(2147483647, -3));
