// caculate the date after the given number of business days
// - use Array.from() to convert the arguments object to an array
// - use Array.prototype.reduce() to sum the days
// - use Date.prototype.setDate() to set the date
// if the current date is on a weekend, update it again by adding either one or two days to make it a weekday.
// note: does not take into account holidays

const addWeekDays = (startDate, count) => {
    const days = Array.from(arguments).slice(1);
    return days.reduce((date, days) => {
        date.setDate(date.getDate() + days);
        if (date.getDay() === 0 || date.getDay() === 6) {
        date.setDate(date.getDate() + 2);
        }
        return date;
    }, new Date(startDate));
    };

console.log(addWeekDays(new Date(2018, 0, 1), 1));
// → Wed Feb 01 2018 00:00:00
console.log(addWeekDays(new Date(2018, 0, 1), 2));
// → Fri Feb 02 2018 00:00:00
console.log(addWeekDays(new Date(2018, 0, 1), 3));
// → Mon Feb 04 2018 00:00:00