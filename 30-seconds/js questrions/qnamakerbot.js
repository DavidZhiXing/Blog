'use strict'

var request = require('request');
var qnastring = require('qnamaker-string');
var restify = require('restify');
var botbuilder = require('botbuilder');

var config = require('nconf').env().argv().file({ file: 'config.json' });

function createChatConnector() {
    let opts = {
        appId: config.get('MicrosoftAppId'),
        appPassword: config.get('MicrosoftAppPassword')
    };
    
    let chatConnector = new botbuilder.ChatConnector(opts);

    let server = restify.createServer();

    server.post('/api/messages', chatConnector.listen());

    server.get('//?.*/', restify.serveStatic({
        directory: __dirname + '/html', 
        default: 'index.html',
    }));
    server.listen(3978, function () {
console.log('%s listening to %s', server.name, server.url);
});

return chatConnector;
}

function qna(q, cb){
    // here where we pass anything the user typed along to the QnA Maker service

    q = querystring.escape(q);
    request(config.get('qnaUri') + q, function (error, response, body) {
        if (error) {
            console.log(error);
            cb(error);
        } else if (response.statusCode !== 200) {
            console.log('Status:', response.statusCode);
            console.log('Headers:', JSON.stringify(response.headers));
            console.log('Response:', body);
            cb(body);
        } else {
            cb(body);
        }
    });
}

function initialBot(bot) {
    // someone set up us the bot

    // Use IntentDialog the watch for messages that match regexes
    let intents = new botbuilder.IntentDialog();

    // any message not matching the previous intents ends up here
    intents.onDefault((session, args, next) => {

        // just throw everything into the qna service
        qna(session.message.text, (err, result) => {
            if (err) {
                session.send(err);
            } else {
                // the qna returns a JSON: { answer: "...", score: 0.5 }
                // where score is a confidence the answer matches the question.
                // Advanced implementations might log lower scored questions and
                // answers since they tend to indicate either gaps in the faq content
                session.send(JSON.parse(result).answer);
            }
        });
    });
    bot.dialog('/', intents);
}

function main() {
    initialBot(new botbuilder.UniversalBot(createChatConnector()));
}

if (module.id === require.main.id) {
    main();
}