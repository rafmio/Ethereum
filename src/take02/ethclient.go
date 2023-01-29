package main

import (
  "fmt"
  "os"
  "log"
)

func main() {
  // Читаем директорию с аккаунтами
  files, err := os.ReadDir("./accounts")
  if err != nil {
    log.Fatal(err)
  }

  // Проверяем на наличие файлов
  if (len(files) == 0) {
    fmt.Println("It seems there are no local accounts here")
  } else {
    fmt.Printf("I found %d account files\n", len(files))
    fmt.Println("Here are they: ")
    for _, f := range files {
      fmt.Println(f.Name())
    }
    fmt.Println("-----------------------------------------")
  }

}


func createAccount()

func logInToAccount()
