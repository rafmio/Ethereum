package main

import (
   "crypto/aes"
   "encoding/hex"
   "crypto/sha256"
   "fmt"
   "os"

   "github.com/ethereum/go-ethereum/common/hexutil"
   "github.com/ethereum/go-ethereum/accounts/keystore"
   "github.com/ethereum/go-ethereum/crypto"
  )

/*
Function for:
1. Ecoding the keyphrase to sha256
2. Pass encoded keyphrase and password itsef to the function,
which will encode the string password and the sha256 keyphrase
for encoding the password to AES
*/
func EncodePassword(password string, keyphrase string) string {
   keyFileName := "key.txt"
   passwordFileName := "password.txt"
   // Generating the hash sum (sha256) for keyphrase:
   key256Hash := sha256.New()
   key256Hash.Write([]byte(keyphrase))

   // Storing the hash in the file
   // Write the keyphrase for separate place for security purposes
   err := os.WriteFile(keyFileName, key256Hash, 0644)
   if err != nil {
     log.Fatal(err)
   }

   // Generating the hash sum (AES) for password
   passwordAESHash := Encrypt(key256Hash, password)
   err = os.WriteFile(passwordFileName, passwordAESHash, 0644)
   if err != nil {
     log.Fatal(err)
   }
   return passwordAESHash
}

// Function for generating the hash sum (AES) for password
func Encrypt(key256Hash []byte, password string) string {
   c, err := aes.NewCipher(key256Hash)
   if err != nil {
      fmt.Errorf("NewCipher(%d bytes) = %s", len(key256Hash), err)
      panic(err)
   }
   out := make([]byte, len(password))
   c.Encrypt(out, []byte(password))

   return hex.EncodeToString(out)
}

func Decrypt(keyByte []byte, passwordStr string) {
   ciphertext, _ := hex.DecodeString(passwordStr)
   c, err := aes.NewCipher(keyByte)
   if err != nil {
      fmt.Errorf("NewCipher(%d bytes) = %s", len(keyByte), err)
      panic(err)
   }
   plain := make([]byte, len(ciphertext))
   c.Decrypt(plain, ciphertext)
   s := string(plain[:])
   fmt.Printf("AES Decrypyed Text:  %s\n", s)
}

// Function for accessing to wallet
func AccessToWallet(walletFileName string) {
  keyFileName := "key.txt"
  passwordFileName := "password.txt"

  walletFile, err := os.ReadFile(walletFileName)
  if err != nil {
    log.Fatal(err)
  }

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
  passwordDecrypted := Decrypt(keyByte, passwordStr) // <=======================

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
