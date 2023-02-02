package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func ChooseWallet(path string) {
	var fileIndex int
	correctScanMarker := int(1)
	minimalRangeValue := 1

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Searching for files...")
	time.Sleep(time.Second * 1)
	if len(files) == 0 { // if directory is empty
		fmt.Println("Directory is empty") // TODO: Ask for creating the new account
	} else {
		fmt.Println("The following files were found:")
		for index, value := range files {
			fmt.Printf("index: %d - name: %s\n", index + 1, value.Name())
		}
		fmt.Printf("Please choose the wallet file's index (1 - %d): ", len(files))
		checkAnswer, _ := fmt.Scanf("%d\n", &fileIndex)
		if checkAnswer != correctScanMarker || fileIndex < minimalRangeValue || fileIndex > len(files) {
			fmt.Println("Incorrect input")
			os.Exit(1)
		} else {
			AccessToWallet(files[fileIndex].Name())
		}
	}
}
