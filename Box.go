package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	box "github.com/elv-zhanhugo/DaoTest/build"
)


func main(){
	client, err := ethclient.Dial("<NODE_IP>")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("<PRIV_KEY>")
	if err != nil {
		log.Fatal(err)
	}

	auth,err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(<NETWORK_NUMBER>))
	if err != nil {
		log.Fatal(err)
	}
	auth.GasLimit = uint64(20000000)


	address, tx, instance, err := box.DeployBox(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())
	_ = instance

}
