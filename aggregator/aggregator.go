// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aggregator

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

// CredentialSumABI is the input ABI used to generate the binding from.
const CredentialSumABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proof\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"aggregatedBlock\",\"type\":\"uint256\"}],\"name\":\"AggregatedRoot\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"digests\",\"type\":\"bytes32[]\"}],\"name\":\"computeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"digests\",\"type\":\"bytes32[]\"}],\"name\":\"verifyRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// CredentialSumBin is the compiled bytecode used for deploying new contracts.
var CredentialSumBin = "0x6105a461003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061006c5760003560e01c8063070f6e8d1461007157806316db33f6146100995780634650570f146100ac5780634a2ebd9b146100c0578063d99e4738146100df578063fae2d12a146100ff575b600080fd5b61008461007f3660046104cf565b610112565b60405190151581526020015b60405180910390f35b6100846100a73660046103dc565b610185565b6100846100ba366004610455565b54151590565b6100d16100ce366004610455565b90565b604051908152602001610090565b8180156100eb57600080fd5b506100d16100fa36600461046d565b6101cd565b6100d161010d3660046103a1565b610286565b60008261011f8154151590565b6101705760405162461bcd60e51b815260206004820152601e60248201527f43726564656e7469616c53756d2f70726f6f66206e6f7420657869737473000060448201526064015b60405180910390fd5b61017983610286565b84541491505092915050565b60006101c383838080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061028692505050565b9093149392505050565b600081600081511161021c5760405162461bcd60e51b815260206004820152601860248201527710dc9959195b9d1a585b14dd5b4bd95b5c1d1e481b1a5cdd60421b6044820152606401610167565b600061022784610286565b80875543600188018190554260028901556040519081529091506001600160a01b03861690339083907feeffbadd8c438905a1fac368e2d0860fc30501f02257d4258bf825cafd34d6dc9060200160405180910390a495945050505050565b60008160008151116102d55760405162461bcd60e51b815260206004820152601860248201527710dc9959195b9d1a585b14dd5b4bd95b5c1d1e481b1a5cdd60421b6044820152606401610167565b826040516020016102e69190610514565b60405160208183030381529060405280519060200120915050919050565b600082601f830112610314578081fd5b8135602067ffffffffffffffff8083111561033157610331610558565b818302604051601f19603f8301168101818110848211171561035557610355610558565b60405284815283810192508684018288018501891015610373578687fd5b8692505b85831015610395578035845292840192600192909201918401610377565b50979650505050505050565b6000602082840312156103b2578081fd5b813567ffffffffffffffff8111156103c8578182fd5b6103d484828501610304565b949350505050565b6000806000604084860312156103f0578182fd5b83359250602084013567ffffffffffffffff8082111561040e578384fd5b818601915086601f830112610421578384fd5b81358181111561042f578485fd5b8760208083028501011115610442578485fd5b6020830194508093505050509250925092565b600060208284031215610466578081fd5b5035919050565b600080600060608486031215610481578283fd5b8335925060208401356001600160a01b038116811461049e578283fd5b9150604084013567ffffffffffffffff8111156104b9578182fd5b6104c586828701610304565b9150509250925092565b600080604083850312156104e1578182fd5b82359150602083013567ffffffffffffffff8111156104fe578182fd5b61050a85828601610304565b9150509250929050565b6020808252825182820181905260009190848201906040850190845b8181101561054c57835183529284019291840191600101610530565b50909695505050505050565b634e487b7160e01b600052604160045260246000fdfea2646970667358221220ac2f61dfec0c08730516da33bad25ed82a8ccc6f04c08fa7a279d99302c7e5ab64736f6c63430008020033"

// DeployCredentialSum deploys a new Ethereum contract, binding an instance of CredentialSum to it.
func DeployCredentialSum(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CredentialSum, error) {
	parsed, err := abi.JSON(strings.NewReader(CredentialSumABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CredentialSumBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CredentialSum{CredentialSumCaller: CredentialSumCaller{contract: contract}, CredentialSumTransactor: CredentialSumTransactor{contract: contract}, CredentialSumFilterer: CredentialSumFilterer{contract: contract}}, nil
}

// CredentialSum is an auto generated Go binding around an Ethereum contract.
type CredentialSum struct {
	CredentialSumCaller     // Read-only binding to the contract
	CredentialSumTransactor // Write-only binding to the contract
	CredentialSumFilterer   // Log filterer for contract events
}

// CredentialSumCaller is an auto generated read-only Go binding around an Ethereum contract.
type CredentialSumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CredentialSumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CredentialSumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CredentialSumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CredentialSumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CredentialSumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CredentialSumSession struct {
	Contract     *CredentialSum    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CredentialSumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CredentialSumCallerSession struct {
	Contract *CredentialSumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CredentialSumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CredentialSumTransactorSession struct {
	Contract     *CredentialSumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CredentialSumRaw is an auto generated low-level Go binding around an Ethereum contract.
type CredentialSumRaw struct {
	Contract *CredentialSum // Generic contract binding to access the raw methods on
}

// CredentialSumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CredentialSumCallerRaw struct {
	Contract *CredentialSumCaller // Generic read-only contract binding to access the raw methods on
}

// CredentialSumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CredentialSumTransactorRaw struct {
	Contract *CredentialSumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCredentialSum creates a new instance of CredentialSum, bound to a specific deployed contract.
func NewCredentialSum(address common.Address, backend bind.ContractBackend) (*CredentialSum, error) {
	contract, err := bindCredentialSum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CredentialSum{CredentialSumCaller: CredentialSumCaller{contract: contract}, CredentialSumTransactor: CredentialSumTransactor{contract: contract}, CredentialSumFilterer: CredentialSumFilterer{contract: contract}}, nil
}

// NewCredentialSumCaller creates a new read-only instance of CredentialSum, bound to a specific deployed contract.
func NewCredentialSumCaller(address common.Address, caller bind.ContractCaller) (*CredentialSumCaller, error) {
	contract, err := bindCredentialSum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CredentialSumCaller{contract: contract}, nil
}

// NewCredentialSumTransactor creates a new write-only instance of CredentialSum, bound to a specific deployed contract.
func NewCredentialSumTransactor(address common.Address, transactor bind.ContractTransactor) (*CredentialSumTransactor, error) {
	contract, err := bindCredentialSum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CredentialSumTransactor{contract: contract}, nil
}

// NewCredentialSumFilterer creates a new log filterer instance of CredentialSum, bound to a specific deployed contract.
func NewCredentialSumFilterer(address common.Address, filterer bind.ContractFilterer) (*CredentialSumFilterer, error) {
	contract, err := bindCredentialSum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CredentialSumFilterer{contract: contract}, nil
}

// bindCredentialSum binds a generic wrapper to an already deployed contract.
func bindCredentialSum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CredentialSumABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CredentialSum *CredentialSumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CredentialSum.Contract.CredentialSumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CredentialSum *CredentialSumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CredentialSum.Contract.CredentialSumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CredentialSum *CredentialSumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CredentialSum.Contract.CredentialSumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CredentialSum *CredentialSumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CredentialSum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CredentialSum *CredentialSumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CredentialSum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CredentialSum *CredentialSumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CredentialSum.Contract.contract.Transact(opts, method, params...)
}

// ComputeRoot is a free data retrieval call binding the contract method 0xfae2d12a.
//
// Solidity: function computeRoot(bytes32[] digests) pure returns(bytes32)
func (_CredentialSum *CredentialSumCaller) ComputeRoot(opts *bind.CallOpts, digests [][32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CredentialSum.contract.Call(opts, &out, "computeRoot", digests)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeRoot is a free data retrieval call binding the contract method 0xfae2d12a.
//
// Solidity: function computeRoot(bytes32[] digests) pure returns(bytes32)
func (_CredentialSum *CredentialSumSession) ComputeRoot(digests [][32]byte) ([32]byte, error) {
	return _CredentialSum.Contract.ComputeRoot(&_CredentialSum.CallOpts, digests)
}

// ComputeRoot is a free data retrieval call binding the contract method 0xfae2d12a.
//
// Solidity: function computeRoot(bytes32[] digests) pure returns(bytes32)
func (_CredentialSum *CredentialSumCallerSession) ComputeRoot(digests [][32]byte) ([32]byte, error) {
	return _CredentialSum.Contract.ComputeRoot(&_CredentialSum.CallOpts, digests)
}

// VerifyRoot is a free data retrieval call binding the contract method 0x16db33f6.
//
// Solidity: function verifyRoot(bytes32 root, bytes32[] digests) pure returns(bool)
func (_CredentialSum *CredentialSumCaller) VerifyRoot(opts *bind.CallOpts, root [32]byte, digests [][32]byte) (bool, error) {
	var out []interface{}
	err := _CredentialSum.contract.Call(opts, &out, "verifyRoot", root, digests)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyRoot is a free data retrieval call binding the contract method 0x16db33f6.
//
// Solidity: function verifyRoot(bytes32 root, bytes32[] digests) pure returns(bool)
func (_CredentialSum *CredentialSumSession) VerifyRoot(root [32]byte, digests [][32]byte) (bool, error) {
	return _CredentialSum.Contract.VerifyRoot(&_CredentialSum.CallOpts, root, digests)
}

// VerifyRoot is a free data retrieval call binding the contract method 0x16db33f6.
//
// Solidity: function verifyRoot(bytes32 root, bytes32[] digests) pure returns(bool)
func (_CredentialSum *CredentialSumCallerSession) VerifyRoot(root [32]byte, digests [][32]byte) (bool, error) {
	return _CredentialSum.Contract.VerifyRoot(&_CredentialSum.CallOpts, root, digests)
}

// CredentialSumAggregatedRootIterator is returned from FilterAggregatedRoot and is used to iterate over the raw logs and unpacked data for AggregatedRoot events raised by the CredentialSum contract.
type CredentialSumAggregatedRootIterator struct {
	Event *CredentialSumAggregatedRoot // Event containing the contract specifics and raw log

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
func (it *CredentialSumAggregatedRootIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CredentialSumAggregatedRoot)
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
		it.Event = new(CredentialSumAggregatedRoot)
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
func (it *CredentialSumAggregatedRootIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CredentialSumAggregatedRootIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CredentialSumAggregatedRoot represents a AggregatedRoot event raised by the CredentialSum contract.
type CredentialSumAggregatedRoot struct {
	Proof           [32]byte
	Aggregator      common.Address
	Subject         common.Address
	AggregatedBlock *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAggregatedRoot is a free log retrieval operation binding the contract event 0xeeffbadd8c438905a1fac368e2d0860fc30501f02257d4258bf825cafd34d6dc.
//
// Solidity: event AggregatedRoot(bytes32 indexed proof, address indexed aggregator, address indexed subject, uint256 aggregatedBlock)
func (_CredentialSum *CredentialSumFilterer) FilterAggregatedRoot(opts *bind.FilterOpts, proof [][32]byte, aggregator []common.Address, subject []common.Address) (*CredentialSumAggregatedRootIterator, error) {

	var proofRule []interface{}
	for _, proofItem := range proof {
		proofRule = append(proofRule, proofItem)
	}
	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}

	logs, sub, err := _CredentialSum.contract.FilterLogs(opts, "AggregatedRoot", proofRule, aggregatorRule, subjectRule)
	if err != nil {
		return nil, err
	}
	return &CredentialSumAggregatedRootIterator{contract: _CredentialSum.contract, event: "AggregatedRoot", logs: logs, sub: sub}, nil
}

// WatchAggregatedRoot is a free log subscription operation binding the contract event 0xeeffbadd8c438905a1fac368e2d0860fc30501f02257d4258bf825cafd34d6dc.
//
// Solidity: event AggregatedRoot(bytes32 indexed proof, address indexed aggregator, address indexed subject, uint256 aggregatedBlock)
func (_CredentialSum *CredentialSumFilterer) WatchAggregatedRoot(opts *bind.WatchOpts, sink chan<- *CredentialSumAggregatedRoot, proof [][32]byte, aggregator []common.Address, subject []common.Address) (event.Subscription, error) {

	var proofRule []interface{}
	for _, proofItem := range proof {
		proofRule = append(proofRule, proofItem)
	}
	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}

	logs, sub, err := _CredentialSum.contract.WatchLogs(opts, "AggregatedRoot", proofRule, aggregatorRule, subjectRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CredentialSumAggregatedRoot)
				if err := _CredentialSum.contract.UnpackLog(event, "AggregatedRoot", log); err != nil {
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

// ParseAggregatedRoot is a log parse operation binding the contract event 0xeeffbadd8c438905a1fac368e2d0860fc30501f02257d4258bf825cafd34d6dc.
//
// Solidity: event AggregatedRoot(bytes32 indexed proof, address indexed aggregator, address indexed subject, uint256 aggregatedBlock)
func (_CredentialSum *CredentialSumFilterer) ParseAggregatedRoot(log types.Log) (*CredentialSumAggregatedRoot, error) {
	event := new(CredentialSumAggregatedRoot)
	if err := _CredentialSum.contract.UnpackLog(event, "AggregatedRoot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
