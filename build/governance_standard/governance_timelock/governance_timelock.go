// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package governance_timelock

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GovernanceTimelockMetaData contains all meta data concerning the GovernanceTimelock contract.
var GovernanceTimelockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minDelay\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"proposers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"executors\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"CallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"CallScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"Cancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"MinDelayChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CANCELLER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXECUTOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROPOSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TIMELOCK_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloads\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"hashOperation\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloads\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"hashOperationBatch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"isOperation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"pending\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"isOperationDone\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"done\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"isOperationPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"pending\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"isOperationReady\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ready\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"schedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloads\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"scheduleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"updateDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620022ad380380620022ad8339810160408190526200003491620003fd565b828282620000526000805160206200222d8339815191528062000222565b6200007c6000805160206200224d8339815191526000805160206200222d83398151915262000222565b620000a66000805160206200226d8339815191526000805160206200222d83398151915262000222565b620000d06000805160206200228d8339815191526000805160206200222d83398151915262000222565b620000eb6000805160206200222d833981519152336200026d565b620001066000805160206200222d833981519152306200026d565b60005b82518110156200018c57620001506000805160206200224d8339815191528483815181106200013c576200013c62000471565b60200260200101516200026d60201b60201c565b620001796000805160206200228d8339815191528483815181106200013c576200013c62000471565b620001848162000487565b905062000109565b5060005b8151811015620001d657620001c36000805160206200226d8339815191528383815181106200013c576200013c62000471565b620001ce8162000487565b905062000190565b5060028390556040805160008152602081018590527f11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5910160405180910390a1505050505050620004af565b600082815260208190526040808220600101805490849055905190918391839186917fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff9190a4505050565b6200027982826200027d565b5050565b6000828152602081815260408083206001600160a01b038516845290915290205460ff1662000279576000828152602081815260408083206001600160a01b03851684529091529020805460ff19166001179055620002d93390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b634e487b7160e01b600052604160045260246000fd5b80516001600160a01b03811681146200034b57600080fd5b919050565b600082601f8301126200036257600080fd5b815160206001600160401b03808311156200038157620003816200031d565b8260051b604051601f19603f83011681018181108482111715620003a957620003a96200031d565b604052938452858101830193838101925087851115620003c857600080fd5b83870191505b84821015620003f257620003e28262000333565b83529183019190830190620003ce565b979650505050505050565b6000806000606084860312156200041357600080fd5b835160208501519093506001600160401b03808211156200043357600080fd5b620004418783880162000350565b935060408601519150808211156200045857600080fd5b50620004678682870162000350565b9150509250925092565b634e487b7160e01b600052603260045260246000fd5b600060018201620004a857634e487b7160e01b600052601160045260246000fd5b5060010190565b611d6e80620004bf6000396000f3fe6080604052600436106101bb5760003560e01c80638065657f116100ec578063bc197c811161008a578063d547741f11610064578063d547741f14610582578063e38335e5146105a2578063f23a6e61146105b5578063f27a0c92146105e157600080fd5b8063bc197c8114610509578063c4d252f514610535578063d45c44351461055557600080fd5b806391d14854116100c657806391d1485414610480578063a217fddf146104a0578063b08e51c0146104b5578063b1c5f427146104e957600080fd5b80638065657f1461040c5780638f2a0bb01461042c5780638f61f4f51461044c57600080fd5b8063248a9ca31161015957806331d507501161013357806331d507501461038c57806336568abe146103ac578063584b153e146103cc57806364d62353146103ec57600080fd5b8063248a9ca31461030b5780632ab0f5291461033b5780632f2ff15d1461036c57600080fd5b80630d3cf6fc116101955780630d3cf6fc14610260578063134008d31461029457806313bc9f20146102a7578063150b7a02146102c757600080fd5b806301d5062a146101c757806301ffc9a7146101e957806307bd02651461021e57600080fd5b366101c257005b600080fd5b3480156101d357600080fd5b506101e76101e2366004611368565b6105f6565b005b3480156101f557600080fd5b506102096102043660046113dc565b61068b565b60405190151581526020015b60405180910390f35b34801561022a57600080fd5b506102527fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e6381565b604051908152602001610215565b34801561026c57600080fd5b506102527f5f58e3a2316349923ce3780f8d587db2d72378aed66a8261c916544fa6846ca581565b6101e76102a2366004611406565b6106b6565b3480156102b357600080fd5b506102096102c2366004611471565b61072e565b3480156102d357600080fd5b506102f26102e236600461153f565b630a85bd0160e11b949350505050565b6040516001600160e01b03199091168152602001610215565b34801561031757600080fd5b50610252610326366004611471565b60009081526020819052604090206001015490565b34801561034757600080fd5b50610209610356366004611471565b6000908152600160208190526040909120541490565b34801561037857600080fd5b506101e76103873660046115a6565b610754565b34801561039857600080fd5b506102096103a7366004611471565b61077e565b3480156103b857600080fd5b506101e76103c73660046115a6565b610797565b3480156103d857600080fd5b506102096103e7366004611471565b61081a565b3480156103f857600080fd5b506101e7610407366004611471565b610830565b34801561041857600080fd5b50610252610427366004611406565b6108d4565b34801561043857600080fd5b506101e7610447366004611616565b610913565b34801561045857600080fd5b506102527fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc181565b34801561048c57600080fd5b5061020961049b3660046115a6565b610a65565b3480156104ac57600080fd5b50610252600081565b3480156104c157600080fd5b506102527ffd643c72710c63c0180259aba6b2d05451e3591a24e58b62239378085726f78381565b3480156104f557600080fd5b506102526105043660046116c7565b610a8e565b34801561051557600080fd5b506102f26105243660046117ee565b63bc197c8160e01b95945050505050565b34801561054157600080fd5b506101e7610550366004611471565b610ad3565b34801561056157600080fd5b50610252610570366004611471565b60009081526001602052604090205490565b34801561058e57600080fd5b506101e761059d3660046115a6565b610ba8565b6101e76105b03660046116c7565b610bcd565b3480156105c157600080fd5b506102f26105d0366004611897565b63f23a6e6160e01b95945050505050565b3480156105ed57600080fd5b50600254610252565b7fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc161062081610d02565b60006106308989898989896108d4565b905061063c8184610d0f565b6000817f4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca8b8b8b8b8b8a60405161067896959493929190611924565b60405180910390a3505050505050505050565b60006001600160e01b03198216630271189760e51b14806106b057506106b082610dfe565b92915050565b7fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e636106e2816000610a65565b6106f0576106f08133610e33565b60006107008888888888886108d4565b905061070c8185610e97565b61071b8160008a8a8a8a610f33565b61072481611047565b5050505050505050565b60008181526001602052604081205460018111801561074d5750428111155b9392505050565b60008281526020819052604090206001015461076f81610d02565b6107798383611080565b505050565b60008181526001602052604081205481905b1192915050565b6001600160a01b038116331461080c5760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b60648201526084015b60405180910390fd5b6108168282611104565b5050565b6000818152600160208190526040822054610790565b3330146108935760405162461bcd60e51b815260206004820152602b60248201527f54696d656c6f636b436f6e74726f6c6c65723a2063616c6c6572206d7573742060448201526a62652074696d656c6f636b60a81b6064820152608401610803565b60025460408051918252602082018390527f11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5910160405180910390a1600255565b60008686868686866040516020016108f196959493929190611924565b6040516020818303038152906040528051906020012090509695505050505050565b7fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc161093d81610d02565b88871461095c5760405162461bcd60e51b815260040161080390611961565b88851461097b5760405162461bcd60e51b815260040161080390611961565b600061098d8b8b8b8b8b8b8b8b610a8e565b90506109998184610d0f565b60005b8a811015610a575780827f4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca8e8e858181106109d9576109d96119a4565b90506020020160208101906109ee91906119ba565b8d8d86818110610a0057610a006119a4565b905060200201358c8c87818110610a1957610a196119a4565b9050602002810190610a2b91906119d5565b8c8b604051610a3f96959493929190611924565b60405180910390a3610a5081611a31565b905061099c565b505050505050505050505050565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b60008888888888888888604051602001610aaf989796959493929190611adb565b60405160208183030381529060405280519060200120905098975050505050505050565b7ffd643c72710c63c0180259aba6b2d05451e3591a24e58b62239378085726f783610afd81610d02565b610b068261081a565b610b6c5760405162461bcd60e51b815260206004820152603160248201527f54696d656c6f636b436f6e74726f6c6c65723a206f7065726174696f6e2063616044820152701b9b9bdd0818994818d85b98d95b1b1959607a1b6064820152608401610803565b6000828152600160205260408082208290555183917fbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb7091a25050565b600082815260208190526040902060010154610bc381610d02565b6107798383611104565b7fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e63610bf9816000610a65565b610c0757610c078133610e33565b878614610c265760405162461bcd60e51b815260040161080390611961565b878414610c455760405162461bcd60e51b815260040161080390611961565b6000610c578a8a8a8a8a8a8a8a610a8e565b9050610c638185610e97565b60005b89811015610cec57610cdc82828d8d85818110610c8557610c856119a4565b9050602002016020810190610c9a91906119ba565b8c8c86818110610cac57610cac6119a4565b905060200201358b8b87818110610cc557610cc56119a4565b9050602002810190610cd791906119d5565b610f33565b610ce581611a31565b9050610c66565b50610cf681611047565b50505050505050505050565b610d0c8133610e33565b50565b610d188261077e565b15610d7d5760405162461bcd60e51b815260206004820152602f60248201527f54696d656c6f636b436f6e74726f6c6c65723a206f7065726174696f6e20616c60448201526e1c9958591e481cd8da19591d5b1959608a1b6064820152608401610803565b600254811015610dde5760405162461bcd60e51b815260206004820152602660248201527f54696d656c6f636b436f6e74726f6c6c65723a20696e73756666696369656e746044820152652064656c617960d01b6064820152608401610803565b610de88142611b86565b6000928352600160205260409092209190915550565b60006001600160e01b03198216637965db0b60e01b14806106b057506301ffc9a760e01b6001600160e01b03198316146106b0565b610e3d8282610a65565b61081657610e55816001600160a01b03166014611169565b610e60836020611169565b604051602001610e71929190611bce565b60408051601f198184030181529082905262461bcd60e51b825261080391600401611c43565b610ea08261072e565b610ebc5760405162461bcd60e51b815260040161080390611c76565b801580610ed85750600081815260016020819052604090912054145b6108165760405162461bcd60e51b815260206004820152602660248201527f54696d656c6f636b436f6e74726f6c6c65723a206d697373696e6720646570656044820152656e64656e637960d01b6064820152608401610803565b6000846001600160a01b0316848484604051610f50929190611cc0565b60006040518083038185875af1925050503d8060008114610f8d576040519150601f19603f3d011682016040523d82523d6000602084013e610f92565b606091505b5050905080610fff5760405162461bcd60e51b815260206004820152603360248201527f54696d656c6f636b436f6e74726f6c6c65723a20756e6465726c79696e6720746044820152721c985b9cd858dd1a5bdb881c995d995c9d1959606a1b6064820152608401610803565b85877fc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58878787876040516110369493929190611cd0565b60405180910390a350505050505050565b6110508161072e565b61106c5760405162461bcd60e51b815260040161080390611c76565b600090815260016020819052604090912055565b61108a8282610a65565b610816576000828152602081815260408083206001600160a01b03851684529091529020805460ff191660011790556110c03390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b61110e8282610a65565b15610816576000828152602081815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b60606000611178836002611d02565b611183906002611b86565b6001600160401b0381111561119a5761119a61148a565b6040519080825280601f01601f1916602001820160405280156111c4576020820181803683370190505b509050600360fc1b816000815181106111df576111df6119a4565b60200101906001600160f81b031916908160001a905350600f60fb1b8160018151811061120e5761120e6119a4565b60200101906001600160f81b031916908160001a9053506000611232846002611d02565b61123d906001611b86565b90505b60018111156112b5576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110611271576112716119a4565b1a60f81b828281518110611287576112876119a4565b60200101906001600160f81b031916908160001a90535060049490941c936112ae81611d21565b9050611240565b50831561074d5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610803565b80356001600160a01b038116811461131b57600080fd5b919050565b60008083601f84011261133257600080fd5b5081356001600160401b0381111561134957600080fd5b60208301915083602082850101111561136157600080fd5b9250929050565b600080600080600080600060c0888a03121561138357600080fd5b61138c88611304565b96506020880135955060408801356001600160401b038111156113ae57600080fd5b6113ba8a828b01611320565b989b979a50986060810135976080820135975060a09091013595509350505050565b6000602082840312156113ee57600080fd5b81356001600160e01b03198116811461074d57600080fd5b60008060008060008060a0878903121561141f57600080fd5b61142887611304565b95506020870135945060408701356001600160401b0381111561144a57600080fd5b61145689828a01611320565b979a9699509760608101359660809091013595509350505050565b60006020828403121561148357600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156114c8576114c861148a565b604052919050565b600082601f8301126114e157600080fd5b81356001600160401b038111156114fa576114fa61148a565b61150d601f8201601f19166020016114a0565b81815284602083860101111561152257600080fd5b816020850160208301376000918101602001919091529392505050565b6000806000806080858703121561155557600080fd5b61155e85611304565b935061156c60208601611304565b92506040850135915060608501356001600160401b0381111561158e57600080fd5b61159a878288016114d0565b91505092959194509250565b600080604083850312156115b957600080fd5b823591506115c960208401611304565b90509250929050565b60008083601f8401126115e457600080fd5b5081356001600160401b038111156115fb57600080fd5b6020830191508360208260051b850101111561136157600080fd5b600080600080600080600080600060c08a8c03121561163457600080fd5b89356001600160401b038082111561164b57600080fd5b6116578d838e016115d2565b909b50995060208c013591508082111561167057600080fd5b61167c8d838e016115d2565b909950975060408c013591508082111561169557600080fd5b506116a28c828d016115d2565b9a9d999c50979a969997986060880135976080810135975060a0013595509350505050565b60008060008060008060008060a0898b0312156116e357600080fd5b88356001600160401b03808211156116fa57600080fd5b6117068c838d016115d2565b909a50985060208b013591508082111561171f57600080fd5b61172b8c838d016115d2565b909850965060408b013591508082111561174457600080fd5b506117518b828c016115d2565b999c989b509699959896976060870135966080013595509350505050565b600082601f83011261178057600080fd5b813560206001600160401b0382111561179b5761179b61148a565b8160051b6117aa8282016114a0565b92835284810182019282810190878511156117c457600080fd5b83870192505b848310156117e3578235825291830191908301906117ca565b979650505050505050565b600080600080600060a0868803121561180657600080fd5b61180f86611304565b945061181d60208701611304565b935060408601356001600160401b038082111561183957600080fd5b61184589838a0161176f565b9450606088013591508082111561185b57600080fd5b61186789838a0161176f565b9350608088013591508082111561187d57600080fd5b5061188a888289016114d0565b9150509295509295909350565b600080600080600060a086880312156118af57600080fd5b6118b886611304565b94506118c660208701611304565b9350604086013592506060860135915060808601356001600160401b038111156118ef57600080fd5b61188a888289016114d0565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60018060a01b038716815285602082015260a06040820152600061194c60a0830186886118fb565b60608301949094525060800152949350505050565b60208082526023908201527f54696d656c6f636b436f6e74726f6c6c65723a206c656e677468206d69736d616040820152620e8c6d60eb1b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b6000602082840312156119cc57600080fd5b61074d82611304565b6000808335601e198436030181126119ec57600080fd5b8301803591506001600160401b03821115611a0657600080fd5b60200191503681900382131561136157600080fd5b634e487b7160e01b600052601160045260246000fd5b600060018201611a4357611a43611a1b565b5060010190565b81835260006020808501808196508560051b810191508460005b87811015611ace5782840389528135601e19883603018112611a8557600080fd5b870185810190356001600160401b03811115611aa057600080fd5b803603821315611aaf57600080fd5b611aba8682846118fb565b9a87019a9550505090840190600101611a64565b5091979650505050505050565b60a0808252810188905260008960c08301825b8b811015611b1c576001600160a01b03611b0784611304565b16825260209283019290910190600101611aee565b5083810360208501528881526001600160fb1b03891115611b3c57600080fd5b8860051b9150818a602083013781810191505060208101600081526020848303016040850152611b6d81888a611a4a565b6060850196909652505050608001529695505050505050565b60008219821115611b9957611b99611a1b565b500190565b60005b83811015611bb9578181015183820152602001611ba1565b83811115611bc8576000848401525b50505050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351611c06816017850160208801611b9e565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351611c37816028840160208801611b9e565b01602801949350505050565b6020815260008251806020840152611c62816040850160208701611b9e565b601f01601f19169190910160400192915050565b6020808252602a908201527f54696d656c6f636b436f6e74726f6c6c65723a206f7065726174696f6e206973604082015269206e6f7420726561647960b01b606082015260800190565b8183823760009101908152919050565b60018060a01b0385168152836020820152606060408201526000611cf86060830184866118fb565b9695505050505050565b6000816000190483118215151615611d1c57611d1c611a1b565b500290565b600081611d3057611d30611a1b565b50600019019056fea2646970667358221220406fea28f569bce2f2d4e29249f092fd12b8ce7160ca67b8369bfb0885d796c464736f6c634300080f00335f58e3a2316349923ce3780f8d587db2d72378aed66a8261c916544fa6846ca5b09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc1d8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e63fd643c72710c63c0180259aba6b2d05451e3591a24e58b62239378085726f783",
}

// GovernanceTimelockABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernanceTimelockMetaData.ABI instead.
var GovernanceTimelockABI = GovernanceTimelockMetaData.ABI

// GovernanceTimelockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GovernanceTimelockMetaData.Bin instead.
var GovernanceTimelockBin = GovernanceTimelockMetaData.Bin

// DeployGovernanceTimelock deploys a new Ethereum contract, binding an instance of GovernanceTimelock to it.
func DeployGovernanceTimelock(auth *bind.TransactOpts, backend bind.ContractBackend, minDelay *big.Int, proposers []common.Address, executors []common.Address) (common.Address, *types.Transaction, *GovernanceTimelock, error) {
	parsed, err := GovernanceTimelockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovernanceTimelockBin), backend, minDelay, proposers, executors)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GovernanceTimelock{GovernanceTimelockCaller: GovernanceTimelockCaller{contract: contract}, GovernanceTimelockTransactor: GovernanceTimelockTransactor{contract: contract}, GovernanceTimelockFilterer: GovernanceTimelockFilterer{contract: contract}}, nil
}

// GovernanceTimelock is an auto generated Go binding around an Ethereum contract.
type GovernanceTimelock struct {
	GovernanceTimelockCaller     // Read-only binding to the contract
	GovernanceTimelockTransactor // Write-only binding to the contract
	GovernanceTimelockFilterer   // Log filterer for contract events
}

// GovernanceTimelockCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceTimelockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTimelockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceTimelockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTimelockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceTimelockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTimelockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceTimelockSession struct {
	Contract     *GovernanceTimelock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GovernanceTimelockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceTimelockCallerSession struct {
	Contract *GovernanceTimelockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// GovernanceTimelockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceTimelockTransactorSession struct {
	Contract     *GovernanceTimelockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// GovernanceTimelockRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceTimelockRaw struct {
	Contract *GovernanceTimelock // Generic contract binding to access the raw methods on
}

// GovernanceTimelockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceTimelockCallerRaw struct {
	Contract *GovernanceTimelockCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceTimelockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceTimelockTransactorRaw struct {
	Contract *GovernanceTimelockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernanceTimelock creates a new instance of GovernanceTimelock, bound to a specific deployed contract.
func NewGovernanceTimelock(address common.Address, backend bind.ContractBackend) (*GovernanceTimelock, error) {
	contract, err := bindGovernanceTimelock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovernanceTimelock{GovernanceTimelockCaller: GovernanceTimelockCaller{contract: contract}, GovernanceTimelockTransactor: GovernanceTimelockTransactor{contract: contract}, GovernanceTimelockFilterer: GovernanceTimelockFilterer{contract: contract}}, nil
}

// NewGovernanceTimelockCaller creates a new read-only instance of GovernanceTimelock, bound to a specific deployed contract.
func NewGovernanceTimelockCaller(address common.Address, caller bind.ContractCaller) (*GovernanceTimelockCaller, error) {
	contract, err := bindGovernanceTimelock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTimelockCaller{contract: contract}, nil
}

// NewGovernanceTimelockTransactor creates a new write-only instance of GovernanceTimelock, bound to a specific deployed contract.
func NewGovernanceTimelockTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceTimelockTransactor, error) {
	contract, err := bindGovernanceTimelock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTimelockTransactor{contract: contract}, nil
}

// NewGovernanceTimelockFilterer creates a new log filterer instance of GovernanceTimelock, bound to a specific deployed contract.
func NewGovernanceTimelockFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceTimelockFilterer, error) {
	contract, err := bindGovernanceTimelock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceTimelockFilterer{contract: contract}, nil
}

// bindGovernanceTimelock binds a generic wrapper to an already deployed contract.
func bindGovernanceTimelock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceTimelockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceTimelock *GovernanceTimelockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceTimelock.Contract.GovernanceTimelockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceTimelock *GovernanceTimelockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.GovernanceTimelockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceTimelock *GovernanceTimelockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.GovernanceTimelockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceTimelock *GovernanceTimelockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceTimelock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceTimelock *GovernanceTimelockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceTimelock *GovernanceTimelockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.contract.Transact(opts, method, params...)
}

// CANCELLERROLE is a free data retrieval call binding the contract method 0xb08e51c0.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (_GovernanceTimelock *GovernanceTimelockCaller) CANCELLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovernanceTimelock.contract.Call(opts, &out, "CANCELLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CANCELLERROLE is a free data retrieval call binding the contract method 0xb08e51c0.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (_GovernanceTimelock *GovernanceTimelockSession) CANCELLERROLE() ([32]byte, error) {
	return _GovernanceTimelock.Contract.CANCELLERROLE(&_GovernanceTimelock.CallOpts)
}

// CANCELLERROLE is a free data retrieval call binding the contract method 0xb08e51c0.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (_GovernanceTimelock *GovernanceTimelockCallerSession) CANCELLERROLE() ([32]byte, error) {
	return _GovernanceTimelock.Contract.CANCELLERROLE(&_GovernanceTimelock.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GovernanceTimelock *GovernanceTimelockCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovernanceTimelock.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GovernanceTimelock *GovernanceTimelockSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GovernanceTimelock.Contract.DEFAULTADMINROLE(&_GovernanceTimelock.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GovernanceTimelock *GovernanceTimelockCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GovernanceTimelock.Contract.DEFAULTADMINROLE(&_GovernanceTimelock.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_GovernanceTimelock *GovernanceTimelockCaller) EXECUTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovernanceTimelock.contract.Call(opts, &out, "EXECUTOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_GovernanceTimelock *GovernanceTimelockSession) EXECUTORROLE() ([32]byte, error) {
	return _GovernanceTimelock.Contract.EXECUTORROLE(&_GovernanceTimelock.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_GovernanceTimelock *GovernanceTimelockCallerSession) EXECUTORROLE() ([32]byte, error) {
	return _GovernanceTimelock.Contract.EXECUTORROLE(&_GovernanceTimelock.CallOpts)
}

// PROPOSERROLE is a free data retrieval call binding the contract method 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (_GovernanceTimelock *GovernanceTimelockCaller) PROPOSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovernanceTimelock.contract.Call(opts, &out, "PROPOSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROPOSERROLE is a free data retrieval call binding the contract method 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (_GovernanceTimelock *GovernanceTimelockSession) PROPOSERROLE() ([32]byte, error) {
	return _GovernanceTimelock.Contract.PROPOSERROLE(&_GovernanceTimelock.CallOpts)
}

// PROPOSERROLE is a free data retrieval call binding the contract method 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (_GovernanceTimelock *GovernanceTimelockCallerSession) PROPOSERROLE() ([32]byte, error) {
	return _GovernanceTimelock.Contract.PROPOSERROLE(&_GovernanceTimelock.CallOpts)
}

// TIMELOCKADMINROLE is a free data retrieval call binding the contract method 0x0d3cf6fc.
//
// Solidity: function TIMELOCK_ADMIN_ROLE() view returns(bytes32)
func (_GovernanceTimelock *GovernanceTimelockCaller) TIMELOCKADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GovernanceTimelock.contract.Call(opts, &out, "TIMELOCK_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TIMELOCKADMINROLE is a free data retrieval call binding the contract method 0x0d3cf6fc.
//
// Solidity: function TIMELOCK_ADMIN_ROLE() view returns(bytes32)
func (_GovernanceTimelock *GovernanceTimelockSession) TIMELOCKADMINROLE() ([32]byte, error) {
	return _GovernanceTimelock.Contract.TIMELOCKADMINROLE(&_GovernanceTimelock.CallOpts)
}

// TIMELOCKADMINROLE is a free data retrieval call binding the contract method 0x0d3cf6fc.
//
// Solidity: function TIMELOCK_ADMIN_ROLE() view returns(bytes32)
func (_GovernanceTimelock *GovernanceTimelockCallerSession) TIMELOCKADMINROLE() ([32]byte, error) {
	return _GovernanceTimelock.Contract.TIMELOCKADMINROLE(&_GovernanceTimelock.CallOpts)
}

// GetMinDelay is a free data retrieval call binding the contract method 0xf27a0c92.
//
// Solidity: function getMinDelay() view returns(uint256 duration)
func (_GovernanceTimelock *GovernanceTimelockCaller) GetMinDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceTimelock.contract.Call(opts, &out, "getMinDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinDelay is a free data retrieval call binding the contract method 0xf27a0c92.
//
// Solidity: function getMinDelay() view returns(uint256 duration)
func (_GovernanceTimelock *GovernanceTimelockSession) GetMinDelay() (*big.Int, error) {
	return _GovernanceTimelock.Contract.GetMinDelay(&_GovernanceTimelock.CallOpts)
}

// GetMinDelay is a free data retrieval call binding the contract method 0xf27a0c92.
//
// Solidity: function getMinDelay() view returns(uint256 duration)
func (_GovernanceTimelock *GovernanceTimelockCallerSession) GetMinDelay() (*big.Int, error) {
	return _GovernanceTimelock.Contract.GetMinDelay(&_GovernanceTimelock.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GovernanceTimelock *GovernanceTimelockCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GovernanceTimelock.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GovernanceTimelock *GovernanceTimelockSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GovernanceTimelock.Contract.GetRoleAdmin(&_GovernanceTimelock.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GovernanceTimelock *GovernanceTimelockCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GovernanceTimelock.Contract.GetRoleAdmin(&_GovernanceTimelock.CallOpts, role)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xd45c4435.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256 timestamp)
func (_GovernanceTimelock *GovernanceTimelockCaller) GetTimestamp(opts *bind.CallOpts, id [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceTimelock.contract.Call(opts, &out, "getTimestamp", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestamp is a free data retrieval call binding the contract method 0xd45c4435.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256 timestamp)
func (_GovernanceTimelock *GovernanceTimelockSession) GetTimestamp(id [32]byte) (*big.Int, error) {
	return _GovernanceTimelock.Contract.GetTimestamp(&_GovernanceTimelock.CallOpts, id)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xd45c4435.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256 timestamp)
func (_GovernanceTimelock *GovernanceTimelockCallerSession) GetTimestamp(id [32]byte) (*big.Int, error) {
	return _GovernanceTimelock.Contract.GetTimestamp(&_GovernanceTimelock.CallOpts, id)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GovernanceTimelock *GovernanceTimelockCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _GovernanceTimelock.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GovernanceTimelock *GovernanceTimelockSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GovernanceTimelock.Contract.HasRole(&_GovernanceTimelock.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GovernanceTimelock *GovernanceTimelockCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GovernanceTimelock.Contract.HasRole(&_GovernanceTimelock.CallOpts, role, account)
}

// HashOperation is a free data retrieval call binding the contract method 0x8065657f.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32 hash)
func (_GovernanceTimelock *GovernanceTimelockCaller) HashOperation(opts *bind.CallOpts, target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GovernanceTimelock.contract.Call(opts, &out, "hashOperation", target, value, data, predecessor, salt)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOperation is a free data retrieval call binding the contract method 0x8065657f.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32 hash)
func (_GovernanceTimelock *GovernanceTimelockSession) HashOperation(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _GovernanceTimelock.Contract.HashOperation(&_GovernanceTimelock.CallOpts, target, value, data, predecessor, salt)
}

// HashOperation is a free data retrieval call binding the contract method 0x8065657f.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32 hash)
func (_GovernanceTimelock *GovernanceTimelockCallerSession) HashOperation(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _GovernanceTimelock.Contract.HashOperation(&_GovernanceTimelock.CallOpts, target, value, data, predecessor, salt)
}

// HashOperationBatch is a free data retrieval call binding the contract method 0xb1c5f427.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32 hash)
func (_GovernanceTimelock *GovernanceTimelockCaller) HashOperationBatch(opts *bind.CallOpts, targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GovernanceTimelock.contract.Call(opts, &out, "hashOperationBatch", targets, values, payloads, predecessor, salt)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOperationBatch is a free data retrieval call binding the contract method 0xb1c5f427.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32 hash)
func (_GovernanceTimelock *GovernanceTimelockSession) HashOperationBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _GovernanceTimelock.Contract.HashOperationBatch(&_GovernanceTimelock.CallOpts, targets, values, payloads, predecessor, salt)
}

// HashOperationBatch is a free data retrieval call binding the contract method 0xb1c5f427.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32 hash)
func (_GovernanceTimelock *GovernanceTimelockCallerSession) HashOperationBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _GovernanceTimelock.Contract.HashOperationBatch(&_GovernanceTimelock.CallOpts, targets, values, payloads, predecessor, salt)
}

// IsOperation is a free data retrieval call binding the contract method 0x31d50750.
//
// Solidity: function isOperation(bytes32 id) view returns(bool pending)
func (_GovernanceTimelock *GovernanceTimelockCaller) IsOperation(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _GovernanceTimelock.contract.Call(opts, &out, "isOperation", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperation is a free data retrieval call binding the contract method 0x31d50750.
//
// Solidity: function isOperation(bytes32 id) view returns(bool pending)
func (_GovernanceTimelock *GovernanceTimelockSession) IsOperation(id [32]byte) (bool, error) {
	return _GovernanceTimelock.Contract.IsOperation(&_GovernanceTimelock.CallOpts, id)
}

// IsOperation is a free data retrieval call binding the contract method 0x31d50750.
//
// Solidity: function isOperation(bytes32 id) view returns(bool pending)
func (_GovernanceTimelock *GovernanceTimelockCallerSession) IsOperation(id [32]byte) (bool, error) {
	return _GovernanceTimelock.Contract.IsOperation(&_GovernanceTimelock.CallOpts, id)
}

// IsOperationDone is a free data retrieval call binding the contract method 0x2ab0f529.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool done)
func (_GovernanceTimelock *GovernanceTimelockCaller) IsOperationDone(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _GovernanceTimelock.contract.Call(opts, &out, "isOperationDone", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperationDone is a free data retrieval call binding the contract method 0x2ab0f529.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool done)
func (_GovernanceTimelock *GovernanceTimelockSession) IsOperationDone(id [32]byte) (bool, error) {
	return _GovernanceTimelock.Contract.IsOperationDone(&_GovernanceTimelock.CallOpts, id)
}

// IsOperationDone is a free data retrieval call binding the contract method 0x2ab0f529.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool done)
func (_GovernanceTimelock *GovernanceTimelockCallerSession) IsOperationDone(id [32]byte) (bool, error) {
	return _GovernanceTimelock.Contract.IsOperationDone(&_GovernanceTimelock.CallOpts, id)
}

// IsOperationPending is a free data retrieval call binding the contract method 0x584b153e.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool pending)
func (_GovernanceTimelock *GovernanceTimelockCaller) IsOperationPending(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _GovernanceTimelock.contract.Call(opts, &out, "isOperationPending", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperationPending is a free data retrieval call binding the contract method 0x584b153e.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool pending)
func (_GovernanceTimelock *GovernanceTimelockSession) IsOperationPending(id [32]byte) (bool, error) {
	return _GovernanceTimelock.Contract.IsOperationPending(&_GovernanceTimelock.CallOpts, id)
}

// IsOperationPending is a free data retrieval call binding the contract method 0x584b153e.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool pending)
func (_GovernanceTimelock *GovernanceTimelockCallerSession) IsOperationPending(id [32]byte) (bool, error) {
	return _GovernanceTimelock.Contract.IsOperationPending(&_GovernanceTimelock.CallOpts, id)
}

// IsOperationReady is a free data retrieval call binding the contract method 0x13bc9f20.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool ready)
func (_GovernanceTimelock *GovernanceTimelockCaller) IsOperationReady(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _GovernanceTimelock.contract.Call(opts, &out, "isOperationReady", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperationReady is a free data retrieval call binding the contract method 0x13bc9f20.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool ready)
func (_GovernanceTimelock *GovernanceTimelockSession) IsOperationReady(id [32]byte) (bool, error) {
	return _GovernanceTimelock.Contract.IsOperationReady(&_GovernanceTimelock.CallOpts, id)
}

// IsOperationReady is a free data retrieval call binding the contract method 0x13bc9f20.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool ready)
func (_GovernanceTimelock *GovernanceTimelockCallerSession) IsOperationReady(id [32]byte) (bool, error) {
	return _GovernanceTimelock.Contract.IsOperationReady(&_GovernanceTimelock.CallOpts, id)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GovernanceTimelock *GovernanceTimelockCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GovernanceTimelock.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GovernanceTimelock *GovernanceTimelockSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GovernanceTimelock.Contract.SupportsInterface(&_GovernanceTimelock.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GovernanceTimelock *GovernanceTimelockCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GovernanceTimelock.Contract.SupportsInterface(&_GovernanceTimelock.CallOpts, interfaceId)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 id) returns()
func (_GovernanceTimelock *GovernanceTimelockTransactor) Cancel(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _GovernanceTimelock.contract.Transact(opts, "cancel", id)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 id) returns()
func (_GovernanceTimelock *GovernanceTimelockSession) Cancel(id [32]byte) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.Cancel(&_GovernanceTimelock.TransactOpts, id)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 id) returns()
func (_GovernanceTimelock *GovernanceTimelockTransactorSession) Cancel(id [32]byte) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.Cancel(&_GovernanceTimelock.TransactOpts, id)
}

// Execute is a paid mutator transaction binding the contract method 0x134008d3.
//
// Solidity: function execute(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) payable returns()
func (_GovernanceTimelock *GovernanceTimelockTransactor) Execute(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _GovernanceTimelock.contract.Transact(opts, "execute", target, value, data, predecessor, salt)
}

// Execute is a paid mutator transaction binding the contract method 0x134008d3.
//
// Solidity: function execute(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) payable returns()
func (_GovernanceTimelock *GovernanceTimelockSession) Execute(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.Execute(&_GovernanceTimelock.TransactOpts, target, value, data, predecessor, salt)
}

// Execute is a paid mutator transaction binding the contract method 0x134008d3.
//
// Solidity: function execute(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) payable returns()
func (_GovernanceTimelock *GovernanceTimelockTransactorSession) Execute(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.Execute(&_GovernanceTimelock.TransactOpts, target, value, data, predecessor, salt)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe38335e5.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (_GovernanceTimelock *GovernanceTimelockTransactor) ExecuteBatch(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _GovernanceTimelock.contract.Transact(opts, "executeBatch", targets, values, payloads, predecessor, salt)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe38335e5.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (_GovernanceTimelock *GovernanceTimelockSession) ExecuteBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.ExecuteBatch(&_GovernanceTimelock.TransactOpts, targets, values, payloads, predecessor, salt)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe38335e5.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (_GovernanceTimelock *GovernanceTimelockTransactorSession) ExecuteBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.ExecuteBatch(&_GovernanceTimelock.TransactOpts, targets, values, payloads, predecessor, salt)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GovernanceTimelock *GovernanceTimelockTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GovernanceTimelock.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GovernanceTimelock *GovernanceTimelockSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.GrantRole(&_GovernanceTimelock.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GovernanceTimelock *GovernanceTimelockTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.GrantRole(&_GovernanceTimelock.TransactOpts, role, account)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_GovernanceTimelock *GovernanceTimelockTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _GovernanceTimelock.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_GovernanceTimelock *GovernanceTimelockSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.OnERC1155BatchReceived(&_GovernanceTimelock.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_GovernanceTimelock *GovernanceTimelockTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.OnERC1155BatchReceived(&_GovernanceTimelock.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_GovernanceTimelock *GovernanceTimelockTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _GovernanceTimelock.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_GovernanceTimelock *GovernanceTimelockSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.OnERC1155Received(&_GovernanceTimelock.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_GovernanceTimelock *GovernanceTimelockTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.OnERC1155Received(&_GovernanceTimelock.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_GovernanceTimelock *GovernanceTimelockTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _GovernanceTimelock.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_GovernanceTimelock *GovernanceTimelockSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.OnERC721Received(&_GovernanceTimelock.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_GovernanceTimelock *GovernanceTimelockTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.OnERC721Received(&_GovernanceTimelock.TransactOpts, arg0, arg1, arg2, arg3)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GovernanceTimelock *GovernanceTimelockTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GovernanceTimelock.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GovernanceTimelock *GovernanceTimelockSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.RenounceRole(&_GovernanceTimelock.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_GovernanceTimelock *GovernanceTimelockTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.RenounceRole(&_GovernanceTimelock.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GovernanceTimelock *GovernanceTimelockTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GovernanceTimelock.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GovernanceTimelock *GovernanceTimelockSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.RevokeRole(&_GovernanceTimelock.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GovernanceTimelock *GovernanceTimelockTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.RevokeRole(&_GovernanceTimelock.TransactOpts, role, account)
}

// Schedule is a paid mutator transaction binding the contract method 0x01d5062a.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_GovernanceTimelock *GovernanceTimelockTransactor) Schedule(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _GovernanceTimelock.contract.Transact(opts, "schedule", target, value, data, predecessor, salt, delay)
}

// Schedule is a paid mutator transaction binding the contract method 0x01d5062a.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_GovernanceTimelock *GovernanceTimelockSession) Schedule(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.Schedule(&_GovernanceTimelock.TransactOpts, target, value, data, predecessor, salt, delay)
}

// Schedule is a paid mutator transaction binding the contract method 0x01d5062a.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_GovernanceTimelock *GovernanceTimelockTransactorSession) Schedule(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.Schedule(&_GovernanceTimelock.TransactOpts, target, value, data, predecessor, salt, delay)
}

// ScheduleBatch is a paid mutator transaction binding the contract method 0x8f2a0bb0.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_GovernanceTimelock *GovernanceTimelockTransactor) ScheduleBatch(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _GovernanceTimelock.contract.Transact(opts, "scheduleBatch", targets, values, payloads, predecessor, salt, delay)
}

// ScheduleBatch is a paid mutator transaction binding the contract method 0x8f2a0bb0.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_GovernanceTimelock *GovernanceTimelockSession) ScheduleBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.ScheduleBatch(&_GovernanceTimelock.TransactOpts, targets, values, payloads, predecessor, salt, delay)
}

// ScheduleBatch is a paid mutator transaction binding the contract method 0x8f2a0bb0.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_GovernanceTimelock *GovernanceTimelockTransactorSession) ScheduleBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.ScheduleBatch(&_GovernanceTimelock.TransactOpts, targets, values, payloads, predecessor, salt, delay)
}

// UpdateDelay is a paid mutator transaction binding the contract method 0x64d62353.
//
// Solidity: function updateDelay(uint256 newDelay) returns()
func (_GovernanceTimelock *GovernanceTimelockTransactor) UpdateDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _GovernanceTimelock.contract.Transact(opts, "updateDelay", newDelay)
}

// UpdateDelay is a paid mutator transaction binding the contract method 0x64d62353.
//
// Solidity: function updateDelay(uint256 newDelay) returns()
func (_GovernanceTimelock *GovernanceTimelockSession) UpdateDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.UpdateDelay(&_GovernanceTimelock.TransactOpts, newDelay)
}

// UpdateDelay is a paid mutator transaction binding the contract method 0x64d62353.
//
// Solidity: function updateDelay(uint256 newDelay) returns()
func (_GovernanceTimelock *GovernanceTimelockTransactorSession) UpdateDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.UpdateDelay(&_GovernanceTimelock.TransactOpts, newDelay)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GovernanceTimelock *GovernanceTimelockTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceTimelock.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GovernanceTimelock *GovernanceTimelockSession) Receive() (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.Receive(&_GovernanceTimelock.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GovernanceTimelock *GovernanceTimelockTransactorSession) Receive() (*types.Transaction, error) {
	return _GovernanceTimelock.Contract.Receive(&_GovernanceTimelock.TransactOpts)
}

// GovernanceTimelockCallExecutedIterator is returned from FilterCallExecuted and is used to iterate over the raw logs and unpacked data for CallExecuted events raised by the GovernanceTimelock contract.
type GovernanceTimelockCallExecutedIterator struct {
	Event *GovernanceTimelockCallExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceTimelockCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceTimelockCallExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceTimelockCallExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceTimelockCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceTimelockCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceTimelockCallExecuted represents a CallExecuted event raised by the GovernanceTimelock contract.
type GovernanceTimelockCallExecuted struct {
	Id     [32]byte
	Index  *big.Int
	Target common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallExecuted is a free log retrieval operation binding the contract event 0xc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (_GovernanceTimelock *GovernanceTimelockFilterer) FilterCallExecuted(opts *bind.FilterOpts, id [][32]byte, index []*big.Int) (*GovernanceTimelockCallExecutedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _GovernanceTimelock.contract.FilterLogs(opts, "CallExecuted", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceTimelockCallExecutedIterator{contract: _GovernanceTimelock.contract, event: "CallExecuted", logs: logs, sub: sub}, nil
}

// WatchCallExecuted is a free log subscription operation binding the contract event 0xc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (_GovernanceTimelock *GovernanceTimelockFilterer) WatchCallExecuted(opts *bind.WatchOpts, sink chan<- *GovernanceTimelockCallExecuted, id [][32]byte, index []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _GovernanceTimelock.contract.WatchLogs(opts, "CallExecuted", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceTimelockCallExecuted)
				if err := _GovernanceTimelock.contract.UnpackLog(event, "CallExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallExecuted is a log parse operation binding the contract event 0xc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (_GovernanceTimelock *GovernanceTimelockFilterer) ParseCallExecuted(log types.Log) (*GovernanceTimelockCallExecuted, error) {
	event := new(GovernanceTimelockCallExecuted)
	if err := _GovernanceTimelock.contract.UnpackLog(event, "CallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceTimelockCallScheduledIterator is returned from FilterCallScheduled and is used to iterate over the raw logs and unpacked data for CallScheduled events raised by the GovernanceTimelock contract.
type GovernanceTimelockCallScheduledIterator struct {
	Event *GovernanceTimelockCallScheduled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceTimelockCallScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceTimelockCallScheduled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceTimelockCallScheduled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceTimelockCallScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceTimelockCallScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceTimelockCallScheduled represents a CallScheduled event raised by the GovernanceTimelock contract.
type GovernanceTimelockCallScheduled struct {
	Id          [32]byte
	Index       *big.Int
	Target      common.Address
	Value       *big.Int
	Data        []byte
	Predecessor [32]byte
	Delay       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCallScheduled is a free log retrieval operation binding the contract event 0x4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (_GovernanceTimelock *GovernanceTimelockFilterer) FilterCallScheduled(opts *bind.FilterOpts, id [][32]byte, index []*big.Int) (*GovernanceTimelockCallScheduledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _GovernanceTimelock.contract.FilterLogs(opts, "CallScheduled", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceTimelockCallScheduledIterator{contract: _GovernanceTimelock.contract, event: "CallScheduled", logs: logs, sub: sub}, nil
}

// WatchCallScheduled is a free log subscription operation binding the contract event 0x4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (_GovernanceTimelock *GovernanceTimelockFilterer) WatchCallScheduled(opts *bind.WatchOpts, sink chan<- *GovernanceTimelockCallScheduled, id [][32]byte, index []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _GovernanceTimelock.contract.WatchLogs(opts, "CallScheduled", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceTimelockCallScheduled)
				if err := _GovernanceTimelock.contract.UnpackLog(event, "CallScheduled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCallScheduled is a log parse operation binding the contract event 0x4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (_GovernanceTimelock *GovernanceTimelockFilterer) ParseCallScheduled(log types.Log) (*GovernanceTimelockCallScheduled, error) {
	event := new(GovernanceTimelockCallScheduled)
	if err := _GovernanceTimelock.contract.UnpackLog(event, "CallScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceTimelockCancelledIterator is returned from FilterCancelled and is used to iterate over the raw logs and unpacked data for Cancelled events raised by the GovernanceTimelock contract.
type GovernanceTimelockCancelledIterator struct {
	Event *GovernanceTimelockCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceTimelockCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceTimelockCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceTimelockCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceTimelockCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceTimelockCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceTimelockCancelled represents a Cancelled event raised by the GovernanceTimelock contract.
type GovernanceTimelockCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCancelled is a free log retrieval operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (_GovernanceTimelock *GovernanceTimelockFilterer) FilterCancelled(opts *bind.FilterOpts, id [][32]byte) (*GovernanceTimelockCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _GovernanceTimelock.contract.FilterLogs(opts, "Cancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceTimelockCancelledIterator{contract: _GovernanceTimelock.contract, event: "Cancelled", logs: logs, sub: sub}, nil
}

// WatchCancelled is a free log subscription operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (_GovernanceTimelock *GovernanceTimelockFilterer) WatchCancelled(opts *bind.WatchOpts, sink chan<- *GovernanceTimelockCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _GovernanceTimelock.contract.WatchLogs(opts, "Cancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceTimelockCancelled)
				if err := _GovernanceTimelock.contract.UnpackLog(event, "Cancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCancelled is a log parse operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (_GovernanceTimelock *GovernanceTimelockFilterer) ParseCancelled(log types.Log) (*GovernanceTimelockCancelled, error) {
	event := new(GovernanceTimelockCancelled)
	if err := _GovernanceTimelock.contract.UnpackLog(event, "Cancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceTimelockMinDelayChangeIterator is returned from FilterMinDelayChange and is used to iterate over the raw logs and unpacked data for MinDelayChange events raised by the GovernanceTimelock contract.
type GovernanceTimelockMinDelayChangeIterator struct {
	Event *GovernanceTimelockMinDelayChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceTimelockMinDelayChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceTimelockMinDelayChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceTimelockMinDelayChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceTimelockMinDelayChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceTimelockMinDelayChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceTimelockMinDelayChange represents a MinDelayChange event raised by the GovernanceTimelock contract.
type GovernanceTimelockMinDelayChange struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMinDelayChange is a free log retrieval operation binding the contract event 0x11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (_GovernanceTimelock *GovernanceTimelockFilterer) FilterMinDelayChange(opts *bind.FilterOpts) (*GovernanceTimelockMinDelayChangeIterator, error) {

	logs, sub, err := _GovernanceTimelock.contract.FilterLogs(opts, "MinDelayChange")
	if err != nil {
		return nil, err
	}
	return &GovernanceTimelockMinDelayChangeIterator{contract: _GovernanceTimelock.contract, event: "MinDelayChange", logs: logs, sub: sub}, nil
}

// WatchMinDelayChange is a free log subscription operation binding the contract event 0x11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (_GovernanceTimelock *GovernanceTimelockFilterer) WatchMinDelayChange(opts *bind.WatchOpts, sink chan<- *GovernanceTimelockMinDelayChange) (event.Subscription, error) {

	logs, sub, err := _GovernanceTimelock.contract.WatchLogs(opts, "MinDelayChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceTimelockMinDelayChange)
				if err := _GovernanceTimelock.contract.UnpackLog(event, "MinDelayChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinDelayChange is a log parse operation binding the contract event 0x11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (_GovernanceTimelock *GovernanceTimelockFilterer) ParseMinDelayChange(log types.Log) (*GovernanceTimelockMinDelayChange, error) {
	event := new(GovernanceTimelockMinDelayChange)
	if err := _GovernanceTimelock.contract.UnpackLog(event, "MinDelayChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceTimelockRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the GovernanceTimelock contract.
type GovernanceTimelockRoleAdminChangedIterator struct {
	Event *GovernanceTimelockRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceTimelockRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceTimelockRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceTimelockRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceTimelockRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceTimelockRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceTimelockRoleAdminChanged represents a RoleAdminChanged event raised by the GovernanceTimelock contract.
type GovernanceTimelockRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GovernanceTimelock *GovernanceTimelockFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GovernanceTimelockRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _GovernanceTimelock.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceTimelockRoleAdminChangedIterator{contract: _GovernanceTimelock.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GovernanceTimelock *GovernanceTimelockFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GovernanceTimelockRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _GovernanceTimelock.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceTimelockRoleAdminChanged)
				if err := _GovernanceTimelock.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GovernanceTimelock *GovernanceTimelockFilterer) ParseRoleAdminChanged(log types.Log) (*GovernanceTimelockRoleAdminChanged, error) {
	event := new(GovernanceTimelockRoleAdminChanged)
	if err := _GovernanceTimelock.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceTimelockRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the GovernanceTimelock contract.
type GovernanceTimelockRoleGrantedIterator struct {
	Event *GovernanceTimelockRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceTimelockRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceTimelockRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceTimelockRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceTimelockRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceTimelockRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceTimelockRoleGranted represents a RoleGranted event raised by the GovernanceTimelock contract.
type GovernanceTimelockRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GovernanceTimelock *GovernanceTimelockFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GovernanceTimelockRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GovernanceTimelock.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceTimelockRoleGrantedIterator{contract: _GovernanceTimelock.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GovernanceTimelock *GovernanceTimelockFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GovernanceTimelockRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GovernanceTimelock.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceTimelockRoleGranted)
				if err := _GovernanceTimelock.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GovernanceTimelock *GovernanceTimelockFilterer) ParseRoleGranted(log types.Log) (*GovernanceTimelockRoleGranted, error) {
	event := new(GovernanceTimelockRoleGranted)
	if err := _GovernanceTimelock.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceTimelockRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the GovernanceTimelock contract.
type GovernanceTimelockRoleRevokedIterator struct {
	Event *GovernanceTimelockRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovernanceTimelockRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceTimelockRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovernanceTimelockRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovernanceTimelockRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceTimelockRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceTimelockRoleRevoked represents a RoleRevoked event raised by the GovernanceTimelock contract.
type GovernanceTimelockRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GovernanceTimelock *GovernanceTimelockFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GovernanceTimelockRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GovernanceTimelock.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceTimelockRoleRevokedIterator{contract: _GovernanceTimelock.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GovernanceTimelock *GovernanceTimelockFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GovernanceTimelockRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GovernanceTimelock.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceTimelockRoleRevoked)
				if err := _GovernanceTimelock.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GovernanceTimelock *GovernanceTimelockFilterer) ParseRoleRevoked(log types.Log) (*GovernanceTimelockRoleRevoked, error) {
	event := new(GovernanceTimelockRoleRevoked)
	if err := _GovernanceTimelock.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
