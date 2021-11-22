// check if the provided predicate function returns true for all elements in a collection

const all = (collection, predicate) => {
    for (let i = 0; i < collection.length; i++) {
        if (!predicate(collection[i])) {
        return false;
        }
    }
    return true;
    }

const all2= (collection, fn = Boolean) => {
    collection.every(fn);
}

// test
console.log(all([1, 2, 3, 4], n => n < 10));
console.log(all([1, 2, 3, 4], n => n < 20));