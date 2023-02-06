package main

import (
	"fmt"
)

func main() {
	keyFileName := "UTC--2023-02-06T07-46-37.469767790Z--f6a4c44f720780d66c94846c9e28dfeb9e7d4a23"
	addr := keyFileName[ (len(keyFileName) - 40) :]
	fmt.Println(addr)
}
