var builder = require('botbuilder');

bot.use({
    botbuilder: function (session, next) {
        if (session.message.text === 'go home') {
            if (session.sessionState.callstack && session.sessionState.callstack.length) {
                session.cancelDialog(0, '/');
            }
        } else {
            session.beginDialog('/');
        }
    }
});

// feeback/ index.js

var feebackBot = (function () {
    var _intents;
    var _lib = new builder.Library('feeback');

    _lib.dialog('hello', [
        function (session, results) {
            session.send('Hello...');
            session.endDialog();
        }
    ]);

    function createIntents() {
        return _lib;
    };

    function initialize (locale) {
        intents = new builder.IntentDialog();
        intents.matches(/^hello/i, 'hello');
        return intents;
    };

    return {
        createIntents: createIntents,
        initialize: initialize
    };
})();

mudule.exports = feebackBot;

// declare the library
bot.library(childbot.createLibrary());

// Initialize the library and receive intents
var botIntents = selectedBot.initialize(locale);

// create a new dialog with the received intents
bot.dialog('/', botIntents);

// enter the dialog when needed 
session.beginDialog('/');

// automatically loading bot

var bots = [];
var botDirectories = getDirectories('./bots');
for (var dirIdx in botDirectories) {
    var dirName = botDirectories[dirIdx];
    var childBot = require(path.join(__dirname, 'bots', dirName, 'index.js'));
    bots.push(childBot);
    bot.library(childBot.createLibrary());
}

// creating and starting a new bot

var botNames = bots.map(function (bot) { return bot.getName(session); });
builder.Prompts.choice(session, 'Which bot would you like to use?', botNames);

var selectedBot = bots.find(function (bot) { return bot.getName(session) === session.dialogData.choice; });
var locale = session.preferredLocale();
var botKey = 'Locale-' + locale + '-' + selectedBot.getName(session);
var localeIntents = selectedBot.initialize(locale);

bot.dialog(botKey, localeIntents);
session.send( selectedBot.welcomeMessage(session) );
session.beginDialog(botKey);