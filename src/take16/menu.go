package main

import (
	"fmt"
	"os"
)

func printMenu() {
	menu = [5]string{
		"1 - Check for wallets",
		"2 - Create a new wallet",
		"3 - Access to a wallet",
		"4 - Delete wallet",
		"0 - Exit the program",
	}
	// for i := 0; i < 2; i++ {
	// 	fmt.Println("\n")
	// }
	fmt.Println("MENU")
	for _, value := range menu {
		fmt.Println(value)
	}
	// for i := 0; i < 2; i++ {
	// 	fmt.Println("\n")
	// }
}

func chooseAction() {
	action := -1
	checkAction := -1
	correctScanMarker := 1

	fmt.Printf("Choose action: ")

	for {
		checkAction, _ = fmt.Scanf("%d\n", &action)
		if checkAction != correctScanMarker || action > len(menu)-1 || action < 0 {
			printMenu()
			fmt.Println("Incorrect input. Pleace enter the menu's item")
		} else {
			switch action {
			case 1:
				CheckForKeystore()
			case 2:
				CreateKeyStore()
			case 3:
				fmt.Println("Access to wallet")
			case 4:
				fmt.Println("Delete wallet")
			case 0:
				fmt.Println("Exiting the program. Bye!")
				os.Exit(0)
			}
			break
		}
	}
}
