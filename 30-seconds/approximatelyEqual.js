

const approximatelyEqual = (a, b, epsilon = 0.00001) => Math.abs(a - b) < epsilon;

test('approximatelyEqual', () => {
    expect(approximatelyEqual(0, 0)).toBe(true);
    expect(approximatelyEqual(0.1, 0.2)).toBe(false);
    expect(approximatelyEqual(0.1, 0.1 + 0.1)).toBe(true);
    expect(approximatelyEqual(0.1, 0.1 + 0.2)).toBe(false);
    });

