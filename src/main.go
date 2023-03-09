package main

import "fmt"

func main() {
	// Hardcoded account number and AES encoded password
	accountNumber := "1002344848"
	AESPassword := "7t6tr57t7tr49d76r9td8783faa8t7598f89gh482f9487t6r"

	// Assign to struct variable values of account number and password
	var newAccount = AccountEntry{AccNumber: accountNumber, Password: AESPassword}

	// Connecting to DB
	ConnectDB(newAccount)

	fmt.Println("The entry account number with password was added to DB")
}
