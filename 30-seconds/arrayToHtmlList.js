const arrayToHtmlList = (array) => {
  return array.map((item) => `<li>${item}</li>`).join('');
};

const arrayToHtmlList = (array, listID) => {
  document.querySelector(`#$(listID)`).innerHTML = array.map((item) => `<li>${item}</li>`).join('');
};

const list = ['apple', 'banana', 'orange'];
const listID = 'myList';

console.log(arrayToHtmlList(list, listID)); // <li>apple</li><li>banana</li><li>orange</li>

test('arrayToHtmlList', () => {
  expect(arrayToHtmlList(['a', 'b', 'c'])).toBe('<li>a</li><li>b</li><li>c</li>');
});


