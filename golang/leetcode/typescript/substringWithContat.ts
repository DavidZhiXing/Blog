let SubStringWithContat = function(s: string | any[], t: string | any[]) {
    let len = s.length;
    let len2 = t.length;
    if (len < len2) {
        return 0;
    }
    let map = new Map();
    for (let i = 0; i < len2; i++) {
        let c = t[i];
        if (map.has(c)) {
            map.set(c, map.get(c) + 1);
        } else {
            map.set(c, 1);
        }
    }
    let start = 0;
    let end = 0;
    let count = 0;
    let charCount = 0;
    while (end < len) {
        let c = s[end];
        if (map.has(c)) {
            let v = map.get(c);
            if (v > 0) {
                v--;
                map.set(c, v);
                charCount++;
            }
        }
        if (charCount == len2) {
            count++;
        }
        while (charCount == len2) {
            let c = s[start];
            if (map.has(c)) {
                let v = map.get(c);
                if (v >= 0) {
                    v++;
                    map.set(c, v);
                    charCount--;
                }
            }
            start++;
        }
        end++;
    }
    return count;
};

// test
let s = "barfoothefoobarman";
let t = "foo";
let r = SubStringWithContat(s, t);
console.log(r);
