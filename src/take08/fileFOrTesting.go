package main

import (
    "fmt"
    "os"
)

func main () { 
    var filename string
    files, err := os.ReadDir("./testdir1")
    if err != nil {
	    fmt.Println("Error occured: ", err.Error())
    } else {

	fmt.Println("Check for files: ")
	for index, value := range files {
	    fmt.Printf("%d = %s\n", index + 1, value.Name())
	}
    } 
    filename = files[0].Name()
    fmt.Println(filename)
}
