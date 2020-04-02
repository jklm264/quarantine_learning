// http:app.js

const http = require('http');

const server = http.createServer((req, res) => {
//These routes look very familiar to Flask!
    if (req.url === '/') {
        res.write('Hello world');
        res.end();
    }

    if (req.url === '/courses'){
        res.write(JSON.stringify([1,2,3,4,5]));
        res.end();
    }
});

server.listen(3000);

console.log('Listening on port 3000...'); 
// now visit localhost:3000 and localhost:3000/courses