var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('beer', { title: 'Take A Swig - Beer', pp: 'Porter', stout: 'Stout' });
});

module.exports = router;
