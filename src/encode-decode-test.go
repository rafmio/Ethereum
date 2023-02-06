package main

import (
  "fmt"
  "crypto/aes"
  "encoding/hex"
  "os"
)

var key []byte = []byte("TZPtSIacEJG18IpqQSkTE6luYmnCNKgR")

func main() {
  // address := "230930hd209uh2309uhr2093h4093h20"
  password := "password00000002"
  EncryptAES(password)
  DecryptAES()
}

func EncryptAES(password string) {
  // key := []byte("TZPtSIacEJG18IpqQSkTE6luYmnCNKgR")
  ciphBlock, err := aes.NewCipher(key)
  if err != nil {
    fmt.Println("creating cipher block: ", err.Error())
  }
  encryptedAESPassword := make([]byte, len(password))
  ciphBlock.Encrypt(encryptedAESPassword, []byte(password))

  fmt.Println("encryptedAESPassword ([]byte): ", encryptedAESPassword)
  fmt.Println("encryptedAESPassword (string): ", hex.EncodeToString(encryptedAESPassword))

  err = os.WriteFile("password.txt", encryptedAESPassword, 0644)
  if err != nil {
    fmt.Println("writing to file: ", err.Error())
  }
}

func DecryptAES() {
  // key := []byte("TZPtSIacEJG18IpqQSkTE6luYmnCNKgR")
  encryptedAESPassword, err := os.ReadFile("password.txt")
  if err != nil {
    fmt.Println("reading file: ", err.Error())
  }

  ciphBlock, err := aes.NewCipher(key)
  if err != nil {
    fmt.Println("creating cipher block: ", err.Error())
  }

  decryptedAESPassword := make([]byte, len(encryptedAESPassword))
  ciphBlock.Decrypt(decryptedAESPassword, encryptedAESPassword)

  fmt.Println("decryptedAESPassword: ", string(decryptedAESPassword))
}
