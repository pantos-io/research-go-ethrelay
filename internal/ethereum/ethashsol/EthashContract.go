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
	Bin: "0x608060405234801561001057600080fd5b5061213a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806329e265df14610046578063c7b81f4f14610073578063c891a29d146100a9575b600080fd5b610059610054366004611e98565b6100be565b604080519283526020830191909152015b60405180910390f35b610099610081366004611f2d565b60009081526020819052604090206102000154151590565b604051901515815260200161006a565b6100bc6100b7366004611f5c565b61019b565b005b600080806100ce6175308c612074565b905060006101438b8b8a8a8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808e0282810182019093528d82529093508d92508c9182918501908490808284376000920191909152508992506102f9915050565b905061015189600019612074565b8111156101845760008082600119141561017057506001905082610177565b5060029050815b909450925061018e915050565b6000809350935050505b9850989650505050505050565b60005b818110156102985760008781526020819052604081206101be8386612088565b61020081106101cf576101cf6120a0565b0154111561023d57337f5cd723400be8430351b9cbaa5ea421b3fb2528c6a7650c493f895e7d97750da16001838661020b8c600160801b6120b6565b6102159190612088565b61021f9190612088565b6040805192835260208301919091520160405180910390a2506102f1565b83818151811061024f5761024f6120a0565b60209081029190910181015160008981529182905260409091206102738386612088565b6102008110610284576102846120a0565b015580610290816120d5565b91505061019e565b5060008681526020818152604080832061020081018990556102010187905580518381529182019290925233917f5cd723400be8430351b9cbaa5ea421b3fb2528c6a7650c493f895e7d97750da1910160405180910390a25b505050505050565b6000610303611d94565b61030b611db3565b610313611dd2565b6040805180820182526000878152602081815292812061020181015483526102000154838301819052888252928190529091819061035b576001199650505050505050610593565b6103658c8c61059c565b9550600091505b601082101561039e57602082028681018051918701918252516102009091015281610396816120d5565b92505061036c565b600091505b60408210156104f85760208084015187516000926103e49186189089906103ca90886120f0565b602081106103da576103da6120a0565b6020020151610638565b6103ee91906120f0565b90506103fa8982610658565b61040e82858e8e89600060200201516106bd565b1461042457600019975050505050505050610593565b600091505b60088210156104e557602060808481028d0182810180519386028a01805163ffffffff9586166301000193918202188616825260408401805161010084018051918916918402919091188816905260608501805161020085018051918a169185029190911889169052949095018051610300909301805193881693909202929092189095169094528051600160201b908190049091528251819004909252805182900490528151049052816104dd816120d5565b925050610429565b50816104f0816120d5565b9250506103a3565b600091505b602082101561057d5761054a61053e610532878560208110610521576105216120a0565b6020020151886103ca876001612088565b876103ca866002612088565b866103ca856003612088565b84610556600485612074565b60088110610566576105666120a0565b6020020152610576600483612088565b91506104fd565b60006105898786610899565b9750505050505050505b95945050505050565b6105a4611d94565b6105ac611df1565b6105b584610b77565b6001600160401b038116825293506105d1600160401b85612074565b6001600160401b038116602083015293506105f0600160401b85612074565b6001600160401b0381166040830152935061060f600160401b85612074565b6001600160401b038116606083015260808201849052935061063081610bc9565b949350505050565b6000816106498463010001936120b6565b1863ffffffff16905092915050565b600082815260208190526040812061020181015483901c90829061067d600284612074565b610200811061068e5761068e6120a0565b0154905061069d6002836120f0565b6106af576001600160801b0316610630565b610593600160801b82612074565b6000806106ca8587610ec9565b6001600160801b031690506000808080806106e66002896120f0565b60029098049711905060006106fb888c6120b6565b905081156107105761070d8b82612088565b90505b60005b8881101561080b5760208282010260208b01015193508c60011660001415610749578695506001600160801b0384169450610759565b6001600160801b03841695508694505b604080516020810188905290810186905260600160408051601f1981840301815291905280516020909101206002909d049c6001600160801b0316965060018d166107af57869550600160801b840494506107bc565b600160801b840495508694505b604080516020810188905290810186905260600160408051601f1981840301815291905280516020909101206002909d049c6001600160801b0316965080610803816120d5565b915050610713565b82156108885760208282010260208b01015193508c6001166000141561083f578695506001600160801b038416945061084f565b6001600160801b03841695508694505b60408051602081018890529081018690526060016040516020818303038152906040528051906020012060001c6001600160801b031696505b50949b9a5050505050505050505050565b60008083600760200201516108b290600160601b6120b6565b60c08501516108c590600160401b6120b6565b60a08601516108d890600160201b6120b6565b60808701516108e79190612088565b6108f19190612088565b6108fb9190612088565b61090990600160801b6120b6565b606085015161091c90600160601b6120b6565b604086015161092f90600160401b6120b6565b602087015161094290600160201b6120b6565b875161094e9190612088565b6109589190612088565b6109629190612088565b61096c9190612088565b9050600084600f602002015161098690600160601b6120b6565b6101c086015161099a90600160401b6120b6565b6101a08701516109ae90600160201b6120b6565b6101808801516109be9190612088565b6109c89190612088565b6109d29190612088565b6109e090600160801b6120b6565b6101608601516109f490600160601b6120b6565b610140870151610a0890600160401b6120b6565b610120880151610a1c90600160201b6120b6565b610100890151610a2c9190612088565b610a369190612088565b610a409190612088565b610a4a9190612088565b905060008460076020020151610a6490600160601b6120b6565b60c0860151610a7790600160401b6120b6565b60a0870151610a8a90600160201b6120b6565b6080880151610a999190612088565b610aa39190612088565b610aad9190612088565b610abb90600160801b6120b6565b6060860151610ace90600160601b6120b6565b6040870151610ae190600160401b6120b6565b6020880151610af490600160201b6120b6565b8851610b009190612088565b610b0a9190612088565b610b149190612088565b610b1e9190612088565b9050610b2983610fca565b610b3283610fca565b610b3b83610fca565b604080516020810194909452830191909152606082015260800160408051601f1981840301815291905280516020909101209695505050505050565b600080805b6020811015610bc257610b91826101006120b6565b9150610ba060ff851683612088565b9150610bae61010085612074565b935080610bba816120d5565b915050610b7c565b5092915050565b610bd1611d94565b610bdd600960086120b6565b604814610c1f5760405162461bcd60e51b815260206004820152600c60248201526b39b837b733b29032b93937b960a11b604482015260640160405180910390fd5b600160a08301526001603f1b610100830152604860086000610c426009836120b6565b9050610c4c611e10565b600080805b610c5b8786612074565b831015610d7e57600091505b6005821015610d61575060005b6005811015610d4f57610c878688612074565b610c928360056120b6565b610c9c9083612088565b1015610d3d5788610cae8360056120b6565b82610cba8660096120b6565b610cc49190612088565b610cce9190612088565b60098110610cde57610cde6120a0565b60200201518483610cf08460056120b6565b610cfa9190612088565b60198110610d0a57610d0a6120a0565b6020020151188483610d1d8460056120b6565b610d279190612088565b60198110610d3757610d376120a0565b60200201525b80610d47816120d5565b915050610c74565b81610d59816120d5565b925050610c67565b610d6a84611015565b935082610d76816120d5565b935050610c51565b610d86611d94565b60005b6010811015610ebb57600093505b6005841015610eb657600092505b6005831015610ea457610db8888a612074565b610dc38560056120b6565b610dcd9085612088565b108015610dda5750601081105b15610e92578584610dec8560056120b6565b610df69190612088565b60198110610e0657610e066120a0565b602002015163ffffffff16828260108110610e2357610e236120a0565b6020020152600160201b8685610e3a8660056120b6565b610e449190612088565b60198110610e5457610e546120a0565b6020020151610e639190612074565b82610e6f836001612088565b60108110610e7f57610e7f6120a0565b6020020152610e8f600282612088565b90505b82610e9c816120d5565b935050610da5565b83610eae816120d5565b945050610d97565b610d89565b509998505050505050505050565b600082610ed78360046120b6565b81518110610ee757610ee76120a0565b602002602001015183836004610efd91906120b6565b610f08906001612088565b81518110610f1857610f186120a0565b602002602001015184846004610f2e91906120b6565b610f39906002612088565b81518110610f4957610f496120a0565b602002602001015185856004610f5f91906120b6565b610f6a906003612088565b81518110610f7a57610f7a6120a0565b6020026020010151604051602001610fab949392919093845260208401929092526040830152606082015260800190565b60408051601f1981840301815291905280516020909101209392505050565b600080805b6020811015610bc257610fe4826101006120b6565b9150610ff360ff851683612088565b915061100161010085612074565b93508061100d816120d5565b915050610fcf565b61101d611e10565b611025611e10565b61102d611e2f565b611035611e2f565b600060405180610300016040528060018152602001618082815260200167800000000000808a8152602001678000000080008000815260200161808b81526020016380000001815260200167800000008000808181526020016780000000000080098152602001608a81526020016088815260200163800080098152602001638000000a8152602001638000808b815260200167800000000000008b8152602001678000000000008089815260200167800000000000800381526020016780000000000080028152602001678000000000000080815260200161800a815260200167800000008000000a81526020016780000000800080818152602001678000000000008080815260200163800000018152602001678000000080008008815250905060005b6018811015611d89576080808801516060808a01516040808c01516020808e01518e511890911890921890931888526101208b01516101008c015160e08d015160c08e015160a08f0151181818189088018190526101c08b01516101a08c01516101808d01516101608e01516101408f015118181818928801929092526102608a01516102408b01516102208c01516102008d01516101e08e015118181818908701526103008901516102e08a01516102c08b01516102a08c01516102808d01511818181891860191909152611236906001603f1b90612074565b60208501516112469060026120b6565b60808601516001600160401b0391909116919091171883526040840151611272906001603f1b90612074565b60408501516112829060026120b6565b85516001600160401b03919091169190911718602084015260608401516112ae906001603f1b90612074565b60608501516112be9060026120b6565b6001600160401b03161784600160200201511883600260200201526001603f1b84600460200201516112f09190612074565b60808501516113009060026120b6565b60408601516001600160401b039190911691909117186060840152835161132c906001603f1b90612074565b84516113399060026120b6565b6060808701516001600160401b03929092169290921718608080860191825285518a5118808b52865160208c81018051909218825288516040808f01805190921890915289518e8801805190911890528951948e0180519095189094528801805160a08e0180519091189052805160c08e0180519091189052805160e08e018051909118905280516101008e0180519091189052516101208d018051909118905291870180516101408d018051909118905280516101608d018051909118905280516101808d018051909118905280516101a08d0180519091189052516101c08c018051909118905292860180516101e08c018051909118905280516102008c018051909118905280516102208c018051909118905280516102408c0180519091189052516102608b018051909118905281516102808b018051909118905281516102a08b018051909118905281516102c08b018051909118905281516102e08b018051909118905290516103008a0180519091189052908652516114c390631000000090612074565b60208801516114d7906410000000006120b6565b6001600160401b03161761010086015260408701516114fb906001603d1b90612074565b604088015161150b9060086120b6565b6001600160401b031617610160860152606087015161152e906280000090612074565b606088015161154390650200000000006120b6565b6001600160401b0316176102608601526080870151611569906540000000000090612074565b608088015161157b90620400006120b6565b6001600160401b0316176102c086015260a087015161159f906001603f1b90612074565b60a08801516115af9060026120b6565b6001600160401b031617604086015260c08701516115d1906210000090612074565b60c08801516115e690651000000000006120b6565b6001600160401b03161760a086015260e087015161160c90664000000000000090612074565b60e088015161161d906104006120b6565b6001600160401b0316176101a0860152610100870151611641906208000090612074565b61010088015161165790652000000000006120b6565b6001600160401b03161761020086015261012087015161167c906001603e1b90612074565b61012088015161168d9060046120b6565b6001600160401b0316176103008601526101408701516116af90600490612074565b6101408801516116c3906001603e1b6120b6565b6001600160401b03161760808601526101608701516116eb9067040000000000000090612074565b6101608801516116fc9060406120b6565b6001600160401b03161760e086015261018087015161171f906220000090612074565b61018088015161173590650800000000006120b6565b6001600160401b0316176101408601526101a087015161175d90660200000000000090612074565b6101a088015161176f906180006120b6565b6001600160401b0316176102408601526101c087015161179190600890612074565b6101c08801516117a5906001603d1b6120b6565b6001600160401b0316176102a08601526101e08701516117cb9064100000000090612074565b6101e08801516117df9063100000006120b6565b6001600160401b0316176020860152610200878101516117ff9190612074565b6102008801516118169066800000000000006120b6565b6001600160401b03161761012086015261022087015161183c9064800000000090612074565b6102208801516118509063020000006120b6565b6001600160401b031617610180860152610240870151611877906508000000000090612074565b61024088015161188a90622000006120b6565b6001600160401b0316176101e08601526102608701516118ad9061010090612074565b6102608801516118c190600160381b6120b6565b6001600160401b0316176102e08601526102808701516118e79064200000000090612074565b6102808801516118fb9063080000006120b6565b6001600160401b03161760608601526102a0870151611921906510000000000090612074565b6102a088015161193490621000006120b6565b6001600160401b03161760c08601526102c087015161195890630200000090612074565b6102c088015161196d906480000000006120b6565b6001600160401b0316176101c08601526102e087015161199290600160381b90612074565b6102e08801516119a4906101006120b6565b6001600160401b0316176102208601526103008701516119cc90660400000000000090612074565b6103008801516119de906140006120b6565b6001600160401b031617856014602002015284600a602002015185600560200201511916856000602002015118876000602002015284600b602002015185600660200201511916856001602002015118876001602002015284600c602002015185600760200201511916856002602002015118876002602002015284600d602002015185600860200201511916856003602002015118876003602002015284600e602002015185600960200201511916856004602002015118876004602002015284600f602002015185600a602002015119168560056020020151188760056020020152846010602002015185600b602002015119168560066020020151188760066020020152846011602002015185600c602002015119168560076020020151188760076020020152846012602002015185600d602002015119168560086020020151188760086020020152846013602002015185600e602002015119168560096020020151188760096020020152846014602002015185600f6020020151191685600a60200201511887600a602002015284601560200201518560106020020151191685600b60200201511887600b602002015284601660200201518560116020020151191685600c60200201511887600c602002015284601760200201518560126020020151191685600d60200201511887600d602002015284601860200201518560136020020151191685600e60200201511887600e602002015284600060200201518560146020020151191685600f60200201511887600f6020020152846001602002015185601560200201511916856010602002015118876010602002015284600260200201518560166020020151191685601160200201511887601160200201528460036020020151856017602002015119168560126020020151188760126020020152846004602002015185601860200201511916856013602002015118876013602002015284600560200201518560006020020151191685601460200201511887601460200201528460066020020151856001602002015119168560156020020151188760156020020152846007602002015185600260200201511916856016602002015118876016602002015284600860200201518560036020020151191685601760200201511887601760200201528460096020020151856004602002015119168560186020020151188760186020020152818160188110611d6d57611d6d6120a0565b6020020151875118875280611d81816120d5565b91505061115b565b509495945050505050565b6040518061020001604052806010906020820280368337509192915050565b6040518061040001604052806020906020820280368337509192915050565b6040518061010001604052806008906020820280368337509192915050565b6040518061012001604052806009906020820280368337509192915050565b6040518061032001604052806019906020820280368337509192915050565b6040518060a001604052806005906020820280368337509192915050565b60008083601f840112611e5f57600080fd5b5081356001600160401b03811115611e7657600080fd5b6020830191508360208260051b8501011115611e9157600080fd5b9250929050565b60008060008060008060008060c0898b031215611eb457600080fd5b8835975060208901359650604089013595506060890135945060808901356001600160401b0380821115611ee757600080fd5b611ef38c838d01611e4d565b909650945060a08b0135915080821115611f0c57600080fd5b50611f198b828c01611e4d565b999c989b5096995094979396929594505050565b600060208284031215611f3f57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b60008060008060008060c08789031215611f7557600080fd5b8635955060208088013595506040880135945060608801356001600160401b0380821115611fa257600080fd5b818a0191508a601f830112611fb657600080fd5b813581811115611fc857611fc8611f46565b8060051b604051601f19603f83011681018181108582111715611fed57611fed611f46565b60405291825284820192508381018501918d83111561200b57600080fd5b938501935b8285101561202957843584529385019392850192612010565b9a9d999c50979a60808101359960a09091013598509650505050505050565b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60008261208357612083612048565b500490565b6000821982111561209b5761209b61205e565b500190565b634e487b7160e01b600052603260045260246000fd5b60008160001904831182151516156120d0576120d061205e565b500290565b60006000198214156120e9576120e961205e565b5060010190565b6000826120ff576120ff612048565b50069056fea2646970667358221220ff7e498ea5a5d7259ea52d40a87009ddfcebafbc9dcd73429fc05251f26231ee64736f6c634300080b0033",
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
