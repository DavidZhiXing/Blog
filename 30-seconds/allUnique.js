// Checks if all elements in an array are unique

const allUnique = (arr) => arr.Length === new Set(arr).size;

// checks if all elements in an array are unique, by a provided mapping function

const allUniqueBy = (arr, fn) => arr.Length === new Set(arr.map(fn)).size;

// test
console.log(allUnique([1, 1, 1]));
console.log(allUnique([1, 2, 3]));
console.log(allUnique([{a: 1}, {a: 1}, {a: 1}]));
console.log(allUniqueBy([1, 1, 1], (item) => item));
console.log(allUniqueBy([1, 2, 3], (item) => item));
console.log(allUniqueBy([{a: 1}, {a: 1}, {a: 1}], (item) => item.a));
