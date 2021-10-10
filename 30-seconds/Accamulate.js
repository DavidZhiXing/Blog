const accamulator = (...nums) => 
    nums.reduce((acc, n) => [...acc, n + + acc.slice(-1)], [])


console.log(accamulator(1, 2, 3, 4))
console.log(accamulator(...[1,2,3,4]))