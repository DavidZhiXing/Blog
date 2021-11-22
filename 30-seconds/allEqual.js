// Checks if all elements in an array are equal

const allEqual = (arr) => {
    return arr.every(item => item === arr[0]);
}

// Tests
console.log(allEqual([1, 1, 1]));
console.log(allEqual([1, 2, 3]));