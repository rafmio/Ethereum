package main

import (
  "fmt"
)

func DeleteWallet() {
  fmt.Println("Let's delete the wallet. Pleace choose wallet you want to delete: ")
  files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("reading directory: ", err)
		fmt.Println("It seems you have no wallets at all. It's nothing to delete")
  }

  
  fmt.Println("Let's delete the wallet. Pleace choose wallet you want to delete: ")
}
