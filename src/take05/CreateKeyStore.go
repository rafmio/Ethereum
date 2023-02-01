package main

import (
  "fmt"
  "log"
  "time"
  "strings"
  "os"

  "github.com/ethereum/go-ethereum/accounts/keystore"
)

func CreateKeyStore(path string){
  filename := "key.txt"

  fmt.Println("Creating a keystore...")
  time.Sleep(time.Second * 1)
  key := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
  fmt.Printf("The keystore '%s' has just been created\n", strings.Split(path, "/")[1])

  var password string
  var keyphrase string

  fmt.Println("And now let's create an account")
  fmt.Printf("Please enter your password: ")
  fmt.Scanf("%s\n", &password)
  fmt.Printf("Please enter your keyphrase: ")
  fmt.Scanf("%s\n", &keyphrase)

  // Write the keyphrase for separate place for security purposes
  // byteKeyphrase := []byte(keyphrase)
  // err := os.WriteFile(filename, keyphrase)

  EncodeDecodeCheckPass(password, keystore)

  password := "password"
  account, err := key.NewAccount(password)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(account.Address)
}


// Initial code:
//
// password := "password"
// account, err := key.NewAccount(password)
// if err != nil {
//   log.Fatal(err)
// }
// fmt.Println(account.Address)
