function executePayment (parameters) {
    console.log('Excuting payment');

    // Appended to the URL by PayPal during the approval step
    let paymentId = parameters.paymentId;
    let payerId = parameters.PayerID;

    // Generate the sample payment execution Json that paypal requires
    let addressEncoded = decodeURIComponent(parameters.addressEncoded);
    let address = JSON.parse(addressEncoded);

    // Finally, execute the payment, and tell the user that we got their payment
    paypal.payment.execute(paymentId, { "payer_id": payerId }, function (error, payment) {
        if (error) {
            console.log(error);
            throw error;
        } else {
            console.log("Payment executed successfully");
            respondToUser(payment, address);
        }
    });
}

function respondToUser (payment, address) {
    let message = builder.Message().address(address).text('Your payment!');

    // asks the bot to send the message we built up above to the user
    bot.send(message.ToMessage());
}

bot.dialog('listFines', [
    function (session, args) {
        console.log('Listing fines');
        session.send('you have 1 outstanding fine: $10');
        session.send('Parking Fine Violation');
        builder.Prompts.Choice(session, 'Do you want to pay the fine?', ['Pay Fine', 'Cancel']);
    },
    function (session, result, next) {
        let choice = result.response.entity;

        if (choice === 'Pay Fine') {
            createAndSendPayment(session);
        } else {
            session.send('Fine payment cancelled');
        }
    }
])