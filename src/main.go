package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	// Hardcoded account number and AES encoded password
	accountNumber := "8437948432"
	AESPassword := "245t57ff88tdfdrf854edfr4r4t75sdf8dd5t78t5948708r4"

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

func menu() {
	var action int
	fmt.Println("Menu: ")
	fmt.Printf("1. Create account\n2. Change password\n")
	fmt.Printf("Choose action (1-2): ")
	fmt.Scanf("%d", &action)

	switch action {
	case 1:
		// pass
	case 2:
		// pass
	default:
		fmt.Println("Incorrect input. Please choose action (1-2)")
	}
}

func generateAcc() Entry {
	var accountNumer [10]string
	for _, digit := range accountNubmer {
		randNum := rand.Intn(9)
		digit := strconv.Itoa(randNum)
	}
}
