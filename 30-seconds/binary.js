const binary = fn => (a, b) => fn(a, b)

test('binary', () => {
    const add = binary((a, b) => a + b)
    const sub = binary((a, b) => a - b)
    const mul = binary((a, b) => a * b)
    const div = binary((a, b) => a / b)
    
    expect(add(1, 2)).toBe(3)
    expect(sub(1, 2)).toBe(-1)
    expect(mul(2, 3)).toBe(6)
    expect(div(6, 2)).toBe(3)
    })