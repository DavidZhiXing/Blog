let init = 0;
let sum = [{x: 1}, {x: 2}, {x: 3}].reduce(
     (pre, current) =>pre + current.x,
    init
)

console.log(sum);