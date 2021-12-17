const bind = (fn, context, ...args) => (...moreArgs) => fn.apply(context, args.concat(moreArgs))

// test
const add = (a, b) => a + b
const add2 = bind(add, null, 2)
console.log(add2(3)) // 5