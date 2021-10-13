const RgbToHex = (r, g, b) =>
    ((r << 16) + (g<<8) + b).toString(16).padStart(6, '0');

console.log(RgbToHex(255, 165, 1))

// Convert given Rgb parameters to hexadecimal string using bitwise left-shift operator
// (<<) and Number.prototype.toString(16).

// Use String.prototype.padStart(6, '0') to get a 6-digit hexadecimal value.