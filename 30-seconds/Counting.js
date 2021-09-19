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

let people = [
    {name: 'Alice', age: 21},
    { name: 'Max', age: 20},
    { name: 'Jane', age: 20}
];

const groudby = (obj, property) => obj.reduce(
    (acc, item) => {
        let key = item[property];
        if (!acc[key]) {
            acc[key] = []
        }
        acc[key].push(item)
        return acc
    }, {}
)

console.log(groudby(people, 'age'))