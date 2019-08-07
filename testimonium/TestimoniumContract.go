// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testimonium

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

// TestimoniumABI is the input ABI used to generate the binding from.
const TestimoniumABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"longestChainEndpoint\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_rlpHeader\",\"type\":\"bytes\"},{\"name\":\"totalDifficulty\",\"type\":\"uint256\"},{\"name\":\"_ethashContractAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"isPoWValid\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"errorCode\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"PoWValidationResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"RemoveBranch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"hashWithoutNonce\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"difficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"parent\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"transactionsRoot\",\"type\":\"bytes32\"}],\"name\":\"SubmitBlockHeader\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"getHeader\",\"outputs\":[{\"name\":\"parent\",\"type\":\"bytes32\"},{\"name\":\"uncleHash\",\"type\":\"bytes32\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"rlpHeaderHashWithoutNonce\",\"type\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"difficulty\",\"type\":\"uint256\"},{\"name\":\"totalDifficulty\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"getHeaderMetaInfo\",\"outputs\":[{\"name\":\"successors\",\"type\":\"bytes32[]\"},{\"name\":\"orderedIndex\",\"type\":\"uint256\"},{\"name\":\"iterableIndex\",\"type\":\"uint256\"},{\"name\":\"latestFork\",\"type\":\"bytes32\"},{\"name\":\"lockedUntil\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNoOfForks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getBlockHashOfEndpoint\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rlpHeader\",\"type\":\"bytes\"}],\"name\":\"submitHeader\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"name\":\"dataSetLookup\",\"type\":\"uint256[]\"},{\"name\":\"witnessForLookup\",\"type\":\"uint256[]\"}],\"name\":\"disputeBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"name\":\"noOfConfirmations\",\"type\":\"uint8\"},{\"name\":\"rlpEncodedTx\",\"type\":\"bytes\"},{\"name\":\"path\",\"type\":\"bytes\"},{\"name\":\"rlpEncodedNodes\",\"type\":\"bytes\"},{\"name\":\"merkleRootHash\",\"type\":\"bytes32\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"name\":\"noOfConfirmations\",\"type\":\"uint8\"},{\"name\":\"rlpEncodedTx\",\"type\":\"bytes\"},{\"name\":\"path\",\"type\":\"bytes\"},{\"name\":\"rlpEncodedNodes\",\"type\":\"bytes\"}],\"name\":\"verifyTransaction\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"name\":\"noOfConfirmations\",\"type\":\"uint8\"},{\"name\":\"rlpEncodedTx\",\"type\":\"bytes\"},{\"name\":\"path\",\"type\":\"bytes\"},{\"name\":\"rlpEncodedNodes\",\"type\":\"bytes\"}],\"name\":\"verifyReceipt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"name\":\"noOfConfirmations\",\"type\":\"uint8\"},{\"name\":\"rlpEncodedTx\",\"type\":\"bytes\"},{\"name\":\"path\",\"type\":\"bytes\"},{\"name\":\"rlpEncodedNodes\",\"type\":\"bytes\"}],\"name\":\"verifyState\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"name\":\"isUnlocked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Testimonium is an auto generated Go binding around an Ethereum contract.
type Testimonium struct {
	TestimoniumCaller     // Read-only binding to the contract
	TestimoniumTransactor // Write-only binding to the contract
	TestimoniumFilterer   // Log filterer for contract events
}

// TestimoniumCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestimoniumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestimoniumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestimoniumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestimoniumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestimoniumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestimoniumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestimoniumSession struct {
	Contract     *Testimonium      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestimoniumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestimoniumCallerSession struct {
	Contract *TestimoniumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TestimoniumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestimoniumTransactorSession struct {
	Contract     *TestimoniumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TestimoniumRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestimoniumRaw struct {
	Contract *Testimonium // Generic contract binding to access the raw methods on
}

// TestimoniumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestimoniumCallerRaw struct {
	Contract *TestimoniumCaller // Generic read-only contract binding to access the raw methods on
}

// TestimoniumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestimoniumTransactorRaw struct {
	Contract *TestimoniumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestimonium creates a new instance of Testimonium, bound to a specific deployed contract.
func NewTestimonium(address common.Address, backend bind.ContractBackend) (*Testimonium, error) {
	contract, err := bindTestimonium(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Testimonium{TestimoniumCaller: TestimoniumCaller{contract: contract}, TestimoniumTransactor: TestimoniumTransactor{contract: contract}, TestimoniumFilterer: TestimoniumFilterer{contract: contract}}, nil
}

// NewTestimoniumCaller creates a new read-only instance of Testimonium, bound to a specific deployed contract.
func NewTestimoniumCaller(address common.Address, caller bind.ContractCaller) (*TestimoniumCaller, error) {
	contract, err := bindTestimonium(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestimoniumCaller{contract: contract}, nil
}

// NewTestimoniumTransactor creates a new write-only instance of Testimonium, bound to a specific deployed contract.
func NewTestimoniumTransactor(address common.Address, transactor bind.ContractTransactor) (*TestimoniumTransactor, error) {
	contract, err := bindTestimonium(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestimoniumTransactor{contract: contract}, nil
}

// NewTestimoniumFilterer creates a new log filterer instance of Testimonium, bound to a specific deployed contract.
func NewTestimoniumFilterer(address common.Address, filterer bind.ContractFilterer) (*TestimoniumFilterer, error) {
	contract, err := bindTestimonium(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestimoniumFilterer{contract: contract}, nil
}

// bindTestimonium binds a generic wrapper to an already deployed contract.
func bindTestimonium(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestimoniumABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testimonium *TestimoniumRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Testimonium.Contract.TestimoniumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testimonium *TestimoniumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testimonium.Contract.TestimoniumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testimonium *TestimoniumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Testimonium.Contract.TestimoniumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testimonium *TestimoniumCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Testimonium.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testimonium *TestimoniumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testimonium.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testimonium *TestimoniumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Testimonium.Contract.contract.Transact(opts, method, params...)
}

// GetBlockHashOfEndpoint is a free data retrieval call binding the contract method 0x84bc44b3.
//
// Solidity: function getBlockHashOfEndpoint(uint256 index) constant returns(bytes32)
func (_Testimonium *TestimoniumCaller) GetBlockHashOfEndpoint(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Testimonium.contract.Call(opts, out, "getBlockHashOfEndpoint", index)
	return *ret0, err
}

// GetBlockHashOfEndpoint is a free data retrieval call binding the contract method 0x84bc44b3.
//
// Solidity: function getBlockHashOfEndpoint(uint256 index) constant returns(bytes32)
func (_Testimonium *TestimoniumSession) GetBlockHashOfEndpoint(index *big.Int) ([32]byte, error) {
	return _Testimonium.Contract.GetBlockHashOfEndpoint(&_Testimonium.CallOpts, index)
}

// GetBlockHashOfEndpoint is a free data retrieval call binding the contract method 0x84bc44b3.
//
// Solidity: function getBlockHashOfEndpoint(uint256 index) constant returns(bytes32)
func (_Testimonium *TestimoniumCallerSession) GetBlockHashOfEndpoint(index *big.Int) ([32]byte, error) {
	return _Testimonium.Contract.GetBlockHashOfEndpoint(&_Testimonium.CallOpts, index)
}

// GetHeader is a free data retrieval call binding the contract method 0xb9615878.
//
// Solidity: function getHeader(bytes32 blockHash) constant returns(bytes32 parent, bytes32 uncleHash, bytes32 stateRoot, bytes32 transactionsRoot, bytes32 receiptsRoot, uint256 blockNumber, bytes32 rlpHeaderHashWithoutNonce, uint256 timestamp, uint256 nonce, uint256 difficulty, uint256 totalDifficulty)
func (_Testimonium *TestimoniumCaller) GetHeader(opts *bind.CallOpts, blockHash [32]byte) (struct {
	Parent                    [32]byte
	UncleHash                 [32]byte
	StateRoot                 [32]byte
	TransactionsRoot          [32]byte
	ReceiptsRoot              [32]byte
	BlockNumber               *big.Int
	RlpHeaderHashWithoutNonce [32]byte
	Timestamp                 *big.Int
	Nonce                     *big.Int
	Difficulty                *big.Int
	TotalDifficulty           *big.Int
}, error) {
	ret := new(struct {
		Parent                    [32]byte
		UncleHash                 [32]byte
		StateRoot                 [32]byte
		TransactionsRoot          [32]byte
		ReceiptsRoot              [32]byte
		BlockNumber               *big.Int
		RlpHeaderHashWithoutNonce [32]byte
		Timestamp                 *big.Int
		Nonce                     *big.Int
		Difficulty                *big.Int
		TotalDifficulty           *big.Int
	})
	out := ret
	err := _Testimonium.contract.Call(opts, out, "getHeader", blockHash)
	return *ret, err
}

// GetHeader is a free data retrieval call binding the contract method 0xb9615878.
//
// Solidity: function getHeader(bytes32 blockHash) constant returns(bytes32 parent, bytes32 uncleHash, bytes32 stateRoot, bytes32 transactionsRoot, bytes32 receiptsRoot, uint256 blockNumber, bytes32 rlpHeaderHashWithoutNonce, uint256 timestamp, uint256 nonce, uint256 difficulty, uint256 totalDifficulty)
func (_Testimonium *TestimoniumSession) GetHeader(blockHash [32]byte) (struct {
	Parent                    [32]byte
	UncleHash                 [32]byte
	StateRoot                 [32]byte
	TransactionsRoot          [32]byte
	ReceiptsRoot              [32]byte
	BlockNumber               *big.Int
	RlpHeaderHashWithoutNonce [32]byte
	Timestamp                 *big.Int
	Nonce                     *big.Int
	Difficulty                *big.Int
	TotalDifficulty           *big.Int
}, error) {
	return _Testimonium.Contract.GetHeader(&_Testimonium.CallOpts, blockHash)
}

// GetHeader is a free data retrieval call binding the contract method 0xb9615878.
//
// Solidity: function getHeader(bytes32 blockHash) constant returns(bytes32 parent, bytes32 uncleHash, bytes32 stateRoot, bytes32 transactionsRoot, bytes32 receiptsRoot, uint256 blockNumber, bytes32 rlpHeaderHashWithoutNonce, uint256 timestamp, uint256 nonce, uint256 difficulty, uint256 totalDifficulty)
func (_Testimonium *TestimoniumCallerSession) GetHeader(blockHash [32]byte) (struct {
	Parent                    [32]byte
	UncleHash                 [32]byte
	StateRoot                 [32]byte
	TransactionsRoot          [32]byte
	ReceiptsRoot              [32]byte
	BlockNumber               *big.Int
	RlpHeaderHashWithoutNonce [32]byte
	Timestamp                 *big.Int
	Nonce                     *big.Int
	Difficulty                *big.Int
	TotalDifficulty           *big.Int
}, error) {
	return _Testimonium.Contract.GetHeader(&_Testimonium.CallOpts, blockHash)
}

// GetHeaderMetaInfo is a free data retrieval call binding the contract method 0x3b93fc10.
//
// Solidity: function getHeaderMetaInfo(bytes32 blockHash) constant returns(bytes32[] successors, uint256 orderedIndex, uint256 iterableIndex, bytes32 latestFork, uint256 lockedUntil)
func (_Testimonium *TestimoniumCaller) GetHeaderMetaInfo(opts *bind.CallOpts, blockHash [32]byte) (struct {
	Successors    [][32]byte
	OrderedIndex  *big.Int
	IterableIndex *big.Int
	LatestFork    [32]byte
	LockedUntil   *big.Int
}, error) {
	ret := new(struct {
		Successors    [][32]byte
		OrderedIndex  *big.Int
		IterableIndex *big.Int
		LatestFork    [32]byte
		LockedUntil   *big.Int
	})
	out := ret
	err := _Testimonium.contract.Call(opts, out, "getHeaderMetaInfo", blockHash)
	return *ret, err
}

// GetHeaderMetaInfo is a free data retrieval call binding the contract method 0x3b93fc10.
//
// Solidity: function getHeaderMetaInfo(bytes32 blockHash) constant returns(bytes32[] successors, uint256 orderedIndex, uint256 iterableIndex, bytes32 latestFork, uint256 lockedUntil)
func (_Testimonium *TestimoniumSession) GetHeaderMetaInfo(blockHash [32]byte) (struct {
	Successors    [][32]byte
	OrderedIndex  *big.Int
	IterableIndex *big.Int
	LatestFork    [32]byte
	LockedUntil   *big.Int
}, error) {
	return _Testimonium.Contract.GetHeaderMetaInfo(&_Testimonium.CallOpts, blockHash)
}

// GetHeaderMetaInfo is a free data retrieval call binding the contract method 0x3b93fc10.
//
// Solidity: function getHeaderMetaInfo(bytes32 blockHash) constant returns(bytes32[] successors, uint256 orderedIndex, uint256 iterableIndex, bytes32 latestFork, uint256 lockedUntil)
func (_Testimonium *TestimoniumCallerSession) GetHeaderMetaInfo(blockHash [32]byte) (struct {
	Successors    [][32]byte
	OrderedIndex  *big.Int
	IterableIndex *big.Int
	LatestFork    [32]byte
	LockedUntil   *big.Int
}, error) {
	return _Testimonium.Contract.GetHeaderMetaInfo(&_Testimonium.CallOpts, blockHash)
}

// GetNoOfForks is a free data retrieval call binding the contract method 0xfbb5df38.
//
// Solidity: function getNoOfForks() constant returns(uint256)
func (_Testimonium *TestimoniumCaller) GetNoOfForks(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Testimonium.contract.Call(opts, out, "getNoOfForks")
	return *ret0, err
}

// GetNoOfForks is a free data retrieval call binding the contract method 0xfbb5df38.
//
// Solidity: function getNoOfForks() constant returns(uint256)
func (_Testimonium *TestimoniumSession) GetNoOfForks() (*big.Int, error) {
	return _Testimonium.Contract.GetNoOfForks(&_Testimonium.CallOpts)
}

// GetNoOfForks is a free data retrieval call binding the contract method 0xfbb5df38.
//
// Solidity: function getNoOfForks() constant returns(uint256)
func (_Testimonium *TestimoniumCallerSession) GetNoOfForks() (*big.Int, error) {
	return _Testimonium.Contract.GetNoOfForks(&_Testimonium.CallOpts)
}

// IsBlock is a free data retrieval call binding the contract method 0x528a309f.
//
// Solidity: function isBlock(bytes32 hash) constant returns(bool)
func (_Testimonium *TestimoniumCaller) IsBlock(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Testimonium.contract.Call(opts, out, "isBlock", hash)
	return *ret0, err
}

// IsBlock is a free data retrieval call binding the contract method 0x528a309f.
//
// Solidity: function isBlock(bytes32 hash) constant returns(bool)
func (_Testimonium *TestimoniumSession) IsBlock(hash [32]byte) (bool, error) {
	return _Testimonium.Contract.IsBlock(&_Testimonium.CallOpts, hash)
}

// IsBlock is a free data retrieval call binding the contract method 0x528a309f.
//
// Solidity: function isBlock(bytes32 hash) constant returns(bool)
func (_Testimonium *TestimoniumCallerSession) IsBlock(hash [32]byte) (bool, error) {
	return _Testimonium.Contract.IsBlock(&_Testimonium.CallOpts, hash)
}

// IsUnlocked is a free data retrieval call binding the contract method 0x3dadc1e1.
//
// Solidity: function isUnlocked(bytes32 blockHash) constant returns(bool)
func (_Testimonium *TestimoniumCaller) IsUnlocked(opts *bind.CallOpts, blockHash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Testimonium.contract.Call(opts, out, "isUnlocked", blockHash)
	return *ret0, err
}

// IsUnlocked is a free data retrieval call binding the contract method 0x3dadc1e1.
//
// Solidity: function isUnlocked(bytes32 blockHash) constant returns(bool)
func (_Testimonium *TestimoniumSession) IsUnlocked(blockHash [32]byte) (bool, error) {
	return _Testimonium.Contract.IsUnlocked(&_Testimonium.CallOpts, blockHash)
}

// IsUnlocked is a free data retrieval call binding the contract method 0x3dadc1e1.
//
// Solidity: function isUnlocked(bytes32 blockHash) constant returns(bool)
func (_Testimonium *TestimoniumCallerSession) IsUnlocked(blockHash [32]byte) (bool, error) {
	return _Testimonium.Contract.IsUnlocked(&_Testimonium.CallOpts, blockHash)
}

// LongestChainEndpoint is a free data retrieval call binding the contract method 0x0aa7fc0f.
//
// Solidity: function longestChainEndpoint() constant returns(bytes32)
func (_Testimonium *TestimoniumCaller) LongestChainEndpoint(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Testimonium.contract.Call(opts, out, "longestChainEndpoint")
	return *ret0, err
}

// LongestChainEndpoint is a free data retrieval call binding the contract method 0x0aa7fc0f.
//
// Solidity: function longestChainEndpoint() constant returns(bytes32)
func (_Testimonium *TestimoniumSession) LongestChainEndpoint() ([32]byte, error) {
	return _Testimonium.Contract.LongestChainEndpoint(&_Testimonium.CallOpts)
}

// LongestChainEndpoint is a free data retrieval call binding the contract method 0x0aa7fc0f.
//
// Solidity: function longestChainEndpoint() constant returns(bytes32)
func (_Testimonium *TestimoniumCallerSession) LongestChainEndpoint() ([32]byte, error) {
	return _Testimonium.Contract.LongestChainEndpoint(&_Testimonium.CallOpts)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0x843b69e9.
//
// Solidity: function verifyMerkleProof(bytes32 blockHash, uint8 noOfConfirmations, bytes rlpEncodedTx, bytes path, bytes rlpEncodedNodes, bytes32 merkleRootHash) constant returns(uint8)
func (_Testimonium *TestimoniumCaller) VerifyMerkleProof(opts *bind.CallOpts, blockHash [32]byte, noOfConfirmations uint8, rlpEncodedTx []byte, path []byte, rlpEncodedNodes []byte, merkleRootHash [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Testimonium.contract.Call(opts, out, "verifyMerkleProof", blockHash, noOfConfirmations, rlpEncodedTx, path, rlpEncodedNodes, merkleRootHash)
	return *ret0, err
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0x843b69e9.
//
// Solidity: function verifyMerkleProof(bytes32 blockHash, uint8 noOfConfirmations, bytes rlpEncodedTx, bytes path, bytes rlpEncodedNodes, bytes32 merkleRootHash) constant returns(uint8)
func (_Testimonium *TestimoniumSession) VerifyMerkleProof(blockHash [32]byte, noOfConfirmations uint8, rlpEncodedTx []byte, path []byte, rlpEncodedNodes []byte, merkleRootHash [32]byte) (uint8, error) {
	return _Testimonium.Contract.VerifyMerkleProof(&_Testimonium.CallOpts, blockHash, noOfConfirmations, rlpEncodedTx, path, rlpEncodedNodes, merkleRootHash)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0x843b69e9.
//
// Solidity: function verifyMerkleProof(bytes32 blockHash, uint8 noOfConfirmations, bytes rlpEncodedTx, bytes path, bytes rlpEncodedNodes, bytes32 merkleRootHash) constant returns(uint8)
func (_Testimonium *TestimoniumCallerSession) VerifyMerkleProof(blockHash [32]byte, noOfConfirmations uint8, rlpEncodedTx []byte, path []byte, rlpEncodedNodes []byte, merkleRootHash [32]byte) (uint8, error) {
	return _Testimonium.Contract.VerifyMerkleProof(&_Testimonium.CallOpts, blockHash, noOfConfirmations, rlpEncodedTx, path, rlpEncodedNodes, merkleRootHash)
}

// VerifyReceipt is a free data retrieval call binding the contract method 0x10fa6e83.
//
// Solidity: function verifyReceipt(bytes32 blockHash, uint8 noOfConfirmations, bytes rlpEncodedTx, bytes path, bytes rlpEncodedNodes) constant returns(uint8)
func (_Testimonium *TestimoniumCaller) VerifyReceipt(opts *bind.CallOpts, blockHash [32]byte, noOfConfirmations uint8, rlpEncodedTx []byte, path []byte, rlpEncodedNodes []byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Testimonium.contract.Call(opts, out, "verifyReceipt", blockHash, noOfConfirmations, rlpEncodedTx, path, rlpEncodedNodes)
	return *ret0, err
}

// VerifyReceipt is a free data retrieval call binding the contract method 0x10fa6e83.
//
// Solidity: function verifyReceipt(bytes32 blockHash, uint8 noOfConfirmations, bytes rlpEncodedTx, bytes path, bytes rlpEncodedNodes) constant returns(uint8)
func (_Testimonium *TestimoniumSession) VerifyReceipt(blockHash [32]byte, noOfConfirmations uint8, rlpEncodedTx []byte, path []byte, rlpEncodedNodes []byte) (uint8, error) {
	return _Testimonium.Contract.VerifyReceipt(&_Testimonium.CallOpts, blockHash, noOfConfirmations, rlpEncodedTx, path, rlpEncodedNodes)
}

// VerifyReceipt is a free data retrieval call binding the contract method 0x10fa6e83.
//
// Solidity: function verifyReceipt(bytes32 blockHash, uint8 noOfConfirmations, bytes rlpEncodedTx, bytes path, bytes rlpEncodedNodes) constant returns(uint8)
func (_Testimonium *TestimoniumCallerSession) VerifyReceipt(blockHash [32]byte, noOfConfirmations uint8, rlpEncodedTx []byte, path []byte, rlpEncodedNodes []byte) (uint8, error) {
	return _Testimonium.Contract.VerifyReceipt(&_Testimonium.CallOpts, blockHash, noOfConfirmations, rlpEncodedTx, path, rlpEncodedNodes)
}

// VerifyState is a free data retrieval call binding the contract method 0xfadfdb81.
//
// Solidity: function verifyState(bytes32 blockHash, uint8 noOfConfirmations, bytes rlpEncodedTx, bytes path, bytes rlpEncodedNodes) constant returns(uint8)
func (_Testimonium *TestimoniumCaller) VerifyState(opts *bind.CallOpts, blockHash [32]byte, noOfConfirmations uint8, rlpEncodedTx []byte, path []byte, rlpEncodedNodes []byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Testimonium.contract.Call(opts, out, "verifyState", blockHash, noOfConfirmations, rlpEncodedTx, path, rlpEncodedNodes)
	return *ret0, err
}

// VerifyState is a free data retrieval call binding the contract method 0xfadfdb81.
//
// Solidity: function verifyState(bytes32 blockHash, uint8 noOfConfirmations, bytes rlpEncodedTx, bytes path, bytes rlpEncodedNodes) constant returns(uint8)
func (_Testimonium *TestimoniumSession) VerifyState(blockHash [32]byte, noOfConfirmations uint8, rlpEncodedTx []byte, path []byte, rlpEncodedNodes []byte) (uint8, error) {
	return _Testimonium.Contract.VerifyState(&_Testimonium.CallOpts, blockHash, noOfConfirmations, rlpEncodedTx, path, rlpEncodedNodes)
}

// VerifyState is a free data retrieval call binding the contract method 0xfadfdb81.
//
// Solidity: function verifyState(bytes32 blockHash, uint8 noOfConfirmations, bytes rlpEncodedTx, bytes path, bytes rlpEncodedNodes) constant returns(uint8)
func (_Testimonium *TestimoniumCallerSession) VerifyState(blockHash [32]byte, noOfConfirmations uint8, rlpEncodedTx []byte, path []byte, rlpEncodedNodes []byte) (uint8, error) {
	return _Testimonium.Contract.VerifyState(&_Testimonium.CallOpts, blockHash, noOfConfirmations, rlpEncodedTx, path, rlpEncodedNodes)
}

// VerifyTransaction is a free data retrieval call binding the contract method 0xa90a51f4.
//
// Solidity: function verifyTransaction(bytes32 blockHash, uint8 noOfConfirmations, bytes rlpEncodedTx, bytes path, bytes rlpEncodedNodes) constant returns(uint8)
func (_Testimonium *TestimoniumCaller) VerifyTransaction(opts *bind.CallOpts, blockHash [32]byte, noOfConfirmations uint8, rlpEncodedTx []byte, path []byte, rlpEncodedNodes []byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Testimonium.contract.Call(opts, out, "verifyTransaction", blockHash, noOfConfirmations, rlpEncodedTx, path, rlpEncodedNodes)
	return *ret0, err
}

// VerifyTransaction is a free data retrieval call binding the contract method 0xa90a51f4.
//
// Solidity: function verifyTransaction(bytes32 blockHash, uint8 noOfConfirmations, bytes rlpEncodedTx, bytes path, bytes rlpEncodedNodes) constant returns(uint8)
func (_Testimonium *TestimoniumSession) VerifyTransaction(blockHash [32]byte, noOfConfirmations uint8, rlpEncodedTx []byte, path []byte, rlpEncodedNodes []byte) (uint8, error) {
	return _Testimonium.Contract.VerifyTransaction(&_Testimonium.CallOpts, blockHash, noOfConfirmations, rlpEncodedTx, path, rlpEncodedNodes)
}

// VerifyTransaction is a free data retrieval call binding the contract method 0xa90a51f4.
//
// Solidity: function verifyTransaction(bytes32 blockHash, uint8 noOfConfirmations, bytes rlpEncodedTx, bytes path, bytes rlpEncodedNodes) constant returns(uint8)
func (_Testimonium *TestimoniumCallerSession) VerifyTransaction(blockHash [32]byte, noOfConfirmations uint8, rlpEncodedTx []byte, path []byte, rlpEncodedNodes []byte) (uint8, error) {
	return _Testimonium.Contract.VerifyTransaction(&_Testimonium.CallOpts, blockHash, noOfConfirmations, rlpEncodedTx, path, rlpEncodedNodes)
}

// DisputeBlock is a paid mutator transaction binding the contract method 0x66962b26.
//
// Solidity: function disputeBlock(bytes32 blockHash, uint256[] dataSetLookup, uint256[] witnessForLookup) returns()
func (_Testimonium *TestimoniumTransactor) DisputeBlock(opts *bind.TransactOpts, blockHash [32]byte, dataSetLookup []*big.Int, witnessForLookup []*big.Int) (*types.Transaction, error) {
	return _Testimonium.contract.Transact(opts, "disputeBlock", blockHash, dataSetLookup, witnessForLookup)
}

// DisputeBlock is a paid mutator transaction binding the contract method 0x66962b26.
//
// Solidity: function disputeBlock(bytes32 blockHash, uint256[] dataSetLookup, uint256[] witnessForLookup) returns()
func (_Testimonium *TestimoniumSession) DisputeBlock(blockHash [32]byte, dataSetLookup []*big.Int, witnessForLookup []*big.Int) (*types.Transaction, error) {
	return _Testimonium.Contract.DisputeBlock(&_Testimonium.TransactOpts, blockHash, dataSetLookup, witnessForLookup)
}

// DisputeBlock is a paid mutator transaction binding the contract method 0x66962b26.
//
// Solidity: function disputeBlock(bytes32 blockHash, uint256[] dataSetLookup, uint256[] witnessForLookup) returns()
func (_Testimonium *TestimoniumTransactorSession) DisputeBlock(blockHash [32]byte, dataSetLookup []*big.Int, witnessForLookup []*big.Int) (*types.Transaction, error) {
	return _Testimonium.Contract.DisputeBlock(&_Testimonium.TransactOpts, blockHash, dataSetLookup, witnessForLookup)
}

// SubmitHeader is a paid mutator transaction binding the contract method 0xc565ba10.
//
// Solidity: function submitHeader(bytes _rlpHeader) returns()
func (_Testimonium *TestimoniumTransactor) SubmitHeader(opts *bind.TransactOpts, _rlpHeader []byte) (*types.Transaction, error) {
	return _Testimonium.contract.Transact(opts, "submitHeader", _rlpHeader)
}

// SubmitHeader is a paid mutator transaction binding the contract method 0xc565ba10.
//
// Solidity: function submitHeader(bytes _rlpHeader) returns()
func (_Testimonium *TestimoniumSession) SubmitHeader(_rlpHeader []byte) (*types.Transaction, error) {
	return _Testimonium.Contract.SubmitHeader(&_Testimonium.TransactOpts, _rlpHeader)
}

// SubmitHeader is a paid mutator transaction binding the contract method 0xc565ba10.
//
// Solidity: function submitHeader(bytes _rlpHeader) returns()
func (_Testimonium *TestimoniumTransactorSession) SubmitHeader(_rlpHeader []byte) (*types.Transaction, error) {
	return _Testimonium.Contract.SubmitHeader(&_Testimonium.TransactOpts, _rlpHeader)
}

// TestimoniumPoWValidationResultIterator is returned from FilterPoWValidationResult and is used to iterate over the raw logs and unpacked data for PoWValidationResult events raised by the Testimonium contract.
type TestimoniumPoWValidationResultIterator struct {
	Event *TestimoniumPoWValidationResult // Event containing the contract specifics and raw log

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
func (it *TestimoniumPoWValidationResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestimoniumPoWValidationResult)
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
		it.Event = new(TestimoniumPoWValidationResult)
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
func (it *TestimoniumPoWValidationResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestimoniumPoWValidationResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestimoniumPoWValidationResult represents a PoWValidationResult event raised by the Testimonium contract.
type TestimoniumPoWValidationResult struct {
	IsPoWValid bool
	ErrorCode  *big.Int
	ErrorInfo  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoWValidationResult is a free log retrieval operation binding the contract event 0x808f84bd298b89733055f3b8294b385ed860b67438b792ff9e312c897a2fcd9f.
//
// Solidity: event PoWValidationResult(bool isPoWValid, uint256 errorCode, uint256 errorInfo)
func (_Testimonium *TestimoniumFilterer) FilterPoWValidationResult(opts *bind.FilterOpts) (*TestimoniumPoWValidationResultIterator, error) {

	logs, sub, err := _Testimonium.contract.FilterLogs(opts, "PoWValidationResult")
	if err != nil {
		return nil, err
	}
	return &TestimoniumPoWValidationResultIterator{contract: _Testimonium.contract, event: "PoWValidationResult", logs: logs, sub: sub}, nil
}

// WatchPoWValidationResult is a free log subscription operation binding the contract event 0x808f84bd298b89733055f3b8294b385ed860b67438b792ff9e312c897a2fcd9f.
//
// Solidity: event PoWValidationResult(bool isPoWValid, uint256 errorCode, uint256 errorInfo)
func (_Testimonium *TestimoniumFilterer) WatchPoWValidationResult(opts *bind.WatchOpts, sink chan<- *TestimoniumPoWValidationResult) (event.Subscription, error) {

	logs, sub, err := _Testimonium.contract.WatchLogs(opts, "PoWValidationResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestimoniumPoWValidationResult)
				if err := _Testimonium.contract.UnpackLog(event, "PoWValidationResult", log); err != nil {
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

// TestimoniumRemoveBranchIterator is returned from FilterRemoveBranch and is used to iterate over the raw logs and unpacked data for RemoveBranch events raised by the Testimonium contract.
type TestimoniumRemoveBranchIterator struct {
	Event *TestimoniumRemoveBranch // Event containing the contract specifics and raw log

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
func (it *TestimoniumRemoveBranchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestimoniumRemoveBranch)
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
		it.Event = new(TestimoniumRemoveBranch)
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
func (it *TestimoniumRemoveBranchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestimoniumRemoveBranchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestimoniumRemoveBranch represents a RemoveBranch event raised by the Testimonium contract.
type TestimoniumRemoveBranch struct {
	Root [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemoveBranch is a free log retrieval operation binding the contract event 0xf96ae1a1e71431cfb86761b9cab725aeddab2afdaf76d40d43fc005bdc6555d4.
//
// Solidity: event RemoveBranch(bytes32 root)
func (_Testimonium *TestimoniumFilterer) FilterRemoveBranch(opts *bind.FilterOpts) (*TestimoniumRemoveBranchIterator, error) {

	logs, sub, err := _Testimonium.contract.FilterLogs(opts, "RemoveBranch")
	if err != nil {
		return nil, err
	}
	return &TestimoniumRemoveBranchIterator{contract: _Testimonium.contract, event: "RemoveBranch", logs: logs, sub: sub}, nil
}

// WatchRemoveBranch is a free log subscription operation binding the contract event 0xf96ae1a1e71431cfb86761b9cab725aeddab2afdaf76d40d43fc005bdc6555d4.
//
// Solidity: event RemoveBranch(bytes32 root)
func (_Testimonium *TestimoniumFilterer) WatchRemoveBranch(opts *bind.WatchOpts, sink chan<- *TestimoniumRemoveBranch) (event.Subscription, error) {

	logs, sub, err := _Testimonium.contract.WatchLogs(opts, "RemoveBranch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestimoniumRemoveBranch)
				if err := _Testimonium.contract.UnpackLog(event, "RemoveBranch", log); err != nil {
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

// TestimoniumSubmitBlockHeaderIterator is returned from FilterSubmitBlockHeader and is used to iterate over the raw logs and unpacked data for SubmitBlockHeader events raised by the Testimonium contract.
type TestimoniumSubmitBlockHeaderIterator struct {
	Event *TestimoniumSubmitBlockHeader // Event containing the contract specifics and raw log

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
func (it *TestimoniumSubmitBlockHeaderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestimoniumSubmitBlockHeader)
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
		it.Event = new(TestimoniumSubmitBlockHeader)
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
func (it *TestimoniumSubmitBlockHeaderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestimoniumSubmitBlockHeaderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestimoniumSubmitBlockHeader represents a SubmitBlockHeader event raised by the Testimonium contract.
type TestimoniumSubmitBlockHeader struct {
	Hash             [32]byte
	HashWithoutNonce [32]byte
	Nonce            *big.Int
	Difficulty       *big.Int
	Parent           [32]byte
	TransactionsRoot [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSubmitBlockHeader is a free log retrieval operation binding the contract event 0x56a7be9848412b60f43fa0c4d4e9bb1d1dec16601d6a291428be09dd76b75951.
//
// Solidity: event SubmitBlockHeader(bytes32 hash, bytes32 hashWithoutNonce, uint256 nonce, uint256 difficulty, bytes32 parent, bytes32 transactionsRoot)
func (_Testimonium *TestimoniumFilterer) FilterSubmitBlockHeader(opts *bind.FilterOpts) (*TestimoniumSubmitBlockHeaderIterator, error) {

	logs, sub, err := _Testimonium.contract.FilterLogs(opts, "SubmitBlockHeader")
	if err != nil {
		return nil, err
	}
	return &TestimoniumSubmitBlockHeaderIterator{contract: _Testimonium.contract, event: "SubmitBlockHeader", logs: logs, sub: sub}, nil
}

// WatchSubmitBlockHeader is a free log subscription operation binding the contract event 0x56a7be9848412b60f43fa0c4d4e9bb1d1dec16601d6a291428be09dd76b75951.
//
// Solidity: event SubmitBlockHeader(bytes32 hash, bytes32 hashWithoutNonce, uint256 nonce, uint256 difficulty, bytes32 parent, bytes32 transactionsRoot)
func (_Testimonium *TestimoniumFilterer) WatchSubmitBlockHeader(opts *bind.WatchOpts, sink chan<- *TestimoniumSubmitBlockHeader) (event.Subscription, error) {

	logs, sub, err := _Testimonium.contract.WatchLogs(opts, "SubmitBlockHeader")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestimoniumSubmitBlockHeader)
				if err := _Testimonium.contract.UnpackLog(event, "SubmitBlockHeader", log); err != nil {
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
