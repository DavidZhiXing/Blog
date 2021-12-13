const arrayToCSV = (array, delimiter = ',') => {
  return array.map(item => item.join(delimiter)).join('\n');
};

test('arrayToCSV', () => {
    expect(arrayToCSV([['a', 'b'], ['c', 'd']])).toBe('a,b\nc,d');
    expect(arrayToCSV([['a', 'b'], ['c', 'd']], ';')).toBe('a;b\nc;d');
    });

