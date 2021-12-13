const average = (...numbers) => numbers.reduce((a, b) => a + b, 0) / numbers.length;

test('average', () => {
    expect(average(1, 2, 3, 4)).toBe(2.5);
    expect(average(1, 2, 3, 4, 5)).toBe(3);
    expect(average(...[1, 2, 3, 4, 5])).toBe(3);
    });

const averageBy = (...numbers, fn) => numbers.map(typeof fn === 'function' ? fn : val => val[fn])
    .reduce((a, b) => a + b, 0) / numbers.length;

test('averageBy', () => {
    expect(averageBy([{n: 1}, {n: 2}, {n: 3}, {n: 4}, {n: 5}], o=>o.n)).toBe(3);
    expect(averageBy([{n: 1}, {n: 2}, {n: 3}, {n: 4}, {n: 5}], 'n')).toBe(3);
});