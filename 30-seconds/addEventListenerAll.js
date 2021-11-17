// Attaches an event listener to all the provided targets
// Use Array.prototype.forEach.call(document.querySelectorAll(selector), function(element) { ... });
// to attach an event listener to all elements matching the selector
// and EventTarget.addEventListener(type, listener, useCapture)

const addEventListenerAll = (targets, type, listener, useCapture) => {
    targets.forEach(target => target.addEventListener(type, listener, useCapture));
    };

// test
addEventListenerAll(document.querySelectorAll('.target'), 'click', event => {
    console.log('target clicked');
});
