package main

import (
	"fmt"
	"math/rand"
	"os"
)

const uADMIN = "admin"
const pADMIN = "admin"

func main() {
	info := Bank_info{ //might change to const()
		name:        "Sasuke's Bank",
		motto:       "Dattebayo!!!",
		established: "established since 1969",
		writteningo: "written in Go so, thanks Google",
	}
	fmt.Printf("\t\tWelcome to %v!\n\t\t\t%v\nWe were %v. This was %v!\n\n", info.name, info.motto, info.established, info.writteningo)
	begin()
}

// Basic info about who you are banking with
type Bank_info struct {
	name, motto, established, writteningo string
}

func begin() {

	var opt int
	fmt.Println("Would you like to do?\n\t1) Login\n\t2) Open an Account\n\t3) Exit")
	fmt.Scanln(&opt)
	switch opt {
	case 1:
		//Goto login
		login()
	case 2:
		//Goto create new account
		fmt.Println("> Let's get you over to our account creation department.\n......")
		create_user()
		begin()
	case 3:
		Uexit()
	default:
		fmt.Println("You did not pick a case so...")
		Uexit()
	}
}

// Print welcome message
// 	- print(bank_info.get())
// 1. Login
// 	1. Check balance
// 		- get(balance)
// 	2. Deposit
// 		- addtobalance()
// 	3. Withdraw
// 	4. Close account
// 		- confirm()
// 	5. Seealldata()
// 2. Open Account
// 	1. fill in from customer.go

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
}

func create_user() {
	fmt.Println("...You are now in account creation.")

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

	fmt.Println(fill_user(fname, lname, uname, pword, cs))
	//TODO: Add user to "map"

}

func fill_user(fname string, lname string, uname string, pword string, cs int) *User {
	user := User{
		first_name:  fname,
		last_name:   lname,
		username:    uname,
		password:    pword,
		creditscore: cs,
		acct:        *create_account(uname),
	}

	return &user
}

func Uexit() {
	fmt.Println("See ya!")
	os.Exit(0)
}

func create_account(uname string) *Account {
	acct := Account{
		username:      uname,
		accountnumber: rand.Int31n(10),
		balance:       0.00,
	}

	return &acct
}

func login() {
	fmt.Println("To access our system any further you must log in.")

	for {
		var uname, passwd string
		fmt.Print("\t%: Username: ")
		fmt.Scanf("%s\n", &uname)
		fmt.Print("\t%: Password: ")
		fmt.Scanf("%s\n", &passwd)

		if uname == uADMIN && passwd == pADMIN {
			fmt.Println("\nCredentials verified.")
			login2(uname)
		}

		fmt.Println("\nIncorrect credentails. Please try again.")
	}
}

func login2(uname string) {
	fmt.Printf("Now that you've entered %s, what would you like to do?\n\t1) Check Balance\n\t2) Deposit Money\n\t3) Withdrawl\n\t4) Close Account\n\t5) Exit", uname)
	var opt int
	switch opt {
	case 1:
	case 2:
	case 3:
	case 4:
	case 5:
		Uexit()
	default:
		fmt.Println("You did not pick a case so...")
		Uexit()
	}
}
