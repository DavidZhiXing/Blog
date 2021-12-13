const bifurcate = (arr, filter) => 
    arr.reduce((acc, val, i) => (acc[filter[i] ? 0 : 1].push(val), acc), [[], []]);

test('bifurcate', () => {
    expect(bifurcate(['beep', 'boop', 'foo', 'bar'], [true, true, false, true])).toEqual([['beep', 'boop', 'bar'], ['foo']]);
    expect(bifurcate([1, 2, 3, 4], [true, false, false, false])).toEqual([[1, 3], [2], [4]]);
    expect(bifurcate(['foo', 'bar', 'baz'], [true, false, true])).toEqual([['foo', 'baz'], ['bar']]);
    expect(bifurcate([1, 2], [true, false])).toEqual([[1], [2]]);
    expect(bifurcate(['foo', 'bar', 'baz', 'qux'], [true, true, true, false])).toEqual([['foo', 'bar', 'baz'], ['qux']]);
    expect(bifurcate([], [])).toEqual([[], []]);
});

const bifurcateBy = (arr, fn) => arr.reduce((acc, val, i) => (acc[fn(val, i) ? 0 : 1].push(val), acc), [[], []]);

test('bifurcateBy', () => {
    expect(bifurcateBy(['beep', 'boop', 'foo', 'bar'], x => x[0] === 'b')).toEqual([['beep', 'boop', ['bar']], ['foo']]);
    expect(bifurcateBy([1, 2, 3, 4], x => x % 2 === 0)).toEqual([[2, 4], [1, 3]]);
    expect(bifurcateBy(['foo', 'bar', 'baz'], x => x[0] === 'b')).toEqual([['bar', 'baz'], ['foo']]);
    expect(bifurcateBy([1, 2], x => x % 2 === 0)).toEqual([[2], [1]]);
    expect(bifurcateBy(['foo', 'bar', 'baz', 'qux'], x => x[0] === 'q')).toEqual([['foo', 'bar', 'baz'], ['qux']]);
    expect(bifurcateBy([], x => x[0] === 'q')).toEqual([[], []]);
});