let flattened = [[0,1],[2,3], [4,5]].reduce(
    (pre, current)=>pre.concat(current),
    []
)

console.log(flattened);