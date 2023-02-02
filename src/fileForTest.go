package main

import (
  "os"
  "fmt"
  "log"
)

func main() {
  files, err := os.ReadDir("./take09")
  if err != nil {
    log.Fatal(err)
  }
  for _, value := range files {
    fmt.Println(value.Name())
  }

  var index int
  fmt.Println("Enter index of file for print: ")
  fmt.Scanf("%d", &index)
  fmt.Println()
  fmt.Println(files[index].Name())
}
