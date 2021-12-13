const attempt = (fn, ...args) => {
  try {
    return fn(...args);
  } catch (e) {
    return e instanceof Error ? e : new Error(e);
  }
};

test('attempt', () => {
    expect(attempt(() => {
        throw new Error('foo');
    })).toEqual(new Error('foo'));
    expect(attempt(() => {
        throw 'bar';
    })).toEqual(new Error('bar'));
    expect(attempt(() => {
        throw {
        toString: () => 'baz'
        };
    })).toEqual(new Error('baz'));
    expect(attempt(() => {
        throw {
        toString: () => {
            throw new Error('qux');
        }
        };
    })).toEqual(new Error('qux'));
    expect(attempt(() => {
        throw {
        toString: () => {
            throw 'quux';
        }
        };
    })).toEqual(new Error('quux'));
    expect(attempt(() => {
        throw {
        toString: () => {
            throw {
            toString: () => 'quuz'
            };
        }
        };
    })).toEqual(new Error('quuz'));
    });