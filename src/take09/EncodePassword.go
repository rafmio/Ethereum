package main

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"


)

/*
Function for:
1. Ecoding the keyphrase to sha256
2. Pass encoded keyphrase and password itsef to the function,
which will encode the string password and the sha256 keyphrase
for encoding the password to AES
*/
func EncodePassword(password string, keyphrase string) {
	keyFileName := "key.txt"
	passwordFileName := "password.txt"

	// Generating the hash sum (sha256) for keyphrase:
	key256Hash := sha256.New()          // return empty hash.Hash (interface)
	key256Hash.Write([]byte(keyphrase)) //
	key256HashString := hex.EncodeToString(key256Hash.Sum(nil))

	// Storing the hash in the file
	// Write the keyphrase for separate place for security purposes
	err := os.WriteFile(keyFileName, []byte(key256HashString), 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Generating the hash sum (AES) for password
	passwordAESHashString := Encrypt([]byte(key256HashString), password)
	err = os.WriteFile(passwordFileName, []byte(passwordAESHashString), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
