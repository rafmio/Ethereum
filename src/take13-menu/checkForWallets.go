package main

import (
	"fmt"
	"os"
	"time"
)

func checkForWallets() {
	path := "./wallet"
	correctScanMarker := 1
	var answerYN string
	// var checkAnswer int

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("reading directory: ", err)
		fmt.Println("It seems to be no keystore yet. Do you want to create a keystore? (y/n): ")
		checkAnswer, _ := fmt.Scanf("%s", &answerYN)
		if checkAnswer != correctScanMarker {
			fmt.Println("Incorrect input. Please check the action")
			printMenu()
			chooseAction()
		} else if answerYN == "n" {
			printMenu()
			chooseAction()
		}
	}

	fmt.Println("Searching for files...")
	time.Sleep(time.Second * 1)
	if len(files) == 0 { // if directory is empty
		fmt.Println("Directory is empty") // TODO: Ask for creating the new account
	} else {
		fmt.Println("The following files were found:")
		for index, value := range files {
			fmt.Printf("index: %d - name: %s\n", index+1, value.Name())
		}
	}

}

// fmt.Printf("Please choose the wallet file's index (1 - %d): ", len(files))
// checkAnswer, _ := fmt.Scanf("%d", &fileIndex)
// if checkAnswer != correctScanMarker || fileIndex < minimalRangeValue || fileIndex > len(files) {
//   fmt.Println("Incorrect input")
//   os.Exit(1)
// }
// }
// walletFileName := path + "/" + files[fileIndex - 1].Name()
// AccessToWallet(walletFileName)
