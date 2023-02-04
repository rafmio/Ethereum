package main

import (
	"fmt"
	"log"
	"strings"
	"time"
	// "encoding/hex"
	// "math/rand"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func CreateKeyStore(path string) {

	fmt.Println("Creating a keystore...")
	time.Sleep(time.Second * 1)
	key := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	fmt.Printf("The keystore '%s' has just been created\n", strings.Split(path, "/")[1])

	var password string

	fmt.Println("And now let's create an account")
	fmt.Printf("Please enter your password: ")
	fmt.Scanf("%s\n", &password)

	EncryptAES(password)

	account, err := key.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The new account was created")
	fmt.Println("Account address: ", account.Address)
}
