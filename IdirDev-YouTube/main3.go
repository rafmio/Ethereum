package main

import (
  "fmt"
  "log"
  "os"

  // "github.com/ethereum/go-ethereum/accounts/keystore"
  "github.com/ethereum/go-ethereum/common/hexutil"
  "github.com/ethereum/go-ethereum/accounts/keystore"
  "github.com/ethereum/go-ethereum/crypto"
)

func main() {
  // key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
  password := "password"
  // a, err := key.NewAccount(password)
  // if err != nil {
  //   log.Fatal(err)
  // }
  // fmt.Println(a.Address)

  b, err := os.ReadFile("./wallet/UTC--2023-01-31T18-25-22.575467888Z--957d96c4df7ce17f0327cefad8112289563d4903")
  if err != nil {
    log.Fatal(err)
  }

  key, err := keystore.DecryptKey(b, password)
  if err != nil {
    log.Fatal(err)
  }
  privData := crypto.FromECDSA(key.PrivateKey)
  fmt.Println("Private:", hexutil.Encode(privData))

  pubData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
  fmt.Println("Public: ", hexutil.Encode(pubData))

  addr := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
  fmt.Println("Add: ", addr)
}
