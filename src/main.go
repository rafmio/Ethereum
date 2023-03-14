package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	menu()
}

func menu() {
	var action int
	fmt.Println("Menu: ")
	fmt.Printf("1. Create account\n2. Change password\n")
	fmt.Printf("Choose action (1-2): ")
	fmt.Scanf("%d", &action)

	switch action {
	case 1:
		entry := generateAcc()
		fmt.Println("The new account and password was generated")
		fmt.Printf("Your account number: %s\nYour password: %s", entry.AccNumber, entry.Password)

		conn, err := EstablishConnectionDB()
		if err != nil {
			fmt.Println("connection to DB failed. Exiting...")
			os.Exit(1)
		}

		err = InsertAccount(conn, entry)
		if err != nil {
			fmt.Println("adding the entry to database failed. Exiting...")
		}

	case 2:
		// pass
		// Extract list of account number and output it for user's choose
		conn, err := EstablishConnectionDB()
		if err != nil {
			fmt.Println("connection to DB failed. Exiting...")
			os.Exit(1)
		}

		rows, err := conn.Query(context.Background(), "select accnumber from accounts")
		if err != nil {
			fmt.Println("extracting list of accounts: ", err.Error())
			os.Exit(1)
		} else {
			fmt.Printf("Type of rows: %T\n", rows)
		}

	default:
		fmt.Println("Incorrect input. Please choose action (1-2)")
		os.Exit(1)
	}
}

func generateAcc() AccountEntry {
	// Generating random account number (10 digits)
	var accountNumber [10]string
	for i := 0; i < 10; i++ {
		rnd := rand.Intn(9)
		accountNumber[i] = strconv.Itoa(rnd)
	}

	var accountNumberStr string
	for i := 0; i < 10; i++ {
		accountNumberStr += accountNumber[i]
	}

	// Generating password (30 sybmols)
	var charset = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+"
	var pswd [30]string

	for i := 0; i < 30; i++ {
		randInx := rand.Intn(len(charset) - 1)
		pswd[i] = string(charset[randInx])
	}

	var pswdStr string
	for i := 0; i < 30; i++ {
		pswdStr += pswd[i]
	}

	entry := AccountEntry{
		AccNumber: accountNumberStr,
		Password:  pswdStr,
	}

	return entry
}
