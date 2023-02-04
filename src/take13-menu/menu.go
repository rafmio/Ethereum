package main

import (
  "fmt"
  "os"
)

func main() {
  printMenu()
  chooseAction()
}

func printMenu() {
  menu := [5]string{
    "1 - Check for wallets",
    "2 - Create a new wallet",
    "3 - Access to a wallet",
    "4 - Delete wallet",
    "0 - Exit the program",
  }
  for i := 0; i < 30; i++ {
    fmt.Println("\n")
  }
  fmt.Println("MENU")
  for _, value := range menu {
    fmt.Println(value)
  }
  for i := 0; i < 9; i++ {
    fmt.Println("\n")
  }
}

func chooseAction(){
  action := 1000
  checkAction := 1000

  fmt.Printf("Choose action: ")

  for ; ; {
    checkAction, _ = fmt.Scanf("%d\n", &action)
    if checkAction != 1 {
      printMenu()
      fmt.Println("Incorrect input. Pleace enter the menu's item")
    } else {
      break
    }
  }
  switch action {
  case 1:
    fmt.Println("Checking for a existing wallets")
  case 0:
    fmt.Println("Exiting the program. Bye!")
    os.Exit(0)
  default:
    fmt.Println("Bla-bbla-bla")

  }
}
