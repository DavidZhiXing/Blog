// add mutiple event listeners with the same handler to an element
// use Array.prototype.forEach to loop through the array of events and EnventTarget.addEventListener to add each event
// with an assigned callback function

const addMultipleEvents = (element, types, listeners, options, useCapture) => {
    types.forEach(type => {
        element.addEventListener(type, listeners, options, useCapture);
    });
    };

// test
addMultipleEvents(document.querySelector('#test'), ['click', 'mouseover'], (e) => {
    console.log(e.type);
}, false);
