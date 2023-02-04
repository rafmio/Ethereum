package main

import (
	"fmt"
	"os"
)

func CheckForKeystore() {

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
		} else if answerYN == "n" || answerYN == "N" {
			printMenu()
			chooseAction()
		} else if answerYN == "y" || answerYN == "Y" {
			CreateKeyStore()
		} else {
			printMenu()
			chooseAction()
		}
	} else {
		CheckForWallets(files)
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
