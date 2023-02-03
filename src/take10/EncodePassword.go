package main

import (
	// "crypto/sha256"
	// "encoding/hex"
	"fmt"
	"log"
	"os"
	"time"
	"math/rand"
)

/*
Function for:
1. Ecoding the keyphrase to sha256
2. Pass encoded keyphrase and password itsef to the function,
which will encode the string password and the sha256 keyphrase
for encoding the password to AES
*/
func EncodePassword(password string) {
	keyFileName := "key.txt"
	passwordFileName := "password.txt"


	fmt.Println("Generating random keyphrase... ")
	time.Sleep(time.Second * 1)

	// Generating random keyphrase:
	keyphraseByte := make([]byte, 16)
	if lenkp, err := rand.Read(keyphraseByte);err != nil {
		panic(err.Error())

	} else {
		fmt.Printf("Keyphrase generated successfully, len - %d bytes\n", lenkp)
	}

	// Storing the hash in the file
	// Write the keyphrase for separate place for security purposes
	err := os.WriteFile(keyFileName, keyphraseByte, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Generating the hash sum (AES) for password
	passwordAESHashString := EncryptAES(keyphraseByte, password)
	err = os.WriteFile(passwordFileName, []byte(passwordAESHashString), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
