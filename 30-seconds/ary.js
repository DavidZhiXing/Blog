const ary = (fn, n) => (...args) => fn(...args.slice(0, n));

test('ary', () => {
    const add = (a, b) => a + b;
    const add3 = ary(add, 3);
    expect(add3(1, 2, 3)).toBe(6);
    expect(add3(1, 2, 3, 4)).toBe(6);
    });