package main

import (
	"crypto/aes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

var key []byte = []byte("TZPtSIacEJG18IpqQSkTE6luYmnCNKgR")

type AccountEntry struct {
	Wallet   string
	Password []byte
}

var Entries []AccountEntry
var passwordsFile string = "passwords.json"

func EncryptAES(password string, accAddr string) {
	fmt.Println("Now let's encrypt our password to AES")
	ciphBlock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("creating cipher block: ", err.Error())
	}
	encryptedAESPassword := make([]byte, len(password))
	ciphBlock.Encrypt(encryptedAESPassword, []byte(password))

	fmt.Println("encryptedAESPassword ([]byte): ", encryptedAESPassword)
	fmt.Println("encryptedAESPassword (string): ", hex.EncodeToString(encryptedAESPassword))

	EncodeJSON(encryptedAESPassword, accAddr)
}

func EncodeJSON(encryptedAESPassword []byte, accAddr string) {
	fmt.Println("Now let's encode encrypted password to JSON and write to file")
	acc := AccountEntry{Wallet: accAddr, Password: encryptedAESPassword}
	Entries := []AccountEntry{acc}
	json_data, err := json.Marshal(Entries)
	if err != nil {
		fmt.Println("marshalling data: ", err.Error())
	}
	file, err := os.ReadFile(passwordsFile)
	if err != nil {
		fmt.Println("reading file", err.Error())
		fmt.Printf("Creting a new %s file\n", passwordsFile)
		err = os.WriteFile(passwordsFile, json_data, 0644)
		if err != nil {
			fmt.Println("writing to file: ", err.Error())
		}
	} else {
    fmt.Printf("File %s is exist\n", passwordsFile)

    var unmarshalEntries []AccountEntry
    err = json.Unmarshal(file, &unmarshalEntries)
    if err != nil {
      fmt.Println("unmarshaling file: ", err.Error())
    }

    fmt.Println("len(unmarshalEntries): ", len(unmarshalEntries))
    if len(unmarshalEntries) > 0 {
      unmarshalEntries = append(unmarshalEntries, acc)
    }

    jsonDataForWrite, err := json.Marshal(unmarshalEntries)
    if err != nil {
      fmt.Println("marshalling data: ", err.Error())
    }

    err = os.WriteFile(passwordsFile, jsonDataForWrite, 0644)
    if err != nil {
      fmt.Println("writing to file: ", err.Error())
    }

    for i, v := range unmarshalEntries {
      fmt.Println(i, ":", v)
    }
  }
}
