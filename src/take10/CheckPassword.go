package main

import (
  "log"
  "os"
)

func CheckPassword (password string) bool {
  keyFileName := "key.txt"
  passwordFileName := "password.txt"

  // Reading key file - []byte(key256HashString) in it
  keyByte, err := os.ReadFile(keyFileName)
  if err != nil {
    log.Fatal(err)
  }

  // Reading password file
  passwordByte, err := os.ReadFile(passwordFileName)
  if err != nil {
    log.Fatal(err)
  }

  decryptedPassword := Decrypt(keyByte, string(passwordByte))

  if password == decryptedPassword {
    return true
  } else {
    return false
  }
}
