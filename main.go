package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/elv-zhanhugo/DaoTest/build/box"
	"github.com/elv-zhanhugo/DaoTest/build/governance_standard/governance_timelock"
	"github.com/elv-zhanhugo/DaoTest/build/governance_standard/governor_contract"
	"github.com/elv-zhanhugo/DaoTest/build/governance_token"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func MoveBlocks(blocks int, client *ethclient.Client, privateKey *ecdsa.PrivateKey) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	value := big.NewInt(1) // wei

	gasLimit := uint64(21000) // units

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < blocks; i++ {
		nonce, err := client.PendingNonceAt(context.Background(), address)
		if err != nil {
			log.Fatal(err)
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		tx := types.NewTransaction(nonce, address, value, gasLimit, gasPrice, nil)

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			log.Fatal(err)
		}

		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	// client, err := ethclient.Dial("http://localhost:8545")
	client, err := ethclient.Dial("https://host-468.contentfabric.io/eth")
	if err != nil {
		log.Fatal(err)
	}

	// privateKey, err := crypto.HexToECDSA("b67bffcebaa19782243b27d8b940ee011cd4e432d40769f788f174fad53f870b")
	privateKey, err := crypto.HexToECDSA("76c59369d6c13f7321af8e5725a76d3b772aaf6b3d28eb631f5572daf4e0de06")
	if err != nil {
		log.Fatal(err)
	}

	// auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(955101))
	//auth,err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(955203))
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(955205))
	if err != nil {
		log.Fatal(err)
	}
	auth.GasLimit = uint64(20000000)

	blkNum, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("current block num:", blkNum)

	// ***************************************
	//  Deploy and configure GovernanceToken
	// ***************************************

	govTknAddr, govTknTx, govTknInst, err := governance_token.DeployGovernanceToken(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	govTknAddr, err = bind.WaitDeployed(context.Background(), client, govTknTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deployed Governance Token,", "address:", govTknAddr.Hex(), "tx hash:", govTknTx.Hash().String())

	// users to delegate to themselves in order to activate checkpoints and have their voting power tracked
	delegateTx, err := govTknInst.Delegate(auth, auth.From)
	if err != nil {
		log.Fatal(err)
	}
	delegateTxRct, err := bind.WaitMined(context.Background(), client, delegateTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Delegate Governance Token to token creator,", "tx:", delegateTxRct.TxHash.String())

	chkpts, err := govTknInst.NumCheckpoints(&bind.CallOpts{}, auth.From)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Check Governance Token checkpoints,", chkpts)
	fmt.Println("=================================================")

	// ***************************************
	// Deploy and configure GovernanceTimelock
	// ***************************************

	var minDelay int64 = 30 // 30 seconds
	govTimelockAddr, govTimelockTx, govTimelockInst, err := governance_timelock.DeployGovernanceTimelock(auth, client, big.NewInt(minDelay), nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	govTimelockAddr, err = bind.WaitDeployed(context.Background(), client, govTimelockTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deployed Governance Timelock,", "address:", govTimelockAddr.Hex(), "tx hash:", govTimelockTx.Hash().String())
	fmt.Println("=================================================")

	// ***************************************
	// Deploy and configure Governor
	// ***************************************

	var quorumPercentage int64 = 4
	var votingPeriod int64 = 5 // blocks
	var votingDelay int64 = 1  // blocks

	governorAddr, governorTx, governorInst, err := governor_contract.DeployGovernorContract(auth, client, govTknAddr, govTimelockAddr, big.NewInt(quorumPercentage), big.NewInt(votingPeriod), big.NewInt(votingDelay))
	if err != nil {
		log.Fatal(err)
	}
	governorAddr, err = bind.WaitDeployed(context.Background(), client, governorTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deployed Governor contract,", "address:", governorAddr.Hex(), "tx hash:", governorTx.Hash().String())

	proposerRole, err := govTimelockInst.PROPOSERROLE(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	executorRole, err := govTimelockInst.EXECUTORROLE(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	timelockAdminRole, err := govTimelockInst.TIMELOCKADMINROLE(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	proposorRoleGrantTx, err := govTimelockInst.GrantRole(auth, proposerRole, governorAddr)
	if err != nil {
		log.Fatal(err)
	}
	proposorRoleGrantTxRct, err := bind.WaitMined(context.Background(), client, proposorRoleGrantTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Set proposer role to governor contract,", "tx:", proposorRoleGrantTxRct.TxHash.String())

	executorRoleGrantTx, err := govTimelockInst.GrantRole(auth, executorRole, common.Address{})
	if err != nil {
		log.Fatal(err)
	}
	executorRoleGrantTxRct, err := bind.WaitMined(context.Background(), client, executorRoleGrantTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Set executor role to null addr,", "tx:", executorRoleGrantTxRct.TxHash.String())

	timelkRoleRevokeTx, err := govTimelockInst.RevokeRole(auth, timelockAdminRole, auth.From)
	if err != nil {
		log.Fatal(err)
	}
	timelkRoleRevokeTxRct, err := bind.WaitMined(context.Background(), client, timelkRoleRevokeTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Revoke timelock admin role from the owner,", "tx:", timelkRoleRevokeTxRct.TxHash.String())
	fmt.Println("=================================================")

	// ***************************************
	// Deploy and configure Box
	// ***************************************

	boxAddr, boxTx, boxInst, err := box.DeployBox(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	boxAddr, err = bind.WaitDeployed(context.Background(), client, boxTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Box contract", "address:", boxAddr.Hex(), "tx hash:", boxTx.Hash().String())

	transferboxOwnerhsipTx, err := boxInst.TransferOwnership(auth, govTimelockAddr)
	if err != nil {
		log.Fatal(err)
	}
	transferboxOwnerhsipTxRct, err := bind.WaitMined(context.Background(), client, transferboxOwnerhsipTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transfer box contract ownership to timelock contract,", "tx:", transferboxOwnerhsipTxRct.TxHash.String())

	// ***************************************
	// Propose new store value for box contract
	// ***************************************

	proposalDescription := "Storing hello world!"
	newStoreValue := []string{"hello world!"}

	boxABI, err := box.BoxMetaData.GetAbi()
	if err != nil {
		log.Fatal(err)
	}
	boxStoreCalldata, err := boxABI.Pack("store", newStoreValue)
	if err != nil {
		log.Fatal(err)
	}
	val := big.NewInt(0)
	govProposeTx, err := governorInst.Propose(auth, []common.Address{boxAddr}, []*big.Int{val}, [][]byte{boxStoreCalldata}, proposalDescription)
	if err != nil {
		log.Fatal(err)
	}
	govProposeTxRct, err := bind.WaitMined(context.Background(), client, govProposeTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Governor proposes new Box contract store value,", "tx:", govProposeTxRct.TxHash.String())

	proposalCreatedEvent, err := governorInst.ParseProposalCreated(*govProposeTxRct.Logs[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Proposal is created, event data:", "proposal_id", proposalCreatedEvent.ProposalId, "desc", proposalCreatedEvent.Description)
	fmt.Println("Proposal is created, event data:", "start_block", proposalCreatedEvent.StartBlock, "end_block", proposalCreatedEvent.EndBlock)

	proposalState, err := governorInst.State(&bind.CallOpts{}, proposalCreatedEvent.ProposalId)
	if err != nil {
		log.Fatal(err)
	}
	proposalVotes, err := governorInst.ProposalVotes(&bind.CallOpts{}, proposalCreatedEvent.ProposalId)
	if err != nil {
		log.Fatal(err)
	}
	proposalDeadline, err := governorInst.ProposalDeadline(&bind.CallOpts{}, proposalCreatedEvent.ProposalId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Proposal Info,", "state", proposalState, "votes", proposalVotes, "deadline", proposalDeadline)
	fmt.Println("=================================================")

	// ***************************************
	// Vote for the proposal
	// ***************************************

	MoveBlocks(int(votingDelay), client, privateKey)

	voteTx, err := governorInst.CastVote(auth, proposalCreatedEvent.ProposalId, 1)
	if err != nil {
		log.Fatal(err)
	}

	MoveBlocks(int(votingPeriod), client, privateKey)

	voteTxRct, err := bind.WaitMined(context.Background(), client, voteTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Vote is casted, event data:", voteTxRct.TxHash.String())

	voteCastEvent, err := governorInst.ParseVoteCast(*voteTxRct.Logs[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Vote support:", voteCastEvent.Support)
	fmt.Println("=================================================")

	// ***************************************
	// Queue the proposal
	// ***************************************

	var proposalDescriptionHash [32]byte
	copy(proposalDescriptionHash[:], solsha3.SoliditySHA3(
		solsha3.String(proposalDescription),
	))

	govQueueTx, err := governorInst.Queue(auth, []common.Address{boxAddr}, []*big.Int{val}, [][]byte{boxStoreCalldata}, proposalDescriptionHash)
	if err != nil {
		log.Fatal(err)
	}
	govQueueTxRct, err := bind.WaitMined(context.Background(), client, govQueueTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Governor queue passed proposal,", "tx:", govQueueTxRct.TxHash.String())

	eta, err := governorInst.ProposalEta(&bind.CallOpts{}, proposalCreatedEvent.ProposalId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Proposal can be executed at", eta, time.Unix(eta.Int64(), 0))
	fmt.Println("Waiting...")
	for {
		if time.Now().Unix() > time.Unix(eta.Int64(), 0).Unix() {
			break
		}
	}
	fmt.Println("Finished queue!")
	fmt.Println("=================================================")

	// ***************************************
	// Execute the proposal
	// ***************************************

	govExecuteTx, err := governorInst.Execute(auth, []common.Address{boxAddr}, []*big.Int{val}, [][]byte{boxStoreCalldata}, proposalDescriptionHash)
	if err != nil {
		log.Fatal(err)
	}
	govExecuteTxRct, err := bind.WaitMined(context.Background(), client, govExecuteTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Governor execute passed proposal,", "tx:", govExecuteTxRct.TxHash.String())

	proposalExecutedEvent, err := governorInst.ParseProposalExecuted(*govExecuteTxRct.Logs[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Proposal is executed, event data:", "proposal_id", proposalExecutedEvent.ProposalId)

	fmt.Println("=================================================")

	// ***************************************
	// Check box value
	// ***************************************

	retrievedValue, err := boxInst.Retrieve(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Box is storing: ", retrievedValue)
}
