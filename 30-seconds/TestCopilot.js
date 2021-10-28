function calculateDaysBetweenDates(startDate, endDate) {
    // Validate input
    if (endDate < startDate)
        return 0;

    // Calculate days between dates
    var millisecondsPerDay = 86400 * 1000; // Day in milliseconds
    startDate.setHours(0, 0, 0, 1);  // Start just after midnight
    endDate.setHours(23, 59, 59, 999);  // End just before midnight
    var diff = endDate - startDate;  // Milliseconds between datetime objects
    var days = Math.ceil(diff / millisecondsPerDay);

    // Subtract two weekend days for every week in between
    var weeks = Math.floor(days / 7);
    days = days - (weeks * 2);

    // Handle special cases
    var startDay = startDate.getDay();
    var endDay = endDate.getDay();

    // Remove weekend not previously removed.
    if (startDay - endDay > 1)
        days = days - 2;
    // Remove start day if span starts on Sunday but ends before Saturday
    if (startDay == 0 && endDay != 6)
        days = days - 1;
    // Remove end day if span ends on Saturday but starts after Sunday
    if (endDay == 6 && startDay != 0)
        days = days - 1;

    return days;
}

function readFile(fileName) {
    var file = File(fileName);
    file.open('r');
    var fileContents = file.read();
    file.close();
    return fileContents;
}

function cvsToArray(csvString) {
    var array = csvString.split(',');
    for (var i = 0; i < array.length; i++) {
        array[i] = array[i].trim();
    }
    return array;
}

function cvsToJson(csvString) {
    var array = csvToArray(csvString);
    var json = {};
    for (var i = 0; i < array.length; i++) {
        var keyValue = array[i].split(':');
        json[keyValue[0]] = keyValue[1];
    }
    return json;
}

function getJsonFromFile(fileName) {
    var fileContents = readFile(fileName);
    return JSON.parse(fileContents);
}

//what is my life purpose?
function getLifePurpose() {
    var lifePurpose = getJsonFromFile('30-seconds/lifePurpose.json');
    var lifePurposeValue = lifePurpose.lifePurpose;
    return lifePurposeValue;
}