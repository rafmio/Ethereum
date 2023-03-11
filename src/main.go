package main

import (
	"fmt"
	"os"
)

func main() {
	// Hardcoded account number and AES encoded password
	accountNumber := "1114949863"
	AESPassword := "24evf7ff88t576rf854ere34r4t759rf8dd5t78t5948707t4"

	// Assign to struct variable values of account number and password
	var newAccount = AccountEntry{AccNumber: accountNumber, Password: AESPassword}

	// Connecting to DB
	conn, err := EstablishConnectionDB()
	if err != nil {
		fmt.Println("connection to DB failed. Exiting...")
		os.Exit(1)
	}

	InsertAccount(conn, newAccount)
}
