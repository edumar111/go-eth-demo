// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package expertise

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ExpertiseABI is the input ABI used to generate the binding from.
const ExpertiseABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getUsers\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"category\",\"type\":\"string\"}],\"name\":\"getGoals\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"levelManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"string\"},{\"name\":\"_key\",\"type\":\"string\"},{\"name\":\"_value\",\"type\":\"string\"}],\"name\":\"setAchievementInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getAllScore\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSubject\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"string\"},{\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"getAchievementInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"category\",\"type\":\"string\"}],\"name\":\"getScore\",\"outputs\":[{\"name\":\"score\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"category\",\"type\":\"string\"}],\"name\":\"getAchievements\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"expertise\",\"outputs\":[{\"name\":\"score\",\"type\":\"uint256\"},{\"name\":\"level\",\"type\":\"string\"},{\"name\":\"rank\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"eternalStorage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOffices\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"getUserInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getAchievementsIds\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ABI_URL\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"category\",\"type\":\"string\"}],\"name\":\"addAchievement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newLevelManager\",\"type\":\"address\"}],\"name\":\"setLevelManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"category\",\"type\":\"string\"}],\"name\":\"getAchievementsIdsCategory\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"categories\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"source\",\"type\":\"string\"}],\"name\":\"stringToBytes32\",\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"string\"},{\"name\":\"category\",\"type\":\"string\"}],\"name\":\"removeAchievement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"string\"},{\"name\":\"_value\",\"type\":\"string\"}],\"name\":\"setUserInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"string\"},{\"name\":\"_key\",\"type\":\"string\"},{\"name\":\"_value\",\"type\":\"bytes32\"}],\"name\":\"setAchievementInfoBytes32\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_levelManager\",\"type\":\"address\"},{\"name\":\"_eternalStorage\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"oldLevelManager\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newLevelManager\",\"type\":\"address\"}],\"name\":\"LevelManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"category\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"}],\"name\":\"AchievementAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"category\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"}],\"name\":\"AchievementRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_idAchievement\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_key\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"string\"}],\"name\":\"AchievementInfoAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_idAchievement\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_key\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"bytes32\"}],\"name\":\"AchievementInfoAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_key\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"string\"}],\"name\":\"InfoAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"}],\"name\":\"UserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"formerOwner\",\"type\":\"address\"}],\"name\":\"OwnerRemoved\",\"type\":\"event\"}]"

// ExpertiseBin is the compiled bytecode used for deploying new contracts.
const ExpertiseBin = `608060405260043610610107576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063014b20d41461010c578063022914a71461016d5780631785f53c146101c85780631cc9c9081461020b57806324d7806c1461026c5780632f54bf6e146102c75780633d6b73011461032257806340129a40146103b2578063429b62e51461041b5780634d134f2f146104765780636add6d1b146104e157806370480275146105675780637065cb48146105aa5780637b2e0046146105ed5780639d84ae69146106565780639faf6fb6146106c7578063b15be2f514610710578063dcb4876f14610727578063f8241d7e1461077a575b600080fd5b34801561011857600080fd5b5061015360048036038101908080356000191690602001909291908035906020019082018035906020019190919293919293905050506107e5565b604051808215151515815260200191505060405180910390f35b34801561017957600080fd5b506101ae600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061083d565b604051808215151515815260200191505060405180910390f35b3480156101d457600080fd5b50610209600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061085d565b005b34801561021757600080fd5b50610252600480360381019080803560001916906020019092919080359060200190820180359060200191909192939192939050505061090f565b604051808215151515815260200191505060405180910390f35b34801561027857600080fd5b506102ad600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610967565b604051808215151515815260200191505060405180910390f35b3480156102d357600080fd5b50610308600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109bd565b604051808215151515815260200191505060405180910390f35b34801561032e57600080fd5b5061035b600480360381019080803590602001908201803590602001919091929391929390505050610a12565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561039e578082015181840152602081019050610383565b505050509050019250505060405180910390f35b3480156103be57600080fd5b506104016004803603810190808035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a54565b604051808215151515815260200191505060405180910390f35b34801561042757600080fd5b5061045c600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a7c565b604051808215151515815260200191505060405180910390f35b34801561048257600080fd5b506104c7600480360381019080803590602001908201803590602001919091929391929390803590602001908201803590602001919091929391929390505050610a9c565b604051808215151515815260200191505060405180910390f35b3480156104ed57600080fd5b506105106004803603810190808035600019169060200190929190505050610b24565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610553578082015181840152602081019050610538565b505050509050019250505060405180910390f35b34801561057357600080fd5b506105a8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b36565b005b3480156105b657600080fd5b506105eb600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610be7565b005b3480156105f957600080fd5b5061063c6004803603810190808035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c98565b604051808215151515815260200191505060405180910390f35b34801561066257600080fd5b506106856004803603810190808035600019169060200190929190505050610cc0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156106d357600080fd5b506106f66004803603810190808035600019169060200190929190505050610cd2565b604051808215151515815260200191505060405180910390f35b34801561071c57600080fd5b50610725610cf8565b005b34801561073357600080fd5b50610760600480360381019080803590602001908201803590602001919091929391929390505050610da8565b604051808215151515815260200191505060405180910390f35b34801561078657600080fd5b506107cb600480360381019080803590602001908201803590602001919091929391929390803590602001908201803590602001919091929391929390505050610dfe565b604051808215151515815260200191505060405180910390f35b60006107f033610967565b15156107fb57600080fd5b61083484848480806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050610e86565b90509392505050565b60006020528060005260406000206000915054906101000a900460ff1681565b610866336109bd565b151561087157600080fd5b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167fa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f60405160405180910390a250565b600061091a33610967565b151561092557600080fd5b61095e848484808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050506110ee565b90509392505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6060610a4c8383808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050506113b2565b905092915050565b6000610a5f33610967565b1515610a6a57600080fd5b610a748383611479565b905092915050565b60016020528060005260406000206000915054906101000a900460ff1681565b6000610aa733610967565b1515610ab257600080fd5b610b1a8585808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050508484808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050506114df565b9050949350505050565b6060610b2f82611574565b9050919050565b610b3f336109bd565b1515610b4a57600080fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167f44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e33960405160405180910390a250565b610bf0336109bd565b1515610bfb57600080fd5b60016000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff167f994a936646fe87ffe4f1e469d3d6aa417d6b855598397f323de5b449f765f0c360405160405180910390a250565b6000610ca333610967565b1515610cae57600080fd5b610cb8838361161d565b905092915050565b6000610ccb826116ab565b9050919050565b6000610cdd33610967565b1515610ce857600080fd5b610cf1826116f0565b9050919050565b610d01336109bd565b1515610d0c57600080fd5b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055503373ffffffffffffffffffffffffffffffffffffffff167f58619076adf5bb0943d100ef88d52d7c3fd691b19d3a9071b555b651fbf418da60405160405180910390a2565b6000610db333610967565b1515610dbe57600080fd5b610df6838380806020026020016040519081016040528093929190818152602001838360200280828437820191505050505050611739565b905092915050565b6000610e0933610967565b1515610e1457600080fd5b610e7c8585808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050508484808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050506117a0565b9050949350505050565b600080606060008060009350600360008860001916600019168152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015610f2d57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610ee3575b50505050509250600091505b85518210156110e157600090505b82518110156110d4578582815181101515610f5e57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff168382815181101515610f8c57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1614156110c75760036000886000191660001916815260200190815260200160002081815481101515610fda57fe5b9060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905582600184510381518110151561101957fe5b906020019060200201516003600089600019166000191681526020019081526020016000208281548110151561104b57fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506003600088600019166000191681526020019081526020016000208054809190600190036110c19190611835565b50600193505b8080600101915050610f47565b8180600101925050610f39565b8394505050505092915050565b6000806000606060008060006001955060019450600360008a6000191660001916815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561119d57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611153575b5050505050935060008451141561126457600092505b875183101561125f57600360008a6000191660001916815260200190815260200160002088848151811015156111e557fe5b9060200190602002015190806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505082806001019350506111b3565b6113a3565b600091505b87518210156113a25760019450600090505b83518110156112f857838181518110151561129257fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1688838151811015156112c057fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1614156112eb57600094505b808060010191505061127b565b841561139557600360008a60001916600019168152602001908152602001600020888381518110151561132757fe5b9060200190602002015190806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b8180600101925050611269565b5b85965050505050505092915050565b606080600083516040519080825280602002602001820160405280156113e75781602001602082028038833980820191505090505b509150600090505b835181101561146f57611418848281518110151561140957fe5b906020019060200201516116ab565b828281518110151561142657fe5b9060200190602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080806001019150506113ef565b8192505050919050565b60008160026000856000191660001916815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001905092915050565b600080600080845186511415156114f9576000935061156b565b60019250600091505b855182101561156757611543868381518110151561151c57fe5b90602001906020020151868481518110151561153457fe5b9060200190602002015161161d565b9050828015611550575080155b1561155a57600092505b8180600101925050611502565b8293505b50505092915050565b606060036000836000191660001916815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561161157602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116115c7575b50505050509050919050565b60008073ffffffffffffffffffffffffffffffffffffffff1660026000856000191660001916815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561169857600090506116a5565b6116a28383611479565b90505b92915050565b600060026000836000191660001916815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b600060026000836000191660001916815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560019050919050565b60008060008060019250600091505b845182101561179557611771858381518110151561176257fe5b906020019060200201516116f0565b905082801561177e575080155b1561178857600092505b8180600101925050611748565b829350505050919050565b600080600080845186511415156117ba576000935061182c565b60019250600091505b85518210156118285761180486838151811015156117dd57fe5b9060200190602002015186848151811015156117f557fe5b90602001906020020151611479565b9050828015611811575080155b1561181b57600092505b81806001019250506117c3565b8293505b50505092915050565b81548183558181111561185c5781836000526020600020918201910161185b9190611861565b5b505050565b61188391905b8082111561187f576000816000905550600101611867565b5090565b905600a165627a7a723058207004473d1929bd1bf52015fdd5b34ba09210f9e240d11fa01889dcadb77b89a60029`

// DeployExpertise deploys a new Ethereum contract, binding an instance of Expertise to it.
func DeployExpertise(auth *bind.TransactOpts, backend bind.ContractBackend, _levelManager common.Address, _eternalStorage common.Address) (common.Address, *types.Transaction, *Expertise, error) {
	parsed, err := abi.JSON(strings.NewReader(ExpertiseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExpertiseBin), backend, _levelManager, _eternalStorage)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Expertise{ExpertiseCaller: ExpertiseCaller{contract: contract}, ExpertiseTransactor: ExpertiseTransactor{contract: contract}, ExpertiseFilterer: ExpertiseFilterer{contract: contract}}, nil
}

// Expertise is an auto generated Go binding around an Ethereum contract.
type Expertise struct {
	ExpertiseCaller     // Read-only binding to the contract
	ExpertiseTransactor // Write-only binding to the contract
	ExpertiseFilterer   // Log filterer for contract events
}

// ExpertiseCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExpertiseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExpertiseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExpertiseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExpertiseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExpertiseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExpertiseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExpertiseSession struct {
	Contract     *Expertise        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExpertiseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExpertiseCallerSession struct {
	Contract *ExpertiseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ExpertiseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExpertiseTransactorSession struct {
	Contract     *ExpertiseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ExpertiseRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExpertiseRaw struct {
	Contract *Expertise // Generic contract binding to access the raw methods on
}

// ExpertiseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExpertiseCallerRaw struct {
	Contract *ExpertiseCaller // Generic read-only contract binding to access the raw methods on
}

// ExpertiseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExpertiseTransactorRaw struct {
	Contract *ExpertiseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExpertise creates a new instance of Expertise, bound to a specific deployed contract.
func NewExpertise(address common.Address, backend bind.ContractBackend) (*Expertise, error) {
	contract, err := bindExpertise(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Expertise{ExpertiseCaller: ExpertiseCaller{contract: contract}, ExpertiseTransactor: ExpertiseTransactor{contract: contract}, ExpertiseFilterer: ExpertiseFilterer{contract: contract}}, nil
}

// NewExpertiseCaller creates a new read-only instance of Expertise, bound to a specific deployed contract.
func NewExpertiseCaller(address common.Address, caller bind.ContractCaller) (*ExpertiseCaller, error) {
	contract, err := bindExpertise(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExpertiseCaller{contract: contract}, nil
}

// NewExpertiseTransactor creates a new write-only instance of Expertise, bound to a specific deployed contract.
func NewExpertiseTransactor(address common.Address, transactor bind.ContractTransactor) (*ExpertiseTransactor, error) {
	contract, err := bindExpertise(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExpertiseTransactor{contract: contract}, nil
}

// NewExpertiseFilterer creates a new log filterer instance of Expertise, bound to a specific deployed contract.
func NewExpertiseFilterer(address common.Address, filterer bind.ContractFilterer) (*ExpertiseFilterer, error) {
	contract, err := bindExpertise(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExpertiseFilterer{contract: contract}, nil
}

// bindExpertise binds a generic wrapper to an already deployed contract.
func bindExpertise(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExpertiseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Expertise *ExpertiseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Expertise.Contract.ExpertiseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Expertise *ExpertiseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Expertise.Contract.ExpertiseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Expertise *ExpertiseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Expertise.Contract.ExpertiseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Expertise *ExpertiseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Expertise.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Expertise *ExpertiseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Expertise.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Expertise *ExpertiseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Expertise.Contract.contract.Transact(opts, method, params...)
}

// ABIURL is a free data retrieval call binding the contract method 0xa2d1183a.
//
// Solidity: function ABI_URL() constant returns(string)
func (_Expertise *ExpertiseCaller) ABIURL(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Expertise.contract.Call(opts, out, "ABI_URL")
	return *ret0, err
}

// ABIURL is a free data retrieval call binding the contract method 0xa2d1183a.
//
// Solidity: function ABI_URL() constant returns(string)
func (_Expertise *ExpertiseSession) ABIURL() (string, error) {
	return _Expertise.Contract.ABIURL(&_Expertise.CallOpts)
}

// ABIURL is a free data retrieval call binding the contract method 0xa2d1183a.
//
// Solidity: function ABI_URL() constant returns(string)
func (_Expertise *ExpertiseCallerSession) ABIURL() (string, error) {
	return _Expertise.Contract.ABIURL(&_Expertise.CallOpts)
}

// Categories is a free data retrieval call binding the contract method 0xc6cdbe5e.
//
// Solidity: function categories(uint256 ) constant returns(string)
func (_Expertise *ExpertiseCaller) Categories(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Expertise.contract.Call(opts, out, "categories", arg0)
	return *ret0, err
}

// Categories is a free data retrieval call binding the contract method 0xc6cdbe5e.
//
// Solidity: function categories(uint256 ) constant returns(string)
func (_Expertise *ExpertiseSession) Categories(arg0 *big.Int) (string, error) {
	return _Expertise.Contract.Categories(&_Expertise.CallOpts, arg0)
}

// Categories is a free data retrieval call binding the contract method 0xc6cdbe5e.
//
// Solidity: function categories(uint256 ) constant returns(string)
func (_Expertise *ExpertiseCallerSession) Categories(arg0 *big.Int) (string, error) {
	return _Expertise.Contract.Categories(&_Expertise.CallOpts, arg0)
}

// EternalStorage is a free data retrieval call binding the contract method 0x98ff9c54.
//
// Solidity: function eternalStorage() constant returns(address)
func (_Expertise *ExpertiseCaller) EternalStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Expertise.contract.Call(opts, out, "eternalStorage")
	return *ret0, err
}

// EternalStorage is a free data retrieval call binding the contract method 0x98ff9c54.
//
// Solidity: function eternalStorage() constant returns(address)
func (_Expertise *ExpertiseSession) EternalStorage() (common.Address, error) {
	return _Expertise.Contract.EternalStorage(&_Expertise.CallOpts)
}

// EternalStorage is a free data retrieval call binding the contract method 0x98ff9c54.
//
// Solidity: function eternalStorage() constant returns(address)
func (_Expertise *ExpertiseCallerSession) EternalStorage() (common.Address, error) {
	return _Expertise.Contract.EternalStorage(&_Expertise.CallOpts)
}

// Expertise is a free data retrieval call binding the contract method 0x9212d567.
//
// Solidity: function expertise(address user) constant returns(uint256 score, string level, string rank)
func (_Expertise *ExpertiseCaller) Expertise(opts *bind.CallOpts, user common.Address) (struct {
	Score *big.Int
	Level string
	Rank  string
}, error) {
	ret := new(struct {
		Score *big.Int
		Level string
		Rank  string
	})
	out := ret
	err := _Expertise.contract.Call(opts, out, "expertise", user)
	return *ret, err
}

// Expertise is a free data retrieval call binding the contract method 0x9212d567.
//
// Solidity: function expertise(address user) constant returns(uint256 score, string level, string rank)
func (_Expertise *ExpertiseSession) Expertise(user common.Address) (struct {
	Score *big.Int
	Level string
	Rank  string
}, error) {
	return _Expertise.Contract.Expertise(&_Expertise.CallOpts, user)
}

// Expertise is a free data retrieval call binding the contract method 0x9212d567.
//
// Solidity: function expertise(address user) constant returns(uint256 score, string level, string rank)
func (_Expertise *ExpertiseCallerSession) Expertise(user common.Address) (struct {
	Score *big.Int
	Level string
	Rank  string
}, error) {
	return _Expertise.Contract.Expertise(&_Expertise.CallOpts, user)
}

// GetAchievementInfo is a free data retrieval call binding the contract method 0x5909b44d.
//
// Solidity: function getAchievementInfo(string _id, string _key) constant returns(bytes32)
func (_Expertise *ExpertiseCaller) GetAchievementInfo(opts *bind.CallOpts, _id string, _key string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Expertise.contract.Call(opts, out, "getAchievementInfo", _id, _key)
	return *ret0, err
}

// GetAchievementInfo is a free data retrieval call binding the contract method 0x5909b44d.
//
// Solidity: function getAchievementInfo(string _id, string _key) constant returns(bytes32)
func (_Expertise *ExpertiseSession) GetAchievementInfo(_id string, _key string) ([32]byte, error) {
	return _Expertise.Contract.GetAchievementInfo(&_Expertise.CallOpts, _id, _key)
}

// GetAchievementInfo is a free data retrieval call binding the contract method 0x5909b44d.
//
// Solidity: function getAchievementInfo(string _id, string _key) constant returns(bytes32)
func (_Expertise *ExpertiseCallerSession) GetAchievementInfo(_id string, _key string) ([32]byte, error) {
	return _Expertise.Contract.GetAchievementInfo(&_Expertise.CallOpts, _id, _key)
}

// GetAchievements is a free data retrieval call binding the contract method 0x7a69c181.
//
// Solidity: function getAchievements(address user, string category) constant returns(bytes32[], bytes32[], bytes32[], uint256[], bytes32[])
func (_Expertise *ExpertiseCaller) GetAchievements(opts *bind.CallOpts, user common.Address, category string) ([][32]byte, [][32]byte, [][32]byte, []*big.Int, [][32]byte, error) {
	var (
		ret0 = new([][32]byte)
		ret1 = new([][32]byte)
		ret2 = new([][32]byte)
		ret3 = new([]*big.Int)
		ret4 = new([][32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Expertise.contract.Call(opts, out, "getAchievements", user, category)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetAchievements is a free data retrieval call binding the contract method 0x7a69c181.
//
// Solidity: function getAchievements(address user, string category) constant returns(bytes32[], bytes32[], bytes32[], uint256[], bytes32[])
func (_Expertise *ExpertiseSession) GetAchievements(user common.Address, category string) ([][32]byte, [][32]byte, [][32]byte, []*big.Int, [][32]byte, error) {
	return _Expertise.Contract.GetAchievements(&_Expertise.CallOpts, user, category)
}

// GetAchievements is a free data retrieval call binding the contract method 0x7a69c181.
//
// Solidity: function getAchievements(address user, string category) constant returns(bytes32[], bytes32[], bytes32[], uint256[], bytes32[])
func (_Expertise *ExpertiseCallerSession) GetAchievements(user common.Address, category string) ([][32]byte, [][32]byte, [][32]byte, []*big.Int, [][32]byte, error) {
	return _Expertise.Contract.GetAchievements(&_Expertise.CallOpts, user, category)
}

// GetAchievementsIds is a free data retrieval call binding the contract method 0xa18d9d43.
//
// Solidity: function getAchievementsIds(address user) constant returns(bytes32[], bytes32[])
func (_Expertise *ExpertiseCaller) GetAchievementsIds(opts *bind.CallOpts, user common.Address) ([][32]byte, [][32]byte, error) {
	var (
		ret0 = new([][32]byte)
		ret1 = new([][32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Expertise.contract.Call(opts, out, "getAchievementsIds", user)
	return *ret0, *ret1, err
}

// GetAchievementsIds is a free data retrieval call binding the contract method 0xa18d9d43.
//
// Solidity: function getAchievementsIds(address user) constant returns(bytes32[], bytes32[])
func (_Expertise *ExpertiseSession) GetAchievementsIds(user common.Address) ([][32]byte, [][32]byte, error) {
	return _Expertise.Contract.GetAchievementsIds(&_Expertise.CallOpts, user)
}

// GetAchievementsIds is a free data retrieval call binding the contract method 0xa18d9d43.
//
// Solidity: function getAchievementsIds(address user) constant returns(bytes32[], bytes32[])
func (_Expertise *ExpertiseCallerSession) GetAchievementsIds(user common.Address) ([][32]byte, [][32]byte, error) {
	return _Expertise.Contract.GetAchievementsIds(&_Expertise.CallOpts, user)
}

// GetAchievementsIdsCategory is a free data retrieval call binding the contract method 0xc58ae177.
//
// Solidity: function getAchievementsIdsCategory(address user, string category) constant returns(bytes32[], bytes32[])
func (_Expertise *ExpertiseCaller) GetAchievementsIdsCategory(opts *bind.CallOpts, user common.Address, category string) ([][32]byte, [][32]byte, error) {
	var (
		ret0 = new([][32]byte)
		ret1 = new([][32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Expertise.contract.Call(opts, out, "getAchievementsIdsCategory", user, category)
	return *ret0, *ret1, err
}

// GetAchievementsIdsCategory is a free data retrieval call binding the contract method 0xc58ae177.
//
// Solidity: function getAchievementsIdsCategory(address user, string category) constant returns(bytes32[], bytes32[])
func (_Expertise *ExpertiseSession) GetAchievementsIdsCategory(user common.Address, category string) ([][32]byte, [][32]byte, error) {
	return _Expertise.Contract.GetAchievementsIdsCategory(&_Expertise.CallOpts, user, category)
}

// GetAchievementsIdsCategory is a free data retrieval call binding the contract method 0xc58ae177.
//
// Solidity: function getAchievementsIdsCategory(address user, string category) constant returns(bytes32[], bytes32[])
func (_Expertise *ExpertiseCallerSession) GetAchievementsIdsCategory(user common.Address, category string) ([][32]byte, [][32]byte, error) {
	return _Expertise.Contract.GetAchievementsIdsCategory(&_Expertise.CallOpts, user, category)
}

// GetAllScore is a free data retrieval call binding the contract method 0x2d896438.
//
// Solidity: function getAllScore(address user) constant returns(uint256)
func (_Expertise *ExpertiseCaller) GetAllScore(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Expertise.contract.Call(opts, out, "getAllScore", user)
	return *ret0, err
}

// GetAllScore is a free data retrieval call binding the contract method 0x2d896438.
//
// Solidity: function getAllScore(address user) constant returns(uint256)
func (_Expertise *ExpertiseSession) GetAllScore(user common.Address) (*big.Int, error) {
	return _Expertise.Contract.GetAllScore(&_Expertise.CallOpts, user)
}

// GetAllScore is a free data retrieval call binding the contract method 0x2d896438.
//
// Solidity: function getAllScore(address user) constant returns(uint256)
func (_Expertise *ExpertiseCallerSession) GetAllScore(user common.Address) (*big.Int, error) {
	return _Expertise.Contract.GetAllScore(&_Expertise.CallOpts, user)
}

// GetGoals is a free data retrieval call binding the contract method 0x0c282d0e.
//
// Solidity: function getGoals(string category) constant returns(bytes32[], bytes32[], bytes32[], bytes32[], uint256[])
func (_Expertise *ExpertiseCaller) GetGoals(opts *bind.CallOpts, category string) ([][32]byte, [][32]byte, [][32]byte, [][32]byte, []*big.Int, error) {
	var (
		ret0 = new([][32]byte)
		ret1 = new([][32]byte)
		ret2 = new([][32]byte)
		ret3 = new([][32]byte)
		ret4 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Expertise.contract.Call(opts, out, "getGoals", category)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetGoals is a free data retrieval call binding the contract method 0x0c282d0e.
//
// Solidity: function getGoals(string category) constant returns(bytes32[], bytes32[], bytes32[], bytes32[], uint256[])
func (_Expertise *ExpertiseSession) GetGoals(category string) ([][32]byte, [][32]byte, [][32]byte, [][32]byte, []*big.Int, error) {
	return _Expertise.Contract.GetGoals(&_Expertise.CallOpts, category)
}

// GetGoals is a free data retrieval call binding the contract method 0x0c282d0e.
//
// Solidity: function getGoals(string category) constant returns(bytes32[], bytes32[], bytes32[], bytes32[], uint256[])
func (_Expertise *ExpertiseCallerSession) GetGoals(category string) ([][32]byte, [][32]byte, [][32]byte, [][32]byte, []*big.Int, error) {
	return _Expertise.Contract.GetGoals(&_Expertise.CallOpts, category)
}

// GetOffices is a free data retrieval call binding the contract method 0x9e718f3a.
//
// Solidity: function getOffices() constant returns(bytes32[])
func (_Expertise *ExpertiseCaller) GetOffices(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Expertise.contract.Call(opts, out, "getOffices")
	return *ret0, err
}

// GetOffices is a free data retrieval call binding the contract method 0x9e718f3a.
//
// Solidity: function getOffices() constant returns(bytes32[])
func (_Expertise *ExpertiseSession) GetOffices() ([][32]byte, error) {
	return _Expertise.Contract.GetOffices(&_Expertise.CallOpts)
}

// GetOffices is a free data retrieval call binding the contract method 0x9e718f3a.
//
// Solidity: function getOffices() constant returns(bytes32[])
func (_Expertise *ExpertiseCallerSession) GetOffices() ([][32]byte, error) {
	return _Expertise.Contract.GetOffices(&_Expertise.CallOpts)
}

// GetScore is a free data retrieval call binding the contract method 0x6339e2e6.
//
// Solidity: function getScore(address user, string category) constant returns(uint256 score)
func (_Expertise *ExpertiseCaller) GetScore(opts *bind.CallOpts, user common.Address, category string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Expertise.contract.Call(opts, out, "getScore", user, category)
	return *ret0, err
}

// GetScore is a free data retrieval call binding the contract method 0x6339e2e6.
//
// Solidity: function getScore(address user, string category) constant returns(uint256 score)
func (_Expertise *ExpertiseSession) GetScore(user common.Address, category string) (*big.Int, error) {
	return _Expertise.Contract.GetScore(&_Expertise.CallOpts, user, category)
}

// GetScore is a free data retrieval call binding the contract method 0x6339e2e6.
//
// Solidity: function getScore(address user, string category) constant returns(uint256 score)
func (_Expertise *ExpertiseCallerSession) GetScore(user common.Address, category string) (*big.Int, error) {
	return _Expertise.Contract.GetScore(&_Expertise.CallOpts, user, category)
}

// GetSubject is a free data retrieval call binding the contract method 0x40f596a8.
//
// Solidity: function getSubject() constant returns(bytes32[])
func (_Expertise *ExpertiseCaller) GetSubject(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Expertise.contract.Call(opts, out, "getSubject")
	return *ret0, err
}

// GetSubject is a free data retrieval call binding the contract method 0x40f596a8.
//
// Solidity: function getSubject() constant returns(bytes32[])
func (_Expertise *ExpertiseSession) GetSubject() ([][32]byte, error) {
	return _Expertise.Contract.GetSubject(&_Expertise.CallOpts)
}

// GetSubject is a free data retrieval call binding the contract method 0x40f596a8.
//
// Solidity: function getSubject() constant returns(bytes32[])
func (_Expertise *ExpertiseCallerSession) GetSubject() ([][32]byte, error) {
	return _Expertise.Contract.GetSubject(&_Expertise.CallOpts)
}

// GetUserInfo is a free data retrieval call binding the contract method 0xa0431815.
//
// Solidity: function getUserInfo(address _user, string _key) constant returns(bytes32[])
func (_Expertise *ExpertiseCaller) GetUserInfo(opts *bind.CallOpts, _user common.Address, _key string) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Expertise.contract.Call(opts, out, "getUserInfo", _user, _key)
	return *ret0, err
}

// GetUserInfo is a free data retrieval call binding the contract method 0xa0431815.
//
// Solidity: function getUserInfo(address _user, string _key) constant returns(bytes32[])
func (_Expertise *ExpertiseSession) GetUserInfo(_user common.Address, _key string) ([][32]byte, error) {
	return _Expertise.Contract.GetUserInfo(&_Expertise.CallOpts, _user, _key)
}

// GetUserInfo is a free data retrieval call binding the contract method 0xa0431815.
//
// Solidity: function getUserInfo(address _user, string _key) constant returns(bytes32[])
func (_Expertise *ExpertiseCallerSession) GetUserInfo(_user common.Address, _key string) ([][32]byte, error) {
	return _Expertise.Contract.GetUserInfo(&_Expertise.CallOpts, _user, _key)
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() constant returns(address[])
func (_Expertise *ExpertiseCaller) GetUsers(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Expertise.contract.Call(opts, out, "getUsers")
	return *ret0, err
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() constant returns(address[])
func (_Expertise *ExpertiseSession) GetUsers() ([]common.Address, error) {
	return _Expertise.Contract.GetUsers(&_Expertise.CallOpts)
}

// GetUsers is a free data retrieval call binding the contract method 0x00ce8e3e.
//
// Solidity: function getUsers() constant returns(address[])
func (_Expertise *ExpertiseCallerSession) GetUsers() ([]common.Address, error) {
	return _Expertise.Contract.GetUsers(&_Expertise.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address addr) constant returns(bool)
func (_Expertise *ExpertiseCaller) IsOwner(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Expertise.contract.Call(opts, out, "isOwner", addr)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address addr) constant returns(bool)
func (_Expertise *ExpertiseSession) IsOwner(addr common.Address) (bool, error) {
	return _Expertise.Contract.IsOwner(&_Expertise.CallOpts, addr)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address addr) constant returns(bool)
func (_Expertise *ExpertiseCallerSession) IsOwner(addr common.Address) (bool, error) {
	return _Expertise.Contract.IsOwner(&_Expertise.CallOpts, addr)
}

// LevelManager is a free data retrieval call binding the contract method 0x2a8d950c.
//
// Solidity: function levelManager() constant returns(address)
func (_Expertise *ExpertiseCaller) LevelManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Expertise.contract.Call(opts, out, "levelManager")
	return *ret0, err
}

// LevelManager is a free data retrieval call binding the contract method 0x2a8d950c.
//
// Solidity: function levelManager() constant returns(address)
func (_Expertise *ExpertiseSession) LevelManager() (common.Address, error) {
	return _Expertise.Contract.LevelManager(&_Expertise.CallOpts)
}

// LevelManager is a free data retrieval call binding the contract method 0x2a8d950c.
//
// Solidity: function levelManager() constant returns(address)
func (_Expertise *ExpertiseCallerSession) LevelManager() (common.Address, error) {
	return _Expertise.Contract.LevelManager(&_Expertise.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) constant returns(bool)
func (_Expertise *ExpertiseCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Expertise.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) constant returns(bool)
func (_Expertise *ExpertiseSession) Owners(arg0 common.Address) (bool, error) {
	return _Expertise.Contract.Owners(&_Expertise.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners(address ) constant returns(bool)
func (_Expertise *ExpertiseCallerSession) Owners(arg0 common.Address) (bool, error) {
	return _Expertise.Contract.Owners(&_Expertise.CallOpts, arg0)
}

// StringToBytes32 is a free data retrieval call binding the contract method 0xcfb51928.
//
// Solidity: function stringToBytes32(string source) constant returns(bytes32 result)
func (_Expertise *ExpertiseCaller) StringToBytes32(opts *bind.CallOpts, source string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Expertise.contract.Call(opts, out, "stringToBytes32", source)
	return *ret0, err
}

// StringToBytes32 is a free data retrieval call binding the contract method 0xcfb51928.
//
// Solidity: function stringToBytes32(string source) constant returns(bytes32 result)
func (_Expertise *ExpertiseSession) StringToBytes32(source string) ([32]byte, error) {
	return _Expertise.Contract.StringToBytes32(&_Expertise.CallOpts, source)
}

// StringToBytes32 is a free data retrieval call binding the contract method 0xcfb51928.
//
// Solidity: function stringToBytes32(string source) constant returns(bytes32 result)
func (_Expertise *ExpertiseCallerSession) StringToBytes32(source string) ([32]byte, error) {
	return _Expertise.Contract.StringToBytes32(&_Expertise.CallOpts, source)
}

// AddAchievement is a paid mutator transaction binding the contract method 0xb21e2607.
//
// Solidity: function addAchievement(string id, string category) returns()
func (_Expertise *ExpertiseTransactor) AddAchievement(opts *bind.TransactOpts, id string, category string) (*types.Transaction, error) {
	return _Expertise.contract.Transact(opts, "addAchievement", id, category)
}

// AddAchievement is a paid mutator transaction binding the contract method 0xb21e2607.
//
// Solidity: function addAchievement(string id, string category) returns()
func (_Expertise *ExpertiseSession) AddAchievement(id string, category string) (*types.Transaction, error) {
	return _Expertise.Contract.AddAchievement(&_Expertise.TransactOpts, id, category)
}

// AddAchievement is a paid mutator transaction binding the contract method 0xb21e2607.
//
// Solidity: function addAchievement(string id, string category) returns()
func (_Expertise *ExpertiseTransactorSession) AddAchievement(id string, category string) (*types.Transaction, error) {
	return _Expertise.Contract.AddAchievement(&_Expertise.TransactOpts, id, category)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address newOwner) returns()
func (_Expertise *ExpertiseTransactor) AddOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Expertise.contract.Transact(opts, "addOwner", newOwner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address newOwner) returns()
func (_Expertise *ExpertiseSession) AddOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Expertise.Contract.AddOwner(&_Expertise.TransactOpts, newOwner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address newOwner) returns()
func (_Expertise *ExpertiseTransactorSession) AddOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Expertise.Contract.AddOwner(&_Expertise.TransactOpts, newOwner)
}

// RemoveAchievement is a paid mutator transaction binding the contract method 0xdf0b2b1e.
//
// Solidity: function removeAchievement(string id, string category) returns()
func (_Expertise *ExpertiseTransactor) RemoveAchievement(opts *bind.TransactOpts, id string, category string) (*types.Transaction, error) {
	return _Expertise.contract.Transact(opts, "removeAchievement", id, category)
}

// RemoveAchievement is a paid mutator transaction binding the contract method 0xdf0b2b1e.
//
// Solidity: function removeAchievement(string id, string category) returns()
func (_Expertise *ExpertiseSession) RemoveAchievement(id string, category string) (*types.Transaction, error) {
	return _Expertise.Contract.RemoveAchievement(&_Expertise.TransactOpts, id, category)
}

// RemoveAchievement is a paid mutator transaction binding the contract method 0xdf0b2b1e.
//
// Solidity: function removeAchievement(string id, string category) returns()
func (_Expertise *ExpertiseTransactorSession) RemoveAchievement(id string, category string) (*types.Transaction, error) {
	return _Expertise.Contract.RemoveAchievement(&_Expertise.TransactOpts, id, category)
}

// Renounce is a paid mutator transaction binding the contract method 0xb15be2f5.
//
// Solidity: function renounce() returns()
func (_Expertise *ExpertiseTransactor) Renounce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Expertise.contract.Transact(opts, "renounce")
}

// Renounce is a paid mutator transaction binding the contract method 0xb15be2f5.
//
// Solidity: function renounce() returns()
func (_Expertise *ExpertiseSession) Renounce() (*types.Transaction, error) {
	return _Expertise.Contract.Renounce(&_Expertise.TransactOpts)
}

// Renounce is a paid mutator transaction binding the contract method 0xb15be2f5.
//
// Solidity: function renounce() returns()
func (_Expertise *ExpertiseTransactorSession) Renounce() (*types.Transaction, error) {
	return _Expertise.Contract.Renounce(&_Expertise.TransactOpts)
}

// SetAchievementInfo is a paid mutator transaction binding the contract method 0x2baf63f6.
//
// Solidity: function setAchievementInfo(string _id, string _key, string _value) returns()
func (_Expertise *ExpertiseTransactor) SetAchievementInfo(opts *bind.TransactOpts, _id string, _key string, _value string) (*types.Transaction, error) {
	return _Expertise.contract.Transact(opts, "setAchievementInfo", _id, _key, _value)
}

// SetAchievementInfo is a paid mutator transaction binding the contract method 0x2baf63f6.
//
// Solidity: function setAchievementInfo(string _id, string _key, string _value) returns()
func (_Expertise *ExpertiseSession) SetAchievementInfo(_id string, _key string, _value string) (*types.Transaction, error) {
	return _Expertise.Contract.SetAchievementInfo(&_Expertise.TransactOpts, _id, _key, _value)
}

// SetAchievementInfo is a paid mutator transaction binding the contract method 0x2baf63f6.
//
// Solidity: function setAchievementInfo(string _id, string _key, string _value) returns()
func (_Expertise *ExpertiseTransactorSession) SetAchievementInfo(_id string, _key string, _value string) (*types.Transaction, error) {
	return _Expertise.Contract.SetAchievementInfo(&_Expertise.TransactOpts, _id, _key, _value)
}

// SetAchievementInfoBytes32 is a paid mutator transaction binding the contract method 0xf6ee3f0e.
//
// Solidity: function setAchievementInfoBytes32(string _id, string _key, bytes32 _value) returns()
func (_Expertise *ExpertiseTransactor) SetAchievementInfoBytes32(opts *bind.TransactOpts, _id string, _key string, _value [32]byte) (*types.Transaction, error) {
	return _Expertise.contract.Transact(opts, "setAchievementInfoBytes32", _id, _key, _value)
}

// SetAchievementInfoBytes32 is a paid mutator transaction binding the contract method 0xf6ee3f0e.
//
// Solidity: function setAchievementInfoBytes32(string _id, string _key, bytes32 _value) returns()
func (_Expertise *ExpertiseSession) SetAchievementInfoBytes32(_id string, _key string, _value [32]byte) (*types.Transaction, error) {
	return _Expertise.Contract.SetAchievementInfoBytes32(&_Expertise.TransactOpts, _id, _key, _value)
}

// SetAchievementInfoBytes32 is a paid mutator transaction binding the contract method 0xf6ee3f0e.
//
// Solidity: function setAchievementInfoBytes32(string _id, string _key, bytes32 _value) returns()
func (_Expertise *ExpertiseTransactorSession) SetAchievementInfoBytes32(_id string, _key string, _value [32]byte) (*types.Transaction, error) {
	return _Expertise.Contract.SetAchievementInfoBytes32(&_Expertise.TransactOpts, _id, _key, _value)
}

// SetLevelManager is a paid mutator transaction binding the contract method 0xbd152bc7.
//
// Solidity: function setLevelManager(address newLevelManager) returns()
func (_Expertise *ExpertiseTransactor) SetLevelManager(opts *bind.TransactOpts, newLevelManager common.Address) (*types.Transaction, error) {
	return _Expertise.contract.Transact(opts, "setLevelManager", newLevelManager)
}

// SetLevelManager is a paid mutator transaction binding the contract method 0xbd152bc7.
//
// Solidity: function setLevelManager(address newLevelManager) returns()
func (_Expertise *ExpertiseSession) SetLevelManager(newLevelManager common.Address) (*types.Transaction, error) {
	return _Expertise.Contract.SetLevelManager(&_Expertise.TransactOpts, newLevelManager)
}

// SetLevelManager is a paid mutator transaction binding the contract method 0xbd152bc7.
//
// Solidity: function setLevelManager(address newLevelManager) returns()
func (_Expertise *ExpertiseTransactorSession) SetLevelManager(newLevelManager common.Address) (*types.Transaction, error) {
	return _Expertise.Contract.SetLevelManager(&_Expertise.TransactOpts, newLevelManager)
}

// SetUserInfo is a paid mutator transaction binding the contract method 0xe21eabed.
//
// Solidity: function setUserInfo(string _key, string _value) returns()
func (_Expertise *ExpertiseTransactor) SetUserInfo(opts *bind.TransactOpts, _key string, _value string) (*types.Transaction, error) {
	return _Expertise.contract.Transact(opts, "setUserInfo", _key, _value)
}

// SetUserInfo is a paid mutator transaction binding the contract method 0xe21eabed.
//
// Solidity: function setUserInfo(string _key, string _value) returns()
func (_Expertise *ExpertiseSession) SetUserInfo(_key string, _value string) (*types.Transaction, error) {
	return _Expertise.Contract.SetUserInfo(&_Expertise.TransactOpts, _key, _value)
}

// SetUserInfo is a paid mutator transaction binding the contract method 0xe21eabed.
//
// Solidity: function setUserInfo(string _key, string _value) returns()
func (_Expertise *ExpertiseTransactorSession) SetUserInfo(_key string, _value string) (*types.Transaction, error) {
	return _Expertise.Contract.SetUserInfo(&_Expertise.TransactOpts, _key, _value)
}

// ExpertiseAchievementAddedIterator is returned from FilterAchievementAdded and is used to iterate over the raw logs and unpacked data for AchievementAdded events raised by the Expertise contract.
type ExpertiseAchievementAddedIterator struct {
	Event *ExpertiseAchievementAdded // Event containing the contract specifics and raw log

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
func (it *ExpertiseAchievementAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpertiseAchievementAdded)
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
		it.Event = new(ExpertiseAchievementAdded)
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
func (it *ExpertiseAchievementAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpertiseAchievementAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpertiseAchievementAdded represents a AchievementAdded event raised by the Expertise contract.
type ExpertiseAchievementAdded struct {
	Id       string
	Category string
	User     common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAchievementAdded is a free log retrieval operation binding the contract event 0xd45ce2af87df3031529ddcc3380987c2cf5c289941d6bd5d3c593225982d0a6a.
//
// Solidity: event AchievementAdded(string id, string category, address user)
func (_Expertise *ExpertiseFilterer) FilterAchievementAdded(opts *bind.FilterOpts) (*ExpertiseAchievementAddedIterator, error) {

	logs, sub, err := _Expertise.contract.FilterLogs(opts, "AchievementAdded")
	if err != nil {
		return nil, err
	}
	return &ExpertiseAchievementAddedIterator{contract: _Expertise.contract, event: "AchievementAdded", logs: logs, sub: sub}, nil
}

// WatchAchievementAdded is a free log subscription operation binding the contract event 0xd45ce2af87df3031529ddcc3380987c2cf5c289941d6bd5d3c593225982d0a6a.
//
// Solidity: event AchievementAdded(string id, string category, address user)
func (_Expertise *ExpertiseFilterer) WatchAchievementAdded(opts *bind.WatchOpts, sink chan<- *ExpertiseAchievementAdded) (event.Subscription, error) {

	logs, sub, err := _Expertise.contract.WatchLogs(opts, "AchievementAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpertiseAchievementAdded)
				if err := _Expertise.contract.UnpackLog(event, "AchievementAdded", log); err != nil {
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

// ExpertiseAchievementInfoAddedIterator is returned from FilterAchievementInfoAdded and is used to iterate over the raw logs and unpacked data for AchievementInfoAdded events raised by the Expertise contract.
type ExpertiseAchievementInfoAddedIterator struct {
	Event *ExpertiseAchievementInfoAdded // Event containing the contract specifics and raw log

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
func (it *ExpertiseAchievementInfoAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpertiseAchievementInfoAdded)
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
		it.Event = new(ExpertiseAchievementInfoAdded)
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
func (it *ExpertiseAchievementInfoAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpertiseAchievementInfoAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpertiseAchievementInfoAdded represents a AchievementInfoAdded event raised by the Expertise contract.
type ExpertiseAchievementInfoAdded struct {
	Addr          common.Address
	IdAchievement string
	Key           string
	Value         [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAchievementInfoAdded is a free log retrieval operation binding the contract event 0x732ff0898ffa64f3eb6f8d1ac992c5f08f714b4fc47a956aac1af1559751c872.
//
// Solidity: event AchievementInfoAdded(address _addr, string _idAchievement, string _key, bytes32 _value)
func (_Expertise *ExpertiseFilterer) FilterAchievementInfoAdded(opts *bind.FilterOpts) (*ExpertiseAchievementInfoAddedIterator, error) {

	logs, sub, err := _Expertise.contract.FilterLogs(opts, "AchievementInfoAdded")
	if err != nil {
		return nil, err
	}
	return &ExpertiseAchievementInfoAddedIterator{contract: _Expertise.contract, event: "AchievementInfoAdded", logs: logs, sub: sub}, nil
}

// WatchAchievementInfoAdded is a free log subscription operation binding the contract event 0x732ff0898ffa64f3eb6f8d1ac992c5f08f714b4fc47a956aac1af1559751c872.
//
// Solidity: event AchievementInfoAdded(address _addr, string _idAchievement, string _key, bytes32 _value)
func (_Expertise *ExpertiseFilterer) WatchAchievementInfoAdded(opts *bind.WatchOpts, sink chan<- *ExpertiseAchievementInfoAdded) (event.Subscription, error) {

	logs, sub, err := _Expertise.contract.WatchLogs(opts, "AchievementInfoAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpertiseAchievementInfoAdded)
				if err := _Expertise.contract.UnpackLog(event, "AchievementInfoAdded", log); err != nil {
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

// ExpertiseAchievementRemovedIterator is returned from FilterAchievementRemoved and is used to iterate over the raw logs and unpacked data for AchievementRemoved events raised by the Expertise contract.
type ExpertiseAchievementRemovedIterator struct {
	Event *ExpertiseAchievementRemoved // Event containing the contract specifics and raw log

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
func (it *ExpertiseAchievementRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpertiseAchievementRemoved)
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
		it.Event = new(ExpertiseAchievementRemoved)
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
func (it *ExpertiseAchievementRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpertiseAchievementRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpertiseAchievementRemoved represents a AchievementRemoved event raised by the Expertise contract.
type ExpertiseAchievementRemoved struct {
	Id       string
	Category string
	User     common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAchievementRemoved is a free log retrieval operation binding the contract event 0x9650be95ae6acabefaa6193f2a4032c2daa4dbc79d54d46d8041678dffe72ccf.
//
// Solidity: event AchievementRemoved(string id, string category, address user)
func (_Expertise *ExpertiseFilterer) FilterAchievementRemoved(opts *bind.FilterOpts) (*ExpertiseAchievementRemovedIterator, error) {

	logs, sub, err := _Expertise.contract.FilterLogs(opts, "AchievementRemoved")
	if err != nil {
		return nil, err
	}
	return &ExpertiseAchievementRemovedIterator{contract: _Expertise.contract, event: "AchievementRemoved", logs: logs, sub: sub}, nil
}

// WatchAchievementRemoved is a free log subscription operation binding the contract event 0x9650be95ae6acabefaa6193f2a4032c2daa4dbc79d54d46d8041678dffe72ccf.
//
// Solidity: event AchievementRemoved(string id, string category, address user)
func (_Expertise *ExpertiseFilterer) WatchAchievementRemoved(opts *bind.WatchOpts, sink chan<- *ExpertiseAchievementRemoved) (event.Subscription, error) {

	logs, sub, err := _Expertise.contract.WatchLogs(opts, "AchievementRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpertiseAchievementRemoved)
				if err := _Expertise.contract.UnpackLog(event, "AchievementRemoved", log); err != nil {
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

// ExpertiseInfoAddedIterator is returned from FilterInfoAdded and is used to iterate over the raw logs and unpacked data for InfoAdded events raised by the Expertise contract.
type ExpertiseInfoAddedIterator struct {
	Event *ExpertiseInfoAdded // Event containing the contract specifics and raw log

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
func (it *ExpertiseInfoAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpertiseInfoAdded)
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
		it.Event = new(ExpertiseInfoAdded)
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
func (it *ExpertiseInfoAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpertiseInfoAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpertiseInfoAdded represents a InfoAdded event raised by the Expertise contract.
type ExpertiseInfoAdded struct {
	Addr  common.Address
	Key   string
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInfoAdded is a free log retrieval operation binding the contract event 0x2b1a13e681efaa0215cb836389223a4f9b0c9a224afecdda91af971c885a66da.
//
// Solidity: event InfoAdded(address _addr, string _key, string _value)
func (_Expertise *ExpertiseFilterer) FilterInfoAdded(opts *bind.FilterOpts) (*ExpertiseInfoAddedIterator, error) {

	logs, sub, err := _Expertise.contract.FilterLogs(opts, "InfoAdded")
	if err != nil {
		return nil, err
	}
	return &ExpertiseInfoAddedIterator{contract: _Expertise.contract, event: "InfoAdded", logs: logs, sub: sub}, nil
}

// WatchInfoAdded is a free log subscription operation binding the contract event 0x2b1a13e681efaa0215cb836389223a4f9b0c9a224afecdda91af971c885a66da.
//
// Solidity: event InfoAdded(address _addr, string _key, string _value)
func (_Expertise *ExpertiseFilterer) WatchInfoAdded(opts *bind.WatchOpts, sink chan<- *ExpertiseInfoAdded) (event.Subscription, error) {

	logs, sub, err := _Expertise.contract.WatchLogs(opts, "InfoAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpertiseInfoAdded)
				if err := _Expertise.contract.UnpackLog(event, "InfoAdded", log); err != nil {
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

// ExpertiseLevelManagerSetIterator is returned from FilterLevelManagerSet and is used to iterate over the raw logs and unpacked data for LevelManagerSet events raised by the Expertise contract.
type ExpertiseLevelManagerSetIterator struct {
	Event *ExpertiseLevelManagerSet // Event containing the contract specifics and raw log

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
func (it *ExpertiseLevelManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpertiseLevelManagerSet)
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
		it.Event = new(ExpertiseLevelManagerSet)
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
func (it *ExpertiseLevelManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpertiseLevelManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpertiseLevelManagerSet represents a LevelManagerSet event raised by the Expertise contract.
type ExpertiseLevelManagerSet struct {
	OldLevelManager common.Address
	NewLevelManager common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLevelManagerSet is a free log retrieval operation binding the contract event 0x04c16e365474ee980a578ba9c054db3805ac81c56b3eae405e7f89f96e52dcb8.
//
// Solidity: event LevelManagerSet(address oldLevelManager, address newLevelManager)
func (_Expertise *ExpertiseFilterer) FilterLevelManagerSet(opts *bind.FilterOpts) (*ExpertiseLevelManagerSetIterator, error) {

	logs, sub, err := _Expertise.contract.FilterLogs(opts, "LevelManagerSet")
	if err != nil {
		return nil, err
	}
	return &ExpertiseLevelManagerSetIterator{contract: _Expertise.contract, event: "LevelManagerSet", logs: logs, sub: sub}, nil
}

// WatchLevelManagerSet is a free log subscription operation binding the contract event 0x04c16e365474ee980a578ba9c054db3805ac81c56b3eae405e7f89f96e52dcb8.
//
// Solidity: event LevelManagerSet(address oldLevelManager, address newLevelManager)
func (_Expertise *ExpertiseFilterer) WatchLevelManagerSet(opts *bind.WatchOpts, sink chan<- *ExpertiseLevelManagerSet) (event.Subscription, error) {

	logs, sub, err := _Expertise.contract.WatchLogs(opts, "LevelManagerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpertiseLevelManagerSet)
				if err := _Expertise.contract.UnpackLog(event, "LevelManagerSet", log); err != nil {
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

// ExpertiseOwnerAddedIterator is returned from FilterOwnerAdded and is used to iterate over the raw logs and unpacked data for OwnerAdded events raised by the Expertise contract.
type ExpertiseOwnerAddedIterator struct {
	Event *ExpertiseOwnerAdded // Event containing the contract specifics and raw log

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
func (it *ExpertiseOwnerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpertiseOwnerAdded)
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
		it.Event = new(ExpertiseOwnerAdded)
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
func (it *ExpertiseOwnerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpertiseOwnerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpertiseOwnerAdded represents a OwnerAdded event raised by the Expertise contract.
type ExpertiseOwnerAdded struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerAdded is a free log retrieval operation binding the contract event 0x994a936646fe87ffe4f1e469d3d6aa417d6b855598397f323de5b449f765f0c3.
//
// Solidity: event OwnerAdded(address indexed newOwner)
func (_Expertise *ExpertiseFilterer) FilterOwnerAdded(opts *bind.FilterOpts, newOwner []common.Address) (*ExpertiseOwnerAddedIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Expertise.contract.FilterLogs(opts, "OwnerAdded", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExpertiseOwnerAddedIterator{contract: _Expertise.contract, event: "OwnerAdded", logs: logs, sub: sub}, nil
}

// WatchOwnerAdded is a free log subscription operation binding the contract event 0x994a936646fe87ffe4f1e469d3d6aa417d6b855598397f323de5b449f765f0c3.
//
// Solidity: event OwnerAdded(address indexed newOwner)
func (_Expertise *ExpertiseFilterer) WatchOwnerAdded(opts *bind.WatchOpts, sink chan<- *ExpertiseOwnerAdded, newOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Expertise.contract.WatchLogs(opts, "OwnerAdded", newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpertiseOwnerAdded)
				if err := _Expertise.contract.UnpackLog(event, "OwnerAdded", log); err != nil {
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

// ExpertiseOwnerRemovedIterator is returned from FilterOwnerRemoved and is used to iterate over the raw logs and unpacked data for OwnerRemoved events raised by the Expertise contract.
type ExpertiseOwnerRemovedIterator struct {
	Event *ExpertiseOwnerRemoved // Event containing the contract specifics and raw log

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
func (it *ExpertiseOwnerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpertiseOwnerRemoved)
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
		it.Event = new(ExpertiseOwnerRemoved)
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
func (it *ExpertiseOwnerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpertiseOwnerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpertiseOwnerRemoved represents a OwnerRemoved event raised by the Expertise contract.
type ExpertiseOwnerRemoved struct {
	FormerOwner common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOwnerRemoved is a free log retrieval operation binding the contract event 0x58619076adf5bb0943d100ef88d52d7c3fd691b19d3a9071b555b651fbf418da.
//
// Solidity: event OwnerRemoved(address indexed formerOwner)
func (_Expertise *ExpertiseFilterer) FilterOwnerRemoved(opts *bind.FilterOpts, formerOwner []common.Address) (*ExpertiseOwnerRemovedIterator, error) {

	var formerOwnerRule []interface{}
	for _, formerOwnerItem := range formerOwner {
		formerOwnerRule = append(formerOwnerRule, formerOwnerItem)
	}

	logs, sub, err := _Expertise.contract.FilterLogs(opts, "OwnerRemoved", formerOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExpertiseOwnerRemovedIterator{contract: _Expertise.contract, event: "OwnerRemoved", logs: logs, sub: sub}, nil
}

// WatchOwnerRemoved is a free log subscription operation binding the contract event 0x58619076adf5bb0943d100ef88d52d7c3fd691b19d3a9071b555b651fbf418da.
//
// Solidity: event OwnerRemoved(address indexed formerOwner)
func (_Expertise *ExpertiseFilterer) WatchOwnerRemoved(opts *bind.WatchOpts, sink chan<- *ExpertiseOwnerRemoved, formerOwner []common.Address) (event.Subscription, error) {

	var formerOwnerRule []interface{}
	for _, formerOwnerItem := range formerOwner {
		formerOwnerRule = append(formerOwnerRule, formerOwnerItem)
	}

	logs, sub, err := _Expertise.contract.WatchLogs(opts, "OwnerRemoved", formerOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpertiseOwnerRemoved)
				if err := _Expertise.contract.UnpackLog(event, "OwnerRemoved", log); err != nil {
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

// ExpertiseUserAddedIterator is returned from FilterUserAdded and is used to iterate over the raw logs and unpacked data for UserAdded events raised by the Expertise contract.
type ExpertiseUserAddedIterator struct {
	Event *ExpertiseUserAdded // Event containing the contract specifics and raw log

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
func (it *ExpertiseUserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExpertiseUserAdded)
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
		it.Event = new(ExpertiseUserAdded)
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
func (it *ExpertiseUserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExpertiseUserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExpertiseUserAdded represents a UserAdded event raised by the Expertise contract.
type ExpertiseUserAdded struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUserAdded is a free log retrieval operation binding the contract event 0x19ef9a4877199f89440a26acb26895ec02ed86f2df1aeaa90dc18041b892f71f.
//
// Solidity: event UserAdded(address user)
func (_Expertise *ExpertiseFilterer) FilterUserAdded(opts *bind.FilterOpts) (*ExpertiseUserAddedIterator, error) {

	logs, sub, err := _Expertise.contract.FilterLogs(opts, "UserAdded")
	if err != nil {
		return nil, err
	}
	return &ExpertiseUserAddedIterator{contract: _Expertise.contract, event: "UserAdded", logs: logs, sub: sub}, nil
}

// WatchUserAdded is a free log subscription operation binding the contract event 0x19ef9a4877199f89440a26acb26895ec02ed86f2df1aeaa90dc18041b892f71f.
//
// Solidity: event UserAdded(address user)
func (_Expertise *ExpertiseFilterer) WatchUserAdded(opts *bind.WatchOpts, sink chan<- *ExpertiseUserAdded) (event.Subscription, error) {

	logs, sub, err := _Expertise.contract.WatchLogs(opts, "UserAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExpertiseUserAdded)
				if err := _Expertise.contract.UnpackLog(event, "UserAdded", log); err != nil {
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
