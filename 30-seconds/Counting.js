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

let friends = [{
    name: 'Anna',
    books: ['Bible', 'Harry Potter'],
    age: 21
  }, {
    name: 'Bob',
    books: ['War and peace', 'Romeo and Juliet'],
    age: 26
  }, {
    name: 'Alice',
    books: ['The Lord of the Rings', 'The Shining'],
    age: 18
  }]

  let allbooks = friends.reduce(
      (pre, current) => [...pre, ...current.books], ['Alphabet']
  )

  console.log(allbooks)