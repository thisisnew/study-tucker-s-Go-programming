package main

import (
	"fmt"
	"tuckers/appB/exB1/bankaccount"
)

func main() {
	account := bankaccount.NewAccount()
	account.Deposit(1000)
	account.Withdraw(500)
	fmt.Println(account.Balance())
}
