const assertValidKeys = (obj, validKeys) => obj.keys().every(key => validKeys.includes(key));

// test
const obj = {
    name: 'Jon',
    age: 30
};

console.log(assertValidKeys(obj, ['name', 'age'])); // true
console.log(assertValidKeys(obj, ['name', 'age', 'location'])); // false