package main

import (
  "crypto/sha256"
  "fmt"
)

func main() {
  h := sha256.New()
  h.Write([]byte("this is a password"))
  fmt.Printf("%x\n", h.Sum(nil))
  fmt.Printf("Type of h: %T\n", h)

  retVal := h.Sum(nil)
  fmt.Println(retVal)
  fmt.Printf("Type of retVal: %T\n", retVal)
  fmt.Println(string(retVal))
}
