package dao

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/elv-zhanhugo/DaoTest/build/MyGovernor"
	"github.com/elv-zhanhugo/DaoTest/build/MyToken"
)

func InvokeMyGovernor(client *ethclient.Client, auth *bind.TransactOpts) error {

	tokenAddr, tx, instance, err := MyToken.DeployMyToken(auth, client)
	if err != nil {
		return err
	}

	tokenAddr, err = bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		return err
	}
	fmt.Println("token contract address:", tokenAddr.Hex())
	fmt.Println("Deployed contract tx hash:", tx.Hash().Hex())

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		return err
	}
	fmt.Println("instance_name:", name)
	fmt.Println()

	// =========

	address, tx, instance2, err := MyGovernor.DeployMyGovernor(auth, client, tokenAddr, common.Address{})
	if err != nil {
		return err
	}
	address, err = bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		return err
	}
	fmt.Println("Governor contract address:", address.Hex())
	fmt.Println("Deployed contract tx hash:", tx.Hash().Hex())

	name, err = instance2.Name(&bind.CallOpts{})
	if err != nil {
		return err
	}
	fmt.Println("instance_name:", name)
	fmt.Println()

	return nil
}
