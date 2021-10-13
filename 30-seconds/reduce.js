// pre 为累加器， current 为当前元素， 源数组为array1, 索引为[1, 4]

const array1 = [1, 2, 3, 4];
const reducer = (pre, current) => pre + current;

console.log(array1.reduce(reducer))

console.log(array1.reduce(reducer, 5))