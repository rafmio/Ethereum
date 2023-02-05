package main

import (
	"fmt"
	// "os"
	"io"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

var passwordFileName string = "passwords.json"

func main() {
  password := "password"
	key := []byte("TZPtSIacEJG18IpqQSkTE6luYmnCNKgR")

	cphr, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	gcm, err := cipher.NewGCM(cphr)
	if err != nil {
		fmt.Println(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

  pwd := gcm.Seal(nonce, nonce, []byte(password), nil)
  fmt.Printf("Type of pwd : %T\n", pwd)

	// err = os.WriteFile(passwordFileName, gcm.Seal(nonce, nonce, []byte(password), nil), 0644)
	// if err != nil {
	// 	fmt.Println("writing to file: ", err.Error())
	// }
}
