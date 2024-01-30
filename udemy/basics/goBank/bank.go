package main

import (
	"fmt"
	fileOPS2 "fourth.com/go-Bank/fileOPS"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, err = fileOPS2.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------")
	}
	fmt.Println("WELCOME TO GO BANK")
	fmt.Println("contact: \n", randomdata.PhoneNumber())
	fmt.Println("First time using, start by depositing or withdrawing money, \n Otherwise system will panic")

	for {
		presentOptions()
		var choice int
		fmt.Println("Your Choice:")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Your Balance Is: ", accountBalance)
			fileOPS2.FileSaver(accountBalance, accountBalanceFile)
		case 2:
			fmt.Println("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount > 0 {
				accountBalance += depositAmount
				fmt.Println("Account Updated, New Balance: ", accountBalance)
				fileOPS2.FileSaver(accountBalance, accountBalanceFile)
			} else {
				fmt.Println("Operation invalid, Cannot Accept Negative Number")
				continue
			}
		case 3:
			fmt.Println("Withdrawal: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)
			if withdrawalAmount < accountBalance && withdrawalAmount > 0 {
				accountBalance -= withdrawalAmount
				fileOPS2.FileSaver(accountBalance, accountBalanceFile)
				fmt.Println("Account Updated, New Balance: ", accountBalance)
			} else {
				fmt.Println("Operation Invalid, Withdrawal Exceeds Bank Balance Or Number Is Negative")
				continue
			}
		default:
			fmt.Println("Have A Good Day!")
			fmt.Println("Thanks For Choosing Our Bank")
			return
		}
	}
}
