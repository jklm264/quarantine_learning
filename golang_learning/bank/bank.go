package main

import (
	"fmt"
	"os"
	"os/exec"
)

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

func Uexit() {
	fmt.Println("See ya!")
	os.Exit(0)
}

func clearscreen(flag int) {
	if flag == 1 {
		for {
			fmt.Println("[When you are finished, hit the c key and enter to continue]")
			var cont string
			fmt.Scanln(&cont)
			if cont == "c" {
				break
			}
		}
	}
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
