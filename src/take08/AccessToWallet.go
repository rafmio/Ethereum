package main

import (
  "fmt"
)

// Function for accessing to wallet
func AccessToWallet(walletFileName string) {
  keyFileName := "key.txt"
  passwordFileName := "password.txt"

// Reading wallet file
  walletFile, err := os.ReadFile(walletFileName)
  if err != nil {
    log.Fatal(err)
  }

// Password request
  var password string
  fmt.Scanf("%s\n", password)

  // передать в функцию Decrypt
  // сравнить результат с веденным паролем

  // Reading key file
  keyByte, err := os.ReadFile(keyFileName)
  if err != nil {
    log.Fatal(err)
  }

  // Reading password file
  passwordByte, err := os.ReadFile(passwordFileName)
  if err != nil {
    log.Fatal(err)
  }
  passwordStr := string(passwordByte[:])

  // Decoding original password from file
  passwordDecryptedByte := Decrypt(keyByte, passwordStr) // <=======================
  passwordDecryptedStr := string(passwordDecryptedByte)
  if (passwordDecryptedStr != password) {
    fmt.Println("Incorrect password")
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
