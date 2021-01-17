# First Project

Make play Payment system
- backend db for payment audits (just use excel with https://github.com/szyhf/go-excel)
- first as cli then gui
- have admin acct that can see all
- should use a currency and password module

## Payment System

Query with switch
    1a. Check status
    2a. Deposit
    3a. Withdraw
    4a. Open Account
    5a. Close Account
    6a. See all data (GDPR)

Structs:

users {
    first_name = str
    last_name = str
    creditscore = int
    username = str
    password = passwd
}

account {
    accountnumber = long
    balance = double
}

// Basic info about who you are banking with
bank_info {
    name = ""
    motto = ""
    established = "established since 1969"
    writteningo = "Written in Go. Thanks Google"
}