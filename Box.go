package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/elv-zhanhugo/DaoTest/build/MyGovernor"
	"github.com/elv-zhanhugo/DaoTest/build/MyToken"
)


func main(){
	//client, err := ethclient.Dial("http://localhost:8545")
	//client, err := ethclient.Dial("https://host-34-105-49-255.testv2.contentfabric.io/eth")
	client, err := ethclient.Dial("https://host-468.contentfabric.io/eth")
	if err != nil {
		log.Fatal(err)
	}

	//privateKey, err := crypto.HexToECDSA("b67bffcebaa19782243b27d8b940ee011cd4e432d40769f788f174fad53f870b")
	//privateKey, err := crypto.HexToECDSA("8c036b3bf6f06a884cd2ab42281a6a15686536090b6ec1ce66c303532f36634c")
	privateKey, err := crypto.HexToECDSA("b67bffcebaa19782243b27d8b940ee011cd4e432d40769f788f174fad53f870b")
	if err != nil {
		log.Fatal(err)
	}

	// auth,err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(955111))
	//auth,err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(955203))
	auth,err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(955205))
	if err != nil {
		log.Fatal(err)
	}
	auth.GasLimit = uint64(20000000)

	tokenAddr, tx, instance,err := MyToken.DeployMyToken(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	

	tokenAddr, err = bind.WaitDeployed(context.Background(), client, tx)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("token contract address:", tokenAddr.Hex())
	fmt.Println("Deployed contract tx hash:", tx.Hash().Hex())
	
	name,err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	address, tx, instance2, err := MyGovernor.DeployMyGovernor(auth,client,tokenAddr,common.Address{})
	if err != nil {
		log.Fatal(err)
	}
	address, err = bind.WaitDeployed(context.Background(), client, tx)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Box contract address:", address.Hex())
	fmt.Println("Deployed contract tx hash:", tx.Hash().Hex())

	name,err = instance2.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)


	// // --- Deploy Box contract ---
	// address, tx, instance, err := box.DeployBox(auth, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// address, err = bind.WaitDeployed(context.Background(), client, tx)
	// if err != nil{
	// 	log.Fatal(err)
	// }
	// fmt.Println("Box contract address:", address.Hex())
	// fmt.Println("Deployed contract tx hash:", tx.Hash().Hex())
	//
	// // --- Invoke owner ---
	// // check if openzeppelin method works
	// owner,err := instance.Owner(&bind.CallOpts{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("owner address:", owner)
	//
	// // --- Invoke store method ---
	// tx,err = instance.Store(auth, []string{"Test","Crypto"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// txrct,err := bind.WaitMined(context.Background(), client, tx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Tx receipt:", txrct.TxHash.String())
	//
	// // -- Check for store event ---
	// evData,err := instance.ParseValueChanged(*txrct.Logs[0])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(evData.NewValue)
	//
	// // --- Retrive data ---
	// data,err := instance.Retrieve(&bind.CallOpts{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(data)
}
