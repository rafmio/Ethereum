package main

import (
    "fmt"
    "os"
)

func main () {
	_, err := os.Stat("./wasllet")
	if err != nil {
	    fmt.Println("Error occured: ", err.Error())
	}

}
