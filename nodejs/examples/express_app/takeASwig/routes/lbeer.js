var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('lbeer', { title: 'Take A Swig - Light Beer'});
});

module.exports = router;
