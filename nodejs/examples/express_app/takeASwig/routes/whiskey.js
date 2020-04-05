var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('whiskey', { title: 'Take A Swig - Whiskey', drink: 'whiskey'});
});

module.exports = router;
