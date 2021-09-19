// the rangee of all output values is [0, 255]
// the range of the input parameters is H: [0,360], S: [0, 100], B: [0, 100]
// use HSB to Rgb conversion formula
const HSBToRGB = (h, s, b) =>{
    s /= 100;
    b /= 100;
    const k = (n) => (n + h / 60) % 6;
    const f = (n) => b * (1 - s * Math.max(0, Math.min(k(n), 4 - k(n), 1)));
    return [255 * f(5), 255 * f(3), 255 * f(1)]
}

HSBToRGB(18, 81, 99); // [252.45, 109.31084999999996, 47.965499999999984]