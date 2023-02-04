package main

import (
  "fmt"
)

func enterPassword() string {
  // text := []byte("Mandalorian is currently the best DisneyPlus show")
  var text string
  fmt.Printf("Enter password: ")
  fmt.Scanf("%s\n", &text)

  return text
}
