bot.dialog('/', function (session, args){
    session.beginDialog('/stubbling');
})

// simple threee step dialog to list 'fines' that a user has received, and allow a user to 'pay' them

bot.dialog('/listfines', [
    function (session, args) {
        session.send("You have received the following fines:")
        session.send("$100 for speeding")
        builder.Prompts.choice(session, "Would you like to pay the fine?", "Yes|No")
    },
    function (session, results, next) {
        let choice = results.response.entity;
        if (choice == "Yes") {
            session.send("You have paid the fine")
        } else {
            session.send("You have not paid the fine")
        }
    },
    function (session, results) {
        session.send("Thank you for using the parking bot")
    }
])
;
