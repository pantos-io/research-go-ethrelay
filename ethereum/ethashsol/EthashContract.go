// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethashsol

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

// EthashsolMetaData contains all meta data concerning the Ethashsol contract.
var EthashsolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"SetEpochData\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochIndex\",\"type\":\"uint256\"}],\"name\":\"isEpochDataSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fullSizeIn128Resultion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"branchDepth\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"merkleNodes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numElems\",\"type\":\"uint256\"}],\"name\":\"setEpochData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"rlpHeaderHashWithoutNonce\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"dataSetLookup\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"witnessForLookup\",\"type\":\"uint256[]\"}],\"name\":\"verifyPoW\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612153806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806329e265df14610046578063c7b81f4f14610073578063c891a29d146100a9575b600080fd5b610059610054366004611eb1565b6100be565b604080519283526020830191909152015b60405180910390f35b610099610081366004611f46565b60009081526020819052604090206102000154151590565b604051901515815260200161006a565b6100bc6100b7366004611f75565b61019b565b005b600080806100ce6175308c61208d565b905060006101438b8b8a8a8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808e0282810182019093528d82529093508d92508c9182918501908490808284376000920191909152508992506102f9915050565b90506101518960001961208d565b8111156101845760008082600119141561017057506001905082610177565b5060029050815b909450925061018e915050565b6000809350935050505b9850989650505050505050565b60005b818110156102985760008781526020819052604081206101be83866120a1565b61020081106101cf576101cf6120b9565b0154111561023d57337f5cd723400be8430351b9cbaa5ea421b3fb2528c6a7650c493f895e7d97750da16001838661020b8c600160801b6120cf565b61021591906120a1565b61021f91906120a1565b6040805192835260208301919091520160405180910390a2506102f1565b83818151811061024f5761024f6120b9565b602090810291909101810151600089815291829052604090912061027383866120a1565b6102008110610284576102846120b9565b015580610290816120ee565b91505061019e565b5060008681526020818152604080832061020081018990556102010187905580518381529182019290925233917f5cd723400be8430351b9cbaa5ea421b3fb2528c6a7650c493f895e7d97750da1910160405180910390a25b505050505050565b6000610303611dad565b61030b611dcc565b610313611deb565b6040805180820182526000878152602081815292812061020181015483526102000154838301819052888252928190529091819061035b5760011996505050505050506105ac565b60208301516103745760001996505050505050506105ac565b61037e8c8c6105b5565b9550600091505b60108210156103b7576020820286810180519187019182525161020090910152816103af816120ee565b925050610385565b600091505b60408210156105115760208084015187516000926103fd9186189089906103e39088612109565b602081106103f3576103f36120b9565b6020020151610651565b6104079190612109565b90506104138982610671565b61042782858e8e89600060200201516106d6565b1461043d576000199750505050505050506105ac565b600091505b60088210156104fe57602060808481028d0182810180519386028a01805163ffffffff9586166301000193918202188616825260408401805161010084018051918916918402919091188816905260608501805161020085018051918a169185029190911889169052949095018051610300909301805193881693909202929092189095169094528051600160201b908190049091528251819004909252805182900490528151049052816104f6816120ee565b925050610442565b5081610509816120ee565b9250506103bc565b600091505b60208210156105965761056361055761054b87856020811061053a5761053a6120b9565b6020020151886103e38760016120a1565b876103e38660026120a1565b866103e38560036120a1565b8461056f60048561208d565b6008811061057f5761057f6120b9565b602002015261058f6004836120a1565b9150610516565b60006105a287866108b2565b9750505050505050505b95945050505050565b6105bd611dad565b6105c5611e0a565b6105ce84610b90565b6001600160401b038116825293506105ea600160401b8561208d565b6001600160401b03811660208301529350610609600160401b8561208d565b6001600160401b03811660408301529350610628600160401b8561208d565b6001600160401b038116606083015260808201849052935061064981610be2565b949350505050565b6000816106628463010001936120cf565b1863ffffffff16905092915050565b600082815260208190526040812061020181015483901c90829061069660028461208d565b61020081106106a7576106a76120b9565b015490506106b6600283612109565b6106c8576001600160801b0316610649565b6105ac600160801b8261208d565b6000806106e38587610ee2565b6001600160801b031690506000808080806106ff600289612109565b6002909804971190506000610714888c6120cf565b90508115610729576107268b826120a1565b90505b60005b888110156108245760208282010260208b01015193508c60011660001415610762578695506001600160801b0384169450610772565b6001600160801b03841695508694505b604080516020810188905290810186905260600160408051601f1981840301815291905280516020909101206002909d049c6001600160801b0316965060018d166107c857869550600160801b840494506107d5565b600160801b840495508694505b604080516020810188905290810186905260600160408051601f1981840301815291905280516020909101206002909d049c6001600160801b031696508061081c816120ee565b91505061072c565b82156108a15760208282010260208b01015193508c60011660001415610858578695506001600160801b0384169450610868565b6001600160801b03841695508694505b60408051602081018890529081018690526060016040516020818303038152906040528051906020012060001c6001600160801b031696505b50949b9a5050505050505050505050565b60008083600760200201516108cb90600160601b6120cf565b60c08501516108de90600160401b6120cf565b60a08601516108f190600160201b6120cf565b608087015161090091906120a1565b61090a91906120a1565b61091491906120a1565b61092290600160801b6120cf565b606085015161093590600160601b6120cf565b604086015161094890600160401b6120cf565b602087015161095b90600160201b6120cf565b875161096791906120a1565b61097191906120a1565b61097b91906120a1565b61098591906120a1565b9050600084600f602002015161099f90600160601b6120cf565b6101c08601516109b390600160401b6120cf565b6101a08701516109c790600160201b6120cf565b6101808801516109d791906120a1565b6109e191906120a1565b6109eb91906120a1565b6109f990600160801b6120cf565b610160860151610a0d90600160601b6120cf565b610140870151610a2190600160401b6120cf565b610120880151610a3590600160201b6120cf565b610100890151610a4591906120a1565b610a4f91906120a1565b610a5991906120a1565b610a6391906120a1565b905060008460076020020151610a7d90600160601b6120cf565b60c0860151610a9090600160401b6120cf565b60a0870151610aa390600160201b6120cf565b6080880151610ab291906120a1565b610abc91906120a1565b610ac691906120a1565b610ad490600160801b6120cf565b6060860151610ae790600160601b6120cf565b6040870151610afa90600160401b6120cf565b6020880151610b0d90600160201b6120cf565b8851610b1991906120a1565b610b2391906120a1565b610b2d91906120a1565b610b3791906120a1565b9050610b4283610fe3565b610b4b83610fe3565b610b5483610fe3565b604080516020810194909452830191909152606082015260800160408051601f1981840301815291905280516020909101209695505050505050565b600080805b6020811015610bdb57610baa826101006120cf565b9150610bb960ff8516836120a1565b9150610bc76101008561208d565b935080610bd3816120ee565b915050610b95565b5092915050565b610bea611dad565b610bf6600960086120cf565b604814610c385760405162461bcd60e51b815260206004820152600c60248201526b39b837b733b29032b93937b960a11b604482015260640160405180910390fd5b600160a08301526001603f1b610100830152604860086000610c5b6009836120cf565b9050610c65611e29565b600080805b610c74878661208d565b831015610d9757600091505b6005821015610d7a575060005b6005811015610d6857610ca0868861208d565b610cab8360056120cf565b610cb590836120a1565b1015610d565788610cc78360056120cf565b82610cd38660096120cf565b610cdd91906120a1565b610ce791906120a1565b60098110610cf757610cf76120b9565b60200201518483610d098460056120cf565b610d1391906120a1565b60198110610d2357610d236120b9565b6020020151188483610d368460056120cf565b610d4091906120a1565b60198110610d5057610d506120b9565b60200201525b80610d60816120ee565b915050610c8d565b81610d72816120ee565b925050610c80565b610d838461102e565b935082610d8f816120ee565b935050610c6a565b610d9f611dad565b60005b6010811015610ed457600093505b6005841015610ecf57600092505b6005831015610ebd57610dd1888a61208d565b610ddc8560056120cf565b610de690856120a1565b108015610df35750601081105b15610eab578584610e058560056120cf565b610e0f91906120a1565b60198110610e1f57610e1f6120b9565b602002015163ffffffff16828260108110610e3c57610e3c6120b9565b6020020152600160201b8685610e538660056120cf565b610e5d91906120a1565b60198110610e6d57610e6d6120b9565b6020020151610e7c919061208d565b82610e888360016120a1565b60108110610e9857610e986120b9565b6020020152610ea86002826120a1565b90505b82610eb5816120ee565b935050610dbe565b83610ec7816120ee565b945050610db0565b610da2565b509998505050505050505050565b600082610ef08360046120cf565b81518110610f0057610f006120b9565b602002602001015183836004610f1691906120cf565b610f219060016120a1565b81518110610f3157610f316120b9565b602002602001015184846004610f4791906120cf565b610f529060026120a1565b81518110610f6257610f626120b9565b602002602001015185856004610f7891906120cf565b610f839060036120a1565b81518110610f9357610f936120b9565b6020026020010151604051602001610fc4949392919093845260208401929092526040830152606082015260800190565b60408051601f1981840301815291905280516020909101209392505050565b600080805b6020811015610bdb57610ffd826101006120cf565b915061100c60ff8516836120a1565b915061101a6101008561208d565b935080611026816120ee565b915050610fe8565b611036611e29565b61103e611e29565b611046611e48565b61104e611e48565b600060405180610300016040528060018152602001618082815260200167800000000000808a8152602001678000000080008000815260200161808b81526020016380000001815260200167800000008000808181526020016780000000000080098152602001608a81526020016088815260200163800080098152602001638000000a8152602001638000808b815260200167800000000000008b8152602001678000000000008089815260200167800000000000800381526020016780000000000080028152602001678000000000000080815260200161800a815260200167800000008000000a81526020016780000000800080818152602001678000000000008080815260200163800000018152602001678000000080008008815250905060005b6018811015611da2576080808801516060808a01516040808c01516020808e01518e511890911890921890931888526101208b01516101008c015160e08d015160c08e015160a08f0151181818189088018190526101c08b01516101a08c01516101808d01516101608e01516101408f015118181818928801929092526102608a01516102408b01516102208c01516102008d01516101e08e015118181818908701526103008901516102e08a01516102c08b01516102a08c01516102808d0151181818189186019190915261124f906001603f1b9061208d565b602085015161125f9060026120cf565b60808601516001600160401b039190911691909117188352604084015161128b906001603f1b9061208d565b604085015161129b9060026120cf565b85516001600160401b03919091169190911718602084015260608401516112c7906001603f1b9061208d565b60608501516112d79060026120cf565b6001600160401b03161784600160200201511883600260200201526001603f1b8460046020020151611309919061208d565b60808501516113199060026120cf565b60408601516001600160401b0391909116919091171860608401528351611345906001603f1b9061208d565b84516113529060026120cf565b6060808701516001600160401b03929092169290921718608080860191825285518a5118808b52865160208c81018051909218825288516040808f01805190921890915289518e8801805190911890528951948e0180519095189094528801805160a08e0180519091189052805160c08e0180519091189052805160e08e018051909118905280516101008e0180519091189052516101208d018051909118905291870180516101408d018051909118905280516101608d018051909118905280516101808d018051909118905280516101a08d0180519091189052516101c08c018051909118905292860180516101e08c018051909118905280516102008c018051909118905280516102208c018051909118905280516102408c0180519091189052516102608b018051909118905281516102808b018051909118905281516102a08b018051909118905281516102c08b018051909118905281516102e08b018051909118905290516103008a0180519091189052908652516114dc9063100000009061208d565b60208801516114f0906410000000006120cf565b6001600160401b0316176101008601526040870151611514906001603d1b9061208d565b60408801516115249060086120cf565b6001600160401b031617610160860152606087015161154790628000009061208d565b606088015161155c90650200000000006120cf565b6001600160401b031617610260860152608087015161158290654000000000009061208d565b608088015161159490620400006120cf565b6001600160401b0316176102c086015260a08701516115b8906001603f1b9061208d565b60a08801516115c89060026120cf565b6001600160401b031617604086015260c08701516115ea90621000009061208d565b60c08801516115ff90651000000000006120cf565b6001600160401b03161760a086015260e08701516116259066400000000000009061208d565b60e0880151611636906104006120cf565b6001600160401b0316176101a086015261010087015161165a90620800009061208d565b61010088015161167090652000000000006120cf565b6001600160401b031617610200860152610120870151611695906001603e1b9061208d565b6101208801516116a69060046120cf565b6001600160401b0316176103008601526101408701516116c89060049061208d565b6101408801516116dc906001603e1b6120cf565b6001600160401b0316176080860152610160870151611704906704000000000000009061208d565b6101608801516117159060406120cf565b6001600160401b03161760e086015261018087015161173890622000009061208d565b61018088015161174e90650800000000006120cf565b6001600160401b0316176101408601526101a08701516117769066020000000000009061208d565b6101a0880151611788906180006120cf565b6001600160401b0316176102408601526101c08701516117aa9060089061208d565b6101c08801516117be906001603d1b6120cf565b6001600160401b0316176102a08601526101e08701516117e4906410000000009061208d565b6101e08801516117f89063100000006120cf565b6001600160401b031617602086015261020087810151611818919061208d565b61020088015161182f9066800000000000006120cf565b6001600160401b031617610120860152610220870151611855906480000000009061208d565b6102208801516118699063020000006120cf565b6001600160401b03161761018086015261024087015161189090650800000000009061208d565b6102408801516118a390622000006120cf565b6001600160401b0316176101e08601526102608701516118c6906101009061208d565b6102608801516118da90600160381b6120cf565b6001600160401b0316176102e0860152610280870151611900906420000000009061208d565b6102808801516119149063080000006120cf565b6001600160401b03161760608601526102a087015161193a90651000000000009061208d565b6102a088015161194d90621000006120cf565b6001600160401b03161760c08601526102c08701516119719063020000009061208d565b6102c0880151611986906480000000006120cf565b6001600160401b0316176101c08601526102e08701516119ab90600160381b9061208d565b6102e08801516119bd906101006120cf565b6001600160401b0316176102208601526103008701516119e59066040000000000009061208d565b6103008801516119f7906140006120cf565b6001600160401b031617856014602002015284600a602002015185600560200201511916856000602002015118876000602002015284600b602002015185600660200201511916856001602002015118876001602002015284600c602002015185600760200201511916856002602002015118876002602002015284600d602002015185600860200201511916856003602002015118876003602002015284600e602002015185600960200201511916856004602002015118876004602002015284600f602002015185600a602002015119168560056020020151188760056020020152846010602002015185600b602002015119168560066020020151188760066020020152846011602002015185600c602002015119168560076020020151188760076020020152846012602002015185600d602002015119168560086020020151188760086020020152846013602002015185600e602002015119168560096020020151188760096020020152846014602002015185600f6020020151191685600a60200201511887600a602002015284601560200201518560106020020151191685600b60200201511887600b602002015284601660200201518560116020020151191685600c60200201511887600c602002015284601760200201518560126020020151191685600d60200201511887600d602002015284601860200201518560136020020151191685600e60200201511887600e602002015284600060200201518560146020020151191685600f60200201511887600f6020020152846001602002015185601560200201511916856010602002015118876010602002015284600260200201518560166020020151191685601160200201511887601160200201528460036020020151856017602002015119168560126020020151188760126020020152846004602002015185601860200201511916856013602002015118876013602002015284600560200201518560006020020151191685601460200201511887601460200201528460066020020151856001602002015119168560156020020151188760156020020152846007602002015185600260200201511916856016602002015118876016602002015284600860200201518560036020020151191685601760200201511887601760200201528460096020020151856004602002015119168560186020020151188760186020020152818160188110611d8657611d866120b9565b6020020151875118875280611d9a816120ee565b915050611174565b509495945050505050565b6040518061020001604052806010906020820280368337509192915050565b6040518061040001604052806020906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b6040518061012001604052806009906020820280368337509192915050565b6040518061032001604052806019906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b60008083601f840112611e7857600080fd5b5081356001600160401b03811115611e8f57600080fd5b6020830191508360208260051b8501011115611eaa57600080fd5b9250929050565b60008060008060008060008060c0898b031215611ecd57600080fd5b8835975060208901359650604089013595506060890135945060808901356001600160401b0380821115611f0057600080fd5b611f0c8c838d01611e66565b909650945060a08b0135915080821115611f2557600080fd5b50611f328b828c01611e66565b999c989b5096995094979396929594505050565b600060208284031215611f5857600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b60008060008060008060c08789031215611f8e57600080fd5b8635955060208088013595506040880135945060608801356001600160401b0380821115611fbb57600080fd5b818a0191508a601f830112611fcf57600080fd5b813581811115611fe157611fe1611f5f565b8060051b604051601f19603f8301168101818110858211171561200657612006611f5f565b60405291825284820192508381018501918d83111561202457600080fd5b938501935b8285101561204257843584529385019392850192612029565b9a9d999c50979a60808101359960a09091013598509650505050505050565b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60008261209c5761209c612061565b500490565b600082198211156120b4576120b4612077565b500190565b634e487b7160e01b600052603260045260246000fd5b60008160001904831182151516156120e9576120e9612077565b500290565b600060001982141561210257612102612077565b5060010190565b60008261211857612118612061565b50069056fea26469706673582212206da6dc0d9b5862576ab879bf536c231ac738f97d85d050a9663f318a8cec8e2364736f6c634300080a0033",
}

// EthashsolABI is the input ABI used to generate the binding from.
// Deprecated: Use EthashsolMetaData.ABI instead.
var EthashsolABI = EthashsolMetaData.ABI

// EthashsolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthashsolMetaData.Bin instead.
var EthashsolBin = EthashsolMetaData.Bin

// DeployEthashsol deploys a new Ethereum contract, binding an instance of Ethashsol to it.
func DeployEthashsol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ethashsol, error) {
	parsed, err := EthashsolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthashsolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ethashsol{EthashsolCaller: EthashsolCaller{contract: contract}, EthashsolTransactor: EthashsolTransactor{contract: contract}, EthashsolFilterer: EthashsolFilterer{contract: contract}}, nil
}

// Ethashsol is an auto generated Go binding around an Ethereum contract.
type Ethashsol struct {
	EthashsolCaller     // Read-only binding to the contract
	EthashsolTransactor // Write-only binding to the contract
	EthashsolFilterer   // Log filterer for contract events
}

// EthashsolCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthashsolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthashsolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthashsolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthashsolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthashsolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthashsolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthashsolSession struct {
	Contract     *Ethashsol        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthashsolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthashsolCallerSession struct {
	Contract *EthashsolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EthashsolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthashsolTransactorSession struct {
	Contract     *EthashsolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EthashsolRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthashsolRaw struct {
	Contract *Ethashsol // Generic contract binding to access the raw methods on
}

// EthashsolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthashsolCallerRaw struct {
	Contract *EthashsolCaller // Generic read-only contract binding to access the raw methods on
}

// EthashsolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthashsolTransactorRaw struct {
	Contract *EthashsolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthashsol creates a new instance of Ethashsol, bound to a specific deployed contract.
func NewEthashsol(address common.Address, backend bind.ContractBackend) (*Ethashsol, error) {
	contract, err := bindEthashsol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethashsol{EthashsolCaller: EthashsolCaller{contract: contract}, EthashsolTransactor: EthashsolTransactor{contract: contract}, EthashsolFilterer: EthashsolFilterer{contract: contract}}, nil
}

// NewEthashsolCaller creates a new read-only instance of Ethashsol, bound to a specific deployed contract.
func NewEthashsolCaller(address common.Address, caller bind.ContractCaller) (*EthashsolCaller, error) {
	contract, err := bindEthashsol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthashsolCaller{contract: contract}, nil
}

// NewEthashsolTransactor creates a new write-only instance of Ethashsol, bound to a specific deployed contract.
func NewEthashsolTransactor(address common.Address, transactor bind.ContractTransactor) (*EthashsolTransactor, error) {
	contract, err := bindEthashsol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthashsolTransactor{contract: contract}, nil
}

// NewEthashsolFilterer creates a new log filterer instance of Ethashsol, bound to a specific deployed contract.
func NewEthashsolFilterer(address common.Address, filterer bind.ContractFilterer) (*EthashsolFilterer, error) {
	contract, err := bindEthashsol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthashsolFilterer{contract: contract}, nil
}

// bindEthashsol binds a generic wrapper to an already deployed contract.
func bindEthashsol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthashsolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethashsol *EthashsolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethashsol.Contract.EthashsolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethashsol *EthashsolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethashsol.Contract.EthashsolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethashsol *EthashsolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethashsol.Contract.EthashsolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethashsol *EthashsolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethashsol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethashsol *EthashsolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethashsol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethashsol *EthashsolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethashsol.Contract.contract.Transact(opts, method, params...)
}

// IsEpochDataSet is a free data retrieval call binding the contract method 0xc7b81f4f.
//
// Solidity: function isEpochDataSet(uint256 epochIndex) view returns(bool)
func (_Ethashsol *EthashsolCaller) IsEpochDataSet(opts *bind.CallOpts, epochIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Ethashsol.contract.Call(opts, &out, "isEpochDataSet", epochIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEpochDataSet is a free data retrieval call binding the contract method 0xc7b81f4f.
//
// Solidity: function isEpochDataSet(uint256 epochIndex) view returns(bool)
func (_Ethashsol *EthashsolSession) IsEpochDataSet(epochIndex *big.Int) (bool, error) {
	return _Ethashsol.Contract.IsEpochDataSet(&_Ethashsol.CallOpts, epochIndex)
}

// IsEpochDataSet is a free data retrieval call binding the contract method 0xc7b81f4f.
//
// Solidity: function isEpochDataSet(uint256 epochIndex) view returns(bool)
func (_Ethashsol *EthashsolCallerSession) IsEpochDataSet(epochIndex *big.Int) (bool, error) {
	return _Ethashsol.Contract.IsEpochDataSet(&_Ethashsol.CallOpts, epochIndex)
}

// VerifyPoW is a free data retrieval call binding the contract method 0x29e265df.
//
// Solidity: function verifyPoW(uint256 blockNumber, bytes32 rlpHeaderHashWithoutNonce, uint256 nonce, uint256 difficulty, uint256[] dataSetLookup, uint256[] witnessForLookup) view returns(uint256, uint256)
func (_Ethashsol *EthashsolCaller) VerifyPoW(opts *bind.CallOpts, blockNumber *big.Int, rlpHeaderHashWithoutNonce [32]byte, nonce *big.Int, difficulty *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Ethashsol.contract.Call(opts, &out, "verifyPoW", blockNumber, rlpHeaderHashWithoutNonce, nonce, difficulty, dataSetLookup, witnessForLookup)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// VerifyPoW is a free data retrieval call binding the contract method 0x29e265df.
//
// Solidity: function verifyPoW(uint256 blockNumber, bytes32 rlpHeaderHashWithoutNonce, uint256 nonce, uint256 difficulty, uint256[] dataSetLookup, uint256[] witnessForLookup) view returns(uint256, uint256)
func (_Ethashsol *EthashsolSession) VerifyPoW(blockNumber *big.Int, rlpHeaderHashWithoutNonce [32]byte, nonce *big.Int, difficulty *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int) (*big.Int, *big.Int, error) {
	return _Ethashsol.Contract.VerifyPoW(&_Ethashsol.CallOpts, blockNumber, rlpHeaderHashWithoutNonce, nonce, difficulty, dataSetLookup, witnessForLookup)
}

// VerifyPoW is a free data retrieval call binding the contract method 0x29e265df.
//
// Solidity: function verifyPoW(uint256 blockNumber, bytes32 rlpHeaderHashWithoutNonce, uint256 nonce, uint256 difficulty, uint256[] dataSetLookup, uint256[] witnessForLookup) view returns(uint256, uint256)
func (_Ethashsol *EthashsolCallerSession) VerifyPoW(blockNumber *big.Int, rlpHeaderHashWithoutNonce [32]byte, nonce *big.Int, difficulty *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int) (*big.Int, *big.Int, error) {
	return _Ethashsol.Contract.VerifyPoW(&_Ethashsol.CallOpts, blockNumber, rlpHeaderHashWithoutNonce, nonce, difficulty, dataSetLookup, witnessForLookup)
}

// SetEpochData is a paid mutator transaction binding the contract method 0xc891a29d.
//
// Solidity: function setEpochData(uint256 epoch, uint256 fullSizeIn128Resultion, uint256 branchDepth, uint256[] merkleNodes, uint256 start, uint256 numElems) returns()
func (_Ethashsol *EthashsolTransactor) SetEpochData(opts *bind.TransactOpts, epoch *big.Int, fullSizeIn128Resultion *big.Int, branchDepth *big.Int, merkleNodes []*big.Int, start *big.Int, numElems *big.Int) (*types.Transaction, error) {
	return _Ethashsol.contract.Transact(opts, "setEpochData", epoch, fullSizeIn128Resultion, branchDepth, merkleNodes, start, numElems)
}

// SetEpochData is a paid mutator transaction binding the contract method 0xc891a29d.
//
// Solidity: function setEpochData(uint256 epoch, uint256 fullSizeIn128Resultion, uint256 branchDepth, uint256[] merkleNodes, uint256 start, uint256 numElems) returns()
func (_Ethashsol *EthashsolSession) SetEpochData(epoch *big.Int, fullSizeIn128Resultion *big.Int, branchDepth *big.Int, merkleNodes []*big.Int, start *big.Int, numElems *big.Int) (*types.Transaction, error) {
	return _Ethashsol.Contract.SetEpochData(&_Ethashsol.TransactOpts, epoch, fullSizeIn128Resultion, branchDepth, merkleNodes, start, numElems)
}

// SetEpochData is a paid mutator transaction binding the contract method 0xc891a29d.
//
// Solidity: function setEpochData(uint256 epoch, uint256 fullSizeIn128Resultion, uint256 branchDepth, uint256[] merkleNodes, uint256 start, uint256 numElems) returns()
func (_Ethashsol *EthashsolTransactorSession) SetEpochData(epoch *big.Int, fullSizeIn128Resultion *big.Int, branchDepth *big.Int, merkleNodes []*big.Int, start *big.Int, numElems *big.Int) (*types.Transaction, error) {
	return _Ethashsol.Contract.SetEpochData(&_Ethashsol.TransactOpts, epoch, fullSizeIn128Resultion, branchDepth, merkleNodes, start, numElems)
}

// EthashsolSetEpochDataIterator is returned from FilterSetEpochData and is used to iterate over the raw logs and unpacked data for SetEpochData events raised by the Ethashsol contract.
type EthashsolSetEpochDataIterator struct {
	Event *EthashsolSetEpochData // Event containing the contract specifics and raw log

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
func (it *EthashsolSetEpochDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthashsolSetEpochData)
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
		it.Event = new(EthashsolSetEpochData)
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
func (it *EthashsolSetEpochDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthashsolSetEpochDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthashsolSetEpochData represents a SetEpochData event raised by the Ethashsol contract.
type EthashsolSetEpochData struct {
	Sender    common.Address
	Error     *big.Int
	ErrorInfo *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetEpochData is a free log retrieval operation binding the contract event 0x5cd723400be8430351b9cbaa5ea421b3fb2528c6a7650c493f895e7d97750da1.
//
// Solidity: event SetEpochData(address indexed sender, uint256 error, uint256 errorInfo)
func (_Ethashsol *EthashsolFilterer) FilterSetEpochData(opts *bind.FilterOpts, sender []common.Address) (*EthashsolSetEpochDataIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ethashsol.contract.FilterLogs(opts, "SetEpochData", senderRule)
	if err != nil {
		return nil, err
	}
	return &EthashsolSetEpochDataIterator{contract: _Ethashsol.contract, event: "SetEpochData", logs: logs, sub: sub}, nil
}

// WatchSetEpochData is a free log subscription operation binding the contract event 0x5cd723400be8430351b9cbaa5ea421b3fb2528c6a7650c493f895e7d97750da1.
//
// Solidity: event SetEpochData(address indexed sender, uint256 error, uint256 errorInfo)
func (_Ethashsol *EthashsolFilterer) WatchSetEpochData(opts *bind.WatchOpts, sink chan<- *EthashsolSetEpochData, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ethashsol.contract.WatchLogs(opts, "SetEpochData", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthashsolSetEpochData)
				if err := _Ethashsol.contract.UnpackLog(event, "SetEpochData", log); err != nil {
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
func (_Ethashsol *EthashsolFilterer) ParseSetEpochData(log types.Log) (*EthashsolSetEpochData, error) {
	event := new(EthashsolSetEpochData)
	if err := _Ethashsol.contract.UnpackLog(event, "SetEpochData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
