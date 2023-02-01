package main

import (
   "crypto/aes"
   "encoding/hex"
   "crypto/sha256"
   "fmt"
   "os"
  )

// получить пароль и ключ
// зашифровать ключ в sha256
// передасть ключ и хэш от ключа в фукнцию AES-шифрования пароля
// получить AES-хэш пароля
// сохранить AES-хэш в файл

func EncodeDecodeCheckPass(password string, keyphrase string) {
   keyFileName := "key.txt"
   passwordFileName := "password.txt"
   // Generating the hash sum (sha256) for keyphrase:
   keyHash := sha256.New()
   keyHash.Write([]byte(keyphrase))

   // Storing the hash in the file
   // Write the keyphrase for separate place for security purposes
   err := os.WriteFile(keyFileName, keyHash, 0644)
   if err != nil {
     log.Fatal(err)
   }

   // Generating the hash sum (AES) for password
   passwordAESHash := Encrypt(keyHash, password)
   err = os.WriteFile(passwordFileName, passwordAESHash, 0644)
   if err != nil {
     log.Fatal(err)
   }
   // Write down the result to the file
   // key := "myverystrongpasswordo32bitlength"
   // plainText := "Hello 8gwifi.org"
   // ct := Encrypt([]byte(key),plainText)
   // fmt.Printf("Original Text:  %s\n",plainText)
   // fmt.Printf("AES Encrypted Text:  %s\n", ct)
   // Decrypt([]byte(key),ct)
}

func Encrypt(key []byte, plaintext string) string {
   c, err := aes.NewCipher(key)
   if err != nil {
      fmt.Errorf("NewCipher(%d bytes) = %s", len(key), err)
      panic(err)
   }
   out := make([]byte, len(plaintext))
   c.Encrypt(out, []byte(plaintext))

   return hex.EncodeToString(out)
}

func Decrypt(key []byte, ct string) {
   ciphertext, _ := hex.DecodeString(ct)
   c, err := aes.NewCipher(key)
   if err != nil {
      fmt.Errorf("NewCipher(%d bytes) = %s", len(key), err)
      panic(err)
   }
   plain := make([]byte, len(ciphertext))
   c.Decrypt(plain, ciphertext)
   s := string(plain[:])
   fmt.Printf("AES Decrypyed Text:  %s\n", s)
}
