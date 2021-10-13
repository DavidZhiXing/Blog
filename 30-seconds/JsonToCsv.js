const JsonToCsv = (arr, col, delimiter = ',') =>
    [
        col.join(delimiter),
        ...arr.map(obj => 
                col.reduce(
                    (acc, key) =>
                        `${acc}${!acc.length?'': delimiter}"${!obj[key] ? '' : obj[key]}"`,''
                )
            ),
    ].join('\n');

console.log(JsonToCsv(
    [{ a: 1, b: 2 }, { a: 3, b: 4, c: 5 }, { a: 6 }, { b: 7 }],
    ['a', 'b']
  ))

// use Array.prototype.join(delimiter) to combine all the names in columns to create
// the fist row

// use Array.prototy.map() and  Array.prototype.reduce() to create
// a row for each object, substituting no-existent values with empty strings and only mapping values in columns

// use Array.prototype.join('\n') to combine all rows into a string.

// omit the third argument, delimeter, to use a default delimiter of ,.