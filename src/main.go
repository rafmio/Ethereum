package main

import (
	"fmt"
	"os"
)

func main() {
	// Hardcoded account number and AES encoded password
	accountNumber := "0003434934"
	AESPassword := "5t6vdd5t58t576rf8547773e4rt759rf8dd5t78t594870gh4"

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
