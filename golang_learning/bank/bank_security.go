package main

import "fmt"

const (
	uADMIN = "admin"
	pADMIN = "admin"
)

// Can also make like this:: var account_db = make(map[string]User, 10)
// var account_db = map[string]User){}
// db[uname] = fname,lname,uname,pword,cs,Account,status
var account_db = make(map[string]*User, 10)

// Basic info about who you are banking with
type Bank_info struct {
	name, motto, established, writteningo string
}

func login() {
	fmt.Println("To access our system any further you must log in.")

	for {
		var uname, passwd string
		fmt.Print("\t%: Username: ")
		fmt.Scanf("%s\n", &uname)
		fmt.Print("\t%: Password: ")
		fmt.Scanf("%s\n", &passwd)

		if (uname == uADMIN && passwd == pADMIN) || (passwd == account_db[uname].password) {
			fmt.Println("\nCredentials verified.")
			login2(uname)
			break
		}

		fmt.Println("\nIncorrect credentails. Please try again.")
	}
}

func login2(uname string) {
	// Utilize the uname as your lookup and then 'get' from there.
	clearscreen(0)

	for {
		var opt int
		fmt.Printf("Now that you've entered, %s, what would you like to do?\n\t1) Check Balance\n\t2) Deposit Money\n\t3) Withdrawl\n\t4) Get all Data\n\t5) Close Account\n\t6) Exit\n", uname)
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			fmt.Println("I will check your balance now.")
			get_account_balance(uname)
			clearscreen(1)
		case 2:
			var funds float32
			fmt.Println("I will now deposit money.")
			fmt.Printf("How much money would you like to deposit? ")
			fmt.Scanln(&funds)
			// TODO: Add error checks
			deposit_funds(uname, funds)
			get_account_balance(uname)
			clearscreen(1)
		case 3:
			var funds float32
			fmt.Println("I will now withdraw money.")
			fmt.Printf("How much money would you like to withdraw? ")
			fmt.Scanln(&funds)
			withdraw_funds(uname, funds)
			get_account_balance(uname)
			clearscreen(1)
		case 4:
			fmt.Println("I will now show you all your account data.")
			show_account(uname)
		case 5:
			fmt.Println("I will now close your account.")
			close_account(uname)
			clearscreen(1)
		case 6:
			Uexit()
		default:
			fmt.Println("You did not pick a valid case so let's try this again.")
			clearscreen(1)
		}
	}
}

func deposit_funds(uname string, funds float32) {
	account_db[uname].acct.balance += funds
}

func withdraw_funds(uname string, funds float32) {
	account_db[uname].acct.balance -= funds
}

func close_account(uname string) {
	account_db[uname].status = false
}

func show_account(uname string) {
	db := account_db[uname]
	fmt.Printf("\n\tFull Name: %s %s\n\tCredit Score: %d\n\tAccount Status: %v\n\tUsername: %s\n\tAccount Number: %v\n\tAccount Balance: %v\n", db.first_name, db.last_name, db.creditscore, db.status, db.acct.username, db.acct.accountnumber, db.acct.balance)
}
