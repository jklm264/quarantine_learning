# My first express-generated app

## ToC

- [Installation/Making](#Setup)
- [File Tree Explination](#File-Tree)
- [Static Serves](#Serving-Static-Files)

## Setup

1. `$ npx express-generator` OR to make with [pug (Templates)](../../expressNotes.md) default: `$ express --view=pug myapp`
2. change directory: `$ cd express_generator_app`
3. install dependencies: `$ npm install`
4. run the app: `$ DEBUG=express-generator-app:* npm start`

In the _package.json_ file there exists a line called start which will run `$node ./bin/www` for you.

[src](https://expressjs.com/en/starter/generator.html)

## File Tree

- `/public` - contains assets. These are static files the html references (images, javascript [ON THE WEBPAGE; not nodejs], css)
- `/routes` - site endpoints
- `/views` - representation of each individual html pages
- `/bin` - contains `www` executable
- `app.js` - entry point of whole app

## Serving Static Files

- To serve static files such as images, CSS files, and JavaScript files, use the express.static built-in middleware function in Express.
- For example, use the following code to serve images, CSS files, and JavaScript files in a directory named public: `app.use(express.static('public'))`

## Routes

- How routes should be organized ([src](https://github.com/expressjs/express/blob/4.13.1/examples/route-separation/index.js#L32-L47)):

```javascript
// General

app.get('/', site.index);

// User

app.get('/users', user.list);
app.all('/user/:id/:op?', user.load);
app.get('/user/:id', user.view);
app.get('/user/:id/view', user.view);
app.get('/user/:id/edit', user.edit);
app.put('/user/:id/edit', user.update);

// Posts

app.get('/posts', post.list);
```

## Sources

- <https://expressjs.com/>
- <https://code.visualstudio.com/docs/nodejs/nodejs-tutorial>
