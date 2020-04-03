//web_app.js

// Start with $npm init
// $npm install express --save will save in the dependencies list
//	Doing --no-save will not add it to the list and only install Express temporarily

// Express vs Express-generator: https://stackoverflow.com/questions/41311348/express-or-express-generator-do-i-need-both
// Installation Instructions: https://expressjs.com/en/starter/installing.html

const express = require('express')
const app = express()
const port = 3000

app.get('/', (req, res) => {
	return res.send('Hello world!')
})

app.listen(port, () => console.log(`Example app listening at http://localhost:${port}`))