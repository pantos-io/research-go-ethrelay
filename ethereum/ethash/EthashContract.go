// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethash

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EthashABI is the input ABI used to generate the binding from.
const EthashABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"SHA3_address\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"SetEpochData\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochIndex\",\"type\":\"uint256\"}],\"name\":\"isEpochDataSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fullSizeIn128Resultion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"branchDepth\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"merkleNodes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numElems\",\"type\":\"uint256\"}],\"name\":\"setEpochData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"rlpHeaderHashWithoutNonce\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"dataSetLookup\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"witnessForLookup\",\"type\":\"uint256[]\"}],\"name\":\"verifyPoW\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EthashBin is the compiled bytecode used for deploying new contracts.
var EthashBin = "0x60806040523480156200001157600080fd5b506040516200260538038062002605833981810160405281019062000037919062000095565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200010f565b6000815190506200008f81620000f5565b92915050565b600060208284031215620000a857600080fd5b6000620000b8848285016200007e565b91505092915050565b6000620000ce82620000d5565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200010081620000c1565b81146200010c57600080fd5b50565b6124e6806200011f6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806329e265df14610046578063c7b81f4f14610077578063c891a29d146100a7575b600080fd5b610060600480360381019061005b9190611dc8565b6100c3565b60405161006e92919061213d565b60405180910390f35b610091600480360381019061008c9190611d9f565b610203565b60405161009e91906120d0565b60405180910390f35b6100c160048036038101906100bc9190611e8b565b610227565b005b60008060006175308b6100d69190612260565b905060006101698b8b8a8a80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050898980806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505086610465565b9050887fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6101979190612260565b8111156101ec576000807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe8314156101d557600191508390506101dd565b600291508290505b818195509550505050506101f6565b6000809350935050505b9850989650505050505050565b60008060016000848152602001908152602001600020610200015414159050919050565b60005b818110156103d2576000600160008981526020019081526020016000206000018285610256919061220a565b610200811061028e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b0154111561031f573373ffffffffffffffffffffffffffffffffffffffff167f5cd723400be8430351b9cbaa5ea421b3fb2528c6a7650c493f895e7d97750da1600183867001000000000000000000000000000000008c6102ef9190612291565b6102f9919061220a565b610303919061220a565b604051610311929190612114565b60405180910390a25061045d565b838181518110610358577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151600160008981526020019081526020016000206000018285610382919061220a565b61020081106103ba577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b018190555080806103ca90612360565b91505061022a565b5084600160008881526020019081526020016000206102000181905550836001600088815260200190815260200160002061020101819055503373ffffffffffffffffffffffffffffffffffffffff167f5cd723400be8430351b9cbaa5ea421b3fb2528c6a7650c493f895e7d97750da16000806040516104549291906120eb565b60405180910390a25b505050505050565b600061046f611b40565b610477611b63565b61047f611b86565b60006040518060400160405280600160008981526020019081526020016000206102010154815260200160016000898152602001908152602001600020610200015481525090506000806104d288610203565b610504577ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe9650505050505050610a3b565b600083600160028110610540577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201511415610579577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9650505050505050610a3b565b6105868c60001c8c610a44565b9550600091505b60108210156105c2576020820280870151818701528087015181610200018701525081806105ba90612360565b92505061058d565b600091505b604082101561087e5760008360016002811061060c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201516106a48860006010811061064e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002015185188860208761066391906123b3565b6020811061069a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020020151610cd3565b6106ae91906123b3565b90506106ba8982610cf5565b61070582858e8e896000600281106106fb577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020020151610dd0565b14610739577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff975050505050505050610a3b565b600091505b600882101561086a5760208b01836080020163ffffffff815116836020028801805163ffffffff8363010001938302181680835260208501945063ffffffff8551169350610100830192508251915063ffffffff84630100019384021816905080835260208501945063ffffffff8551169350610100830192508251915063ffffffff84630100019384021816905080835260208501945063ffffffff8551169350610100830192508251915063ffffffff84630100019384021816905080835250505050506080830260208c01016401000000008151048082526020820191506401000000008251049050808252602082019150640100000000825104905080825260208201915064010000000082510490508082525050818061086290612360565b92505061073e565b50818061087690612360565b9250506105c7565b600091505b6020821015610a22576109c06109716109228785602081106108ce577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020020151886001876108e1919061220a565b60208110610918577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020020151610cd3565b87600286610930919061220a565b60208110610967577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020020151610cd3565b8660038561097f919061220a565b602081106109b6577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020020151610cd3565b846004846109ce9190612260565b60088110610a05577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002018181525050600482610a1b919061220a565b9150610883565b6000610a2e878661101c565b9050809750505050505050505b95945050505050565b610a4c611b40565b610a54611ba9565b610a5d846118f0565b935067ffffffffffffffff841681600060098110610aa4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020020181815250506801000000000000000084610ac29190612260565b935067ffffffffffffffff841681600160098110610b09577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020020181815250506801000000000000000084610b279190612260565b935067ffffffffffffffff841681600260098110610b6e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020020181815250506801000000000000000084610b8c9190612260565b935067ffffffffffffffff841681600360098110610bd3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020020181815250508281600460098110610c17577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201818152505060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b2df8456826040518263ffffffff1660e01b8152600401610c7991906120b4565b6102006040518083038186803b158015610c9257600080fd5b505afa158015610ca6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cca9190611d75565b91505092915050565b600063ffffffff82630100019385610ceb9190612291565b1816905092915050565b60008060016000858152602001908152602001600020610201015483901c9050600060016000868152602001908152602001600020600001600283610d3a9190612260565b6102008110610d72577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b015490506000600283610d8591906123b3565b1415610da5576fffffffffffffffffffffffffffffffff81169050610dc5565b70010000000000000000000000000000000081610dc29190612260565b90505b809250505092915050565b6000806fffffffffffffffffffffffffffffffff610dee8688611951565b1690506000806000806000600288610e0691906123b3565b1190506002870496506000878b610e1d9190612291565b90508115610e34578a81610e31919061220a565b90505b60005b88811015610f725760208282010260208b0101519350600060018e161415610e76578695506fffffffffffffffffffffffffffffffff84169450610e8f565b6fffffffffffffffffffffffffffffffff841695508694505b6fffffffffffffffffffffffffffffffff8686604051602001610eb3929190611ffd565b6040516020818303038152906040528051906020012060001c16965060028d049c50600060018e161415610eff5786955070010000000000000000000000000000000084049450610f19565b700100000000000000000000000000000000840495508694505b6fffffffffffffffffffffffffffffffff8686604051602001610f3d929190611ffd565b6040516020818303038152906040528051906020012060001c16965060028d049c508080610f6a90612360565b915050610e37565b82156110095760208282010260208b0101519350600060018e161415610faf578695506fffffffffffffffffffffffffffffffff84169450610fc8565b6fffffffffffffffffffffffffffffffff841695508694505b6fffffffffffffffffffffffffffffffff8686604051602001610fec929190611ffd565b6040516020818303038152906040528051906020012060001c1696505b8697505050505050505095945050505050565b6000807001000000000000000000000000000000006c0100000000000000000000000085600760108110611079577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201516110889190612291565b68010000000000000000866006601081106110cc577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201516110db9190612291565b6401000000008760056010811061111b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002015161112a9190612291565b87600460108110611164577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020020151611173919061220a565b61117d919061220a565b611187919061220a565b6111919190612291565b6c01000000000000000000000000856003601081106111d9577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201516111e89190612291565b680100000000000000008660026010811061122c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002015161123b9190612291565b6401000000008760016010811061127b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002015161128a9190612291565b876000601081106112c4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201516112d3919061220a565b6112dd919061220a565b6112e7919061220a565b6112f1919061220a565b905060007001000000000000000000000000000000006c0100000000000000000000000086600f6010811061134f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002015161135e9190612291565b6801000000000000000087600e601081106113a2577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201516113b19190612291565b64010000000088600d601081106113f1577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201516114009190612291565b88600c6010811061143a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020020151611449919061220a565b611453919061220a565b61145d919061220a565b6114679190612291565b6c0100000000000000000000000086600b601081106114af577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201516114be9190612291565b6801000000000000000087600a60108110611502577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201516115119190612291565b64010000000088600960108110611551577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201516115609190612291565b8860086010811061159a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201516115a9919061220a565b6115b3919061220a565b6115bd919061220a565b6115c7919061220a565b905060007001000000000000000000000000000000006c0100000000000000000000000086600760088110611625577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201516116349190612291565b6801000000000000000087600660088110611678577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201516116879190612291565b640100000000886005600881106116c7577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201516116d69190612291565b88600460088110611710577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002015161171f919061220a565b611729919061220a565b611733919061220a565b61173d9190612291565b6c0100000000000000000000000086600360088110611785577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201516117949190612291565b68010000000000000000876002600881106117d8577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201516117e79190612291565b64010000000088600160088110611827577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200201516118369190612291565b88600060088110611870577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002015161187f919061220a565b611889919061220a565b611893919061220a565b61189d919061220a565b90506118a883611adf565b6118b183611adf565b6118ba83611adf565b6040516020016118cc93929190612029565b6040516020818303038152906040528051906020012060001c935050505092915050565b6000806000905060005b602081101561194757610100826119119190612291565b915060ff841682611922919061220a565b9150610100846119329190612260565b9350808061193f90612360565b9150506118fa565b5080915050919050565b6000828260046119619190612291565b81518110611998577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518360018460046119b09190612291565b6119ba919061220a565b815181106119f1577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151846002856004611a099190612291565b611a13919061220a565b81518110611a4a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151856003866004611a629190612291565b611a6c919061220a565b81518110611aa3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151604051602001611abe9493929190612066565b6040516020818303038152906040528051906020012060001c905092915050565b6000806000905060005b6020811015611b365761010082611b009190612291565b915060ff841682611b11919061220a565b915061010084611b219190612260565b93508080611b2e90612360565b915050611ae9565b5080915050919050565b604051806102000160405280601090602082028036833780820191505090505090565b604051806104000160405280602090602082028036833780820191505090505090565b604051806101000160405280600890602082028036833780820191505090505090565b604051806101200160405280600990602082028036833780820191505090505090565b6000611bdf611bda8461218b565b612166565b90508082856020860282011115611bf557600080fd5b60005b85811015611c255781611c0b8882611d60565b845260208401935060208301925050600181019050611bf8565b5050509392505050565b6000611c42611c3d846121b1565b612166565b90508083825260208201905082856020860282011115611c6157600080fd5b60005b85811015611c915781611c778882611d4b565b845260208401935060208301925050600181019050611c64565b5050509392505050565b600082601f830112611cac57600080fd5b6010611cb9848285611bcc565b91505092915050565b60008083601f840112611cd457600080fd5b8235905067ffffffffffffffff811115611ced57600080fd5b602083019150836020820283011115611d0557600080fd5b9250929050565b600082601f830112611d1d57600080fd5b8135611d2d848260208601611c2f565b91505092915050565b600081359050611d4581612482565b92915050565b600081359050611d5a81612499565b92915050565b600081519050611d6f81612499565b92915050565b60006102008284031215611d8857600080fd5b6000611d9684828501611c9b565b91505092915050565b600060208284031215611db157600080fd5b6000611dbf84828501611d4b565b91505092915050565b60008060008060008060008060c0898b031215611de457600080fd5b6000611df28b828c01611d4b565b9850506020611e038b828c01611d36565b9750506040611e148b828c01611d4b565b9650506060611e258b828c01611d4b565b955050608089013567ffffffffffffffff811115611e4257600080fd5b611e4e8b828c01611cc2565b945094505060a089013567ffffffffffffffff811115611e6d57600080fd5b611e798b828c01611cc2565b92509250509295985092959890939650565b60008060008060008060c08789031215611ea457600080fd5b6000611eb289828a01611d4b565b9650506020611ec389828a01611d4b565b9550506040611ed489828a01611d4b565b945050606087013567ffffffffffffffff811115611ef157600080fd5b611efd89828a01611d0c565b9350506080611f0e89828a01611d4b565b92505060a0611f1f89828a01611d4b565b9150509295509295509295565b6000611f388383611fc8565b60208301905092915050565b611f4d816121e7565b611f5781846121ff565b9250611f62826121dd565b8060005b83811015611f93578151611f7a8782611f2c565b9650611f85836121f2565b925050600181019050611f66565b505050505050565b611fa4816122eb565b82525050565b611fb38161230b565b82525050565b611fc28161231d565b82525050565b611fd181612301565b82525050565b611fe081612301565b82525050565b611ff7611ff282612301565b6123a9565b82525050565b60006120098285611fe6565b6020820191506120198284611fe6565b6020820191508190509392505050565b60006120358286611fe6565b6020820191506120458285611fe6565b6020820191506120558284611fe6565b602082019150819050949350505050565b60006120728287611fe6565b6020820191506120828286611fe6565b6020820191506120928285611fe6565b6020820191506120a28284611fe6565b60208201915081905095945050505050565b6000610120820190506120ca6000830184611f44565b92915050565b60006020820190506120e56000830184611f9b565b92915050565b60006040820190506121006000830185611faa565b61210d6020830184611faa565b9392505050565b60006040820190506121296000830185611fb9565b6121366020830184611fd7565b9392505050565b60006040820190506121526000830185611fd7565b61215f6020830184611fd7565b9392505050565b6000612170612181565b905061217c828261232f565b919050565b6000604051905090565b600067ffffffffffffffff8211156121a6576121a5612442565b5b602082029050919050565b600067ffffffffffffffff8211156121cc576121cb612442565b5b602082029050602081019050919050565b6000819050919050565b600060099050919050565b6000602082019050919050565b600081905092915050565b600061221582612301565b915061222083612301565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115612255576122546123e4565b5b828201905092915050565b600061226b82612301565b915061227683612301565b92508261228657612285612413565b5b828204905092915050565b600061229c82612301565b91506122a783612301565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156122e0576122df6123e4565b5b828202905092915050565b60008115159050919050565b6000819050919050565b6000819050919050565b600061231682612301565b9050919050565b600061232882612301565b9050919050565b61233882612471565b810181811067ffffffffffffffff8211171561235757612356612442565b5b80604052505050565b600061236b82612301565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561239e5761239d6123e4565b5b600182019050919050565b6000819050919050565b60006123be82612301565b91506123c983612301565b9250826123d9576123d8612413565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b61248b816122f7565b811461249657600080fd5b50565b6124a281612301565b81146124ad57600080fd5b5056fea2646970667358221220aa6b825af397cb84f40a2d6a274231c243d44fac944e7c7e5d13cd322dc2e86064736f6c63430008040033"

// DeployEthash deploys a new Ethereum contract, binding an instance of Ethash to it.
func DeployEthash(auth *bind.TransactOpts, backend bind.ContractBackend, SHA3_address common.Address) (common.Address, *types.Transaction, *Ethash, error) {
	parsed, err := abi.JSON(strings.NewReader(EthashABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthashBin), backend, SHA3_address)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ethash{EthashCaller: EthashCaller{contract: contract}, EthashTransactor: EthashTransactor{contract: contract}, EthashFilterer: EthashFilterer{contract: contract}}, nil
}

// Ethash is an auto generated Go binding around an Ethereum contract.
type Ethash struct {
	EthashCaller     // Read-only binding to the contract
	EthashTransactor // Write-only binding to the contract
	EthashFilterer   // Log filterer for contract events
}

// EthashCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthashTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthashFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthashSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthashSession struct {
	Contract     *Ethash           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthashCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthashCallerSession struct {
	Contract *EthashCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EthashTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthashTransactorSession struct {
	Contract     *EthashTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthashRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthashRaw struct {
	Contract *Ethash // Generic contract binding to access the raw methods on
}

// EthashCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthashCallerRaw struct {
	Contract *EthashCaller // Generic read-only contract binding to access the raw methods on
}

// EthashTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthashTransactorRaw struct {
	Contract *EthashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthash creates a new instance of Ethash, bound to a specific deployed contract.
func NewEthash(address common.Address, backend bind.ContractBackend) (*Ethash, error) {
	contract, err := bindEthash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethash{EthashCaller: EthashCaller{contract: contract}, EthashTransactor: EthashTransactor{contract: contract}, EthashFilterer: EthashFilterer{contract: contract}}, nil
}

// NewEthashCaller creates a new read-only instance of Ethash, bound to a specific deployed contract.
func NewEthashCaller(address common.Address, caller bind.ContractCaller) (*EthashCaller, error) {
	contract, err := bindEthash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthashCaller{contract: contract}, nil
}

// NewEthashTransactor creates a new write-only instance of Ethash, bound to a specific deployed contract.
func NewEthashTransactor(address common.Address, transactor bind.ContractTransactor) (*EthashTransactor, error) {
	contract, err := bindEthash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthashTransactor{contract: contract}, nil
}

// NewEthashFilterer creates a new log filterer instance of Ethash, bound to a specific deployed contract.
func NewEthashFilterer(address common.Address, filterer bind.ContractFilterer) (*EthashFilterer, error) {
	contract, err := bindEthash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthashFilterer{contract: contract}, nil
}

// bindEthash binds a generic wrapper to an already deployed contract.
func bindEthash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthashABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethash *EthashRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ethash.Contract.EthashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethash *EthashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethash.Contract.EthashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethash *EthashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethash.Contract.EthashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethash *EthashCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ethash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethash *EthashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethash *EthashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethash.Contract.contract.Transact(opts, method, params...)
}

// IsEpochDataSet is a free data retrieval call binding the contract method 0xc7b81f4f.
//
// Solidity: function isEpochDataSet(uint256 epochIndex) view returns(bool)
func (_Ethash *EthashCaller) IsEpochDataSet(opts *bind.CallOpts, epochIndex *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ethash.contract.Call(opts, out, "isEpochDataSet", epochIndex)
	return *ret0, err
}

// IsEpochDataSet is a free data retrieval call binding the contract method 0xc7b81f4f.
//
// Solidity: function isEpochDataSet(uint256 epochIndex) view returns(bool)
func (_Ethash *EthashSession) IsEpochDataSet(epochIndex *big.Int) (bool, error) {
	return _Ethash.Contract.IsEpochDataSet(&_Ethash.CallOpts, epochIndex)
}

// IsEpochDataSet is a free data retrieval call binding the contract method 0xc7b81f4f.
//
// Solidity: function isEpochDataSet(uint256 epochIndex) view returns(bool)
func (_Ethash *EthashCallerSession) IsEpochDataSet(epochIndex *big.Int) (bool, error) {
	return _Ethash.Contract.IsEpochDataSet(&_Ethash.CallOpts, epochIndex)
}

// VerifyPoW is a free data retrieval call binding the contract method 0x29e265df.
//
// Solidity: function verifyPoW(uint256 blockNumber, bytes32 rlpHeaderHashWithoutNonce, uint256 nonce, uint256 difficulty, uint256[] dataSetLookup, uint256[] witnessForLookup) view returns(uint256, uint256)
func (_Ethash *EthashCaller) VerifyPoW(opts *bind.CallOpts, blockNumber *big.Int, rlpHeaderHashWithoutNonce [32]byte, nonce *big.Int, difficulty *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Ethash.contract.Call(opts, out, "verifyPoW", blockNumber, rlpHeaderHashWithoutNonce, nonce, difficulty, dataSetLookup, witnessForLookup)
	return *ret0, *ret1, err
}

// VerifyPoW is a free data retrieval call binding the contract method 0x29e265df.
//
// Solidity: function verifyPoW(uint256 blockNumber, bytes32 rlpHeaderHashWithoutNonce, uint256 nonce, uint256 difficulty, uint256[] dataSetLookup, uint256[] witnessForLookup) view returns(uint256, uint256)
func (_Ethash *EthashSession) VerifyPoW(blockNumber *big.Int, rlpHeaderHashWithoutNonce [32]byte, nonce *big.Int, difficulty *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int) (*big.Int, *big.Int, error) {
	return _Ethash.Contract.VerifyPoW(&_Ethash.CallOpts, blockNumber, rlpHeaderHashWithoutNonce, nonce, difficulty, dataSetLookup, witnessForLookup)
}

// VerifyPoW is a free data retrieval call binding the contract method 0x29e265df.
//
// Solidity: function verifyPoW(uint256 blockNumber, bytes32 rlpHeaderHashWithoutNonce, uint256 nonce, uint256 difficulty, uint256[] dataSetLookup, uint256[] witnessForLookup) view returns(uint256, uint256)
func (_Ethash *EthashCallerSession) VerifyPoW(blockNumber *big.Int, rlpHeaderHashWithoutNonce [32]byte, nonce *big.Int, difficulty *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int) (*big.Int, *big.Int, error) {
	return _Ethash.Contract.VerifyPoW(&_Ethash.CallOpts, blockNumber, rlpHeaderHashWithoutNonce, nonce, difficulty, dataSetLookup, witnessForLookup)
}

// SetEpochData is a paid mutator transaction binding the contract method 0xc891a29d.
//
// Solidity: function setEpochData(uint256 epoch, uint256 fullSizeIn128Resultion, uint256 branchDepth, uint256[] merkleNodes, uint256 start, uint256 numElems) returns()
func (_Ethash *EthashTransactor) SetEpochData(opts *bind.TransactOpts, epoch *big.Int, fullSizeIn128Resultion *big.Int, branchDepth *big.Int, merkleNodes []*big.Int, start *big.Int, numElems *big.Int) (*types.Transaction, error) {
	return _Ethash.contract.Transact(opts, "setEpochData", epoch, fullSizeIn128Resultion, branchDepth, merkleNodes, start, numElems)
}

// SetEpochData is a paid mutator transaction binding the contract method 0xc891a29d.
//
// Solidity: function setEpochData(uint256 epoch, uint256 fullSizeIn128Resultion, uint256 branchDepth, uint256[] merkleNodes, uint256 start, uint256 numElems) returns()
func (_Ethash *EthashSession) SetEpochData(epoch *big.Int, fullSizeIn128Resultion *big.Int, branchDepth *big.Int, merkleNodes []*big.Int, start *big.Int, numElems *big.Int) (*types.Transaction, error) {
	return _Ethash.Contract.SetEpochData(&_Ethash.TransactOpts, epoch, fullSizeIn128Resultion, branchDepth, merkleNodes, start, numElems)
}

// SetEpochData is a paid mutator transaction binding the contract method 0xc891a29d.
//
// Solidity: function setEpochData(uint256 epoch, uint256 fullSizeIn128Resultion, uint256 branchDepth, uint256[] merkleNodes, uint256 start, uint256 numElems) returns()
func (_Ethash *EthashTransactorSession) SetEpochData(epoch *big.Int, fullSizeIn128Resultion *big.Int, branchDepth *big.Int, merkleNodes []*big.Int, start *big.Int, numElems *big.Int) (*types.Transaction, error) {
	return _Ethash.Contract.SetEpochData(&_Ethash.TransactOpts, epoch, fullSizeIn128Resultion, branchDepth, merkleNodes, start, numElems)
}

// EthashSetEpochDataIterator is returned from FilterSetEpochData and is used to iterate over the raw logs and unpacked data for SetEpochData events raised by the Ethash contract.
type EthashSetEpochDataIterator struct {
	Event *EthashSetEpochData // Event containing the contract specifics and raw log

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
func (it *EthashSetEpochDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthashSetEpochData)
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
		it.Event = new(EthashSetEpochData)
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
func (it *EthashSetEpochDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthashSetEpochDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthashSetEpochData represents a SetEpochData event raised by the Ethash contract.
type EthashSetEpochData struct {
	Sender    common.Address
	Error     *big.Int
	ErrorInfo *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetEpochData is a free log retrieval operation binding the contract event 0x5cd723400be8430351b9cbaa5ea421b3fb2528c6a7650c493f895e7d97750da1.
//
// Solidity: event SetEpochData(address indexed sender, uint256 error, uint256 errorInfo)
func (_Ethash *EthashFilterer) FilterSetEpochData(opts *bind.FilterOpts, sender []common.Address) (*EthashSetEpochDataIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ethash.contract.FilterLogs(opts, "SetEpochData", senderRule)
	if err != nil {
		return nil, err
	}
	return &EthashSetEpochDataIterator{contract: _Ethash.contract, event: "SetEpochData", logs: logs, sub: sub}, nil
}

// WatchSetEpochData is a free log subscription operation binding the contract event 0x5cd723400be8430351b9cbaa5ea421b3fb2528c6a7650c493f895e7d97750da1.
//
// Solidity: event SetEpochData(address indexed sender, uint256 error, uint256 errorInfo)
func (_Ethash *EthashFilterer) WatchSetEpochData(opts *bind.WatchOpts, sink chan<- *EthashSetEpochData, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ethash.contract.WatchLogs(opts, "SetEpochData", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthashSetEpochData)
				if err := _Ethash.contract.UnpackLog(event, "SetEpochData", log); err != nil {
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

// ParseSetEpochData is a log parse operation binding the contract event 0x5cd723400be8430351b9cbaa5ea421b3fb2528c6a7650c493f895e7d97750da1.
//
// Solidity: event SetEpochData(address indexed sender, uint256 error, uint256 errorInfo)
func (_Ethash *EthashFilterer) ParseSetEpochData(log types.Log) (*EthashSetEpochData, error) {
	event := new(EthashSetEpochData)
	if err := _Ethash.contract.UnpackLog(event, "SetEpochData", log); err != nil {
		return nil, err
	}
	return event, nil
}
