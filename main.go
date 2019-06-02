package main

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis"
)

//var client *redis.Client

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	reader := bufio.NewReader(os.Stdin) //std i/o
	for true {
		conn, err := ethclient.Dial("http://localhost:7545")

		if err != nil {
			log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		}

		var a = "0x4C5Dab94C9C166EF1D99c3cC14B74A77763f46d9"

		fmt.Print("Enter Private key  :")
		prik, _ := reader.ReadString('\n')
		prik = strings.TrimRight(prik, "\n")

		privateKey, err := crypto.HexToECDSA(prik)
		if err != nil {
			log.Fatal(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("error casting public key to ECDSA")
		}

		var pk = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

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

		fmt.Print("enter username  :")
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\n")

		var b = ""
		b, err = entity.GetEntity(&bind.CallOpts{}, text)
		if err != nil {
			log.Fatalf("Failed to get the entity: %v", err)
		} else {
			err := client.Set(b, text, 0).Err() //redis set
			if err != nil {
				fmt.Print("error in redis:", err)
			}
		}

		_, err = entity.SetEntity(auth, pk, text)
		if err != nil {
			log.Fatalf("Failed to set a new entity: %v", err)
		} else {
			fmt.Println("set entity status : true")
		}

		fmt.Print("key-value for entity:", text, " address: ", b)

		fmt.Print("\n To exit enter x else press any key to continue \n")
		status, _ := reader.ReadString('\n')
		status = strings.TrimRight(status, "\n")
		if strings.ToLower(status) == "x" {
			fmt.Print("exiting...")
			break
		}
		fmt.Print("\n")

	}
}
