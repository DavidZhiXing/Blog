const UrlJoin = (...args) =>
    args
    .join('/')
    .replace(/[\/]+/g, '/')
    .replace(/^(.+):\//, '$1://')
    .replace(/^file:/, 'file:/')
    .replace(/\/(\?|&|#[^!])/g, '$1')
    .replace(/\?/g, '&')
    .replace('&', '?')

console.log(UrlJoin('http://www.google.com', 'a', '/b/cd', '?foo=123', '?bar=foo'))

// Use String.prototype.join('/') to combine Url segments.

// use a series of String.prototype.replace() calls with various regexps to normalize the
// resulting url (remove double slashes, add proper slashes for protocol, remove slashes)
// before parameters, combine parameters with '&' and mormalize first parameter delimeter).