//Before moving 'database' to excel, start with map/dict
package main

import (
	"fmt"
)

type Account struct {
	username      string
	accountnumber int32
	balance       float32
}

// func get_balance() {
// 	return balance
// }

// init(){
// 	balance = 0
// }

//For testing
func Create_account() {
	fmt.Println("You are now in customer_data.go.")
}
