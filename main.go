package main

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/elv-zhanhugo/DaoTest/dao"
)

func main() {
	//client, err := ethclient.Dial("http://localhost:8545")
	//client, err := ethclient.Dial("https://host-34-105-49-255.testv2.contentfabric.io/eth")
	client, err := ethclient.Dial("https://host-468.contentfabric.io/eth")
	if err != nil {
		log.Fatal(err)
	}

	//privateKey, err := crypto.HexToECDSA("b67bffcebaa19782243b27d8b940ee011cd4e432d40769f788f174fad53f870b")
	//privateKey, err := crypto.HexToECDSA("8c036b3bf6f06a884cd2ab42281a6a15686536090b6ec1ce66c303532f36634c")
	//privateKey, err := crypto.HexToECDSA("b67bffcebaa19782243b27d8b940ee011cd4e432d40769f788f174fad53f870b")
	privateKey, err := crypto.HexToECDSA("76c59369d6c13f7321af8e5725a76d3b772aaf6b3d28eb631f5572daf4e0de06")
	if err != nil {
		log.Fatal(err)
	}

	//auth,err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(955101))
	//auth,err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(955203))
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(955205))
	if err != nil {
		log.Fatal(err)
	}
	auth.GasLimit = uint64(20000000)

	err = dao.InvokeBox(client, auth, common.HexToAddress("0xd2043ee9d45363981e6f2f312b61c3a4a79ecad6"))
	if err != nil {
		log.Fatal(err)
	}

}
