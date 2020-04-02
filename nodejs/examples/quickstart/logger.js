// logger.js
const EventEmitter = require('events'); //Class is camel-case convention

var url = 'https://jmussman.net'

class Logger extends EventEmitter {
    // Don't need function keyword 
    // in class because its now considered a method.
    log(message) { 
        // Send an HTTP request
        console.log(message);
    
        // Raise an event
        // Note 'this' keyword
        this.emit('messageLogged', {id: 1, url: 'http://'});
    }
}

module.exports = Logger;