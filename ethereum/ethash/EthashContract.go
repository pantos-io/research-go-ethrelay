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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EthashABI is the input ABI used to generate the binding from.
const EthashABI = "[{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"SetEpochData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"errorCode\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"VerifyPoW\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"epochIndex\",\"type\":\"uint256\"}],\"name\":\"isEpochDataSet\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\"},{\"name\":\"fullSizeIn128Resultion\",\"type\":\"uint256\"},{\"name\":\"branchDepth\",\"type\":\"uint256\"},{\"name\":\"merkleNodes\",\"type\":\"uint256[]\"},{\"name\":\"start\",\"type\":\"uint256\"},{\"name\":\"numElems\",\"type\":\"uint256\"}],\"name\":\"setEpochData\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes32\"},{\"name\":\"nonceLe\",\"type\":\"bytes8\"},{\"name\":\"dataSetLookup\",\"type\":\"uint256[]\"},{\"name\":\"witnessForLookup\",\"type\":\"uint256[]\"},{\"name\":\"epochIndex\",\"type\":\"uint256\"}],\"name\":\"hashimoto\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"rlpHeaderHashWithoutNonce\",\"type\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"difficulty\",\"type\":\"uint256\"},{\"name\":\"dataSetLookup\",\"type\":\"uint256[]\"},{\"name\":\"witnessForLookup\",\"type\":\"uint256[]\"}],\"name\":\"verifyPoW\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"rlpHeaderHashWithoutNonce\",\"type\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"difficulty\",\"type\":\"uint256\"},{\"name\":\"dataSetLookup\",\"type\":\"uint256[]\"},{\"name\":\"witnessForLookup\",\"type\":\"uint256[]\"}],\"name\":\"test\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"}]"

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

// Hashimoto is a free data retrieval call binding the contract method 0x40ed79f4.
//
// Solidity: function hashimoto(bytes32 header, bytes8 nonceLe, uint256[] dataSetLookup, uint256[] witnessForLookup, uint256 epochIndex) constant returns(uint256)
func (_Ethash *EthashCaller) Hashimoto(opts *bind.CallOpts, header [32]byte, nonceLe [8]byte, dataSetLookup []*big.Int, witnessForLookup []*big.Int, epochIndex *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Ethash.contract.Call(opts, out, "hashimoto", header, nonceLe, dataSetLookup, witnessForLookup, epochIndex)
	return *ret0, err
}

// Hashimoto is a free data retrieval call binding the contract method 0x40ed79f4.
//
// Solidity: function hashimoto(bytes32 header, bytes8 nonceLe, uint256[] dataSetLookup, uint256[] witnessForLookup, uint256 epochIndex) constant returns(uint256)
func (_Ethash *EthashSession) Hashimoto(header [32]byte, nonceLe [8]byte, dataSetLookup []*big.Int, witnessForLookup []*big.Int, epochIndex *big.Int) (*big.Int, error) {
	return _Ethash.Contract.Hashimoto(&_Ethash.CallOpts, header, nonceLe, dataSetLookup, witnessForLookup, epochIndex)
}

// Hashimoto is a free data retrieval call binding the contract method 0x40ed79f4.
//
// Solidity: function hashimoto(bytes32 header, bytes8 nonceLe, uint256[] dataSetLookup, uint256[] witnessForLookup, uint256 epochIndex) constant returns(uint256)
func (_Ethash *EthashCallerSession) Hashimoto(header [32]byte, nonceLe [8]byte, dataSetLookup []*big.Int, witnessForLookup []*big.Int, epochIndex *big.Int) (*big.Int, error) {
	return _Ethash.Contract.Hashimoto(&_Ethash.CallOpts, header, nonceLe, dataSetLookup, witnessForLookup, epochIndex)
}

// IsEpochDataSet is a free data retrieval call binding the contract method 0xc7b81f4f.
//
// Solidity: function isEpochDataSet(uint256 epochIndex) constant returns(bool)
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
// Solidity: function isEpochDataSet(uint256 epochIndex) constant returns(bool)
func (_Ethash *EthashSession) IsEpochDataSet(epochIndex *big.Int) (bool, error) {
	return _Ethash.Contract.IsEpochDataSet(&_Ethash.CallOpts, epochIndex)
}

// IsEpochDataSet is a free data retrieval call binding the contract method 0xc7b81f4f.
//
// Solidity: function isEpochDataSet(uint256 epochIndex) constant returns(bool)
func (_Ethash *EthashCallerSession) IsEpochDataSet(epochIndex *big.Int) (bool, error) {
	return _Ethash.Contract.IsEpochDataSet(&_Ethash.CallOpts, epochIndex)
}

// Test is a free data retrieval call binding the contract method 0x1adda5fe.
//
// Solidity: function test(uint256 blockNumber, bytes32 rlpHeaderHashWithoutNonce, uint256 nonce, uint256 difficulty, uint256[] dataSetLookup, uint256[] witnessForLookup) constant returns(bool)
func (_Ethash *EthashCaller) Test(opts *bind.CallOpts, blockNumber *big.Int, rlpHeaderHashWithoutNonce [32]byte, nonce *big.Int, difficulty *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ethash.contract.Call(opts, out, "test", blockNumber, rlpHeaderHashWithoutNonce, nonce, difficulty, dataSetLookup, witnessForLookup)
	return *ret0, err
}

// Test is a free data retrieval call binding the contract method 0x1adda5fe.
//
// Solidity: function test(uint256 blockNumber, bytes32 rlpHeaderHashWithoutNonce, uint256 nonce, uint256 difficulty, uint256[] dataSetLookup, uint256[] witnessForLookup) constant returns(bool)
func (_Ethash *EthashSession) Test(blockNumber *big.Int, rlpHeaderHashWithoutNonce [32]byte, nonce *big.Int, difficulty *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int) (bool, error) {
	return _Ethash.Contract.Test(&_Ethash.CallOpts, blockNumber, rlpHeaderHashWithoutNonce, nonce, difficulty, dataSetLookup, witnessForLookup)
}

// Test is a free data retrieval call binding the contract method 0x1adda5fe.
//
// Solidity: function test(uint256 blockNumber, bytes32 rlpHeaderHashWithoutNonce, uint256 nonce, uint256 difficulty, uint256[] dataSetLookup, uint256[] witnessForLookup) constant returns(bool)
func (_Ethash *EthashCallerSession) Test(blockNumber *big.Int, rlpHeaderHashWithoutNonce [32]byte, nonce *big.Int, difficulty *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int) (bool, error) {
	return _Ethash.Contract.Test(&_Ethash.CallOpts, blockNumber, rlpHeaderHashWithoutNonce, nonce, difficulty, dataSetLookup, witnessForLookup)
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

// VerifyPoW is a paid mutator transaction binding the contract method 0x29e265df.
//
// Solidity: function verifyPoW(uint256 blockNumber, bytes32 rlpHeaderHashWithoutNonce, uint256 nonce, uint256 difficulty, uint256[] dataSetLookup, uint256[] witnessForLookup) returns(bool, uint256, uint256)
func (_Ethash *EthashTransactor) VerifyPoW(opts *bind.TransactOpts, blockNumber *big.Int, rlpHeaderHashWithoutNonce [32]byte, nonce *big.Int, difficulty *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int) (*types.Transaction, error) {
	return _Ethash.contract.Transact(opts, "verifyPoW", blockNumber, rlpHeaderHashWithoutNonce, nonce, difficulty, dataSetLookup, witnessForLookup)
}

// VerifyPoW is a paid mutator transaction binding the contract method 0x29e265df.
//
// Solidity: function verifyPoW(uint256 blockNumber, bytes32 rlpHeaderHashWithoutNonce, uint256 nonce, uint256 difficulty, uint256[] dataSetLookup, uint256[] witnessForLookup) returns(bool, uint256, uint256)
func (_Ethash *EthashSession) VerifyPoW(blockNumber *big.Int, rlpHeaderHashWithoutNonce [32]byte, nonce *big.Int, difficulty *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int) (*types.Transaction, error) {
	return _Ethash.Contract.VerifyPoW(&_Ethash.TransactOpts, blockNumber, rlpHeaderHashWithoutNonce, nonce, difficulty, dataSetLookup, witnessForLookup)
}

// VerifyPoW is a paid mutator transaction binding the contract method 0x29e265df.
//
// Solidity: function verifyPoW(uint256 blockNumber, bytes32 rlpHeaderHashWithoutNonce, uint256 nonce, uint256 difficulty, uint256[] dataSetLookup, uint256[] witnessForLookup) returns(bool, uint256, uint256)
func (_Ethash *EthashTransactorSession) VerifyPoW(blockNumber *big.Int, rlpHeaderHashWithoutNonce [32]byte, nonce *big.Int, difficulty *big.Int, dataSetLookup []*big.Int, witnessForLookup []*big.Int) (*types.Transaction, error) {
	return _Ethash.Contract.VerifyPoW(&_Ethash.TransactOpts, blockNumber, rlpHeaderHashWithoutNonce, nonce, difficulty, dataSetLookup, witnessForLookup)
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

// EthashVerifyPoWIterator is returned from FilterVerifyPoW and is used to iterate over the raw logs and unpacked data for VerifyPoW events raised by the Ethash contract.
type EthashVerifyPoWIterator struct {
	Event *EthashVerifyPoW // Event containing the contract specifics and raw log

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
func (it *EthashVerifyPoWIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthashVerifyPoW)
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
		it.Event = new(EthashVerifyPoW)
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
func (it *EthashVerifyPoWIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthashVerifyPoWIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthashVerifyPoW represents a VerifyPoW event raised by the Ethash contract.
type EthashVerifyPoW struct {
	Sender    common.Address
	ErrorCode *big.Int
	ErrorInfo *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVerifyPoW is a free log retrieval operation binding the contract event 0x1bdc8073b72ce69f515135b100be07a0e029bb336d555c1a53d80715de4279c3.
//
// Solidity: event VerifyPoW(address indexed sender, uint256 errorCode, uint256 errorInfo)
func (_Ethash *EthashFilterer) FilterVerifyPoW(opts *bind.FilterOpts, sender []common.Address) (*EthashVerifyPoWIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ethash.contract.FilterLogs(opts, "VerifyPoW", senderRule)
	if err != nil {
		return nil, err
	}
	return &EthashVerifyPoWIterator{contract: _Ethash.contract, event: "VerifyPoW", logs: logs, sub: sub}, nil
}

// WatchVerifyPoW is a free log subscription operation binding the contract event 0x1bdc8073b72ce69f515135b100be07a0e029bb336d555c1a53d80715de4279c3.
//
// Solidity: event VerifyPoW(address indexed sender, uint256 errorCode, uint256 errorInfo)
func (_Ethash *EthashFilterer) WatchVerifyPoW(opts *bind.WatchOpts, sink chan<- *EthashVerifyPoW, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ethash.contract.WatchLogs(opts, "VerifyPoW", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthashVerifyPoW)
				if err := _Ethash.contract.UnpackLog(event, "VerifyPoW", log); err != nil {
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
