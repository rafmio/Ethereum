package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func UserInteract() {
	path := "./wallet"
	var answer string

	// Starting directory search
	_, err := os.Stat(path)

	fmt.Printf("Searching for %s keystore directory...\n", strings.Split(path, "/")[1])
	time.Sleep(time.Second * 1)

	if err != nil {
		fmt.Println("Error occured: ", err.Error())
		fmt.Println("No keystore was found")
		fmt.Printf("Do you want to create a keystore? (y/n): ") // ask for creating keystore
		fmt.Scanf("%s\n", &answer)
		if answer == "y" {

			fmt.Println("Let's create a keystore") // call create createKeyStore
			CreateKeyStore(path)

		} else if answer == "n" {
			fmt.Println("Let's finish the jorney")
			os.Exit(0)
		} else {
			fmt.Println("Unknown answer. The programm will be terminated")
			os.Exit(1)
		}
	} else {
		fmt.Printf("The %s keystore directory found\n", strings.Split(path, "/")[1])
	}

		fmt.Printf("Do you want to accsess the new wallet? (y/n): ")
		fmt.Scanf("%s\n", &answer)
		if answer == "y" {
			ChooseWallet(path)
		} else if answer == "n" {
			fmt.Println("Let's finish the jorney")
			os.Exit(0)
		} else {
			fmt.Println("Unknown answer. The programm will be terminated")
			os.Exit(1)
		}

}
