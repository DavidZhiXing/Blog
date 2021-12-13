const atob = (str) => Buffer.from(str, 'base64').toString('binary');

test('atob', () => {
    expect(atob('SGVsbG8gV29ybGQ=')).toBe('Hello World');
    });