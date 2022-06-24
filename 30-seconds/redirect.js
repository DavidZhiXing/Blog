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

// show notification

const appId = 'electron-windows-notification';
const {ToastNotification} = require('electron-windows-notification');

let notification = new ToastNotification({
    appId: appId,
    template: '<toast><visual><binding template="ToastText01"><text id="1">Hello</text></binding></visual></toast>',
    strings: ['Hello'],
});

notification.on('activated', () => {console.log('activated')});
notification.show();
