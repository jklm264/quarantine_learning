var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('bourbon', { title: 'Take A Swig - Bourbon', drink: 'Bourbon' });
});

module.exports = router;
