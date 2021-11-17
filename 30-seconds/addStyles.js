// Add the provided styles to the given element.
// use Object.assign to merge the styles and 
// ElementCSSInlineStyle.style to add the styles to the element.

const addStyles = (element, styles) => {
  Object.assign(element.style, styles);
};

//test
const element = document.querySelector('#test');
addStyles(element, {
    color: 'red',
    backgroundColor: 'blue',
    fontSize: '20px'
});