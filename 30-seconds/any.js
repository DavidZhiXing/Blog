const any = (arr, fn) => arr.some(fn);

// test
const arr = [1, 2, 3, 4, 5];
const fn = n => n > 3;
console.log(any(arr, fn)); // true