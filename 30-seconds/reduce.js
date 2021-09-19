const array1 = [1, 2, 3, 4];
const reducer = (pre, current) => pre + current;

console.log(array1.reduce(reducer))

console.log(array1.reduce(reducer, 5))