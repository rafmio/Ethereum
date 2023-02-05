package main

import (
  "fmt"
  "encoding/json"
  "os"
)

type AccountEntry struct {
  Wallet string
  Password string
}

var Entries []AccountEntry

var passwordsFile string = "passwords.json"

func main() {
  acc1 := AccountEntry{Wallet: "3jr3f93hr9fu3943", Password: "passwordnew"}
  file, err := os.ReadFile(passwordsFile)
  if err != nil {
    fmt.Println("reading file", err.Error())
    fmt.Printf("Creting a new %s file\n", passwordsFile)
    Entries := []AccountEntry{acc1}
    json_data, err := json.Marshal(Entries)
    if err != nil {
      fmt.Println("marshalling data: ", err.Error())
    }
    err = os.WriteFile(passwordsFile, json_data, 0644)
    if err != nil {
      fmt.Println("writing to file: ", err.Error())
    }
  } else {
    fmt.Printf("File %s is exist\n", passwordsFile)

    var unmarshalEntries []AccountEntry
    err = json.Unmarshal(file, &unmarshalEntries)
    if err != nil {
      fmt.Println("unmarshaling file: ", err.Error())
    }

    fmt.Println("len(unmarshalEntries): ", len(unmarshalEntries))
    if len(unmarshalEntries) > 0 {
      unmarshalEntries = append(unmarshalEntries, acc1)
    }

    jsonDataForWrite, err := json.Marshal(unmarshalEntries)
    if err != nil {
      fmt.Println("marshalling data: ", err.Error())
    }

    err = os.WriteFile(passwordsFile, jsonDataForWrite, 0644)
    if err != nil {
      fmt.Println("writing to file: ", err.Error())
    }

    for i, v := range unmarshalEntries {
      fmt.Println(i, ":", v)
    }
  }

}
