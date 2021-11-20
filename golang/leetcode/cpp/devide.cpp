#include <stdio.h>

int devideTwoInteger(int dividend, int divisor) {
    if (divisor == 0) {
        return INT_MAX;
    }
    if (dividend == 0) {
        return 0;
    }
    if (dividend == INT_MIN && divisor == -1) {
        return INT_MAX;
    }
    int sign = (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) ? 1 : -1;
    long long dvd = labs(dividend);
    long long dvs = labs(divisor);
    int res = 0;
    while (dvd >= dvs) {
        long long temp = dvs, multiple = 1;
        while (dvd >= (temp << 1)) {
            temp <<= 1;
            multiple <<= 1;
        }
        dvd -= temp;
        res += multiple;
    }
    return sign == 1 ? res : -res;
}

long long labs(long long x) {
    return x > 0 ? x : -x;
}

const int INT_MAX = 2147483647;
const int INT_MIN = -2147483648;

// test
int main() {
    int dividend = -2147483648;
    int divisor = -1;
    int res = devideTwoInteger(dividend, divisor);
    printf("%d\n", res);
    return 0;
}