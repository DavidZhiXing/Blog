let names = ['Alice', 'Bob', 'Tiff', 'Bruce', 'Alice'];

let count = names.reduce(
    (pre, name)=>{
        if(name in pre)
            pre[name] ++;
        else
            pre[name] = 1;

        return pre;
    }, {}
)

console.log(count);