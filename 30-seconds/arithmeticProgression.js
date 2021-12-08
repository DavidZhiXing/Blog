const arithmeticProgression = (n, lim) => {
  Array.from({ length: Math.ceil(lim / n) }, (_, i) => i * n);
};

test("arithmeticProgression", () => {
  expect(arithmeticProgression(2, 10)).toEqual([0, 2, 4, 6, 8, 10]);
  expect(arithmeticProgression(3, 10)).toEqual([0, 3, 6, 9]);
  expect(arithmeticProgression(4, 10)).toEqual([0, 4, 8]);
  expect(arithmeticProgression(5, 10)).toEqual([0, 5]);
  expect(arithmeticProgression(6, 10)).toEqual([0]);
});