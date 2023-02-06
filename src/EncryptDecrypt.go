package main

import (
	"fmt"
	// "math/rand"
	"io/ioutil"
	"io"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func EncryptAES(password string) {
	passwordFileName := "password.txt"
	// keyFileName := "key.txt"
	key := []byte("TZPtSIacEJG18IpqQSkTE6luYmnCNKgR")
	// err := os.WriteFile(keyFileName, key, 0644)
	// if err != nil {
	// 	fmt.Println("writing to file: ", err.Error())
	// 	os.Exit(1)
	// }

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
	err = ioutil.WriteFile(passwordFileName, gcm.Seal(nonce, nonce, []byte(password), nil), 0644)
	if err != nil {
		fmt.Println("writing to file: ", err.Error())
	}
}

func DecryptAES(key []byte, passwordFileName string) string {
  ciphertext, err := ioutil.ReadFile(passwordFileName)
  if err != nil {
    fmt.Println(err)
  }
  c, err := aes.NewCipher(key)
  if err != nil {
    fmt.Println(err)
  }

  gcmDecrypt, err := cipher.NewGCM(c)
  if err != nil {
    fmt.Println(err)
  }

  nonceSize := gcmDecrypt.NonceSize()
  if len(ciphertext) < nonceSize {
    fmt.Println(err)
  }

  nonce, encryptedMessage := ciphertext[:nonceSize], ciphertext[nonceSize :]
  plaintext, err := gcmDecrypt.Open(nil, nonce, encryptedMessage, nil)
  if err != nil {
    fmt.Println(err)
  }

	return string(plaintext)
}
