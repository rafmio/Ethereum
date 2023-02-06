package main

import (
  "crypto/aes"
  "encoding/json"
  "crypto/cipher"
  "strings"
  // "crypto/rand"
  "fmt"
  "os"
  // "io"
)

var key []byte = []byte("TZPtSIacEJG18IpqQSkTE6luYmnCNKgR")

type AccountEntry struct {
	Wallet   string
	Password []byte
}

var Entries []AccountEntry
var passwordsFile string = "passwords.json"

func DecryptAES(addr string) string {
  fmt.Println("Let's decrypt AES-password for wallet's address: ", addr)
  encodedAESPassword := DecodeJSON(addr)

  block, err := aes.NewCipher(key)
  if err != nil {
    fmt.Println("creating new cipher block: ", err.Error())
  }

  gcmDecrypt, err := cipher.NewGCM(block)
  if err != nil {
    fmt.Println("wrape the AES in Galios Counter Mode: ", err.Error())
  }

  nonceSize := gcmDecrypt.NonceSize()
  if len(encodedAESPassword) < nonceSize {
    fmt.Println(err.Error())
  }

  nonce, encryptedMessage := encodedAESPassword[:nonceSize], encodedAESPassword[nonceSize :]
  decryptedPassword, err := gcmDecrypt.Open(nil, nonce, encryptedMessage, nil)
  if err != nil {
    fmt.Println(err)
  }

  return string(decryptedPassword)
}

func DecodeJSON(addr string) []byte {
  addrFormatted := "0x" + addr
  var encodedAESPassword []byte
  fmt.Println("Decoding JSON...")
  file, err := os.ReadFile(passwordsFile)
	if err != nil {
		fmt.Println("reading file", err.Error())
  }

  var unmarshalEntries []AccountEntry
  err = json.Unmarshal(file, &unmarshalEntries)
    if err != nil {
      fmt.Println("unmarshaling file: ", err.Error())
    }

  for _, value := range unmarshalEntries {
    fmt.Println("Wallet: ", strings.ToLower(value.Wallet), " Password: ", value.Password)
  }
  for _, value := range unmarshalEntries {
    if strings.ToLower(value.Wallet) == addrFormatted {
      encodedAESPassword = value.Password
      fmt.Println("Password was found")
    }
  }
  fmt.Println("JSON file was decoded")
  fmt.Println("encodedAESPassword - ", encodedAESPassword)
  return encodedAESPassword
}
