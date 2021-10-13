const fs = require('fs');

const JsonToFile = (obj, filename) =>
    fs.writeFileSync(`${filename}.json`, JSON.stringify(obj, null, 2));

JsonToFile({ test: 'is passed'}, 'testJsonFile')