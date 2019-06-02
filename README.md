## Go binding to Ethereum contract
$ abigen --sol=aicumen.sol --pkg=main --out=aicumen.go
## Generate abi
$ solcjs --abi aicumen.sol
## Generate binary
$ solcjs --bin aicumen.sol 
## Connect to network
$ testrpc
## Attach geth to the network
$ geth attach http://localhost:8545
## In geth console
> abi = <aicumen_sol_aicumen.abi>

> bytecode = <"aicumen_sol_aicumen.bin">

> aicumen = eth.contract(abi).new({from: eth.accounts[0], data: bytecode, gas: 3000000})

> aicumen.address

## Connect to go-redis
$ redis-server //terminal 1

$ redis-cli //terminal 2

## Insert data
$ go run main.go aicumen.go //from "contracts directory"

## Fetch the data on go-redis 
$ get <<Public Key>> //from redis-cli terminal 2


