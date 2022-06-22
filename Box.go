package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	box "github.com/elv-zhanhugo/DaoTest/build"
)


func main(){
	client, err := ethclient.Dial("<NET_IP>")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("<PK>")
	if err != nil {
		log.Fatal(err)
	}

	auth,err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(<NETWORK_ID>))
	if err != nil {
		log.Fatal(err)
	}
	auth.GasLimit = uint64(20000000)

	// --- Deploy Box contract ---
	address, tx, instance, err := box.DeployBox(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	address, err = bind.WaitDeployed(context.Background(), client, tx)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Box contract address:", address.Hex())
	fmt.Println("Deployed contract tx hash:", tx.Hash().Hex())

	// --- Invoke owner ---
	// check if openzeppelin method works
	owner,err := instance.Owner(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("owner address:", owner)

	// --- Invoke store method ---
	tx,err = instance.Store(auth, []string{"Test","Crypto"})
	if err != nil {
		log.Fatal(err)
	}
	txrct,err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tx receipt:", txrct.TxHash.String())

	// -- Check for store event ---
	evData,err := instance.ParseValueChanged(*txrct.Logs[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(evData.NewValue)

	// --- Retrive data ---
	data,err := instance.Retrieve(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
