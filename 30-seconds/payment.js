function createPaymentJson (returnUrl, cancelUrl){
    return {
        "intent": "sale",
        "payer": {
            "payment_method": "paypal"
        },
        "redirect_urls": {
            "return_url": returnUrl,
            "cancel_url": cancelUrl,
        },
        "transactions": [{
            "item_list": {
                "items": [{
                    "name": "30 Seconds",
                    "sku": "30-seconds",
                    "price": "0.00",
                    "currency": "USD",
                    "quantity": 1
                }]
            },
            "amount": {
                "total": "0.01",
                "currency": "USD"
            },
            "description": "This is the payment transaction description."
        }]
    };
}

function createAndSendPayment (session){
    console.log("createAndSendPayment");

    let paymentJson = createPaymentJson(session.returnUrl, session.cancelUrl);

    paypal.payment.create(paymentJson, function (error, payment) {
        if (error) {
            console.log(error);
            throw error;
        } else {
            for (var index = 0; index < payment.links.length; index++) {
                if (payment.links[index].rel === 'approval_url') {
                    session.paymentUrl = payment.links[index].href;
                    session.save();
                }
            }
        }
    });
}