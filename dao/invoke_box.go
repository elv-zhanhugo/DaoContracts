package dao

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/elv-zhanhugo/DaoTest/build/box"
)

func InvokeBox(client *ethclient.Client, auth *bind.TransactOpts, contractAddr common.Address) error {

	inst, err := NewBoxInstance(client, contractAddr)
	if err != nil {
		return err
	}

	// --- Invoke owner ---
	// check if openzeppelin method works
	owner, err := inst.Owner(&bind.CallOpts{})
	if err != nil {
		return err
	}
	fmt.Println("owner address:", owner)

	// --- Invoke store method ---
	tx, err := inst.Store(auth, []string{"Test", "Crypto"})
	if err != nil {
		return err
	}
	txrct, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return err
	}
	fmt.Println("Tx receipt:", txrct.TxHash.String())

	// -- Check for store event ---
	evData, err := inst.ParseValueChanged(*txrct.Logs[0])
	if err != nil {
		return err
	}
	fmt.Println("event data:", evData.NewValue)

	// --- Retrive data ---
	data, err := inst.Retrieve(&bind.CallOpts{})
	if err != nil {
		return err
	}
	fmt.Println("data:", data)
	return nil
}

func NewBoxInstance(client *ethclient.Client, address common.Address) (*box.Box, error) {
	boxInst, err := box.NewBox(address, client)
	if err != nil {
		return nil, err
	}

	fmt.Println("Box contract address:", address.Hex())
	return boxInst, nil
}

func DeployBoxInstance(client *ethclient.Client, auth *bind.TransactOpts) (*box.Box, common.Address, error) {
	// // --- Deploy Box contract ---
	address, tx, instance, err := box.DeployBox(auth, client)
	if err != nil {
		return nil, common.Address{}, err
	}
	address, err = bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		return nil, common.Address{}, err
	}

	fmt.Println("Box contract address:", address.Hex())
	fmt.Println("Deployed box contract tx hash:", tx.Hash().Hex())

	return instance, address, nil

}
