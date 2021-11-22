// Checks if all elements in an array are equal, based on a given mapping function

const allEqualBy = (arr, fn) => {
    return arr.every(item => fn(item) === fn(arr[0]));
}

// test
console.log(allEqualBy([1, 1, 1], (item) => item));
console.log(allEqualBy([1, 2, 3], (item) => item));
console.log(allEqualBy([{a: 1}, {a: 1}, {a: 1}], (item) => item.a));
console.log(allEqualBy([{a: 1}, {a: 2}, {a: 3}], (item) => item.a));