// https://appdividend.com/2019/12/05/golang-encryption-decryption-example-aes-encryption-in-go/
package main
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {
	// text := []byte("Mandalorian is currently the best DisneyPlus show")
	key := []byte("TZPtSIacEJG18IpqQSkTE6luYmnCNKgR")
	text := []byte(enterPassword())

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
	err = ioutil.WriteFile("app.txt", gcm.Seal(nonce, nonce, text, nil), 0644)
	if err != nil {
		fmt.Println(err)
	}

  decryptAES(key, text)
}

func decryptAES(key []byte, text []byte) {
  ciphertext, err := ioutil.ReadFile("app.txt")
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

  fmt.Printf("Type of plaintext is : %T\n", plaintext)
  fmt.Println(string(plaintext))

  if string(plaintext) != string(text) {
    fmt.Println("Fail!")
  } else {
    fmt.Println("GOTCHA!!!")
  }
}
