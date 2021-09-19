let init = 0;
let sum = [{x: 1}, {x: 2}, {x: 3}].reduce(
    function (pre, current) {
        return pre + current.x;
    },
    init
)

console.log(sum);