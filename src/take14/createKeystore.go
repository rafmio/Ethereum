package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func CreateKeyStore() {

	fmt.Println("Creating a keystore...")
	time.Sleep(time.Second * 1)
	key := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	fmt.Printf("The keystore '%s' has just been created\n", strings.Split(path, "/")[1])
  fmt.Println(key)

	var password string

	fmt.Println("And now let's create an account (wallet)")
	fmt.Printf("Please enter your password: ")
	fmt.Scanf("%s\n", &password)

	// EncryptAES(password)
  fmt.Println("Here we call EncryptAES() function")

	account, err := key.NewAccount(password)
	if err != nil {
		fmt.Println("creating a new account: ", err.Error())
	} else {
    printMenu()
    fmt.Println("The new account was created")
    fmt.Println("Account address: ", account.Address)
  }

  chooseAction()
}
