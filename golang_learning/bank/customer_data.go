//Before moving 'database' to excel, start with map/dict
package main

import (
	"fmt"
	"math/rand"
)

//----------------------------------------------------------------//
//Before moving 'database' to excel, start with map/dict
type Account struct {
	username      string
	accountnumber int32
	balance       float32
}

type User struct {
	first_name, last_name, username string
	creditscore                     int
	password                        string //password
	acct                            Account
	status                          bool
}

func create_account(uname string) Account {
	acct := Account{
		username:      uname,
		accountnumber: rand.Int31n(10),
		balance:       0.00,
	}

	return acct
}

func create_user() {
	fmt.Println("... You are now in account creation.")

	var fname, lname, uname, pword string
	var cs int
	for {
		fmt.Print("\n%: What is your name? <First> <Last> ")
		fmt.Scanf("%s %s\n", &fname, &lname)
		fmt.Print("%: And what would you like your username to be? ")
		fmt.Scanf("%s\n", &uname)
		fmt.Print("%: Password? ")
		fmt.Scanf("%s\n", &pword)
		fmt.Print("%: How about your credit score? ")
		fmt.Scanf("%d\n", &cs)

		fmt.Println("\nPlease confirm the following information:")
		fmt.Printf("\tFirst name: %s\n", fname)
		fmt.Printf("\tLast name: %s\n", lname)
		fmt.Printf("\tUsername: %s\n", uname)
		fmt.Printf("\tPassword: %s\n", "***********")
		fmt.Printf("\tCredit Score: %v\n", cs)

		var correct rune
		fmt.Print("\nIs all this correct? (Y/N) ")
		fmt.Scanf("%c\n", &correct)
		if correct == 'Y' || correct == 'y' {
			break
		}
	}

	// Add to db
	user := fill_user(fname, lname, uname, pword, cs)
	account_db[uname] = &user
}

func fill_user(fname string, lname string, uname string, pword string, cs int) User {
	user := User{
		first_name:  fname,
		last_name:   lname,
		username:    uname,
		password:    pword,
		creditscore: cs,
		acct:        create_account(uname),
		status:      true,
	}

	return user
}

func get_account_balance(uname string) {
	fmt.Printf("> You have $%v in your account, %s\n", account_db[uname].acct.balance, uname)
}
