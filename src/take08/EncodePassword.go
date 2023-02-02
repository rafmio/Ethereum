package main

import (
  "log"
  "crypto/sha256"
  "os"
)

/*
Function for:
1. Ecoding the keyphrase to sha256
2. Pass encoded keyphrase and password itsef to the function,
which will encode the string password and the sha256 keyphrase
for encoding the password to AES
*/
func EncodePassword(password string, keyphrase string) string {
   keyFileName := "key.txt"
   passwordFileName := "password.txt"

   // Generating the hash sum (sha256) for keyphrase:
   key256Hash := sha256.New()             // return empty hash.Hash (interface)
   key256Hash.Write([]byte(keyphrase))    //


   // Storing the hash in the file
   // Write the keyphrase for separate place for security purposes
   err := os.WriteFile(keyFileName, []byte(key256HashString), 0644)
   if err != nil {
     log.Fatal(err)
   }

   // Generating the hash sum (AES) for password
   passwordAESHash := Encrypt(key256Hash, password)
   err = os.WriteFile(passwordFileName, passwordAESHash, 0644)
   if err != nil {
     log.Fatal(err)
   }
   return passwordAESHash
}
