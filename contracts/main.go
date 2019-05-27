package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial("http://localhost:8545")
	// /home/adi/.ethereum/geth.ipc

	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	entity, err := NewAicumen(common.HexToAddress("0x6cd320d51a901c00481da32fa853f73c59913262"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a entity contract: %v", err)
	}

	_, err = entity.NewEntity(&bind.TransactOpts{}, common.HexToAddress("0x6cd320d51a901c00481da32fa853f73c59913262"), "aditya")
	fmt.Println("entity:", err)

	_, err = entity.GetEntity(&bind.CallOpts{}, common.HexToAddress("0x6cd320d51a901c00481da32fa853f73c59913262"))
	fmt.Println("entity:", err)

	name, err := entity.EntityStructs(&bind.CallOpts{}, common.HexToAddress("0x6cd320d51a901c00481da32fa853f73c59913262"))
	if err != nil {
		log.Fatalf("Failed to retrieve entity: %v", err)
	}
	fmt.Println("entity name:", name.IsEntity)
}
