package main

func CheckPassword (password string) bool {
  // keyFileName := "key.txt"
  key := []byte("TZPtSIacEJG18IpqQSkTE6luYmnCNKgR")
  passwordFileName := "password.txt"

  decryptedPassword := DecryptAES(key, passwordFileName)

  if password == decryptedPassword {
    return true
  } else {
    return false
  }
}
