module hello

go 1.15

// To overwrite commonplace modules
replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000
