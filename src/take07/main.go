package main

import (
  "fmt"
  "os"
  "time"
  "strings"


)

func main() {
  path := "./wallet"

  _, err := os.Stat(path)

  fmt.Printf("Searching for %s keystore directory...\n", strings.Split(path, "/")[1])
  time.Sleep(time.Second * 1)

  if err != nil {
    fmt.Println("Error occured: ", err.Error())
    fmt.Println("No keystore was found")
    fmt.Printf("Do you want to create a keystore? (y/n): ")   // ask for creating keystore
    var answer string
    fmt.Scanf("%s\n", &answer)
    if answer == "y" {

      fmt.Println("Let's create a keystore")               // call create createKeyStore
      CreateKeyStore(path)

    } else if answer == "n" {
      fmt.Println("Let's finish the jorney")
      os.Exit(1)
    } else {
      fmt.Println("Unknown answer. The programm will be terminated")
      os.Exit(1)
    }
  } else {
    fmt.Printf("The %s keystore directory found\n", strings.Split(path, "/")[1])
    fmt.Printf("Searching for files in %s directory...\n", strings.Split(path, "/")[1])

    // If keystore directory was found - searching for files in the directory:
    files, err := os.ReadDir(path)
    fmt.Println("Searching for account files...")
    time.Sleep(time.Second * 1)
    var walletFileNameByte []byte
    var checkWalletIndex int
    var walletFileName string

    if err != nil {
      fmt.Println("Error occured: ", err.Error())
      } else {                                    // if files in path directory was found
        fmt.Println("files: ", files)             // change to loop range []files
        if len(files) == 0 {                      // if directory is empty
          fmt.Println("Directory is empty")
        } else {
          fmt.Println("Check the account: ")
          for index, value := range files {
            fmt.Printf("%d - %s\n", index + 1, value)
          }

          fmt.Printf("Enter the index of the wallet list: ")
          fmt.Scanf("%d\n", checkWalletIndex)
          walletFileNameByte = files[checkWalletIndex]
          walletFileName = string(walletFileNameByte)
        }
        // If wallet files is exist - ask for password
        AccessToWallet(walletFileName)
      }
  }

}
