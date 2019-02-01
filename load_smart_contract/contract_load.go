package main

import (
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    store "./contracts" // for demo
)

func main() {
    client, err := ethclient.Dial("https://rinkeby.infura.io")
    if err != nil {
        log.Fatal(err)
    }

	//fmt.Println(address.Hex())   // 0xC0C0f55923D53471838E66a2FbB7b911c666315e
    address := common.HexToAddress("0xC0C0f55923D53471838E66a2FbB7b911c666315e")
    instance, err := store.NewStore(address, client)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("contract is loaded")
	// version, err := instance.Version(nil)
    // if err != nil {
    //     log.Fatal(err)
    // }

    // fmt.Println(version) // "1.0"

    key := [32]byte{}
   
    copy(key[:], []byte("foo"))
  
    result, err := instance.Items(nil, key)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(result[:])) // "bar"

}