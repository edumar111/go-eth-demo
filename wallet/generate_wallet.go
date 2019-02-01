package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	// generate private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	// bytes to hex
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hexutil.Encode(privateKeyBytes)
	privateKeyHexClean := privateKeyHex[2:]
	fmt.Println("=====Private key=========")
	fmt.Println(privateKeyBytes)
	fmt.Println(privateKeyHex)
	fmt.Println(privateKeyHexClean)

	// generate public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	// publicKey byte to hex
	publicKeybytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyHex := hexutil.Encode(publicKeybytes)
	publicKeyHexClean := publicKeyHex[4:]
	fmt.Println("=====Public key=========")
	fmt.Println(publicKeybytes)
	fmt.Println(publicKeyHex)
	fmt.Println(publicKeyHexClean)

	// Generate address
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("=====Public address=========")
	fmt.Println(address)
	//generate legacy
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeybytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))

}
