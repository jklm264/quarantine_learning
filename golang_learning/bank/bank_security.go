package main

import "fmt"

const (
	uADMIN = "admin"
	pADMIN = "admin"
)

// Can also make like this:: var account_db = make(map[string]User, 10)
// var account_db = map[string]User){}
// db[uname] = fname,lname,uname,pword,cs,Account,status
var account_db = make(map[string]User, 10)

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
			fmt.Println("I will now deposit money.")
			var funds float32
			fmt.Printf("How much money would you like to deposit? ")
			fmt.Scanln(&funds)
			// TODO: Add error check
			deposit_funds(uname, funds)
			get_account_balance(uname)
			clearscreen(1)
		case 3:
			fmt.Println("I will now withdraw money.")
		case 4:
			fmt.Println("I will now show you all your account data.")
		case 5:
			fmt.Println("I will now close your account.")
			//Simply change status to false
		case 6:
			Uexit()
		default:
			fmt.Println("You did not pick a valid case so let's try this again.")
			clearscreen(1)
		}
	}
}

func deposit_funds(uname string, funds float32) {
	fmt.Printf("TESTING: %v\n", funds)
	fmt.Println(account_db[uname].acct.balance)
	account_db[uname].acct.balance = funds
}
