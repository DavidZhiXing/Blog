//use array.prototype.slice and array.prototype.indexof and string.prototype.split to 
// separate the first row (title row) into values.
//Use String.prototype.split to create a string for each row, then Array.prototype.map() and 
// String.prototype.split(delimeter)  to separate the values in each row.
// use Array.prototype.reduce() to create an object for each row's values, with
// the keys parsed from the title row.
// Omit the second argument, delimiter, to use a default delimeter of ,.
const CSVToJSON = (data, delimiter = ',') => {
    const titles = data.slice(0, data.indexOf('\n')).split(delimiter);
    return data
      .slice(data.indexOf('\n') + 1)
      .split('\n')
      .map(v => {
        const values = v.split(delimiter);
        return titles.reduce(
          (obj, title, index) => ((obj[title] = values[index]), obj),
          {}
        );
      });
  };
  CSVToJSON('col1,col2\na,b\nc,d');
  CSVToJSON('col1;col2\na;b\nc;d', ';');
  