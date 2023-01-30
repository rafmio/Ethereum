// https://golang.company/blog/how-to-generate-ethereum-wallets-using-golang
package main

import (
  "fmt"
  "log"

  "github.com/ethereum/go-ethereum/crypto"
  "github.com/ethereum/go-ethereum/common/hexutil"
)

func main() {
  // generate the new wallet
  privateKey, err := crypto.GenerateKey()

  if err != nil {
    log.Fatal("Not generating new keys")
  }
  fmt.Println(privateKey)
  fmt.Printf("Type of privateKey: %T\n", privateKey)
  fmt.Println("")
  // convert the privateKey that is returned to us from the ECDSA format
  // into a bytes format
  // return a hexadecimal formatted private key
  privateKeyBytes := crypto.FromECDSA(privateKey)
  fmt.Println(hexutil.Encode(privateKeyBytes))
}

// go mod init github.com/golang-company/ethWall
// go get github.com/ethereum/go-ethereum/crypto
// go get github.com/ethereum/go-ethereum/common/hexutil
