package main

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial("http://localhost:7545")
	// /home/adi/.ethereum/geth.ipc

	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	var a = "0x319a59deC27B83F233C187D9926900a76226EF9F"
	// Instantiate the contract and display its name

	// fmt.Println("init")

	// var pk = "0x5caf8eb09302dab80df496639e0f5627710f8ee7"
	privateKey, err := crypto.HexToECDSA("ce8bc723f8d6b873902f225fce6d3a6c8b58e14e9cfa2013775f17f92e85c748")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	// var pk = hexutil.Encode(publicKeyECDSA)[4:]
	var pk = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	// var pk = hex.EncodeToString(privk)

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := conn.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	entity, err := NewAicumen(common.HexToAddress(a), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a entity contract: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("enter username  :")
	text, _ := reader.ReadString('\n')

	_, err = entity.SetEntity(auth, pk, text)
	if err != nil {
		log.Fatalf("Failed to set a new entity: %v", err)
	} else {
		fmt.Println("set entity status : true")
	}

	var b = ""
	b, err = entity.GetEntity(&bind.CallOpts{}, text)
	if err != nil {
		log.Fatalf("Failed to get the entity: %v", err)
	}

	fmt.Print("key-value for entity:", text, " address: ", b)

}
