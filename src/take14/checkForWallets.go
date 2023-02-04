package main

import (
	"fmt"
	"time"
	"os"
)

func CheckForWallets(files []os.DirEntry) {
	fmt.Println("Searching for files...")
	time.Sleep(time.Second * 1)
	if len(files) == 0 { // if directory is empty
		fmt.Println("Directory is empty")
	} else {
		fmt.Println("The following files were found:")
		printMenu()
		for index, value := range files {
			fmt.Printf("index: %d - name: %s\n", index + 1, value.Name())
		}
	}
}
