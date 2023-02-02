package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

// Function for generating the hash sum (AES) for password
func Encrypt(key256Hash []byte, password string) string {
	c, err := aes.NewCipher(key256Hash)
	if err != nil {
		fmt.Errorf("NewCipher(%d bytes) = %s", len(key256Hash), err)
		panic(err)
	}
	out := make([]byte, len(password))
	c.Encrypt(out, []byte(password))

	return hex.EncodeToString(out)
}

// Function for decrypt password-key - reverse
func Decrypt(keyByte []byte, passwordStr string) string {
	// Convert password from string to []byte:
	passwordByte, _ := hex.DecodeString(passwordStr)

	// Convert key from []byte to cipher.Block interface:
	keyByteCipherBlock, err := aes.NewCipher(keyByte)
	if err != nil {
		fmt.Errorf("NewCipher(%d bytes) = %s", len(keyByte), err)
		panic(err)
	}

	plain := make([]byte, len(passwordByte))

	// Decrypt password-key
	keyByteCipherBlock.Decrypt(plain, passwordByte)
	s := string(plain[:])
	fmt.Printf("AES Decrypyed Text:  %s\n", s)

	return string(plain)
}
