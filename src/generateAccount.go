package main

import (
	"math/rand"
	"strconv"
)

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
