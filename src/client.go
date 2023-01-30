package main

import (
  "fmt"
  "log"

  "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
  client, err := ethclient.Dial("https://cloudflare-eth.com")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("we have a connection")
  fmt.Printf("Type of client: %T\n", client)
  _ = client
}


// func Dial(rewurl string) (*Client, error)
// Dial connects a client to the given URL
// If you hae a local instance of geth you may pass the path to the IPC endpoint
// client, err := ethclient.Dial("/home/user/.ethereum/geth.ipc")
