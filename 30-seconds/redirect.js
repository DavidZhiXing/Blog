server.get('approvalComplete', function (req, res, next) {
    console.log('User approvalComplete');
    res.send(200);
});

function createRetUrl() {
    let url = require('url');

    let urlObject = {
        protocol: 'http',
        hostname: 'localhost',
        port: configuration.PORT,
        pathname: 'approvalComplete',
    }
    return url.format(urlObject);
};

function wcreateAndSendPayment (session) {
    let returnUrl = createReturnUrl();
    let paymentJson = createPaymentJson(retureUrl);
}
