// caculates the date of n minutes from the given date
// returning its string representation
// use new Date() to create a date object
// use Date.porototype.getTime() to get the number of milliseconds since January 1, 1970, 00:00:00 UTC
// use Date.prototype.setTime() to set the number of milliseconds since January 1, 1970, 00:00:00 UTC
// use toISOString() to get the string representation
// use new Date(string) to create a date object from the string
// use String.porotype.split() to split the string into an array of substrings
// use String.prototype.replace() to replace the substring

const addMinutesToDate = (date, minutes) => {
  const dateObj = new Date(date);
  const newDate = new Date(dateObj.getTime() + minutes * 60000);
  return newDate.toISOString().split('.')[0].replace('T', ' ');
}

// test
console.log(addMinutesToDate('2020-01-02 00:00:00', 10));
console.log(addMinutesToDate('2020-01-01 00:00:00', 20));
console.log(addMinutesToDate('2020-01-01 00:00:00', -30));