package main

import (
	"fmt"
	"zcontractowner"
)

func main() {
	rpcServer := "http://domain-or-ip-address:8545"
	contractAddress := "0x_THE_CONTRACT_THAT_HAS_OWNER"

	owner := zcontractowner.GetContractOwner(rpcServer, contractAddress)
	fmt.Println(owner)
}
