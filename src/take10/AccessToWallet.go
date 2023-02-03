package main

import (
	"fmt"
	"os"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Function for accessing to wallet
func AccessToWallet(walletFileName string) {

	// Reading wallet file
	fmt.Println("Opening the wallet...")
	fmt.Println("Wallet's file name: ", walletFileName)
	walletFile, err := os.ReadFile(walletFileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The wallet file was read successfuly")
	fmt.Printf("Please enter wallet's password: ")
	// Password request
	var password string
	fmt.Scanf("%s\n", &password)

	checkPasswordResult := CheckPassword(password)
	if checkPasswordResult == false {
		fmt.Println("The password is incorrect")
		os.Exit(1)
	} 

	key, err := keystore.DecryptKey(walletFile, password)
	if err != nil {
		log.Fatal(err)
	}
	privData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("Private:", hexutil.Encode(privData))

	pubData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Public: ", hexutil.Encode(pubData))

	addr := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
	fmt.Println("Add: ", addr)
}
