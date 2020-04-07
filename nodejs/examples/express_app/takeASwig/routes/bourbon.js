var express = require('express');
var router = express.Router();


/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('bourbon', { title: 'Take A Swig - Bourbon', drink: 'Bourbon' });
});

router.post('/', (req,res,next) => {
  res.json({message: 'hooray! Welcome to our api!' });
  var ip = req.connection.remoteAddress;
  console.log(`Recieved POST request from ${ip}`);
})

module.exports = router;
