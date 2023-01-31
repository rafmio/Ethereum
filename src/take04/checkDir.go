package main

import (
  "fmt"
  "os"
  "time"
  "strings"

  "github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
  path := "./testdir2"
  _, err := os.Stat(path)

  fmt.Printf("Searching for %s keystore directory...\n", strings.Split(path, "/")[1])
  time.Sleep(time.Second * 1)

  if err != nil {
    fmt.Println("Error occured: ", err.Error())
    fmt.Println("No keystore was found")
    fmt.Printf("Do you want to create a keystore? (y/n): ")
    var answer string
    fmt.Scanf("%s\n", &answer)
    if answer == "y" {
      fmt.Println("Let's create the keystore")
    } else if answer == "n" {
      fmt.Println("Let's finish the jorney")
      os.Exit(1)
    } else {
      fmt.Println("Unknown answer. The programm will be terminated")
      os.Exit(1)
    }
  } else {
    fmt.Printf("The %s keystore directory found\n", strings.Split(path, "/")[1])
  }

  files, err := os.ReadDir(path)
  fmt.Println("Searching for account files...")
  time.Sleep(time.Second * 1)

  if err != nil {
    fmt.Println("Error occured: ", err.Error())
  } else {
    fmt.Println("files: ", files)
      if len(files) == 0 {
        fmt.Println("Yes!")
      }
    }
}

func createKeyStore(path string){
  key := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
  password := "password"
  account, err := key.NewAccount(password)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(account.Address)

}
