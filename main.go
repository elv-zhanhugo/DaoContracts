package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/elv-zhanhugo/DaoContracts/build/box"
	"github.com/elv-zhanhugo/DaoContracts/build/dex"
	"github.com/elv-zhanhugo/DaoContracts/build/governance_timelock"
	"github.com/elv-zhanhugo/DaoContracts/build/governance_token"
	"github.com/elv-zhanhugo/DaoContracts/build/governor_contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func MoveBlocks(blocks int, client *ethclient.Client, privateKey *ecdsa.PrivateKey) {
	networkId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Waiting for", blocks, "block(s)")
	if networkId == big.NewInt(4) {
		curr, err := client.BlockNumber(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		end := curr + uint64(blocks) + 1
		for {
			curr, err = client.BlockNumber(context.Background())
			if err != nil {
				log.Fatal(err)
			}
			if curr > end {
				break
			}
		}
	} else {
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

		for i := 0; i < blocks+1; i++ {
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

			_, err = bind.WaitMined(context.Background(), client, signedTx)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func createDeploymentsJSON(fileName string, contractABI string) {
	filePath := fileName + ".json"
	contractFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatal(err)
	}
	contractFile.Close()
	contracrABIString, _ := json.Marshal(map[string]interface{}{"abi": contractABI})
	os.WriteFile(filePath, contracrABIString, os.ModePerm)
}

func selectNetwork(network string) (*ethclient.Client, *ecdsa.PrivateKey, *bind.TransactOpts) {
	var url string
	var privateKeyString string
	var networkId *big.Int

	switch network {
	case "local":
		url = "http://localhost:8545"
		privateKeyString = "b67bffcebaa19782243b27d8b940ee011cd4e432d40769f788f174fad53f870b"
		networkId = big.NewInt(955101)
	case "demo":
		url = "https://host-76-74-28-234.contentfabric.io/eth/"
		privateKeyString = "76c59369d6c13f7321af8e5725a76d3b772aaf6b3d28eb631f5572daf4e0de06"
		networkId = big.NewInt(955210)
	case "tv4":
		url = "https://host-468.contentfabric.io/eth"
		privateKeyString = "76c59369d6c13f7321af8e5725a76d3b772aaf6b3d28eb631f5572daf4e0de06"
		networkId = big.NewInt(955205)

	case "rinkeby":
		url = "https://rinkeby.infura.io/v3/" + os.Getenv("WEB3_INFURA_PROJECT_ID")
		privateKeyString = os.Getenv("PRIVATE_KEY")
		networkId = big.NewInt(4)
	default:
		log.Fatal("invalid network")
	}

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, networkId)
	if err != nil {
		log.Fatal(err)
	}

	return client, privateKey, auth
}

func main() {
	godotenv.Load()

	if len(os.Args) == 1 {
		log.Fatal("Please choose a network")
	}

	client, privateKey, auth := selectNetwork(os.Args[1])

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
	// Deploy and configure Dex
	// ***************************************

	dexAddr, dexTx, dexInst, err := dex.DeployDex(auth, client, govTknAddr)
	if err != nil {
		log.Fatal(err)
	}
	dexAddr, err = bind.WaitDeployed(context.Background(), client, dexTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deployed Dex,", "address:", dexAddr.Hex(), "tx hash:", dexTx.Hash().String())

	totalSupply, err := govTknInst.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	var quorumPercentage int64 = 4
	quorumAmount := big.NewInt(0).Add(big.NewInt(0), totalSupply)
	quorumFraction := big.NewInt(0).Div(big.NewInt(100), big.NewInt(quorumPercentage))
	quorumAmount.Div(quorumAmount, quorumFraction)
	dexSupply := big.NewInt(0).Sub(totalSupply, quorumAmount)
	transferTx, err := govTknInst.Transfer(&bind.TransactOpts{From: auth.From, Signer: auth.Signer}, dexAddr, dexSupply)
	if err != nil {
		log.Fatal(err)
	}
	transferTxRct, err := bind.WaitMined(context.Background(), client, transferTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transferred", dexSupply, "of GT to Dex,", "tx hash:", transferTxRct.TxHash.String())
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

	buyTx, err := dexInst.Buy(&bind.TransactOpts{From: auth.From, Signer: auth.Signer, Value: big.NewInt(1)})
	if err != nil {
		log.Fatal(err)
	}
	buyTxRct, err := bind.WaitMined(context.Background(), client, buyTx)
	if err != nil {
		log.Fatal(err)
	}
	votes, err := govTknInst.GetVotes(&bind.CallOpts{}, auth.From)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bought GT from Dex, now account contains", votes, "GT", "tx hash:", buyTxRct.TxHash.String())

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

	MoveBlocks(int(votingPeriod), client, privateKey)

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
	time.Sleep(time.Duration(minDelay) * time.Second)

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
	fmt.Println("=================================================")

	// ***************************************
	// Create/Modify deployments files
	// ***************************************
	networkIdBigInt, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll("deployments/"+networkIdBigInt.String(), 0700)
	if err != nil {
		log.Fatal(err)
	}

	// Create contractAddr.json files in the networkId's directory
	createDeploymentsJSON("deployments/"+networkIdBigInt.String()+"/"+boxAddr.String(), box.BoxMetaData.ABI)
	createDeploymentsJSON("deployments/"+networkIdBigInt.String()+"/"+dexAddr.String(), dex.DexMetaData.ABI)
	createDeploymentsJSON("deployments/"+networkIdBigInt.String()+"/"+governorAddr.String(), governor_contract.GovernorContractMetaData.ABI)
	createDeploymentsJSON("deployments/"+networkIdBigInt.String()+"/"+govTimelockAddr.String(), governance_timelock.GovernanceTimelockMetaData.ABI)
	createDeploymentsJSON("deployments/"+networkIdBigInt.String()+"/"+govTknAddr.String(), governance_token.GovernanceTokenMetaData.ABI)

	// Create map.json if it doesn't exist, and add deployment information
	mapFile, err := os.OpenFile("deployments/map.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatal(err)
	}
	mapFile.Close()

	mapBytes, _ := os.ReadFile("deployments/map.json")
	var mapData map[int]map[string][]string
	err = json.Unmarshal(mapBytes, &mapData)

	if mapData == nil {
		mapData = map[int]map[string][]string{}
	}

	networkId := int(networkIdBigInt.Int64())
	if mapData[networkId] == nil {
		mapData[networkId] = map[string][]string{
			"Box":                {},
			"Dex":                {},
			"GovernanceTimeLock": {},
			"GovernanceToken":    {},
			"GovernorContract":   {},
		}
	}

	mapData[networkId]["Box"] = append(mapData[networkId]["Box"], boxAddr.String())
	mapData[networkId]["Dex"] = append(mapData[networkId]["Dex"], dexAddr.String())
	mapData[networkId]["GovernanceTimeLock"] = append(mapData[networkId]["GovernanceTimeLock"], govTimelockAddr.String())
	mapData[networkId]["GovernanceToken"] = append(mapData[networkId]["GovernanceToken"], govTknAddr.String())
	mapData[networkId]["GovernorContract"] = append(mapData[networkId]["GovernorContract"], governorAddr.String())

	mapDataString, _ := json.Marshal(mapData)
	os.WriteFile("deployments/map.json", mapDataString, os.ModePerm)

	fmt.Println("Contract addresses and abi are now stored in the deployments directory")
	fmt.Println("=================================================")
}
