package main

import (
  "crypto/sha256"
  "fmt"
)

func main() {
  h := sha256.New()
  h.Write([]byte("this is a password"))
  fmt.Printf("%x", h.Sum(nil))
}
