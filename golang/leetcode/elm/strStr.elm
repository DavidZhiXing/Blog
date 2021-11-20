strStr: string, p: string, q: string): string {
    if (p.length > q.length) {
        return this.strStr(strStr, q, p);
    }
    if (p.length === 0) {
        return strStr;
    }
    if (strStr.length < p.length) {
        return "";
    }
    let i = 0;
    while (i <= strStr.length - p.length) {
        if (strStr.substr(i, p.length) === p) {
            return strStr.substr(i, p.length);
        }
        i++;
    }
    return "";
}

// test
console.log(this.strStr("hello", "ll"));
console.log(this.strStr("aaaaa", "bba"));
console.log(this.strStr("aaaaa", ""));
console.log(this.strStr("", ""));
console.log(this.strStr("", "a"));
console.log(this.strStr("a", "a"));
