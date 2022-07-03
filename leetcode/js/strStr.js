const strStr = function(haystack, needle) {
    if (needle.length === 0) {
        return 0;
    }
    if (haystack.length === 0) {
        return -1;
    }
    if (needle.length > haystack.length) {
        return -1;
    }
    let i = 0;
    let j = 0;
    while (i < haystack.length) {
        if (haystack[i] === needle[j]) {
            if (j === needle.length - 1) {
                return i - j;
            }
            i++;
            j++;
        } else {
            i = i - j + 1;
            j = 0;
        }
    }
    return -1;
};

// test
console.log(strStr("hello", "ll"));
console.log(strStr("aaaaa", "bba"));
console.log(strStr("aaaaa", ""));
console.log(strStr("", ""));
console.log(strStr("", "a"));
console.log(strStr("a", "a"));
console.log(strStr("a", "b"));
console.log(strStr("a", ""));
console.log(strStr("", "a"));