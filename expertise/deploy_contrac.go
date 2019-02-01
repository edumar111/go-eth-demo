package main

import (
    "fmt"
	"log"

	//"encoding/hex"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"

    expertise "./contracts" // for demo
)

func main() {
    client, err := ethclient.Dial("http://35.231.124.87:8545")
    if err != nil {
        log.Fatal(err)
    }

	//fmt.Println(address.Hex())   // 0xC0C0f55923D53471838E66a2FbB7b911c666315e
    address := common.HexToAddress("0xad52c5e11867d3edf506905cf08e57dc3d2a4477")
    instance, err := expertise.NewExpertise(address, client)
    if err != nil {
        log.Fatal(err)
    }

    //fmt.Println("contract is loaded")
	// version, err := instance.Version(nil)
    // if err != nil {
    //     log.Fatal(err)
    // }

    // fmt.Println(version) // "1.0"

  
    result, err := instance.GetUsers(nil)
    if err != nil {
        log.Fatal(err)
    }
	
	//fmt.Println(result) // "bar"
	count := len(result)
	for idx := int(0); idx < count; idx++ {
		achivements,goals, err := instance .GetAchievementsIds(nil,result[idx] )
		if err != nil {
			log.Fatal(err)
		}
		count2 := len(achivements)
		for idx2 := int(0); idx2 < count2; idx2++ {
		
			fmt.Println( result[idx].Hex() ,string(achivements[idx2][:]), string(goals[idx2][:]))
			
		}
		
		
	}
	
}
