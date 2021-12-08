// create an array of n-tuples of consecutive elements

const aperture = (n, arr) => 
    n > arr.length ? [] : arr.slice(0, n).map((_, i) => arr.slice(i, i + n));

// test
const arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
console.log(aperture(3, arr)); // [[1, 2, 3], [2, 3, 4], [3, 4, 5], [4, 5, 6], [5, 6, 7], [6, 7, 8], [7, 8, 9], [8, 9, 10]]
console.log(aperture(4, arr)); // [[1, 2, 3, 4], [2, 3, 4, 5], [3, 4, 5, 6], [4, 5, 6, 7], [5, 6, 7, 8], [6, 7, 8, 9], [7, 8, 9, 10]]