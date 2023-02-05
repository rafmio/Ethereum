package main

import (
  "fmt"
  "crypto/aes"
  "encoding/hex"
  "os"
  "encoding/json"
)

var key []byte = []byte("TZPtSIacEJG18IpqQSkTE6luYmnCNKgR")
type AccountEntry struct {
  Wallet string
  Password []byte
}
var Entries []AccountEntry
var passwordsFile string = "passwords.json"

func EncryptAES(password string, accAddr string) {
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

func EncodeJSON(encryptedAESPassword string, accAddr string) {
  acc := AccountEntry{Wallet: accAddr, Password: accAddr}
}
