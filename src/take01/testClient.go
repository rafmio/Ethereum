package main

import (
  "fmt"
  "log"

  "github.com/ethereum/go-ethereum/ethclient"
  // "github.com/ethereum/go-ethereum/accounts"
  "github.com/ethereum/go-ethereum/accounts/keystore"
  // "github.com/ethereum/go-ethereum/common"
)

type KeyStore struct {
  Handle *keystore.KeyStore
}



func main() {
  // Создаем клиент
  client, err := ethclient.Dial("https://cloudflare-eth.com")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("Type of client: %T\n", client)

  keydir := "./accounts"
  // Создаем хранилище
  ks := &KeyStore{}
  ks.Handle = keystore.NewKeyStore(keydir, keystore.LightScryptN, keystore.LightScryptP)

  var password string
  fmt.Printf("Enter password: ")
  fmt.Scanf("%s\n", &password)

  account, err := ks.Handle.NewAccount(password)
  if err != nil {
    fmt.Println("Error creating the new account: ", err.Error())
  }

  fmt.Printf("Type of account: %T\n", account)


}
