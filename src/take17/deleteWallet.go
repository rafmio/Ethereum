package main

import (
  "fmt"
  "os"
  "encoding/json"
  "strings"
)

func DeleteWallet() {
  var fileIndex int
  var addr string
  correctScanMarker := 1
  minimalRangeValue := 1

  fmt.Println()
  fmt.Println("Let's delete the wallet")
  files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("reading directory: ", err)
		fmt.Println("It seems you have no wallets at all. It's nothing to delete")
    printMenu()
    chooseAction()
  } else if len(files) == 0 {
    fmt.Println("It seems you have no wallets at all. It's nothing to delete")
    printMenu()
    chooseAction()
  } else {
    fmt.Println("Pleace choose wallet to delete")
    fmt.Printf("Index N  Address\n")
    for index, value := range files {
      fileName := value.Name()
      addr = fileName[(len(fileName) - 40) : ]
      // fmt.Printf("index: %d - address: %s\n", index + 1, addr)
      fmt.Printf("%d \t %s\n", index + 1, addr)
    }

    fmt.Printf("Please choose the wallet file's index to delete (1 - %d): ", len(files))

    for {
      checkAnswer, _ := fmt.Scanf("%d", &fileIndex)
      if checkAnswer != correctScanMarker || fileIndex < minimalRangeValue || fileIndex > len(files) {
        fmt.Printf("Incorrect input. Please choose the wallet file's index (1 - %d): ", len(files))
        } else {
          walletFileName := path + "/" + files[fileIndex - 1].Name()
          fmt.Println("Deleting: ", walletFileName)

          fmt.Println("Deleting password")
          DeletePassword(addr)
          fmt.Println("Password deleted")
        }
      }
  }
}

func DeletePassword(addr string) {
  // var sliceForJSON []AccountEntry
  var indexForDelete int

  addrFormatted := "0x" + addr
  // fmt.Println("addrFormatted: ", addrFormatted)

  file, err := os.ReadFile(passwordsFile)
	if err != nil {
		fmt.Println("reading file", err.Error())
  }

  var unmarshalEntries []AccountEntry
  err = json.Unmarshal(file, &unmarshalEntries)
   if err != nil {
     fmt.Println("unmarshaling file: ", err.Error())
   }

   for i, value := range unmarshalEntries {
     if strings.ToLower(value.Wallet) == addrFormatted {
       indexForDelete = i
       fmt.Println("indexForDelete: ", indexForDelete)
       // break
     //   fmt.Println("GOTCHA!")
     //   fmt.Println(i, " - ", strings.ToLower(value.Wallet))
     //   fmt.Println("Desired value: ", strings.ToLower(value.Wallet))
     } else {
       fmt.Printf("No any entries for address %s\n", addrFormatted)
     }
     // tmpSlice1 := unmarshalEntries[:(i - 1)]
     // fmt.Println(tmpSlice1)
     // tmpSlice2 := unmarshalEntries[(i + 1):]
     // fmt.Println(tmpSlice2)
     // sliceForJSON = append(tmpSlice1[:], tmpSlice2[:] ...)
   }


   // fmt.Println(sliceForJSON)

   // jsonDataForWrite, err := json.Marshal(sliceForJSON)
   // if err != nil {
   //   fmt.Println("unmarshaling data: ", err.Error())
   // }


   // err = os.WriteFile(passwordsFile, jsonDataForWrite, 0644)
   // if err != nil {
   //   fmt.Println("writing to file: ", err.Error())
   // }

}
