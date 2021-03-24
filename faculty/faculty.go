// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package faculty

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

// CredentialSumRoot is an auto generated low-level Go binding around an user-defined struct.
type CredentialSumRoot struct {
	Proof          [32]byte
	InsertedBlock  *big.Int
	BlockTimestamp *big.Int
}

// NotaryCredentialProof is an auto generated low-level Go binding around an user-defined struct.
type NotaryCredentialProof struct {
	Signed         *big.Int
	InsertedBlock  *big.Int
	BlockTimestamp *big.Int
	Nonce          *big.Int
	Digest         [32]byte
	Approved       bool
	Registrar      common.Address
	Subject        common.Address
	Witnesses      []common.Address
	EvidenceRoot   [32]byte
}

// NotaryRevocationProof is an auto generated low-level Go binding around an user-defined struct.
type NotaryRevocationProof struct {
	Registrar    common.Address
	Subject      common.Address
	RevokedBlock *big.Int
	Reason       [32]byte
}

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

// ERC165ABI is the input ABI used to generate the binding from.
const ERC165ABI = "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ERC165 is an auto generated Go binding around an Ethereum contract.
type ERC165 struct {
	ERC165Caller     // Read-only binding to the contract
	ERC165Transactor // Write-only binding to the contract
	ERC165Filterer   // Log filterer for contract events
}

// ERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165Session struct {
	Contract     *ERC165           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165CallerSession struct {
	Contract *ERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165TransactorSession struct {
	Contract     *ERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165Raw struct {
	Contract *ERC165 // Generic contract binding to access the raw methods on
}

// ERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165CallerRaw struct {
	Contract *ERC165Caller // Generic read-only contract binding to access the raw methods on
}

// ERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165TransactorRaw struct {
	Contract *ERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165 creates a new instance of ERC165, bound to a specific deployed contract.
func NewERC165(address common.Address, backend bind.ContractBackend) (*ERC165, error) {
	contract, err := bindERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// NewERC165Caller creates a new read-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Caller(address common.Address, caller bind.ContractCaller) (*ERC165Caller, error) {
	contract, err := bindERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Caller{contract: contract}, nil
}

// NewERC165Transactor creates a new write-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC165Transactor, error) {
	contract, err := bindERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Transactor{contract: contract}, nil
}

// NewERC165Filterer creates a new log filterer instance of ERC165, bound to a specific deployed contract.
func NewERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC165Filterer, error) {
	contract, err := bindERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165Filterer{contract: contract}, nil
}

// bindERC165 binds a generic wrapper to an already deployed contract.
func bindERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.ERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_ERC165 *ERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_ERC165 *ERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_ERC165 *ERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// ERC165CheckerABI is the input ABI used to generate the binding from.
const ERC165CheckerABI = "[]"

// ERC165CheckerBin is the compiled bytecode used for deploying new contracts.
var ERC165CheckerBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205f5412c719b586053d3d26ea99ed0357ec020c6de5712b96ac32639e9a10d6e564736f6c63430008020033"

// DeployERC165Checker deploys a new Ethereum contract, binding an instance of ERC165Checker to it.
func DeployERC165Checker(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC165Checker, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165CheckerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC165CheckerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC165Checker{ERC165CheckerCaller: ERC165CheckerCaller{contract: contract}, ERC165CheckerTransactor: ERC165CheckerTransactor{contract: contract}, ERC165CheckerFilterer: ERC165CheckerFilterer{contract: contract}}, nil
}

// ERC165Checker is an auto generated Go binding around an Ethereum contract.
type ERC165Checker struct {
	ERC165CheckerCaller     // Read-only binding to the contract
	ERC165CheckerTransactor // Write-only binding to the contract
	ERC165CheckerFilterer   // Log filterer for contract events
}

// ERC165CheckerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165CheckerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165CheckerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165CheckerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165CheckerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165CheckerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165CheckerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165CheckerSession struct {
	Contract     *ERC165Checker    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165CheckerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165CheckerCallerSession struct {
	Contract *ERC165CheckerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC165CheckerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165CheckerTransactorSession struct {
	Contract     *ERC165CheckerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC165CheckerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165CheckerRaw struct {
	Contract *ERC165Checker // Generic contract binding to access the raw methods on
}

// ERC165CheckerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165CheckerCallerRaw struct {
	Contract *ERC165CheckerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC165CheckerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165CheckerTransactorRaw struct {
	Contract *ERC165CheckerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165Checker creates a new instance of ERC165Checker, bound to a specific deployed contract.
func NewERC165Checker(address common.Address, backend bind.ContractBackend) (*ERC165Checker, error) {
	contract, err := bindERC165Checker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165Checker{ERC165CheckerCaller: ERC165CheckerCaller{contract: contract}, ERC165CheckerTransactor: ERC165CheckerTransactor{contract: contract}, ERC165CheckerFilterer: ERC165CheckerFilterer{contract: contract}}, nil
}

// NewERC165CheckerCaller creates a new read-only instance of ERC165Checker, bound to a specific deployed contract.
func NewERC165CheckerCaller(address common.Address, caller bind.ContractCaller) (*ERC165CheckerCaller, error) {
	contract, err := bindERC165Checker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165CheckerCaller{contract: contract}, nil
}

// NewERC165CheckerTransactor creates a new write-only instance of ERC165Checker, bound to a specific deployed contract.
func NewERC165CheckerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC165CheckerTransactor, error) {
	contract, err := bindERC165Checker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165CheckerTransactor{contract: contract}, nil
}

// NewERC165CheckerFilterer creates a new log filterer instance of ERC165Checker, bound to a specific deployed contract.
func NewERC165CheckerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC165CheckerFilterer, error) {
	contract, err := bindERC165Checker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165CheckerFilterer{contract: contract}, nil
}

// bindERC165Checker binds a generic wrapper to an already deployed contract.
func bindERC165Checker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165CheckerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165Checker *ERC165CheckerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165Checker.Contract.ERC165CheckerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165Checker *ERC165CheckerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165Checker.Contract.ERC165CheckerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165Checker *ERC165CheckerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165Checker.Contract.ERC165CheckerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165Checker *ERC165CheckerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165Checker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165Checker *ERC165CheckerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165Checker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165Checker *ERC165CheckerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165Checker.Contract.contract.Transact(opts, method, params...)
}

// FacultyABI is the input ABI used to generate the binding from.
const FacultyABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"quorum\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registrar\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"insertedBlock\",\"type\":\"uint256\"}],\"name\":\"CredentialIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"revoker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revokedBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"reason\",\"type\":\"bytes32\"}],\"name\":\"CredentialRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"signedBlock\",\"type\":\"uint256\"}],\"name\":\"CredentialSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"createdBy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"NodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registar\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"semester\",\"type\":\"bytes32\"}],\"name\":\"SemesterRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"addChild\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"digests\",\"type\":\"bytes32[]\"}],\"name\":\"aggregateCredentials\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"approveCredential\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChildren\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"semester\",\"type\":\"bytes32\"}],\"name\":\"getCoursesBySemester\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"getCredentialProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"signed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"insertedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"registrar\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"witnesses\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"evidenceRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structNotary.CredentialProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"getCredentialSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"getDigests\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"getEvidenceRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"getProof\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"proof\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"insertedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structCredentialSum.Root\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"getRevoked\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"getRevokedProof\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"registrar\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"revokedBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"reason\",\"type\":\"bytes32\"}],\"internalType\":\"structNotary.RevocationProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRole\",\"outputs\":[{\"internalType\":\"enumRole\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"getWitnesses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"hasRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"isApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isChild\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLeaf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"isQuorumSigned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"isRevoked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isSigned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"myParent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownersCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"recordExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"witnesses\",\"type\":\"address[]\"}],\"name\":\"registerCredential\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"semester\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"courses\",\"type\":\"address[]\"}],\"name\":\"registerSemester\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"reason\",\"type\":\"bytes32\"}],\"name\":\"revokeCredential\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"revokedCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"semester\",\"type\":\"bytes32\"}],\"name\":\"semesterExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"verifyCredential\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyCredentialRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"verifyCredentialTree\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"verifyIssuedCredentials\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"digests\",\"type\":\"bytes32[]\"}],\"name\":\"verifyRootOf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"witnessesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FacultyBin is the compiled bytecode used for deploying new contracts.
var FacultyBin = "0x60c0604052635041329d60e01b60a09081526200002190600b90600162000328565b503480156200002f57600080fd5b506040516200463e3803806200463e833981016040819052620000529162000479565b6001828281818181600082511180156200006e5750815160ff10155b620000c05760405162461bcd60e51b815260206004820152601860248201527f4f776e6572732f6e6f7420656e6f756768206f776e657273000000000000000060448201526064015b60405180910390fd5b60008160ff16118015620000d8575081518160ff1611155b620001265760405162461bcd60e51b815260206004820152601a60248201527f4f776e6572732f71756f72756d206f7574206f662072616e67650000000000006044820152606401620000b7565b60005b82518160ff1610156200025b5760026000848360ff16815181106200015e57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff16158015620001d1575060006001600160a01b0316838260ff1681518110620001bd57634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b031614155b620001ec57634e487b7160e01b600052600160045260246000fd5b600160026000858460ff16815181106200021657634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055620002538162000560565b905062000129565b50815162000271906001906020850190620003db565b5090516000805461ff00191661010060ff938416021760ff19169190921617905550339050620002e45760405162461bcd60e51b815260206004820152601760248201527f4e6f64652f73656e6465722063616e6e6f7420626520300000000000000000006044820152606401620000b7565b3360601b608052600c805484919060ff1916600183818111156200031857634e487b7160e01b600052602160045260246000fd5b02179055505050505050620005a3565b82805482825590600052602060002090600701600890048101928215620003c95791602002820160005b838211156200039557835183826101000a81548163ffffffff021916908360e01c0217905550926020019260040160208160030104928301926001030262000352565b8015620003c75782816101000a81549063ffffffff021916905560040160208160030104928301926001030262000395565b505b50620003d792915062000433565b5090565b828054828255906000526020600020908101928215620003c9579160200282015b82811115620003c957825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620003fc565b5b80821115620003d7576000815560010162000434565b80516001600160a01b03811681146200046257600080fd5b919050565b805160ff811681146200046257600080fd5b600080604083850312156200048c578182fd5b82516001600160401b0380821115620004a3578384fd5b818501915085601f830112620004b7578384fd5b8151602082821115620004ce57620004ce6200058d565b808202604051601f19603f83011681018181108682111715620004f557620004f56200058d565b604052838152828101945085830182870184018b101562000514578889fd5b8896505b8487101562000541576200052c816200044a565b86526001969096019594830194830162000518565b50965062000553905087820162000467565b9450505050509250929050565b600060ff821660ff8114156200058457634e487b7160e01b82526011600452602482fd5b60010192915050565b634e487b7160e01b600052604160045260246000fd5b60805160601c61407c620005c26000396000610609015261407c6000f3fe608060405234801561001057600080fd5b50600436106102535760003560e01c8063762b77bb11610146578063b9488546116100c3578063eedb973c11610087578063eedb973c14610646578063f537ff9d14610659578063f79fe5381461067c578063f936cc151461068f578063fc91a897146106a2578063fde08174146106c557610253565b8063b9488546146105b8578063c8bf608a146105c9578063dac46a62146105dc578063de07fc45146105fc578063e2d02cea1461063357610253565b8063a355f4a51161010a578063a355f4a514610523578063a6f9dae114610536578063abbff98414610549578063ae02c0cc1461059d578063affe39c1146105b057610253565b8063762b77bb146104c25780637f8d61e0146104e257806386509dd6146104f55780638ebd96231461050857806399d1dbc01461051b57610253565b80633610d3ca116101d45780634892e8e8116101985780634892e8e81461046e57806348aefc32146104765780635f889e17146104895780636f28eca61461049c5780637532ed26146104af57610253565b80633610d3ca1461038c5780633ac82e191461039f5780633eea79d1146103b25780634294857f146104385780634302802d1461044b57610253565b80631eee993a1161021b5780631eee993a146102ed5780632adfe8b2146103025780632cf0695a146103245780632f54bf6e14610337578063342577fa1461036357610253565b806301ffc9a714610258578063038ec1b814610280578063079cf76e146102a05780630ae07028146102c15780631703a018146102d4575b600080fd5b61026b610266366004613adf565b6106d8565b60405190151581526020015b60405180910390f35b61029361028e366004613a1a565b610711565b6040516102779190613d0e565b6102b36102ae3660046138bd565b61087a565b604051908152602001610277565b61026b6102cf36600461397e565b610898565b60005460ff165b60405160ff9091168152602001610277565b6103006102fb3660046138bd565b6108ad565b005b61026b610310366004613a1a565b6000908152600f6020526040902054151590565b61026b6103323660046138bd565b610a89565b61026b6103453660046138bd565b6001600160a01b031660009081526002602052604090205460ff1690565b6102b36103713660046138bd565b6001600160a01b031660009081526008602052604090205490565b61030061039a366004613abe565b610ad2565b6103006103ad3660046139a9565b610b0f565b6104166103c03660046138bd565b6040805160608082018352600080835260208084018290529284018190526001600160a01b03949094168452600a8252928290208251938401835280548452600181015491840191909152600201549082015290565b6040805182518152602080840151908201529181015190820152606001610277565b61026b610446366004613a1a565b610f3d565b6102b3610459366004613a1a565b60009081526006602052604090206008015490565b610293610fcb565b61026b610484366004613a1a565b61102d565b61026b61049736600461397e565b61106f565b6102936104aa366004613a1a565b61107d565b61026b6104bd3660046138bd565b6110ec565b6104d56104d03660046138bd565b61113d565b6040516102779190613d34565b61026b6104f0366004613a4a565b6111a8565b61026b6105033660046138d9565b611245565b6104d56105163660046138bd565b611295565b61026b611448565b610293610531366004613a1a565b611476565b6103006105443660046138bd565b611535565b61055c610557366004613a1a565b611843565b604051610277919081516001600160a01b03908116825260208084015190911690820152604080830151908201526060918201519181019190915260800190565b61026b6105ab366004613a1a565b6118ed565b61029361193a565b6102db600054610100900460ff1690565b6103006105d7366004613a79565b61199a565b6105ef6105ea366004613a1a565b611ad3565b6040516102779190613e5e565b6040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152602001610277565b6102b36106413660046138d9565b611bad565b610300610654366004613a1a565b611be6565b6102b3610667366004613a1a565b60009081526006602052604090206007015490565b61026b61068a366004613a1a565b611bf2565b61026b61069d3660046138bd565b611c34565b61026b6106b03660046138bd565b600e6020526000908152604090205460ff1681565b600c5460ff166040516102779190613d91565b60006001600160e01b031982166301ffc9a760e01b148061070957506001600160e01b03198216635041329d60e01b145b90505b919050565b60015460609060009067ffffffffffffffff81111561074057634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610769578160200160208202803683370190505b5090506000805b60015481101561087157600085815260096020526040812060018054919291849081106107ad57634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff161561085f57600181815481106107fb57634e487b7160e01b600052603260045260246000fd5b9060005260206000200160009054906101000a90046001600160a01b031683838151811061083957634e487b7160e01b600052603260045260246000fd5b6001600160a01b03909216602092830291909101909101528161085b81613fca565b9250505b8061086981613fca565b915050610770565b50909392505050565b6001600160a01b0381166000908152600a6020526040812054610709565b60006108a48383611e6b565b90505b92915050565b3360009081526002602052604090205460ff166108e55760405162461bcd60e51b81526004016108dc90613df0565b60405180910390fd5b6001600c5460ff16600181111561090c57634e487b7160e01b600052602160045260246000fd5b146109595760405162461bcd60e51b815260206004820152601760248201527f4e6f64652f6e6f6465206d75737420626520496e6e657200000000000000000060448201526064016108dc565b60006109e230600b8054806020026020016040519081016040528092919081815260200182805480156109d857602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b0319168152602001906004019060208260030104928301926001038202915080841161099a5790505b5050505050611f78565b9050806109ff57634e487b7160e01b600052600160045260246000fd5b60008290506000816001600160a01b031663fde081746040518163ffffffff1660e01b815260040160206040518083038186803b158015610a3f57600080fd5b505afa158015610a53573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a779190613b07565b9050610a838482611fec565b50505050565b6001600160a01b0381166000908152600560205260408120548290610ac05760405162461bcd60e51b81526004016108dc90613db9565b610acb6003846121e3565b9392505050565b3360009081526002602052604090205460ff16610b015760405162461bcd60e51b81526004016108dc90613df0565b610b0b8282612256565b5050565b3360009081526002602052604090205460ff16610b3e5760405162461bcd60e51b81526004016108dc90613df0565b6000600c5460ff166001811115610b6557634e487b7160e01b600052602160045260246000fd5b1415610bdd57805115610bba5760405162461bcd60e51b815260206004820152601f60248201527f4e6f64652f4c6561662063616e6e6f742068617665207769746e65737365730060448201526064016108dc565b60408051600080825260208201909252610bd8918591859190612326565b610f38565b6001600c5460ff166001811115610c0457634e487b7160e01b600052602160045260246000fd5b14610c1f57634e487b7160e01b600052600160045260246000fd5b6000815111610c695760405162461bcd60e51b8152602060048201526016602482015275139bd9194bddda5d1b995cdcc81b9bdd08199bdd5b9960521b60448201526064016108dc565b6000815167ffffffffffffffff811115610c9357634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610cbc578160200160208202803683370190505b50905060005b8251811015610e9c576000838281518110610ced57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b0381166000908152600e90925260409091205490915060ff16610d655760405162461bcd60e51b815260206004820152601b60248201527f4e6f64652f61646472657373206e6f7420617574686f72697a6564000000000060448201526064016108dc565b6000610d7882635041329d60e01b6123f3565b905080610d9557634e487b7160e01b600052600160045260246000fd5b6040516303ce7bb760e11b81526001600160a01b038881166004830152839160009183169063079cf76e9060240160206040518083038186803b158015610ddb57600080fd5b505afa158015610def573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e139190613a32565b905080610e585760405162461bcd60e51b8152602060048201526013602482015272139bd9194bdc9bdbdd081b9bdd08199bdd5b99606a1b60448201526064016108dc565b80868681518110610e7957634e487b7160e01b600052603260045260246000fd5b602002602001018181525050505050508080610e9490613fca565b915050610cc2565b50604051637d71689560e11b815260009073__$9176519fc5fb44025266bd62b274a40525$__9063fae2d12a90610ed7908590600401613d21565b60206040518083038186803b158015610eef57600080fd5b505af4158015610f03573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f279190613a32565b9050610f3585858386612326565b50505b505050565b604051631999f00560e21b8152600360048201526024810182905260009073__$1d2b6a72c2c8582bd594e19b02f34b62b6$__90636667c014906044015b60206040518083038186803b158015610f9357600080fd5b505af4158015610fa7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107099190613a00565b6060600d80548060200260200160405190810160405280929190818152602001828054801561102357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611005575b5050505050905090565b6040516312f4761160e31b8152600360048201526024810182905260009073__$1d2b6a72c2c8582bd594e19b02f34b62b6$__906397a3b08890604401610f7b565b60006108a46003848461240f565b6000818152600660209081526040918290206007018054835181840281018401909452808452606093928301828280156110e057602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116110c2575b50505050509050919050565b6001600160a01b0381166000908152600a60205260408082209051634650570f60e01b8152600481019190915273__$9176519fc5fb44025266bd62b274a40525$__90634650570f90602401610f7b565b6001600160a01b0381166000908152600560209081526040918290208054835181840281018401909452808452606093928301828280156110e057602002820191906000526020600020905b8154815260200190600101908083116111895750505050509050919050565b60405163a8b0bcfd60e01b815260036004820152602481018390526001600160a01b038216604482015260009073__$1d2b6a72c2c8582bd594e19b02f34b62b6$__9063a8b0bcfd906064015b60206040518083038186803b15801561120d57600080fd5b505af4158015611221573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108a49190613a00565b6001600160a01b0382166000908152600a6020526040808220905163070f6e8d60e01b815273__$9176519fc5fb44025266bd62b274a40525$__9163070f6e8d916111f591908690600401613d78565b6001600160a01b03811660009081526005602052604090205460609082906112cf5760405162461bcd60e51b81526004016108dc90613db9565b6001600160a01b03831660009081526008602052604081205467ffffffffffffffff81111561130e57634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015611337578160200160208202803683370190505b5090506000805b6001600160a01b03861660009081526005602052604090205481101561143e576001600160a01b038616600090815260056020526040902080546113ab91908390811061139b57634e487b7160e01b600052603260045260246000fd5b9060005260206000200154610f3d565b1561142c576001600160a01b03861660009081526005602052604090208054829081106113e857634e487b7160e01b600052603260045260246000fd5b906000526020600020015483838151811061141357634e487b7160e01b600052603260045260246000fd5b60209081029190910101528161142881613fca565b9250505b8061143681613fca565b91505061133e565b5090949350505050565b600080600c5460ff16600181111561147057634e487b7160e01b600052602160045260246000fd5b14905090565b6000818152600f60205260409020546060906114cd5760405162461bcd60e51b8152602060048201526016602482015275119858dd5b1d1e4bdb9bdd081c9959da5cdd195c995960521b60448201526064016108dc565b6000828152600f6020908152604091829020805483518184028101840190945280845290918301828280156110e0576020028201919060005260206000209081546001600160a01b031681526001909101906020018083116110c25750505050509050919050565b3360009081526002602052604090205460ff166115645760405162461bcd60e51b81526004016108dc90613df0565b6001600160a01b03811660009081526002602052604090205460ff1615801561159557506001600160a01b03811615155b6115e15760405162461bcd60e51b815260206004820152601c60248201527f4f776e6572732f696e76616c6964206164647265737320676976656e0000000060448201526064016108dc565b600154158015906115f5575060015460ff10155b61160f57634e487b7160e01b600052600160045260246000fd5b60015460009067ffffffffffffffff81111561163b57634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015611664578160200160208202803683370190505b50905060005b60015460ff8216101561179e57336001600160a01b031660018260ff16815481106116a557634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b0316146117495760018160ff16815481106116e357634e487b7160e01b600052603260045260246000fd5b9060005260206000200160009054906101000a90046001600160a01b0316828260ff168151811061172457634e487b7160e01b600052603260045260246000fd5b60200260200101906001600160a01b031690816001600160a01b03168152505061178e565b82828260ff168151811061176d57634e487b7160e01b600052603260045260246000fd5b60200260200101906001600160a01b031690816001600160a01b0316815250505b61179781613fe5565b905061166a565b50600054815160ff909116146117c457634e487b7160e01b600052600160045260246000fd5b6040516001600160a01b0383169033907fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c90600090a3805161180d906001906020840190613750565b50506001600160a01b0316600090815260026020526040808220805460ff19908116600117909155338352912080549091169055565b6040805160808101825260008082526020820181905281830181905260608201529051636dbae62960e11b8152600360048201526024810183905273__$1d2b6a72c2c8582bd594e19b02f34b62b6$__9063db75cc529060440160806040518083038186803b1580156118b557600080fd5b505af41580156118c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107099190613c0b565b60008054604051633e01fd5d60e21b8152600360048201526024810184905260ff909116604482015273__$1d2b6a72c2c8582bd594e19b02f34b62b6$__9063f807f57490606401610f7b565b60606001805480602002602001604051908101604052809291908181526020018280548015611023576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611005575050505050905090565b3360009081526002602052604090205460ff166119c95760405162461bcd60e51b81526004016108dc90613df0565b6000828152600f602052604090205415611a255760405162461bcd60e51b815260206004820152601a60248201527f466163756c74792f616c7265616479207265676973746572656400000000000060448201526064016108dc565b6000815111611a765760405162461bcd60e51b815260206004820152601a60248201527f466163756c74792f656d70747920636f7572736573206c69737400000000000060448201526064016108dc565b6000828152600f602090815260409091208251611a9592840190613750565b5060408051338152602081018490527fcb185a4a072cf4e8ef43d02535b01a5ca5a4eeb8dcbc242651888e159e6c8bf0910160405180910390a15050565b604080516101408101825260008082526020820181905281830181905260608083018290526080830182905260a0830182905260c0830182905260e0830182905261010083015261012082015290516301c30fd760e71b8152600360048201526024810183905273__$1d2b6a72c2c8582bd594e19b02f34b62b6$__9063e187eb809060440160006040518083038186803b158015611b7157600080fd5b505af4158015611b85573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526107099190810190613b26565b3360009081526002602052604081205460ff16611bdc5760405162461bcd60e51b81526004016108dc90613df0565b6108a48383612517565b611bef81612678565b50565b60405163143ceab160e21b8152600360048201526024810182905260009073__$1d2b6a72c2c8582bd594e19b02f34b62b6$__906350f3aac490604401610f7b565b600080611c408361113d565b90506000815111611c935760405162461bcd60e51b815260206004820152601960248201527f4e6f64652f63726564656e7469616c206e6f7420666f756e640000000000000060448201526064016108dc565b611c9c836110ec565b15611cb957611cab8382611245565b611cb957600091505061070c565b60005b8151811015611e6157611cf5828281518110611ce857634e487b7160e01b600052603260045260246000fd5b6020026020010151611bf2565b611d0f57634e487b7160e01b600052600160045260246000fd5b611d4084838381518110611d3357634e487b7160e01b600052603260045260246000fd5b602002602001015161106f565b611d4f5760009250505061070c565b6001600c5460ff166001811115611d7657634e487b7160e01b600052602160045260246000fd5b148015611dc257506000611dc0838381518110611da357634e487b7160e01b600052603260045260246000fd5b602002602001015160009081526006602052604090206007015490565b115b15611e4f57611e4084611e0b848481518110611dee57634e487b7160e01b600052603260045260246000fd5b602002602001015160009081526006602052604090206008015490565b611e3b858581518110611e2e57634e487b7160e01b600052603260045260246000fd5b602002602001015161107d565b61274b565b611e4f5760009250505061070c565b80611e5981613fca565b915050611cbc565b5060019392505050565b6001600160a01b0382166000908152600560205260408120548390611ea25760405162461bcd60e51b81526004016108dc90613db9565b6001600160a01b0384166000908152600a60209081526040808320600590925291829020915163070f6e8d60e01b815273__$9176519fc5fb44025266bd62b274a40525$__9263070f6e8d92611efd92909190600401613f32565b60206040518083038186803b158015611f1557600080fd5b505af4158015611f29573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f4d9190613a00565b8015611f7057506001600160a01b0384166000908152600a602052604090205483145b949350505050565b6000611f8383612bea565b611f8f575060006108a7565b60005b8251811015611e6157611fcc84848381518110611fbf57634e487b7160e01b600052603260045260246000fd5b6020026020010151612c1d565b611fda5760009150506108a7565b80611fe481613fca565b915050611f92565b306001600160a01b038316141561203e5760405162461bcd60e51b81526020600482015260166024820152752737b23297b1b0b73737ba1030b2321034ba39b2b63360511b60448201526064016108dc565b6001600160a01b0382166000908152600e602052604090205460ff16156120a75760405162461bcd60e51b815260206004820152601760248201527f4e6f64652f6e6f646520616c726561647920616464656400000000000000000060448201526064016108dc565b60008160018111156120c957634e487b7160e01b600052602160045260246000fd5b14806120f4575060018160018111156120f257634e487b7160e01b600052602160045260246000fd5b145b6121405760405162461bcd60e51b815260206004820152601760248201527f4e6f64652f696e76616c6964206368696c6420726f6c6500000000000000000060448201526064016108dc565b6001600160a01b0382166000818152600e6020526040808220805460ff19166001908117909155600d805491820181559092527fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb590910180546001600160a01b031916831790555133907f202458425cc4a92ea08231128a0a57aa24cfa42968e8090aa7d151f3f0a4eb3f906121d7908590613d91565b60405180910390a35050565b6001600160a01b0381166000908152600283016020908152604080832080548251818502810185019093528083526108a493879387939092909183018282801561224c57602002820191906000526020600020905b815481526020019060010190808311612238575b5050505050612c40565b8161226081610f3d565b1561227d5760405162461bcd60e51b81526004016108dc90613e27565b600083815260066020819052604090912001546001600160a01b03166122bb336001600160a01b031660009081526002602052604090205460ff1690565b806122ce57506001600160a01b03811633145b61231a5760405162461bcd60e51b815260206004820152601c60248201527f4973737565722f73656e646572206e6f7420617574686f72697a65640000000060448201526064016108dc565b610a8360038585612caa565b3360009081526002602052604090205460ff166123555760405162461bcd60e51b81526004016108dc90613df0565b8261235f81610f3d565b1561237c5760405162461bcd60e51b81526004016108dc90613e27565b6001600160a01b03851660009081526002602052604090205460ff16156123e55760405162461bcd60e51b815260206004820152601a60248201527f4973737565722f666f7262696464656e2072656769737472617200000000000060448201526064016108dc565b610f35600386868686612e59565b60006123fe83612bea565b80156108a457506108a48383612c1d565b600081815260038401602052604081206001015461246f5760405162461bcd60e51b815260206004820152601b60248201527f4e6f746172792f63726564656e7469616c206e6f7420666f756e64000000000060448201526064016108dc565b60008281526003850160205260409020600601546001600160a01b038481169116146124dd5760405162461bcd60e51b815260206004820152601b60248201527f4e6f746172792f6e6f74206f776e6564206279207375626a656374000000000060448201526064016108dc565b600082815260038501602052604090206005015460ff168015611f7057505060009081526004929092016020525060409020600201541590565b3360009081526002602052604081205460ff166125465760405162461bcd60e51b81526004016108dc90613df0565b6001600160a01b038316600090815260056020526040902054839061257d5760405162461bcd60e51b81526004016108dc90613db9565b61258960038585612c40565b6125d55760405162461bcd60e51b815260206004820152601e60248201527f4973737565722f68617320696e76616c69642063726564656e7469616c73000060448201526064016108dc565b6001600160a01b0384166000908152600a6020526040908190209051631b33c8e760e31b815273__$9176519fc5fb44025266bd62b274a40525$__9163d99e473891612628919088908890600401613f08565b60206040518083038186803b15801561264057600080fd5b505af4158015612654573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f709190613a32565b8061268281610f3d565b1561269f5760405162461bcd60e51b81526004016108dc90613e27565b60005460ff166126ea5760405162461bcd60e51b8152602060048201526016602482015275125cdcdd595c8bdb9bc81c5d5bdc9d5b48199bdd5b9960521b60448201526064016108dc565b612706826126fa60005460ff1690565b6003919060ff166134ad565b610b0b5760405162461bcd60e51b8152602060048201526016602482015275125cdcdd595c8bd85c1c1c9bdd985b0819985a5b195960521b60448201526064016108dc565b60008261279a5760405162461bcd60e51b815260206004820152601860248201527f4e6f64652f726f6f742063616e6e6f74206265206e756c6c000000000000000060448201526064016108dc565b6000825167ffffffffffffffff8111156127c457634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156127ed578160200160208202803683370190505b50905060005b8351811015612b5757600084828151811061281e57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b0381166000908152600e90925260409091205490915060ff166128965760405162461bcd60e51b815260206004820152601b60248201527f4e6f64652f61646472657373206e6f7420617574686f72697a6564000000000060448201526064016108dc565b60006128a982635041329d60e01b6123f3565b9050806128c657634e487b7160e01b600052600160045260246000fd5b6000829050806001600160a01b03166399d1dbc06040518163ffffffff1660e01b815260040160206040518083038186803b15801561290457600080fd5b505afa158015612918573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061293c9190613a00565b15612ab6576040516303ce7bb760e11b81526001600160a01b038a8116600483015282169063079cf76e9060240160206040518083038186803b15801561298257600080fd5b505afa158015612996573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129ba9190613a32565b8585815181106129da57634e487b7160e01b600052603260045260246000fd5b602002602001018181525050806001600160a01b0316630ae070288a878781518110612a1657634e487b7160e01b600052603260045260246000fd5b60200260200101516040518363ffffffff1660e01b8152600401612a4f9291906001600160a01b03929092168252602082015260400190565b60206040518083038186803b158015612a6757600080fd5b505afa158015612a7b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a9f9190613a00565b612ab157600095505050505050610acb565b612b41565b60405163f936cc1560e01b81526001600160a01b038a8116600483015282169063f936cc159060240160206040518083038186803b158015612af757600080fd5b505afa158015612b0b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b2f9190613a00565b612b4157600095505050505050610acb565b5050508080612b4f90613fca565b9150506127f3565b50604051630b6d99fb60e11b815273__$9176519fc5fb44025266bd62b274a40525$__906316db33f690612b919087908590600401613d78565b60206040518083038186803b158015612ba957600080fd5b505af4158015612bbd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612be19190613a00565b95945050505050565b6000612bfd826301ffc9a760e01b612c1d565b80156107095750612c16826001600160e01b0319612c1d565b1592915050565b6000806000612c2c8585613665565b91509150818015612be15750949350505050565b6000805b8251811015612c9f57612c7f8585858481518110612c7257634e487b7160e01b600052603260045260246000fd5b602002602001015161240f565b612c8d576000915050610acb565b80612c9781613fca565b915050612c44565b506001949350505050565b6000828152600384016020526040902060010154612d0a5760405162461bcd60e51b815260206004820152601b60248201527f4e6f746172792f63726564656e7469616c206e6f7420666f756e64000000000060448201526064016108dc565b60008281526003840160205260409020600601546001600160a01b031680612d745760405162461bcd60e51b815260206004820152601d60248201527f4e6f746172792f7375626a6563742063616e6e6f74206265207a65726f00000060448201526064016108dc565b6001600160a01b038116600090815260058501602052604081208054909190612d9c90613fca565b9091555060408051608081018252338082526001600160a01b03848116602080850182815243868801818152606088018b815260008d815260048f0186528a9020985189549088166001600160a01b0319918216178a55935160018a0180549190981694169390931790955593516002870155516003909501949094558451918252928101869052909286917fb53d448559d4a1f97e1fd48262b64223f5bdafd54d429a33599236f73e900b1f910160405180910390a450505050565b6000838152600686016020908152604080832033845290915290205460ff1615612ec55760405162461bcd60e51b815260206004820152601c60248201527f4e6f746172792f73656e64657220616c7265616479207369676e65640000000060448201526064016108dc565b6000838152600386016020526040902060010154613235576001600160a01b038416600090815260018601602052604090205415613090576001600160a01b0384166000908152600180870160209081526040808420548452600389019091529091200154612f4457634e487b7160e01b600052600160045260246000fd5b6001600160a01b0380851660009081526001808801602090815260408084205484526003808b01835281852082516101408101845281548152948101548585015260028101548584015290810154606085015260048101546080850152600581015460ff8116151560a086015261010090819004871660c0860152600682015490961660e085015260078101805483518186028101860190945280845295969495919486019390919083018282801561302657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613008575b5050505050815260200160088201548152505090504381602001511061308e5760405162461bcd60e51b815260206004820152601d60248201527f4e6f746172792f626c6f636b206e756d6265722076696f6c6174696f6e00000060448201526064016108dc565b505b60405180610140016040528060018152602001438152602001428152602001866000016000876001600160a01b03166001600160a01b03168152602001908152602001600020600081546130e390613fca565b91829055508152602080820186905260006040808401829052336060808601919091526001600160a01b038a811660808088019190915260a08088018a905260c09788018b90528b865260038e81018852958590208951815589880151600182015594890151600286015592880151948401949094559286015160048301558501516005820180549587015160ff1990961691151591909117610100600160a81b031916610100958416860217905560e08501516006820180546001600160a01b03191691909316179091559183015180516131c59260078501920190613750565b5061012091909101516008909101556001600160a01b0384166000818152600187016020526040908190208590555133919085907fc6f6a1702b44006e35bd83045fd933510f3a924c1bb3877bff1061c14eb1c6be906132289043815260200190565b60405180910390a461344b565b60008381526003808701602090815260408084208151610140810183528154815260018201548185015260028201548184015293810154606085015260048101546080850152600581015460ff8116151560a08601526001600160a01b0361010091829004811660c087015260068301541660e086015260078201805484518187028101870190955280855292949186019392909183018282801561330357602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116132e5575b505050505081526020016008820154815250509050846001600160a01b03168160e001516001600160a01b03161461337d5760405162461bcd60e51b815260206004820181905260248201527f4e6f746172792f64696765737420616c7265616479207265676973746572656460448201526064016108dc565b82816101200151146133d15760405162461bcd60e51b815260206004820152601f60248201527f4e6f746172792f6d69736d6174636865642065766964656e636520726f6f740060448201526064016108dc565b815181610100015151146134275760405162461bcd60e51b815260206004820152601b60248201527f4e6f746172792f6d69736d617463686564207769746e6573736573000000000060448201526064016108dc565b60008481526003870160205260408120805490919061344590613fca565b90915550505b600083815260068601602090815260408083203380855290835292819020805460ff19166001179055514381528592917f3c4c49076c09b6a1214e1ba437ac70fd0cff1c6ea5baf6df4c5f6ba937c4bf12910160405180910390a35050505050565b60008281526003840160205260408120600601546001600160a01b03163381146135105760405162461bcd60e51b8152602060048201526014602482015273139bdd185c9e4bdddc9bdb99c81cdd589a9958dd60621b60448201526064016108dc565b600084815260038601602052604090206005015460ff16156135745760405162461bcd60e51b815260206004820181905260248201527f4e6f746172792f63726564656e7469616c20616c7265616479207369676e656460448201526064016108dc565b60008481526003860160205260409020548311156135d45760405162461bcd60e51b815260206004820152601e60248201527f4e6f746172792f6e6f2071756f72756d206f66207369676e617475726573000060448201526064016108dc565b60008481526003860160209081526040808320600501805460ff191660019081179091556001600160a01b0385168452600289018352818420805491820181558452919092200185905551849033907f3c4c49076c09b6a1214e1ba437ac70fd0cff1c6ea5baf6df4c5f6ba937c4bf12906136529043815260200190565b60405180910390a3506001949350505050565b604080516001600160e01b0319831660248083019190915282518083039091018152604490910182526020810180516001600160e01b03166301ffc9a760e01b17905290516000918291829081906001600160a01b03881690617530906136cd908690613cd5565b6000604051808303818686fa925050503d8060008114613709576040519150601f19603f3d011682016040523d82523d6000602084013e61370e565b606091505b509150915060208151101561372c5760008094509450505050613749565b81818060200190518101906137419190613a00565b945094505050505b9250929050565b8280548282559060005260206000209081019282156137a5579160200282015b828111156137a557825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613770565b506137b19291506137b5565b5090565b5b808211156137b157600081556001016137b6565b805161070c81614031565b600082601f8301126137e5578081fd5b813560206137fa6137f583613fa6565b613f75565b8281528181019085830183850287018401881015613816578586fd5b855b8581101561383d57813561382b81614031565b84529284019290840190600101613818565b5090979650505050505050565b600082601f83011261385a578081fd5b8151602061386a6137f583613fa6565b8281528181019085830183850287018401881015613886578586fd5b855b8581101561383d57815161389b81614031565b84529284019290840190600101613888565b8051801515811461070c57600080fd5b6000602082840312156138ce578081fd5b8135610acb81614031565b600080604083850312156138eb578081fd5b82356138f681614031565b915060208381013567ffffffffffffffff811115613912578283fd5b8401601f81018613613922578283fd5b80356139306137f582613fa6565b81815283810190838501858402850186018a101561394c578687fd5b8694505b8385101561396e578035835260019490940193918501918501613950565b5080955050505050509250929050565b60008060408385031215613990578182fd5b823561399b81614031565b946020939093013593505050565b6000806000606084860312156139bd578081fd5b83356139c881614031565b925060208401359150604084013567ffffffffffffffff8111156139ea578182fd5b6139f6868287016137d5565b9150509250925092565b600060208284031215613a11578081fd5b6108a4826138ad565b600060208284031215613a2b578081fd5b5035919050565b600060208284031215613a43578081fd5b5051919050565b60008060408385031215613a5c578182fd5b823591506020830135613a6e81614031565b809150509250929050565b60008060408385031215613a8b578182fd5b82359150602083013567ffffffffffffffff811115613aa8578182fd5b613ab4858286016137d5565b9150509250929050565b60008060408385031215613ad0578182fd5b50508035926020909101359150565b600060208284031215613af0578081fd5b81356001600160e01b031981168114610acb578182fd5b600060208284031215613b18578081fd5b815160028110610acb578182fd5b600060208284031215613b37578081fd5b815167ffffffffffffffff80821115613b4e578283fd5b8184019150610140808387031215613b64578384fd5b613b6d81613f75565b90508251815260208301516020820152604083015160408201526060830151606082015260808301516080820152613ba760a084016138ad565b60a0820152613bb860c084016137ca565b60c0820152613bc960e084016137ca565b60e08201526101008084015183811115613be1578586fd5b613bed8882870161384a565b91830191909152506101209283015192810192909252509392505050565b600060808284031215613c1c578081fd5b613c266080613f75565b8251613c3181614031565b81526020830151613c4181614031565b6020820152604083810151908201526060928301519281019290925250919050565b6000815180845260208085019450808401835b83811015613c9b5781516001600160a01b031687529582019590820190600101613c76565b509495945050505050565b6000815180845260208085019450808401835b83811015613c9b57815187529582019590820190600101613cb9565b60008251815b81811015613cf55760208186018101518583015201613cdb565b81811115613d035782828501525b509190910192915050565b6000602082526108a46020830184613c63565b6000602082526108a46020830184613ca6565b6020808252825182820181905260009190848201906040850190845b81811015613d6c57835183529284019291840191600101613d50565b50909695505050505050565b600083825260406020830152611f706040830184613ca6565b6020810160028310613db357634e487b7160e01b600052602160045260246000fd5b91905290565b6020808252601f908201527f4973737565722f746865726520617265206e6f2063726564656e7469616c7300604082015260600190565b6020808252601d908201527f4f776e6572732f73656e646572206973206e6f7420616e206f776e6572000000604082015260600190565b60208082526019908201527f4973737565722f63726564656e7469616c207265766f6b656400000000000000604082015260600190565b60006020825282516020830152602083015160408301526040830151606083015260608301516080830152608083015160a083015260a0830151613ea660c084018215159052565b5060c08301516001600160a01b03811660e08401525060e0830151610100613ed8818501836001600160a01b03169052565b808501519150506101406101208181860152613ef8610160860184613c63565b9501519301929092525090919050565b8381526001600160a01b0383166020820152606060408201819052600090612be190830184613ca6565b60006040820184835260206040818501528185548084526060860191508685528285209350845b8181101561383d57845483526001948501949284019201613f59565b604051601f8201601f1916810167ffffffffffffffff81118282101715613f9e57613f9e61401b565b604052919050565b600067ffffffffffffffff821115613fc057613fc061401b565b5060209081020190565b6000600019821415613fde57613fde614005565b5060010190565b600060ff821660ff811415613ffc57613ffc614005565b60010192915050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114611bef57600080fdfea2646970667358221220480cd05b9369ed671e507eaff98bbed09902abc0e87600691b0570bb046e558864736f6c63430008020033"

// DeployFaculty deploys a new Ethereum contract, binding an instance of Faculty to it.
func DeployFaculty(auth *bind.TransactOpts, backend bind.ContractBackend, owners []common.Address, quorum uint8) (common.Address, *types.Transaction, *Faculty, error) {
	parsed, err := abi.JSON(strings.NewReader(FacultyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	notaryAddr, _, _, _ := DeployNotary(auth, backend)
	FacultyBin = strings.Replace(FacultyBin, "__$1d2b6a72c2c8582bd594e19b02f34b62b6$__", notaryAddr.String()[2:], -1)

	credentialSumAddr, _, _, _ := DeployCredentialSum(auth, backend)
	FacultyBin = strings.Replace(FacultyBin, "__$9176519fc5fb44025266bd62b274a40525$__", credentialSumAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FacultyBin), backend, owners, quorum)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Faculty{FacultyCaller: FacultyCaller{contract: contract}, FacultyTransactor: FacultyTransactor{contract: contract}, FacultyFilterer: FacultyFilterer{contract: contract}}, nil
}

// Faculty is an auto generated Go binding around an Ethereum contract.
type Faculty struct {
	FacultyCaller     // Read-only binding to the contract
	FacultyTransactor // Write-only binding to the contract
	FacultyFilterer   // Log filterer for contract events
}

// FacultyCaller is an auto generated read-only Go binding around an Ethereum contract.
type FacultyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FacultyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FacultyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FacultyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FacultyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FacultySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FacultySession struct {
	Contract     *Faculty          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FacultyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FacultyCallerSession struct {
	Contract *FacultyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FacultyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FacultyTransactorSession struct {
	Contract     *FacultyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FacultyRaw is an auto generated low-level Go binding around an Ethereum contract.
type FacultyRaw struct {
	Contract *Faculty // Generic contract binding to access the raw methods on
}

// FacultyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FacultyCallerRaw struct {
	Contract *FacultyCaller // Generic read-only contract binding to access the raw methods on
}

// FacultyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FacultyTransactorRaw struct {
	Contract *FacultyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFaculty creates a new instance of Faculty, bound to a specific deployed contract.
func NewFaculty(address common.Address, backend bind.ContractBackend) (*Faculty, error) {
	contract, err := bindFaculty(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Faculty{FacultyCaller: FacultyCaller{contract: contract}, FacultyTransactor: FacultyTransactor{contract: contract}, FacultyFilterer: FacultyFilterer{contract: contract}}, nil
}

// NewFacultyCaller creates a new read-only instance of Faculty, bound to a specific deployed contract.
func NewFacultyCaller(address common.Address, caller bind.ContractCaller) (*FacultyCaller, error) {
	contract, err := bindFaculty(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FacultyCaller{contract: contract}, nil
}

// NewFacultyTransactor creates a new write-only instance of Faculty, bound to a specific deployed contract.
func NewFacultyTransactor(address common.Address, transactor bind.ContractTransactor) (*FacultyTransactor, error) {
	contract, err := bindFaculty(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FacultyTransactor{contract: contract}, nil
}

// NewFacultyFilterer creates a new log filterer instance of Faculty, bound to a specific deployed contract.
func NewFacultyFilterer(address common.Address, filterer bind.ContractFilterer) (*FacultyFilterer, error) {
	contract, err := bindFaculty(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FacultyFilterer{contract: contract}, nil
}

// bindFaculty binds a generic wrapper to an already deployed contract.
func bindFaculty(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FacultyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Faculty *FacultyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Faculty.Contract.FacultyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Faculty *FacultyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faculty.Contract.FacultyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Faculty *FacultyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Faculty.Contract.FacultyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Faculty *FacultyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Faculty.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Faculty *FacultyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faculty.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Faculty *FacultyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Faculty.Contract.contract.Transact(opts, method, params...)
}

// GetChildren is a free data retrieval call binding the contract method 0x4892e8e8.
//
// Solidity: function getChildren() view returns(address[])
func (_Faculty *FacultyCaller) GetChildren(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "getChildren")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetChildren is a free data retrieval call binding the contract method 0x4892e8e8.
//
// Solidity: function getChildren() view returns(address[])
func (_Faculty *FacultySession) GetChildren() ([]common.Address, error) {
	return _Faculty.Contract.GetChildren(&_Faculty.CallOpts)
}

// GetChildren is a free data retrieval call binding the contract method 0x4892e8e8.
//
// Solidity: function getChildren() view returns(address[])
func (_Faculty *FacultyCallerSession) GetChildren() ([]common.Address, error) {
	return _Faculty.Contract.GetChildren(&_Faculty.CallOpts)
}

// GetCoursesBySemester is a free data retrieval call binding the contract method 0xa355f4a5.
//
// Solidity: function getCoursesBySemester(bytes32 semester) view returns(address[])
func (_Faculty *FacultyCaller) GetCoursesBySemester(opts *bind.CallOpts, semester [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "getCoursesBySemester", semester)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCoursesBySemester is a free data retrieval call binding the contract method 0xa355f4a5.
//
// Solidity: function getCoursesBySemester(bytes32 semester) view returns(address[])
func (_Faculty *FacultySession) GetCoursesBySemester(semester [32]byte) ([]common.Address, error) {
	return _Faculty.Contract.GetCoursesBySemester(&_Faculty.CallOpts, semester)
}

// GetCoursesBySemester is a free data retrieval call binding the contract method 0xa355f4a5.
//
// Solidity: function getCoursesBySemester(bytes32 semester) view returns(address[])
func (_Faculty *FacultyCallerSession) GetCoursesBySemester(semester [32]byte) ([]common.Address, error) {
	return _Faculty.Contract.GetCoursesBySemester(&_Faculty.CallOpts, semester)
}

// GetCredentialProof is a free data retrieval call binding the contract method 0xdac46a62.
//
// Solidity: function getCredentialProof(bytes32 digest) view returns((uint256,uint256,uint256,uint256,bytes32,bool,address,address,address[],bytes32))
func (_Faculty *FacultyCaller) GetCredentialProof(opts *bind.CallOpts, digest [32]byte) (NotaryCredentialProof, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "getCredentialProof", digest)

	if err != nil {
		return *new(NotaryCredentialProof), err
	}

	out0 := *abi.ConvertType(out[0], new(NotaryCredentialProof)).(*NotaryCredentialProof)

	return out0, err

}

// GetCredentialProof is a free data retrieval call binding the contract method 0xdac46a62.
//
// Solidity: function getCredentialProof(bytes32 digest) view returns((uint256,uint256,uint256,uint256,bytes32,bool,address,address,address[],bytes32))
func (_Faculty *FacultySession) GetCredentialProof(digest [32]byte) (NotaryCredentialProof, error) {
	return _Faculty.Contract.GetCredentialProof(&_Faculty.CallOpts, digest)
}

// GetCredentialProof is a free data retrieval call binding the contract method 0xdac46a62.
//
// Solidity: function getCredentialProof(bytes32 digest) view returns((uint256,uint256,uint256,uint256,bytes32,bool,address,address,address[],bytes32))
func (_Faculty *FacultyCallerSession) GetCredentialProof(digest [32]byte) (NotaryCredentialProof, error) {
	return _Faculty.Contract.GetCredentialProof(&_Faculty.CallOpts, digest)
}

// GetCredentialSigners is a free data retrieval call binding the contract method 0x038ec1b8.
//
// Solidity: function getCredentialSigners(bytes32 digest) view returns(address[])
func (_Faculty *FacultyCaller) GetCredentialSigners(opts *bind.CallOpts, digest [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "getCredentialSigners", digest)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCredentialSigners is a free data retrieval call binding the contract method 0x038ec1b8.
//
// Solidity: function getCredentialSigners(bytes32 digest) view returns(address[])
func (_Faculty *FacultySession) GetCredentialSigners(digest [32]byte) ([]common.Address, error) {
	return _Faculty.Contract.GetCredentialSigners(&_Faculty.CallOpts, digest)
}

// GetCredentialSigners is a free data retrieval call binding the contract method 0x038ec1b8.
//
// Solidity: function getCredentialSigners(bytes32 digest) view returns(address[])
func (_Faculty *FacultyCallerSession) GetCredentialSigners(digest [32]byte) ([]common.Address, error) {
	return _Faculty.Contract.GetCredentialSigners(&_Faculty.CallOpts, digest)
}

// GetDigests is a free data retrieval call binding the contract method 0x762b77bb.
//
// Solidity: function getDigests(address subject) view returns(bytes32[])
func (_Faculty *FacultyCaller) GetDigests(opts *bind.CallOpts, subject common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "getDigests", subject)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetDigests is a free data retrieval call binding the contract method 0x762b77bb.
//
// Solidity: function getDigests(address subject) view returns(bytes32[])
func (_Faculty *FacultySession) GetDigests(subject common.Address) ([][32]byte, error) {
	return _Faculty.Contract.GetDigests(&_Faculty.CallOpts, subject)
}

// GetDigests is a free data retrieval call binding the contract method 0x762b77bb.
//
// Solidity: function getDigests(address subject) view returns(bytes32[])
func (_Faculty *FacultyCallerSession) GetDigests(subject common.Address) ([][32]byte, error) {
	return _Faculty.Contract.GetDigests(&_Faculty.CallOpts, subject)
}

// GetEvidenceRoot is a free data retrieval call binding the contract method 0x4302802d.
//
// Solidity: function getEvidenceRoot(bytes32 digest) view returns(bytes32)
func (_Faculty *FacultyCaller) GetEvidenceRoot(opts *bind.CallOpts, digest [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "getEvidenceRoot", digest)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEvidenceRoot is a free data retrieval call binding the contract method 0x4302802d.
//
// Solidity: function getEvidenceRoot(bytes32 digest) view returns(bytes32)
func (_Faculty *FacultySession) GetEvidenceRoot(digest [32]byte) ([32]byte, error) {
	return _Faculty.Contract.GetEvidenceRoot(&_Faculty.CallOpts, digest)
}

// GetEvidenceRoot is a free data retrieval call binding the contract method 0x4302802d.
//
// Solidity: function getEvidenceRoot(bytes32 digest) view returns(bytes32)
func (_Faculty *FacultyCallerSession) GetEvidenceRoot(digest [32]byte) ([32]byte, error) {
	return _Faculty.Contract.GetEvidenceRoot(&_Faculty.CallOpts, digest)
}

// GetProof is a free data retrieval call binding the contract method 0x3eea79d1.
//
// Solidity: function getProof(address subject) view returns((bytes32,uint256,uint256))
func (_Faculty *FacultyCaller) GetProof(opts *bind.CallOpts, subject common.Address) (CredentialSumRoot, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "getProof", subject)

	if err != nil {
		return *new(CredentialSumRoot), err
	}

	out0 := *abi.ConvertType(out[0], new(CredentialSumRoot)).(*CredentialSumRoot)

	return out0, err

}

// GetProof is a free data retrieval call binding the contract method 0x3eea79d1.
//
// Solidity: function getProof(address subject) view returns((bytes32,uint256,uint256))
func (_Faculty *FacultySession) GetProof(subject common.Address) (CredentialSumRoot, error) {
	return _Faculty.Contract.GetProof(&_Faculty.CallOpts, subject)
}

// GetProof is a free data retrieval call binding the contract method 0x3eea79d1.
//
// Solidity: function getProof(address subject) view returns((bytes32,uint256,uint256))
func (_Faculty *FacultyCallerSession) GetProof(subject common.Address) (CredentialSumRoot, error) {
	return _Faculty.Contract.GetProof(&_Faculty.CallOpts, subject)
}

// GetRevoked is a free data retrieval call binding the contract method 0x8ebd9623.
//
// Solidity: function getRevoked(address subject) view returns(bytes32[])
func (_Faculty *FacultyCaller) GetRevoked(opts *bind.CallOpts, subject common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "getRevoked", subject)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetRevoked is a free data retrieval call binding the contract method 0x8ebd9623.
//
// Solidity: function getRevoked(address subject) view returns(bytes32[])
func (_Faculty *FacultySession) GetRevoked(subject common.Address) ([][32]byte, error) {
	return _Faculty.Contract.GetRevoked(&_Faculty.CallOpts, subject)
}

// GetRevoked is a free data retrieval call binding the contract method 0x8ebd9623.
//
// Solidity: function getRevoked(address subject) view returns(bytes32[])
func (_Faculty *FacultyCallerSession) GetRevoked(subject common.Address) ([][32]byte, error) {
	return _Faculty.Contract.GetRevoked(&_Faculty.CallOpts, subject)
}

// GetRevokedProof is a free data retrieval call binding the contract method 0xabbff984.
//
// Solidity: function getRevokedProof(bytes32 digest) view returns((address,address,uint256,bytes32))
func (_Faculty *FacultyCaller) GetRevokedProof(opts *bind.CallOpts, digest [32]byte) (NotaryRevocationProof, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "getRevokedProof", digest)

	if err != nil {
		return *new(NotaryRevocationProof), err
	}

	out0 := *abi.ConvertType(out[0], new(NotaryRevocationProof)).(*NotaryRevocationProof)

	return out0, err

}

// GetRevokedProof is a free data retrieval call binding the contract method 0xabbff984.
//
// Solidity: function getRevokedProof(bytes32 digest) view returns((address,address,uint256,bytes32))
func (_Faculty *FacultySession) GetRevokedProof(digest [32]byte) (NotaryRevocationProof, error) {
	return _Faculty.Contract.GetRevokedProof(&_Faculty.CallOpts, digest)
}

// GetRevokedProof is a free data retrieval call binding the contract method 0xabbff984.
//
// Solidity: function getRevokedProof(bytes32 digest) view returns((address,address,uint256,bytes32))
func (_Faculty *FacultyCallerSession) GetRevokedProof(digest [32]byte) (NotaryRevocationProof, error) {
	return _Faculty.Contract.GetRevokedProof(&_Faculty.CallOpts, digest)
}

// GetRole is a free data retrieval call binding the contract method 0xfde08174.
//
// Solidity: function getRole() view returns(uint8)
func (_Faculty *FacultyCaller) GetRole(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "getRole")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetRole is a free data retrieval call binding the contract method 0xfde08174.
//
// Solidity: function getRole() view returns(uint8)
func (_Faculty *FacultySession) GetRole() (uint8, error) {
	return _Faculty.Contract.GetRole(&_Faculty.CallOpts)
}

// GetRole is a free data retrieval call binding the contract method 0xfde08174.
//
// Solidity: function getRole() view returns(uint8)
func (_Faculty *FacultyCallerSession) GetRole() (uint8, error) {
	return _Faculty.Contract.GetRole(&_Faculty.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x079cf76e.
//
// Solidity: function getRoot(address subject) view returns(bytes32)
func (_Faculty *FacultyCaller) GetRoot(opts *bind.CallOpts, subject common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "getRoot", subject)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x079cf76e.
//
// Solidity: function getRoot(address subject) view returns(bytes32)
func (_Faculty *FacultySession) GetRoot(subject common.Address) ([32]byte, error) {
	return _Faculty.Contract.GetRoot(&_Faculty.CallOpts, subject)
}

// GetRoot is a free data retrieval call binding the contract method 0x079cf76e.
//
// Solidity: function getRoot(address subject) view returns(bytes32)
func (_Faculty *FacultyCallerSession) GetRoot(subject common.Address) ([32]byte, error) {
	return _Faculty.Contract.GetRoot(&_Faculty.CallOpts, subject)
}

// GetWitnesses is a free data retrieval call binding the contract method 0x6f28eca6.
//
// Solidity: function getWitnesses(bytes32 digest) view returns(address[])
func (_Faculty *FacultyCaller) GetWitnesses(opts *bind.CallOpts, digest [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "getWitnesses", digest)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWitnesses is a free data retrieval call binding the contract method 0x6f28eca6.
//
// Solidity: function getWitnesses(bytes32 digest) view returns(address[])
func (_Faculty *FacultySession) GetWitnesses(digest [32]byte) ([]common.Address, error) {
	return _Faculty.Contract.GetWitnesses(&_Faculty.CallOpts, digest)
}

// GetWitnesses is a free data retrieval call binding the contract method 0x6f28eca6.
//
// Solidity: function getWitnesses(bytes32 digest) view returns(address[])
func (_Faculty *FacultyCallerSession) GetWitnesses(digest [32]byte) ([]common.Address, error) {
	return _Faculty.Contract.GetWitnesses(&_Faculty.CallOpts, digest)
}

// HasRoot is a free data retrieval call binding the contract method 0x7532ed26.
//
// Solidity: function hasRoot(address subject) view returns(bool)
func (_Faculty *FacultyCaller) HasRoot(opts *bind.CallOpts, subject common.Address) (bool, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "hasRoot", subject)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRoot is a free data retrieval call binding the contract method 0x7532ed26.
//
// Solidity: function hasRoot(address subject) view returns(bool)
func (_Faculty *FacultySession) HasRoot(subject common.Address) (bool, error) {
	return _Faculty.Contract.HasRoot(&_Faculty.CallOpts, subject)
}

// HasRoot is a free data retrieval call binding the contract method 0x7532ed26.
//
// Solidity: function hasRoot(address subject) view returns(bool)
func (_Faculty *FacultyCallerSession) HasRoot(subject common.Address) (bool, error) {
	return _Faculty.Contract.HasRoot(&_Faculty.CallOpts, subject)
}

// IsApproved is a free data retrieval call binding the contract method 0x48aefc32.
//
// Solidity: function isApproved(bytes32 digest) view returns(bool)
func (_Faculty *FacultyCaller) IsApproved(opts *bind.CallOpts, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "isApproved", digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0x48aefc32.
//
// Solidity: function isApproved(bytes32 digest) view returns(bool)
func (_Faculty *FacultySession) IsApproved(digest [32]byte) (bool, error) {
	return _Faculty.Contract.IsApproved(&_Faculty.CallOpts, digest)
}

// IsApproved is a free data retrieval call binding the contract method 0x48aefc32.
//
// Solidity: function isApproved(bytes32 digest) view returns(bool)
func (_Faculty *FacultyCallerSession) IsApproved(digest [32]byte) (bool, error) {
	return _Faculty.Contract.IsApproved(&_Faculty.CallOpts, digest)
}

// IsChild is a free data retrieval call binding the contract method 0xfc91a897.
//
// Solidity: function isChild(address ) view returns(bool)
func (_Faculty *FacultyCaller) IsChild(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "isChild", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChild is a free data retrieval call binding the contract method 0xfc91a897.
//
// Solidity: function isChild(address ) view returns(bool)
func (_Faculty *FacultySession) IsChild(arg0 common.Address) (bool, error) {
	return _Faculty.Contract.IsChild(&_Faculty.CallOpts, arg0)
}

// IsChild is a free data retrieval call binding the contract method 0xfc91a897.
//
// Solidity: function isChild(address ) view returns(bool)
func (_Faculty *FacultyCallerSession) IsChild(arg0 common.Address) (bool, error) {
	return _Faculty.Contract.IsChild(&_Faculty.CallOpts, arg0)
}

// IsLeaf is a free data retrieval call binding the contract method 0x99d1dbc0.
//
// Solidity: function isLeaf() view returns(bool)
func (_Faculty *FacultyCaller) IsLeaf(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "isLeaf")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLeaf is a free data retrieval call binding the contract method 0x99d1dbc0.
//
// Solidity: function isLeaf() view returns(bool)
func (_Faculty *FacultySession) IsLeaf() (bool, error) {
	return _Faculty.Contract.IsLeaf(&_Faculty.CallOpts)
}

// IsLeaf is a free data retrieval call binding the contract method 0x99d1dbc0.
//
// Solidity: function isLeaf() view returns(bool)
func (_Faculty *FacultyCallerSession) IsLeaf() (bool, error) {
	return _Faculty.Contract.IsLeaf(&_Faculty.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_Faculty *FacultyCaller) IsOwner(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "isOwner", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_Faculty *FacultySession) IsOwner(account common.Address) (bool, error) {
	return _Faculty.Contract.IsOwner(&_Faculty.CallOpts, account)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_Faculty *FacultyCallerSession) IsOwner(account common.Address) (bool, error) {
	return _Faculty.Contract.IsOwner(&_Faculty.CallOpts, account)
}

// IsQuorumSigned is a free data retrieval call binding the contract method 0xae02c0cc.
//
// Solidity: function isQuorumSigned(bytes32 digest) view returns(bool)
func (_Faculty *FacultyCaller) IsQuorumSigned(opts *bind.CallOpts, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "isQuorumSigned", digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsQuorumSigned is a free data retrieval call binding the contract method 0xae02c0cc.
//
// Solidity: function isQuorumSigned(bytes32 digest) view returns(bool)
func (_Faculty *FacultySession) IsQuorumSigned(digest [32]byte) (bool, error) {
	return _Faculty.Contract.IsQuorumSigned(&_Faculty.CallOpts, digest)
}

// IsQuorumSigned is a free data retrieval call binding the contract method 0xae02c0cc.
//
// Solidity: function isQuorumSigned(bytes32 digest) view returns(bool)
func (_Faculty *FacultyCallerSession) IsQuorumSigned(digest [32]byte) (bool, error) {
	return _Faculty.Contract.IsQuorumSigned(&_Faculty.CallOpts, digest)
}

// IsRevoked is a free data retrieval call binding the contract method 0x4294857f.
//
// Solidity: function isRevoked(bytes32 digest) view returns(bool)
func (_Faculty *FacultyCaller) IsRevoked(opts *bind.CallOpts, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "isRevoked", digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRevoked is a free data retrieval call binding the contract method 0x4294857f.
//
// Solidity: function isRevoked(bytes32 digest) view returns(bool)
func (_Faculty *FacultySession) IsRevoked(digest [32]byte) (bool, error) {
	return _Faculty.Contract.IsRevoked(&_Faculty.CallOpts, digest)
}

// IsRevoked is a free data retrieval call binding the contract method 0x4294857f.
//
// Solidity: function isRevoked(bytes32 digest) view returns(bool)
func (_Faculty *FacultyCallerSession) IsRevoked(digest [32]byte) (bool, error) {
	return _Faculty.Contract.IsRevoked(&_Faculty.CallOpts, digest)
}

// IsSigned is a free data retrieval call binding the contract method 0x7f8d61e0.
//
// Solidity: function isSigned(bytes32 digest, address account) view returns(bool)
func (_Faculty *FacultyCaller) IsSigned(opts *bind.CallOpts, digest [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "isSigned", digest, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigned is a free data retrieval call binding the contract method 0x7f8d61e0.
//
// Solidity: function isSigned(bytes32 digest, address account) view returns(bool)
func (_Faculty *FacultySession) IsSigned(digest [32]byte, account common.Address) (bool, error) {
	return _Faculty.Contract.IsSigned(&_Faculty.CallOpts, digest, account)
}

// IsSigned is a free data retrieval call binding the contract method 0x7f8d61e0.
//
// Solidity: function isSigned(bytes32 digest, address account) view returns(bool)
func (_Faculty *FacultyCallerSession) IsSigned(digest [32]byte, account common.Address) (bool, error) {
	return _Faculty.Contract.IsSigned(&_Faculty.CallOpts, digest, account)
}

// MyParent is a free data retrieval call binding the contract method 0xde07fc45.
//
// Solidity: function myParent() view returns(address)
func (_Faculty *FacultyCaller) MyParent(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "myParent")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MyParent is a free data retrieval call binding the contract method 0xde07fc45.
//
// Solidity: function myParent() view returns(address)
func (_Faculty *FacultySession) MyParent() (common.Address, error) {
	return _Faculty.Contract.MyParent(&_Faculty.CallOpts)
}

// MyParent is a free data retrieval call binding the contract method 0xde07fc45.
//
// Solidity: function myParent() view returns(address)
func (_Faculty *FacultyCallerSession) MyParent() (common.Address, error) {
	return _Faculty.Contract.MyParent(&_Faculty.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address[])
func (_Faculty *FacultyCaller) Owners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "owners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address[])
func (_Faculty *FacultySession) Owners() ([]common.Address, error) {
	return _Faculty.Contract.Owners(&_Faculty.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address[])
func (_Faculty *FacultyCallerSession) Owners() ([]common.Address, error) {
	return _Faculty.Contract.Owners(&_Faculty.CallOpts)
}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() view returns(uint8)
func (_Faculty *FacultyCaller) OwnersCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "ownersCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() view returns(uint8)
func (_Faculty *FacultySession) OwnersCount() (uint8, error) {
	return _Faculty.Contract.OwnersCount(&_Faculty.CallOpts)
}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() view returns(uint8)
func (_Faculty *FacultyCallerSession) OwnersCount() (uint8, error) {
	return _Faculty.Contract.OwnersCount(&_Faculty.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint8)
func (_Faculty *FacultyCaller) Quorum(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "quorum")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint8)
func (_Faculty *FacultySession) Quorum() (uint8, error) {
	return _Faculty.Contract.Quorum(&_Faculty.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint8)
func (_Faculty *FacultyCallerSession) Quorum() (uint8, error) {
	return _Faculty.Contract.Quorum(&_Faculty.CallOpts)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 digest) view returns(bool)
func (_Faculty *FacultyCaller) RecordExists(opts *bind.CallOpts, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "recordExists", digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 digest) view returns(bool)
func (_Faculty *FacultySession) RecordExists(digest [32]byte) (bool, error) {
	return _Faculty.Contract.RecordExists(&_Faculty.CallOpts, digest)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 digest) view returns(bool)
func (_Faculty *FacultyCallerSession) RecordExists(digest [32]byte) (bool, error) {
	return _Faculty.Contract.RecordExists(&_Faculty.CallOpts, digest)
}

// RevokedCounter is a free data retrieval call binding the contract method 0x342577fa.
//
// Solidity: function revokedCounter(address subject) view returns(uint256)
func (_Faculty *FacultyCaller) RevokedCounter(opts *bind.CallOpts, subject common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "revokedCounter", subject)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RevokedCounter is a free data retrieval call binding the contract method 0x342577fa.
//
// Solidity: function revokedCounter(address subject) view returns(uint256)
func (_Faculty *FacultySession) RevokedCounter(subject common.Address) (*big.Int, error) {
	return _Faculty.Contract.RevokedCounter(&_Faculty.CallOpts, subject)
}

// RevokedCounter is a free data retrieval call binding the contract method 0x342577fa.
//
// Solidity: function revokedCounter(address subject) view returns(uint256)
func (_Faculty *FacultyCallerSession) RevokedCounter(subject common.Address) (*big.Int, error) {
	return _Faculty.Contract.RevokedCounter(&_Faculty.CallOpts, subject)
}

// SemesterExists is a free data retrieval call binding the contract method 0x2adfe8b2.
//
// Solidity: function semesterExists(bytes32 semester) view returns(bool)
func (_Faculty *FacultyCaller) SemesterExists(opts *bind.CallOpts, semester [32]byte) (bool, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "semesterExists", semester)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SemesterExists is a free data retrieval call binding the contract method 0x2adfe8b2.
//
// Solidity: function semesterExists(bytes32 semester) view returns(bool)
func (_Faculty *FacultySession) SemesterExists(semester [32]byte) (bool, error) {
	return _Faculty.Contract.SemesterExists(&_Faculty.CallOpts, semester)
}

// SemesterExists is a free data retrieval call binding the contract method 0x2adfe8b2.
//
// Solidity: function semesterExists(bytes32 semester) view returns(bool)
func (_Faculty *FacultyCallerSession) SemesterExists(semester [32]byte) (bool, error) {
	return _Faculty.Contract.SemesterExists(&_Faculty.CallOpts, semester)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Faculty *FacultyCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Faculty *FacultySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Faculty.Contract.SupportsInterface(&_Faculty.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Faculty *FacultyCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Faculty.Contract.SupportsInterface(&_Faculty.CallOpts, interfaceId)
}

// VerifyCredential is a free data retrieval call binding the contract method 0x5f889e17.
//
// Solidity: function verifyCredential(address subject, bytes32 digest) view returns(bool)
func (_Faculty *FacultyCaller) VerifyCredential(opts *bind.CallOpts, subject common.Address, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "verifyCredential", subject, digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyCredential is a free data retrieval call binding the contract method 0x5f889e17.
//
// Solidity: function verifyCredential(address subject, bytes32 digest) view returns(bool)
func (_Faculty *FacultySession) VerifyCredential(subject common.Address, digest [32]byte) (bool, error) {
	return _Faculty.Contract.VerifyCredential(&_Faculty.CallOpts, subject, digest)
}

// VerifyCredential is a free data retrieval call binding the contract method 0x5f889e17.
//
// Solidity: function verifyCredential(address subject, bytes32 digest) view returns(bool)
func (_Faculty *FacultyCallerSession) VerifyCredential(subject common.Address, digest [32]byte) (bool, error) {
	return _Faculty.Contract.VerifyCredential(&_Faculty.CallOpts, subject, digest)
}

// VerifyCredentialRoot is a free data retrieval call binding the contract method 0x0ae07028.
//
// Solidity: function verifyCredentialRoot(address subject, bytes32 root) view returns(bool)
func (_Faculty *FacultyCaller) VerifyCredentialRoot(opts *bind.CallOpts, subject common.Address, root [32]byte) (bool, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "verifyCredentialRoot", subject, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyCredentialRoot is a free data retrieval call binding the contract method 0x0ae07028.
//
// Solidity: function verifyCredentialRoot(address subject, bytes32 root) view returns(bool)
func (_Faculty *FacultySession) VerifyCredentialRoot(subject common.Address, root [32]byte) (bool, error) {
	return _Faculty.Contract.VerifyCredentialRoot(&_Faculty.CallOpts, subject, root)
}

// VerifyCredentialRoot is a free data retrieval call binding the contract method 0x0ae07028.
//
// Solidity: function verifyCredentialRoot(address subject, bytes32 root) view returns(bool)
func (_Faculty *FacultyCallerSession) VerifyCredentialRoot(subject common.Address, root [32]byte) (bool, error) {
	return _Faculty.Contract.VerifyCredentialRoot(&_Faculty.CallOpts, subject, root)
}

// VerifyCredentialTree is a free data retrieval call binding the contract method 0xf936cc15.
//
// Solidity: function verifyCredentialTree(address subject) view returns(bool)
func (_Faculty *FacultyCaller) VerifyCredentialTree(opts *bind.CallOpts, subject common.Address) (bool, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "verifyCredentialTree", subject)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyCredentialTree is a free data retrieval call binding the contract method 0xf936cc15.
//
// Solidity: function verifyCredentialTree(address subject) view returns(bool)
func (_Faculty *FacultySession) VerifyCredentialTree(subject common.Address) (bool, error) {
	return _Faculty.Contract.VerifyCredentialTree(&_Faculty.CallOpts, subject)
}

// VerifyCredentialTree is a free data retrieval call binding the contract method 0xf936cc15.
//
// Solidity: function verifyCredentialTree(address subject) view returns(bool)
func (_Faculty *FacultyCallerSession) VerifyCredentialTree(subject common.Address) (bool, error) {
	return _Faculty.Contract.VerifyCredentialTree(&_Faculty.CallOpts, subject)
}

// VerifyIssuedCredentials is a free data retrieval call binding the contract method 0x2cf0695a.
//
// Solidity: function verifyIssuedCredentials(address subject) view returns(bool)
func (_Faculty *FacultyCaller) VerifyIssuedCredentials(opts *bind.CallOpts, subject common.Address) (bool, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "verifyIssuedCredentials", subject)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyIssuedCredentials is a free data retrieval call binding the contract method 0x2cf0695a.
//
// Solidity: function verifyIssuedCredentials(address subject) view returns(bool)
func (_Faculty *FacultySession) VerifyIssuedCredentials(subject common.Address) (bool, error) {
	return _Faculty.Contract.VerifyIssuedCredentials(&_Faculty.CallOpts, subject)
}

// VerifyIssuedCredentials is a free data retrieval call binding the contract method 0x2cf0695a.
//
// Solidity: function verifyIssuedCredentials(address subject) view returns(bool)
func (_Faculty *FacultyCallerSession) VerifyIssuedCredentials(subject common.Address) (bool, error) {
	return _Faculty.Contract.VerifyIssuedCredentials(&_Faculty.CallOpts, subject)
}

// VerifyRootOf is a free data retrieval call binding the contract method 0x86509dd6.
//
// Solidity: function verifyRootOf(address subject, bytes32[] digests) view returns(bool)
func (_Faculty *FacultyCaller) VerifyRootOf(opts *bind.CallOpts, subject common.Address, digests [][32]byte) (bool, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "verifyRootOf", subject, digests)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyRootOf is a free data retrieval call binding the contract method 0x86509dd6.
//
// Solidity: function verifyRootOf(address subject, bytes32[] digests) view returns(bool)
func (_Faculty *FacultySession) VerifyRootOf(subject common.Address, digests [][32]byte) (bool, error) {
	return _Faculty.Contract.VerifyRootOf(&_Faculty.CallOpts, subject, digests)
}

// VerifyRootOf is a free data retrieval call binding the contract method 0x86509dd6.
//
// Solidity: function verifyRootOf(address subject, bytes32[] digests) view returns(bool)
func (_Faculty *FacultyCallerSession) VerifyRootOf(subject common.Address, digests [][32]byte) (bool, error) {
	return _Faculty.Contract.VerifyRootOf(&_Faculty.CallOpts, subject, digests)
}

// WitnessesLength is a free data retrieval call binding the contract method 0xf537ff9d.
//
// Solidity: function witnessesLength(bytes32 digest) view returns(uint256)
func (_Faculty *FacultyCaller) WitnessesLength(opts *bind.CallOpts, digest [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Faculty.contract.Call(opts, &out, "witnessesLength", digest)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WitnessesLength is a free data retrieval call binding the contract method 0xf537ff9d.
//
// Solidity: function witnessesLength(bytes32 digest) view returns(uint256)
func (_Faculty *FacultySession) WitnessesLength(digest [32]byte) (*big.Int, error) {
	return _Faculty.Contract.WitnessesLength(&_Faculty.CallOpts, digest)
}

// WitnessesLength is a free data retrieval call binding the contract method 0xf537ff9d.
//
// Solidity: function witnessesLength(bytes32 digest) view returns(uint256)
func (_Faculty *FacultyCallerSession) WitnessesLength(digest [32]byte) (*big.Int, error) {
	return _Faculty.Contract.WitnessesLength(&_Faculty.CallOpts, digest)
}

// AddChild is a paid mutator transaction binding the contract method 0x1eee993a.
//
// Solidity: function addChild(address nodeAddress) returns()
func (_Faculty *FacultyTransactor) AddChild(opts *bind.TransactOpts, nodeAddress common.Address) (*types.Transaction, error) {
	return _Faculty.contract.Transact(opts, "addChild", nodeAddress)
}

// AddChild is a paid mutator transaction binding the contract method 0x1eee993a.
//
// Solidity: function addChild(address nodeAddress) returns()
func (_Faculty *FacultySession) AddChild(nodeAddress common.Address) (*types.Transaction, error) {
	return _Faculty.Contract.AddChild(&_Faculty.TransactOpts, nodeAddress)
}

// AddChild is a paid mutator transaction binding the contract method 0x1eee993a.
//
// Solidity: function addChild(address nodeAddress) returns()
func (_Faculty *FacultyTransactorSession) AddChild(nodeAddress common.Address) (*types.Transaction, error) {
	return _Faculty.Contract.AddChild(&_Faculty.TransactOpts, nodeAddress)
}

// AggregateCredentials is a paid mutator transaction binding the contract method 0xe2d02cea.
//
// Solidity: function aggregateCredentials(address subject, bytes32[] digests) returns(bytes32)
func (_Faculty *FacultyTransactor) AggregateCredentials(opts *bind.TransactOpts, subject common.Address, digests [][32]byte) (*types.Transaction, error) {
	return _Faculty.contract.Transact(opts, "aggregateCredentials", subject, digests)
}

// AggregateCredentials is a paid mutator transaction binding the contract method 0xe2d02cea.
//
// Solidity: function aggregateCredentials(address subject, bytes32[] digests) returns(bytes32)
func (_Faculty *FacultySession) AggregateCredentials(subject common.Address, digests [][32]byte) (*types.Transaction, error) {
	return _Faculty.Contract.AggregateCredentials(&_Faculty.TransactOpts, subject, digests)
}

// AggregateCredentials is a paid mutator transaction binding the contract method 0xe2d02cea.
//
// Solidity: function aggregateCredentials(address subject, bytes32[] digests) returns(bytes32)
func (_Faculty *FacultyTransactorSession) AggregateCredentials(subject common.Address, digests [][32]byte) (*types.Transaction, error) {
	return _Faculty.Contract.AggregateCredentials(&_Faculty.TransactOpts, subject, digests)
}

// ApproveCredential is a paid mutator transaction binding the contract method 0xeedb973c.
//
// Solidity: function approveCredential(bytes32 digest) returns()
func (_Faculty *FacultyTransactor) ApproveCredential(opts *bind.TransactOpts, digest [32]byte) (*types.Transaction, error) {
	return _Faculty.contract.Transact(opts, "approveCredential", digest)
}

// ApproveCredential is a paid mutator transaction binding the contract method 0xeedb973c.
//
// Solidity: function approveCredential(bytes32 digest) returns()
func (_Faculty *FacultySession) ApproveCredential(digest [32]byte) (*types.Transaction, error) {
	return _Faculty.Contract.ApproveCredential(&_Faculty.TransactOpts, digest)
}

// ApproveCredential is a paid mutator transaction binding the contract method 0xeedb973c.
//
// Solidity: function approveCredential(bytes32 digest) returns()
func (_Faculty *FacultyTransactorSession) ApproveCredential(digest [32]byte) (*types.Transaction, error) {
	return _Faculty.Contract.ApproveCredential(&_Faculty.TransactOpts, digest)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Faculty *FacultyTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Faculty.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Faculty *FacultySession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Faculty.Contract.ChangeOwner(&_Faculty.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Faculty *FacultyTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Faculty.Contract.ChangeOwner(&_Faculty.TransactOpts, newOwner)
}

// RegisterCredential is a paid mutator transaction binding the contract method 0x3ac82e19.
//
// Solidity: function registerCredential(address subject, bytes32 digest, address[] witnesses) returns()
func (_Faculty *FacultyTransactor) RegisterCredential(opts *bind.TransactOpts, subject common.Address, digest [32]byte, witnesses []common.Address) (*types.Transaction, error) {
	return _Faculty.contract.Transact(opts, "registerCredential", subject, digest, witnesses)
}

// RegisterCredential is a paid mutator transaction binding the contract method 0x3ac82e19.
//
// Solidity: function registerCredential(address subject, bytes32 digest, address[] witnesses) returns()
func (_Faculty *FacultySession) RegisterCredential(subject common.Address, digest [32]byte, witnesses []common.Address) (*types.Transaction, error) {
	return _Faculty.Contract.RegisterCredential(&_Faculty.TransactOpts, subject, digest, witnesses)
}

// RegisterCredential is a paid mutator transaction binding the contract method 0x3ac82e19.
//
// Solidity: function registerCredential(address subject, bytes32 digest, address[] witnesses) returns()
func (_Faculty *FacultyTransactorSession) RegisterCredential(subject common.Address, digest [32]byte, witnesses []common.Address) (*types.Transaction, error) {
	return _Faculty.Contract.RegisterCredential(&_Faculty.TransactOpts, subject, digest, witnesses)
}

// RegisterSemester is a paid mutator transaction binding the contract method 0xc8bf608a.
//
// Solidity: function registerSemester(bytes32 semester, address[] courses) returns()
func (_Faculty *FacultyTransactor) RegisterSemester(opts *bind.TransactOpts, semester [32]byte, courses []common.Address) (*types.Transaction, error) {
	return _Faculty.contract.Transact(opts, "registerSemester", semester, courses)
}

// RegisterSemester is a paid mutator transaction binding the contract method 0xc8bf608a.
//
// Solidity: function registerSemester(bytes32 semester, address[] courses) returns()
func (_Faculty *FacultySession) RegisterSemester(semester [32]byte, courses []common.Address) (*types.Transaction, error) {
	return _Faculty.Contract.RegisterSemester(&_Faculty.TransactOpts, semester, courses)
}

// RegisterSemester is a paid mutator transaction binding the contract method 0xc8bf608a.
//
// Solidity: function registerSemester(bytes32 semester, address[] courses) returns()
func (_Faculty *FacultyTransactorSession) RegisterSemester(semester [32]byte, courses []common.Address) (*types.Transaction, error) {
	return _Faculty.Contract.RegisterSemester(&_Faculty.TransactOpts, semester, courses)
}

// RevokeCredential is a paid mutator transaction binding the contract method 0x3610d3ca.
//
// Solidity: function revokeCredential(bytes32 digest, bytes32 reason) returns()
func (_Faculty *FacultyTransactor) RevokeCredential(opts *bind.TransactOpts, digest [32]byte, reason [32]byte) (*types.Transaction, error) {
	return _Faculty.contract.Transact(opts, "revokeCredential", digest, reason)
}

// RevokeCredential is a paid mutator transaction binding the contract method 0x3610d3ca.
//
// Solidity: function revokeCredential(bytes32 digest, bytes32 reason) returns()
func (_Faculty *FacultySession) RevokeCredential(digest [32]byte, reason [32]byte) (*types.Transaction, error) {
	return _Faculty.Contract.RevokeCredential(&_Faculty.TransactOpts, digest, reason)
}

// RevokeCredential is a paid mutator transaction binding the contract method 0x3610d3ca.
//
// Solidity: function revokeCredential(bytes32 digest, bytes32 reason) returns()
func (_Faculty *FacultyTransactorSession) RevokeCredential(digest [32]byte, reason [32]byte) (*types.Transaction, error) {
	return _Faculty.Contract.RevokeCredential(&_Faculty.TransactOpts, digest, reason)
}

// FacultyCredentialIssuedIterator is returned from FilterCredentialIssued and is used to iterate over the raw logs and unpacked data for CredentialIssued events raised by the Faculty contract.
type FacultyCredentialIssuedIterator struct {
	Event *FacultyCredentialIssued // Event containing the contract specifics and raw log

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
func (it *FacultyCredentialIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FacultyCredentialIssued)
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
		it.Event = new(FacultyCredentialIssued)
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
func (it *FacultyCredentialIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FacultyCredentialIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FacultyCredentialIssued represents a CredentialIssued event raised by the Faculty contract.
type FacultyCredentialIssued struct {
	Digest        [32]byte
	Subject       common.Address
	Registrar     common.Address
	InsertedBlock *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCredentialIssued is a free log retrieval operation binding the contract event 0xc6f6a1702b44006e35bd83045fd933510f3a924c1bb3877bff1061c14eb1c6be.
//
// Solidity: event CredentialIssued(bytes32 indexed digest, address indexed subject, address indexed registrar, uint256 insertedBlock)
func (_Faculty *FacultyFilterer) FilterCredentialIssued(opts *bind.FilterOpts, digest [][32]byte, subject []common.Address, registrar []common.Address) (*FacultyCredentialIssuedIterator, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}

	logs, sub, err := _Faculty.contract.FilterLogs(opts, "CredentialIssued", digestRule, subjectRule, registrarRule)
	if err != nil {
		return nil, err
	}
	return &FacultyCredentialIssuedIterator{contract: _Faculty.contract, event: "CredentialIssued", logs: logs, sub: sub}, nil
}

// WatchCredentialIssued is a free log subscription operation binding the contract event 0xc6f6a1702b44006e35bd83045fd933510f3a924c1bb3877bff1061c14eb1c6be.
//
// Solidity: event CredentialIssued(bytes32 indexed digest, address indexed subject, address indexed registrar, uint256 insertedBlock)
func (_Faculty *FacultyFilterer) WatchCredentialIssued(opts *bind.WatchOpts, sink chan<- *FacultyCredentialIssued, digest [][32]byte, subject []common.Address, registrar []common.Address) (event.Subscription, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}

	logs, sub, err := _Faculty.contract.WatchLogs(opts, "CredentialIssued", digestRule, subjectRule, registrarRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FacultyCredentialIssued)
				if err := _Faculty.contract.UnpackLog(event, "CredentialIssued", log); err != nil {
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

// ParseCredentialIssued is a log parse operation binding the contract event 0xc6f6a1702b44006e35bd83045fd933510f3a924c1bb3877bff1061c14eb1c6be.
//
// Solidity: event CredentialIssued(bytes32 indexed digest, address indexed subject, address indexed registrar, uint256 insertedBlock)
func (_Faculty *FacultyFilterer) ParseCredentialIssued(log types.Log) (*FacultyCredentialIssued, error) {
	event := new(FacultyCredentialIssued)
	if err := _Faculty.contract.UnpackLog(event, "CredentialIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FacultyCredentialRevokedIterator is returned from FilterCredentialRevoked and is used to iterate over the raw logs and unpacked data for CredentialRevoked events raised by the Faculty contract.
type FacultyCredentialRevokedIterator struct {
	Event *FacultyCredentialRevoked // Event containing the contract specifics and raw log

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
func (it *FacultyCredentialRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FacultyCredentialRevoked)
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
		it.Event = new(FacultyCredentialRevoked)
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
func (it *FacultyCredentialRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FacultyCredentialRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FacultyCredentialRevoked represents a CredentialRevoked event raised by the Faculty contract.
type FacultyCredentialRevoked struct {
	Digest       [32]byte
	Subject      common.Address
	Revoker      common.Address
	RevokedBlock *big.Int
	Reason       [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCredentialRevoked is a free log retrieval operation binding the contract event 0xb53d448559d4a1f97e1fd48262b64223f5bdafd54d429a33599236f73e900b1f.
//
// Solidity: event CredentialRevoked(bytes32 indexed digest, address indexed subject, address indexed revoker, uint256 revokedBlock, bytes32 reason)
func (_Faculty *FacultyFilterer) FilterCredentialRevoked(opts *bind.FilterOpts, digest [][32]byte, subject []common.Address, revoker []common.Address) (*FacultyCredentialRevokedIterator, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var revokerRule []interface{}
	for _, revokerItem := range revoker {
		revokerRule = append(revokerRule, revokerItem)
	}

	logs, sub, err := _Faculty.contract.FilterLogs(opts, "CredentialRevoked", digestRule, subjectRule, revokerRule)
	if err != nil {
		return nil, err
	}
	return &FacultyCredentialRevokedIterator{contract: _Faculty.contract, event: "CredentialRevoked", logs: logs, sub: sub}, nil
}

// WatchCredentialRevoked is a free log subscription operation binding the contract event 0xb53d448559d4a1f97e1fd48262b64223f5bdafd54d429a33599236f73e900b1f.
//
// Solidity: event CredentialRevoked(bytes32 indexed digest, address indexed subject, address indexed revoker, uint256 revokedBlock, bytes32 reason)
func (_Faculty *FacultyFilterer) WatchCredentialRevoked(opts *bind.WatchOpts, sink chan<- *FacultyCredentialRevoked, digest [][32]byte, subject []common.Address, revoker []common.Address) (event.Subscription, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var revokerRule []interface{}
	for _, revokerItem := range revoker {
		revokerRule = append(revokerRule, revokerItem)
	}

	logs, sub, err := _Faculty.contract.WatchLogs(opts, "CredentialRevoked", digestRule, subjectRule, revokerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FacultyCredentialRevoked)
				if err := _Faculty.contract.UnpackLog(event, "CredentialRevoked", log); err != nil {
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

// ParseCredentialRevoked is a log parse operation binding the contract event 0xb53d448559d4a1f97e1fd48262b64223f5bdafd54d429a33599236f73e900b1f.
//
// Solidity: event CredentialRevoked(bytes32 indexed digest, address indexed subject, address indexed revoker, uint256 revokedBlock, bytes32 reason)
func (_Faculty *FacultyFilterer) ParseCredentialRevoked(log types.Log) (*FacultyCredentialRevoked, error) {
	event := new(FacultyCredentialRevoked)
	if err := _Faculty.contract.UnpackLog(event, "CredentialRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FacultyCredentialSignedIterator is returned from FilterCredentialSigned and is used to iterate over the raw logs and unpacked data for CredentialSigned events raised by the Faculty contract.
type FacultyCredentialSignedIterator struct {
	Event *FacultyCredentialSigned // Event containing the contract specifics and raw log

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
func (it *FacultyCredentialSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FacultyCredentialSigned)
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
		it.Event = new(FacultyCredentialSigned)
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
func (it *FacultyCredentialSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FacultyCredentialSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FacultyCredentialSigned represents a CredentialSigned event raised by the Faculty contract.
type FacultyCredentialSigned struct {
	Signer      common.Address
	Digest      [32]byte
	SignedBlock *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCredentialSigned is a free log retrieval operation binding the contract event 0x3c4c49076c09b6a1214e1ba437ac70fd0cff1c6ea5baf6df4c5f6ba937c4bf12.
//
// Solidity: event CredentialSigned(address indexed signer, bytes32 indexed digest, uint256 signedBlock)
func (_Faculty *FacultyFilterer) FilterCredentialSigned(opts *bind.FilterOpts, signer []common.Address, digest [][32]byte) (*FacultyCredentialSignedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _Faculty.contract.FilterLogs(opts, "CredentialSigned", signerRule, digestRule)
	if err != nil {
		return nil, err
	}
	return &FacultyCredentialSignedIterator{contract: _Faculty.contract, event: "CredentialSigned", logs: logs, sub: sub}, nil
}

// WatchCredentialSigned is a free log subscription operation binding the contract event 0x3c4c49076c09b6a1214e1ba437ac70fd0cff1c6ea5baf6df4c5f6ba937c4bf12.
//
// Solidity: event CredentialSigned(address indexed signer, bytes32 indexed digest, uint256 signedBlock)
func (_Faculty *FacultyFilterer) WatchCredentialSigned(opts *bind.WatchOpts, sink chan<- *FacultyCredentialSigned, signer []common.Address, digest [][32]byte) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _Faculty.contract.WatchLogs(opts, "CredentialSigned", signerRule, digestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FacultyCredentialSigned)
				if err := _Faculty.contract.UnpackLog(event, "CredentialSigned", log); err != nil {
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

// ParseCredentialSigned is a log parse operation binding the contract event 0x3c4c49076c09b6a1214e1ba437ac70fd0cff1c6ea5baf6df4c5f6ba937c4bf12.
//
// Solidity: event CredentialSigned(address indexed signer, bytes32 indexed digest, uint256 signedBlock)
func (_Faculty *FacultyFilterer) ParseCredentialSigned(log types.Log) (*FacultyCredentialSigned, error) {
	event := new(FacultyCredentialSigned)
	if err := _Faculty.contract.UnpackLog(event, "CredentialSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FacultyNodeAddedIterator is returned from FilterNodeAdded and is used to iterate over the raw logs and unpacked data for NodeAdded events raised by the Faculty contract.
type FacultyNodeAddedIterator struct {
	Event *FacultyNodeAdded // Event containing the contract specifics and raw log

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
func (it *FacultyNodeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FacultyNodeAdded)
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
		it.Event = new(FacultyNodeAdded)
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
func (it *FacultyNodeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FacultyNodeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FacultyNodeAdded represents a NodeAdded event raised by the Faculty contract.
type FacultyNodeAdded struct {
	CreatedBy   common.Address
	NodeAddress common.Address
	Role        uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeAdded is a free log retrieval operation binding the contract event 0x202458425cc4a92ea08231128a0a57aa24cfa42968e8090aa7d151f3f0a4eb3f.
//
// Solidity: event NodeAdded(address indexed createdBy, address indexed nodeAddress, uint8 role)
func (_Faculty *FacultyFilterer) FilterNodeAdded(opts *bind.FilterOpts, createdBy []common.Address, nodeAddress []common.Address) (*FacultyNodeAddedIterator, error) {

	var createdByRule []interface{}
	for _, createdByItem := range createdBy {
		createdByRule = append(createdByRule, createdByItem)
	}
	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _Faculty.contract.FilterLogs(opts, "NodeAdded", createdByRule, nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &FacultyNodeAddedIterator{contract: _Faculty.contract, event: "NodeAdded", logs: logs, sub: sub}, nil
}

// WatchNodeAdded is a free log subscription operation binding the contract event 0x202458425cc4a92ea08231128a0a57aa24cfa42968e8090aa7d151f3f0a4eb3f.
//
// Solidity: event NodeAdded(address indexed createdBy, address indexed nodeAddress, uint8 role)
func (_Faculty *FacultyFilterer) WatchNodeAdded(opts *bind.WatchOpts, sink chan<- *FacultyNodeAdded, createdBy []common.Address, nodeAddress []common.Address) (event.Subscription, error) {

	var createdByRule []interface{}
	for _, createdByItem := range createdBy {
		createdByRule = append(createdByRule, createdByItem)
	}
	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _Faculty.contract.WatchLogs(opts, "NodeAdded", createdByRule, nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FacultyNodeAdded)
				if err := _Faculty.contract.UnpackLog(event, "NodeAdded", log); err != nil {
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

// ParseNodeAdded is a log parse operation binding the contract event 0x202458425cc4a92ea08231128a0a57aa24cfa42968e8090aa7d151f3f0a4eb3f.
//
// Solidity: event NodeAdded(address indexed createdBy, address indexed nodeAddress, uint8 role)
func (_Faculty *FacultyFilterer) ParseNodeAdded(log types.Log) (*FacultyNodeAdded, error) {
	event := new(FacultyNodeAdded)
	if err := _Faculty.contract.UnpackLog(event, "NodeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FacultyOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Faculty contract.
type FacultyOwnerChangedIterator struct {
	Event *FacultyOwnerChanged // Event containing the contract specifics and raw log

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
func (it *FacultyOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FacultyOwnerChanged)
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
		it.Event = new(FacultyOwnerChanged)
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
func (it *FacultyOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FacultyOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FacultyOwnerChanged represents a OwnerChanged event raised by the Faculty contract.
type FacultyOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_Faculty *FacultyFilterer) FilterOwnerChanged(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*FacultyOwnerChangedIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Faculty.contract.FilterLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FacultyOwnerChangedIterator{contract: _Faculty.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_Faculty *FacultyFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *FacultyOwnerChanged, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Faculty.contract.WatchLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FacultyOwnerChanged)
				if err := _Faculty.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_Faculty *FacultyFilterer) ParseOwnerChanged(log types.Log) (*FacultyOwnerChanged, error) {
	event := new(FacultyOwnerChanged)
	if err := _Faculty.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FacultySemesterRegisteredIterator is returned from FilterSemesterRegistered and is used to iterate over the raw logs and unpacked data for SemesterRegistered events raised by the Faculty contract.
type FacultySemesterRegisteredIterator struct {
	Event *FacultySemesterRegistered // Event containing the contract specifics and raw log

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
func (it *FacultySemesterRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FacultySemesterRegistered)
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
		it.Event = new(FacultySemesterRegistered)
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
func (it *FacultySemesterRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FacultySemesterRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FacultySemesterRegistered represents a SemesterRegistered event raised by the Faculty contract.
type FacultySemesterRegistered struct {
	Registar common.Address
	Semester [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSemesterRegistered is a free log retrieval operation binding the contract event 0xcb185a4a072cf4e8ef43d02535b01a5ca5a4eeb8dcbc242651888e159e6c8bf0.
//
// Solidity: event SemesterRegistered(address registar, bytes32 semester)
func (_Faculty *FacultyFilterer) FilterSemesterRegistered(opts *bind.FilterOpts) (*FacultySemesterRegisteredIterator, error) {

	logs, sub, err := _Faculty.contract.FilterLogs(opts, "SemesterRegistered")
	if err != nil {
		return nil, err
	}
	return &FacultySemesterRegisteredIterator{contract: _Faculty.contract, event: "SemesterRegistered", logs: logs, sub: sub}, nil
}

// WatchSemesterRegistered is a free log subscription operation binding the contract event 0xcb185a4a072cf4e8ef43d02535b01a5ca5a4eeb8dcbc242651888e159e6c8bf0.
//
// Solidity: event SemesterRegistered(address registar, bytes32 semester)
func (_Faculty *FacultyFilterer) WatchSemesterRegistered(opts *bind.WatchOpts, sink chan<- *FacultySemesterRegistered) (event.Subscription, error) {

	logs, sub, err := _Faculty.contract.WatchLogs(opts, "SemesterRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FacultySemesterRegistered)
				if err := _Faculty.contract.UnpackLog(event, "SemesterRegistered", log); err != nil {
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

// ParseSemesterRegistered is a log parse operation binding the contract event 0xcb185a4a072cf4e8ef43d02535b01a5ca5a4eeb8dcbc242651888e159e6c8bf0.
//
// Solidity: event SemesterRegistered(address registar, bytes32 semester)
func (_Faculty *FacultyFilterer) ParseSemesterRegistered(log types.Log) (*FacultySemesterRegistered, error) {
	event := new(FacultySemesterRegistered)
	if err := _Faculty.contract.UnpackLog(event, "SemesterRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IssuerABI is the input ABI used to generate the binding from.
const IssuerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registrar\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"insertedBlock\",\"type\":\"uint256\"}],\"name\":\"CredentialIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"revoker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revokedBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"reason\",\"type\":\"bytes32\"}],\"name\":\"CredentialRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"signedBlock\",\"type\":\"uint256\"}],\"name\":\"CredentialSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"getCredentialProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"signed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"insertedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"registrar\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"witnesses\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"evidenceRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structNotary.CredentialProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"getCredentialSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"getDigests\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"getEvidenceRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"getProof\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"proof\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"insertedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structCredentialSum.Root\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"getRevoked\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"getRevokedProof\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"registrar\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"revokedBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"reason\",\"type\":\"bytes32\"}],\"internalType\":\"structNotary.RevocationProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"getWitnesses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"hasRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"isApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"isQuorumSigned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"isRevoked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isSigned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownersCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"recordExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"revokedCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"verifyCredential\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"verifyIssuedCredentials\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"digests\",\"type\":\"bytes32[]\"}],\"name\":\"verifyRootOf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"witnessesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Issuer is an auto generated Go binding around an Ethereum contract.
type Issuer struct {
	IssuerCaller     // Read-only binding to the contract
	IssuerTransactor // Write-only binding to the contract
	IssuerFilterer   // Log filterer for contract events
}

// IssuerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IssuerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IssuerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IssuerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IssuerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IssuerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IssuerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IssuerSession struct {
	Contract     *Issuer           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IssuerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IssuerCallerSession struct {
	Contract *IssuerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IssuerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IssuerTransactorSession struct {
	Contract     *IssuerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IssuerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IssuerRaw struct {
	Contract *Issuer // Generic contract binding to access the raw methods on
}

// IssuerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IssuerCallerRaw struct {
	Contract *IssuerCaller // Generic read-only contract binding to access the raw methods on
}

// IssuerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IssuerTransactorRaw struct {
	Contract *IssuerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIssuer creates a new instance of Issuer, bound to a specific deployed contract.
func NewIssuer(address common.Address, backend bind.ContractBackend) (*Issuer, error) {
	contract, err := bindIssuer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Issuer{IssuerCaller: IssuerCaller{contract: contract}, IssuerTransactor: IssuerTransactor{contract: contract}, IssuerFilterer: IssuerFilterer{contract: contract}}, nil
}

// NewIssuerCaller creates a new read-only instance of Issuer, bound to a specific deployed contract.
func NewIssuerCaller(address common.Address, caller bind.ContractCaller) (*IssuerCaller, error) {
	contract, err := bindIssuer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IssuerCaller{contract: contract}, nil
}

// NewIssuerTransactor creates a new write-only instance of Issuer, bound to a specific deployed contract.
func NewIssuerTransactor(address common.Address, transactor bind.ContractTransactor) (*IssuerTransactor, error) {
	contract, err := bindIssuer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IssuerTransactor{contract: contract}, nil
}

// NewIssuerFilterer creates a new log filterer instance of Issuer, bound to a specific deployed contract.
func NewIssuerFilterer(address common.Address, filterer bind.ContractFilterer) (*IssuerFilterer, error) {
	contract, err := bindIssuer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IssuerFilterer{contract: contract}, nil
}

// bindIssuer binds a generic wrapper to an already deployed contract.
func bindIssuer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IssuerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Issuer *IssuerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Issuer.Contract.IssuerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Issuer *IssuerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Issuer.Contract.IssuerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Issuer *IssuerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Issuer.Contract.IssuerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Issuer *IssuerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Issuer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Issuer *IssuerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Issuer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Issuer *IssuerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Issuer.Contract.contract.Transact(opts, method, params...)
}

// GetCredentialProof is a free data retrieval call binding the contract method 0xdac46a62.
//
// Solidity: function getCredentialProof(bytes32 digest) view returns((uint256,uint256,uint256,uint256,bytes32,bool,address,address,address[],bytes32))
func (_Issuer *IssuerCaller) GetCredentialProof(opts *bind.CallOpts, digest [32]byte) (NotaryCredentialProof, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "getCredentialProof", digest)

	if err != nil {
		return *new(NotaryCredentialProof), err
	}

	out0 := *abi.ConvertType(out[0], new(NotaryCredentialProof)).(*NotaryCredentialProof)

	return out0, err

}

// GetCredentialProof is a free data retrieval call binding the contract method 0xdac46a62.
//
// Solidity: function getCredentialProof(bytes32 digest) view returns((uint256,uint256,uint256,uint256,bytes32,bool,address,address,address[],bytes32))
func (_Issuer *IssuerSession) GetCredentialProof(digest [32]byte) (NotaryCredentialProof, error) {
	return _Issuer.Contract.GetCredentialProof(&_Issuer.CallOpts, digest)
}

// GetCredentialProof is a free data retrieval call binding the contract method 0xdac46a62.
//
// Solidity: function getCredentialProof(bytes32 digest) view returns((uint256,uint256,uint256,uint256,bytes32,bool,address,address,address[],bytes32))
func (_Issuer *IssuerCallerSession) GetCredentialProof(digest [32]byte) (NotaryCredentialProof, error) {
	return _Issuer.Contract.GetCredentialProof(&_Issuer.CallOpts, digest)
}

// GetCredentialSigners is a free data retrieval call binding the contract method 0x038ec1b8.
//
// Solidity: function getCredentialSigners(bytes32 digest) view returns(address[])
func (_Issuer *IssuerCaller) GetCredentialSigners(opts *bind.CallOpts, digest [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "getCredentialSigners", digest)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCredentialSigners is a free data retrieval call binding the contract method 0x038ec1b8.
//
// Solidity: function getCredentialSigners(bytes32 digest) view returns(address[])
func (_Issuer *IssuerSession) GetCredentialSigners(digest [32]byte) ([]common.Address, error) {
	return _Issuer.Contract.GetCredentialSigners(&_Issuer.CallOpts, digest)
}

// GetCredentialSigners is a free data retrieval call binding the contract method 0x038ec1b8.
//
// Solidity: function getCredentialSigners(bytes32 digest) view returns(address[])
func (_Issuer *IssuerCallerSession) GetCredentialSigners(digest [32]byte) ([]common.Address, error) {
	return _Issuer.Contract.GetCredentialSigners(&_Issuer.CallOpts, digest)
}

// GetDigests is a free data retrieval call binding the contract method 0x762b77bb.
//
// Solidity: function getDigests(address subject) view returns(bytes32[])
func (_Issuer *IssuerCaller) GetDigests(opts *bind.CallOpts, subject common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "getDigests", subject)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetDigests is a free data retrieval call binding the contract method 0x762b77bb.
//
// Solidity: function getDigests(address subject) view returns(bytes32[])
func (_Issuer *IssuerSession) GetDigests(subject common.Address) ([][32]byte, error) {
	return _Issuer.Contract.GetDigests(&_Issuer.CallOpts, subject)
}

// GetDigests is a free data retrieval call binding the contract method 0x762b77bb.
//
// Solidity: function getDigests(address subject) view returns(bytes32[])
func (_Issuer *IssuerCallerSession) GetDigests(subject common.Address) ([][32]byte, error) {
	return _Issuer.Contract.GetDigests(&_Issuer.CallOpts, subject)
}

// GetEvidenceRoot is a free data retrieval call binding the contract method 0x4302802d.
//
// Solidity: function getEvidenceRoot(bytes32 digest) view returns(bytes32)
func (_Issuer *IssuerCaller) GetEvidenceRoot(opts *bind.CallOpts, digest [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "getEvidenceRoot", digest)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEvidenceRoot is a free data retrieval call binding the contract method 0x4302802d.
//
// Solidity: function getEvidenceRoot(bytes32 digest) view returns(bytes32)
func (_Issuer *IssuerSession) GetEvidenceRoot(digest [32]byte) ([32]byte, error) {
	return _Issuer.Contract.GetEvidenceRoot(&_Issuer.CallOpts, digest)
}

// GetEvidenceRoot is a free data retrieval call binding the contract method 0x4302802d.
//
// Solidity: function getEvidenceRoot(bytes32 digest) view returns(bytes32)
func (_Issuer *IssuerCallerSession) GetEvidenceRoot(digest [32]byte) ([32]byte, error) {
	return _Issuer.Contract.GetEvidenceRoot(&_Issuer.CallOpts, digest)
}

// GetProof is a free data retrieval call binding the contract method 0x3eea79d1.
//
// Solidity: function getProof(address subject) view returns((bytes32,uint256,uint256))
func (_Issuer *IssuerCaller) GetProof(opts *bind.CallOpts, subject common.Address) (CredentialSumRoot, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "getProof", subject)

	if err != nil {
		return *new(CredentialSumRoot), err
	}

	out0 := *abi.ConvertType(out[0], new(CredentialSumRoot)).(*CredentialSumRoot)

	return out0, err

}

// GetProof is a free data retrieval call binding the contract method 0x3eea79d1.
//
// Solidity: function getProof(address subject) view returns((bytes32,uint256,uint256))
func (_Issuer *IssuerSession) GetProof(subject common.Address) (CredentialSumRoot, error) {
	return _Issuer.Contract.GetProof(&_Issuer.CallOpts, subject)
}

// GetProof is a free data retrieval call binding the contract method 0x3eea79d1.
//
// Solidity: function getProof(address subject) view returns((bytes32,uint256,uint256))
func (_Issuer *IssuerCallerSession) GetProof(subject common.Address) (CredentialSumRoot, error) {
	return _Issuer.Contract.GetProof(&_Issuer.CallOpts, subject)
}

// GetRevoked is a free data retrieval call binding the contract method 0x8ebd9623.
//
// Solidity: function getRevoked(address subject) view returns(bytes32[])
func (_Issuer *IssuerCaller) GetRevoked(opts *bind.CallOpts, subject common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "getRevoked", subject)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetRevoked is a free data retrieval call binding the contract method 0x8ebd9623.
//
// Solidity: function getRevoked(address subject) view returns(bytes32[])
func (_Issuer *IssuerSession) GetRevoked(subject common.Address) ([][32]byte, error) {
	return _Issuer.Contract.GetRevoked(&_Issuer.CallOpts, subject)
}

// GetRevoked is a free data retrieval call binding the contract method 0x8ebd9623.
//
// Solidity: function getRevoked(address subject) view returns(bytes32[])
func (_Issuer *IssuerCallerSession) GetRevoked(subject common.Address) ([][32]byte, error) {
	return _Issuer.Contract.GetRevoked(&_Issuer.CallOpts, subject)
}

// GetRevokedProof is a free data retrieval call binding the contract method 0xabbff984.
//
// Solidity: function getRevokedProof(bytes32 digest) view returns((address,address,uint256,bytes32))
func (_Issuer *IssuerCaller) GetRevokedProof(opts *bind.CallOpts, digest [32]byte) (NotaryRevocationProof, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "getRevokedProof", digest)

	if err != nil {
		return *new(NotaryRevocationProof), err
	}

	out0 := *abi.ConvertType(out[0], new(NotaryRevocationProof)).(*NotaryRevocationProof)

	return out0, err

}

// GetRevokedProof is a free data retrieval call binding the contract method 0xabbff984.
//
// Solidity: function getRevokedProof(bytes32 digest) view returns((address,address,uint256,bytes32))
func (_Issuer *IssuerSession) GetRevokedProof(digest [32]byte) (NotaryRevocationProof, error) {
	return _Issuer.Contract.GetRevokedProof(&_Issuer.CallOpts, digest)
}

// GetRevokedProof is a free data retrieval call binding the contract method 0xabbff984.
//
// Solidity: function getRevokedProof(bytes32 digest) view returns((address,address,uint256,bytes32))
func (_Issuer *IssuerCallerSession) GetRevokedProof(digest [32]byte) (NotaryRevocationProof, error) {
	return _Issuer.Contract.GetRevokedProof(&_Issuer.CallOpts, digest)
}

// GetWitnesses is a free data retrieval call binding the contract method 0x6f28eca6.
//
// Solidity: function getWitnesses(bytes32 digest) view returns(address[])
func (_Issuer *IssuerCaller) GetWitnesses(opts *bind.CallOpts, digest [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "getWitnesses", digest)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWitnesses is a free data retrieval call binding the contract method 0x6f28eca6.
//
// Solidity: function getWitnesses(bytes32 digest) view returns(address[])
func (_Issuer *IssuerSession) GetWitnesses(digest [32]byte) ([]common.Address, error) {
	return _Issuer.Contract.GetWitnesses(&_Issuer.CallOpts, digest)
}

// GetWitnesses is a free data retrieval call binding the contract method 0x6f28eca6.
//
// Solidity: function getWitnesses(bytes32 digest) view returns(address[])
func (_Issuer *IssuerCallerSession) GetWitnesses(digest [32]byte) ([]common.Address, error) {
	return _Issuer.Contract.GetWitnesses(&_Issuer.CallOpts, digest)
}

// HasRoot is a free data retrieval call binding the contract method 0x7532ed26.
//
// Solidity: function hasRoot(address subject) view returns(bool)
func (_Issuer *IssuerCaller) HasRoot(opts *bind.CallOpts, subject common.Address) (bool, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "hasRoot", subject)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRoot is a free data retrieval call binding the contract method 0x7532ed26.
//
// Solidity: function hasRoot(address subject) view returns(bool)
func (_Issuer *IssuerSession) HasRoot(subject common.Address) (bool, error) {
	return _Issuer.Contract.HasRoot(&_Issuer.CallOpts, subject)
}

// HasRoot is a free data retrieval call binding the contract method 0x7532ed26.
//
// Solidity: function hasRoot(address subject) view returns(bool)
func (_Issuer *IssuerCallerSession) HasRoot(subject common.Address) (bool, error) {
	return _Issuer.Contract.HasRoot(&_Issuer.CallOpts, subject)
}

// IsApproved is a free data retrieval call binding the contract method 0x48aefc32.
//
// Solidity: function isApproved(bytes32 digest) view returns(bool)
func (_Issuer *IssuerCaller) IsApproved(opts *bind.CallOpts, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "isApproved", digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0x48aefc32.
//
// Solidity: function isApproved(bytes32 digest) view returns(bool)
func (_Issuer *IssuerSession) IsApproved(digest [32]byte) (bool, error) {
	return _Issuer.Contract.IsApproved(&_Issuer.CallOpts, digest)
}

// IsApproved is a free data retrieval call binding the contract method 0x48aefc32.
//
// Solidity: function isApproved(bytes32 digest) view returns(bool)
func (_Issuer *IssuerCallerSession) IsApproved(digest [32]byte) (bool, error) {
	return _Issuer.Contract.IsApproved(&_Issuer.CallOpts, digest)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_Issuer *IssuerCaller) IsOwner(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "isOwner", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_Issuer *IssuerSession) IsOwner(account common.Address) (bool, error) {
	return _Issuer.Contract.IsOwner(&_Issuer.CallOpts, account)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_Issuer *IssuerCallerSession) IsOwner(account common.Address) (bool, error) {
	return _Issuer.Contract.IsOwner(&_Issuer.CallOpts, account)
}

// IsQuorumSigned is a free data retrieval call binding the contract method 0xae02c0cc.
//
// Solidity: function isQuorumSigned(bytes32 digest) view returns(bool)
func (_Issuer *IssuerCaller) IsQuorumSigned(opts *bind.CallOpts, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "isQuorumSigned", digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsQuorumSigned is a free data retrieval call binding the contract method 0xae02c0cc.
//
// Solidity: function isQuorumSigned(bytes32 digest) view returns(bool)
func (_Issuer *IssuerSession) IsQuorumSigned(digest [32]byte) (bool, error) {
	return _Issuer.Contract.IsQuorumSigned(&_Issuer.CallOpts, digest)
}

// IsQuorumSigned is a free data retrieval call binding the contract method 0xae02c0cc.
//
// Solidity: function isQuorumSigned(bytes32 digest) view returns(bool)
func (_Issuer *IssuerCallerSession) IsQuorumSigned(digest [32]byte) (bool, error) {
	return _Issuer.Contract.IsQuorumSigned(&_Issuer.CallOpts, digest)
}

// IsRevoked is a free data retrieval call binding the contract method 0x4294857f.
//
// Solidity: function isRevoked(bytes32 digest) view returns(bool)
func (_Issuer *IssuerCaller) IsRevoked(opts *bind.CallOpts, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "isRevoked", digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRevoked is a free data retrieval call binding the contract method 0x4294857f.
//
// Solidity: function isRevoked(bytes32 digest) view returns(bool)
func (_Issuer *IssuerSession) IsRevoked(digest [32]byte) (bool, error) {
	return _Issuer.Contract.IsRevoked(&_Issuer.CallOpts, digest)
}

// IsRevoked is a free data retrieval call binding the contract method 0x4294857f.
//
// Solidity: function isRevoked(bytes32 digest) view returns(bool)
func (_Issuer *IssuerCallerSession) IsRevoked(digest [32]byte) (bool, error) {
	return _Issuer.Contract.IsRevoked(&_Issuer.CallOpts, digest)
}

// IsSigned is a free data retrieval call binding the contract method 0x7f8d61e0.
//
// Solidity: function isSigned(bytes32 digest, address account) view returns(bool)
func (_Issuer *IssuerCaller) IsSigned(opts *bind.CallOpts, digest [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "isSigned", digest, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigned is a free data retrieval call binding the contract method 0x7f8d61e0.
//
// Solidity: function isSigned(bytes32 digest, address account) view returns(bool)
func (_Issuer *IssuerSession) IsSigned(digest [32]byte, account common.Address) (bool, error) {
	return _Issuer.Contract.IsSigned(&_Issuer.CallOpts, digest, account)
}

// IsSigned is a free data retrieval call binding the contract method 0x7f8d61e0.
//
// Solidity: function isSigned(bytes32 digest, address account) view returns(bool)
func (_Issuer *IssuerCallerSession) IsSigned(digest [32]byte, account common.Address) (bool, error) {
	return _Issuer.Contract.IsSigned(&_Issuer.CallOpts, digest, account)
}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address[])
func (_Issuer *IssuerCaller) Owners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "owners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address[])
func (_Issuer *IssuerSession) Owners() ([]common.Address, error) {
	return _Issuer.Contract.Owners(&_Issuer.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address[])
func (_Issuer *IssuerCallerSession) Owners() ([]common.Address, error) {
	return _Issuer.Contract.Owners(&_Issuer.CallOpts)
}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() view returns(uint8)
func (_Issuer *IssuerCaller) OwnersCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "ownersCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() view returns(uint8)
func (_Issuer *IssuerSession) OwnersCount() (uint8, error) {
	return _Issuer.Contract.OwnersCount(&_Issuer.CallOpts)
}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() view returns(uint8)
func (_Issuer *IssuerCallerSession) OwnersCount() (uint8, error) {
	return _Issuer.Contract.OwnersCount(&_Issuer.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint8)
func (_Issuer *IssuerCaller) Quorum(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "quorum")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint8)
func (_Issuer *IssuerSession) Quorum() (uint8, error) {
	return _Issuer.Contract.Quorum(&_Issuer.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint8)
func (_Issuer *IssuerCallerSession) Quorum() (uint8, error) {
	return _Issuer.Contract.Quorum(&_Issuer.CallOpts)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 digest) view returns(bool)
func (_Issuer *IssuerCaller) RecordExists(opts *bind.CallOpts, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "recordExists", digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 digest) view returns(bool)
func (_Issuer *IssuerSession) RecordExists(digest [32]byte) (bool, error) {
	return _Issuer.Contract.RecordExists(&_Issuer.CallOpts, digest)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 digest) view returns(bool)
func (_Issuer *IssuerCallerSession) RecordExists(digest [32]byte) (bool, error) {
	return _Issuer.Contract.RecordExists(&_Issuer.CallOpts, digest)
}

// RevokedCounter is a free data retrieval call binding the contract method 0x342577fa.
//
// Solidity: function revokedCounter(address subject) view returns(uint256)
func (_Issuer *IssuerCaller) RevokedCounter(opts *bind.CallOpts, subject common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "revokedCounter", subject)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RevokedCounter is a free data retrieval call binding the contract method 0x342577fa.
//
// Solidity: function revokedCounter(address subject) view returns(uint256)
func (_Issuer *IssuerSession) RevokedCounter(subject common.Address) (*big.Int, error) {
	return _Issuer.Contract.RevokedCounter(&_Issuer.CallOpts, subject)
}

// RevokedCounter is a free data retrieval call binding the contract method 0x342577fa.
//
// Solidity: function revokedCounter(address subject) view returns(uint256)
func (_Issuer *IssuerCallerSession) RevokedCounter(subject common.Address) (*big.Int, error) {
	return _Issuer.Contract.RevokedCounter(&_Issuer.CallOpts, subject)
}

// VerifyCredential is a free data retrieval call binding the contract method 0x5f889e17.
//
// Solidity: function verifyCredential(address subject, bytes32 digest) view returns(bool)
func (_Issuer *IssuerCaller) VerifyCredential(opts *bind.CallOpts, subject common.Address, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "verifyCredential", subject, digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyCredential is a free data retrieval call binding the contract method 0x5f889e17.
//
// Solidity: function verifyCredential(address subject, bytes32 digest) view returns(bool)
func (_Issuer *IssuerSession) VerifyCredential(subject common.Address, digest [32]byte) (bool, error) {
	return _Issuer.Contract.VerifyCredential(&_Issuer.CallOpts, subject, digest)
}

// VerifyCredential is a free data retrieval call binding the contract method 0x5f889e17.
//
// Solidity: function verifyCredential(address subject, bytes32 digest) view returns(bool)
func (_Issuer *IssuerCallerSession) VerifyCredential(subject common.Address, digest [32]byte) (bool, error) {
	return _Issuer.Contract.VerifyCredential(&_Issuer.CallOpts, subject, digest)
}

// VerifyIssuedCredentials is a free data retrieval call binding the contract method 0x2cf0695a.
//
// Solidity: function verifyIssuedCredentials(address subject) view returns(bool)
func (_Issuer *IssuerCaller) VerifyIssuedCredentials(opts *bind.CallOpts, subject common.Address) (bool, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "verifyIssuedCredentials", subject)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyIssuedCredentials is a free data retrieval call binding the contract method 0x2cf0695a.
//
// Solidity: function verifyIssuedCredentials(address subject) view returns(bool)
func (_Issuer *IssuerSession) VerifyIssuedCredentials(subject common.Address) (bool, error) {
	return _Issuer.Contract.VerifyIssuedCredentials(&_Issuer.CallOpts, subject)
}

// VerifyIssuedCredentials is a free data retrieval call binding the contract method 0x2cf0695a.
//
// Solidity: function verifyIssuedCredentials(address subject) view returns(bool)
func (_Issuer *IssuerCallerSession) VerifyIssuedCredentials(subject common.Address) (bool, error) {
	return _Issuer.Contract.VerifyIssuedCredentials(&_Issuer.CallOpts, subject)
}

// VerifyRootOf is a free data retrieval call binding the contract method 0x86509dd6.
//
// Solidity: function verifyRootOf(address subject, bytes32[] digests) view returns(bool)
func (_Issuer *IssuerCaller) VerifyRootOf(opts *bind.CallOpts, subject common.Address, digests [][32]byte) (bool, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "verifyRootOf", subject, digests)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyRootOf is a free data retrieval call binding the contract method 0x86509dd6.
//
// Solidity: function verifyRootOf(address subject, bytes32[] digests) view returns(bool)
func (_Issuer *IssuerSession) VerifyRootOf(subject common.Address, digests [][32]byte) (bool, error) {
	return _Issuer.Contract.VerifyRootOf(&_Issuer.CallOpts, subject, digests)
}

// VerifyRootOf is a free data retrieval call binding the contract method 0x86509dd6.
//
// Solidity: function verifyRootOf(address subject, bytes32[] digests) view returns(bool)
func (_Issuer *IssuerCallerSession) VerifyRootOf(subject common.Address, digests [][32]byte) (bool, error) {
	return _Issuer.Contract.VerifyRootOf(&_Issuer.CallOpts, subject, digests)
}

// WitnessesLength is a free data retrieval call binding the contract method 0xf537ff9d.
//
// Solidity: function witnessesLength(bytes32 digest) view returns(uint256)
func (_Issuer *IssuerCaller) WitnessesLength(opts *bind.CallOpts, digest [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Issuer.contract.Call(opts, &out, "witnessesLength", digest)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WitnessesLength is a free data retrieval call binding the contract method 0xf537ff9d.
//
// Solidity: function witnessesLength(bytes32 digest) view returns(uint256)
func (_Issuer *IssuerSession) WitnessesLength(digest [32]byte) (*big.Int, error) {
	return _Issuer.Contract.WitnessesLength(&_Issuer.CallOpts, digest)
}

// WitnessesLength is a free data retrieval call binding the contract method 0xf537ff9d.
//
// Solidity: function witnessesLength(bytes32 digest) view returns(uint256)
func (_Issuer *IssuerCallerSession) WitnessesLength(digest [32]byte) (*big.Int, error) {
	return _Issuer.Contract.WitnessesLength(&_Issuer.CallOpts, digest)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Issuer *IssuerTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Issuer.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Issuer *IssuerSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Issuer.Contract.ChangeOwner(&_Issuer.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Issuer *IssuerTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Issuer.Contract.ChangeOwner(&_Issuer.TransactOpts, newOwner)
}

// IssuerCredentialIssuedIterator is returned from FilterCredentialIssued and is used to iterate over the raw logs and unpacked data for CredentialIssued events raised by the Issuer contract.
type IssuerCredentialIssuedIterator struct {
	Event *IssuerCredentialIssued // Event containing the contract specifics and raw log

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
func (it *IssuerCredentialIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IssuerCredentialIssued)
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
		it.Event = new(IssuerCredentialIssued)
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
func (it *IssuerCredentialIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IssuerCredentialIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IssuerCredentialIssued represents a CredentialIssued event raised by the Issuer contract.
type IssuerCredentialIssued struct {
	Digest        [32]byte
	Subject       common.Address
	Registrar     common.Address
	InsertedBlock *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCredentialIssued is a free log retrieval operation binding the contract event 0xc6f6a1702b44006e35bd83045fd933510f3a924c1bb3877bff1061c14eb1c6be.
//
// Solidity: event CredentialIssued(bytes32 indexed digest, address indexed subject, address indexed registrar, uint256 insertedBlock)
func (_Issuer *IssuerFilterer) FilterCredentialIssued(opts *bind.FilterOpts, digest [][32]byte, subject []common.Address, registrar []common.Address) (*IssuerCredentialIssuedIterator, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}

	logs, sub, err := _Issuer.contract.FilterLogs(opts, "CredentialIssued", digestRule, subjectRule, registrarRule)
	if err != nil {
		return nil, err
	}
	return &IssuerCredentialIssuedIterator{contract: _Issuer.contract, event: "CredentialIssued", logs: logs, sub: sub}, nil
}

// WatchCredentialIssued is a free log subscription operation binding the contract event 0xc6f6a1702b44006e35bd83045fd933510f3a924c1bb3877bff1061c14eb1c6be.
//
// Solidity: event CredentialIssued(bytes32 indexed digest, address indexed subject, address indexed registrar, uint256 insertedBlock)
func (_Issuer *IssuerFilterer) WatchCredentialIssued(opts *bind.WatchOpts, sink chan<- *IssuerCredentialIssued, digest [][32]byte, subject []common.Address, registrar []common.Address) (event.Subscription, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}

	logs, sub, err := _Issuer.contract.WatchLogs(opts, "CredentialIssued", digestRule, subjectRule, registrarRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IssuerCredentialIssued)
				if err := _Issuer.contract.UnpackLog(event, "CredentialIssued", log); err != nil {
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

// ParseCredentialIssued is a log parse operation binding the contract event 0xc6f6a1702b44006e35bd83045fd933510f3a924c1bb3877bff1061c14eb1c6be.
//
// Solidity: event CredentialIssued(bytes32 indexed digest, address indexed subject, address indexed registrar, uint256 insertedBlock)
func (_Issuer *IssuerFilterer) ParseCredentialIssued(log types.Log) (*IssuerCredentialIssued, error) {
	event := new(IssuerCredentialIssued)
	if err := _Issuer.contract.UnpackLog(event, "CredentialIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IssuerCredentialRevokedIterator is returned from FilterCredentialRevoked and is used to iterate over the raw logs and unpacked data for CredentialRevoked events raised by the Issuer contract.
type IssuerCredentialRevokedIterator struct {
	Event *IssuerCredentialRevoked // Event containing the contract specifics and raw log

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
func (it *IssuerCredentialRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IssuerCredentialRevoked)
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
		it.Event = new(IssuerCredentialRevoked)
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
func (it *IssuerCredentialRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IssuerCredentialRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IssuerCredentialRevoked represents a CredentialRevoked event raised by the Issuer contract.
type IssuerCredentialRevoked struct {
	Digest       [32]byte
	Subject      common.Address
	Revoker      common.Address
	RevokedBlock *big.Int
	Reason       [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCredentialRevoked is a free log retrieval operation binding the contract event 0xb53d448559d4a1f97e1fd48262b64223f5bdafd54d429a33599236f73e900b1f.
//
// Solidity: event CredentialRevoked(bytes32 indexed digest, address indexed subject, address indexed revoker, uint256 revokedBlock, bytes32 reason)
func (_Issuer *IssuerFilterer) FilterCredentialRevoked(opts *bind.FilterOpts, digest [][32]byte, subject []common.Address, revoker []common.Address) (*IssuerCredentialRevokedIterator, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var revokerRule []interface{}
	for _, revokerItem := range revoker {
		revokerRule = append(revokerRule, revokerItem)
	}

	logs, sub, err := _Issuer.contract.FilterLogs(opts, "CredentialRevoked", digestRule, subjectRule, revokerRule)
	if err != nil {
		return nil, err
	}
	return &IssuerCredentialRevokedIterator{contract: _Issuer.contract, event: "CredentialRevoked", logs: logs, sub: sub}, nil
}

// WatchCredentialRevoked is a free log subscription operation binding the contract event 0xb53d448559d4a1f97e1fd48262b64223f5bdafd54d429a33599236f73e900b1f.
//
// Solidity: event CredentialRevoked(bytes32 indexed digest, address indexed subject, address indexed revoker, uint256 revokedBlock, bytes32 reason)
func (_Issuer *IssuerFilterer) WatchCredentialRevoked(opts *bind.WatchOpts, sink chan<- *IssuerCredentialRevoked, digest [][32]byte, subject []common.Address, revoker []common.Address) (event.Subscription, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var revokerRule []interface{}
	for _, revokerItem := range revoker {
		revokerRule = append(revokerRule, revokerItem)
	}

	logs, sub, err := _Issuer.contract.WatchLogs(opts, "CredentialRevoked", digestRule, subjectRule, revokerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IssuerCredentialRevoked)
				if err := _Issuer.contract.UnpackLog(event, "CredentialRevoked", log); err != nil {
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

// ParseCredentialRevoked is a log parse operation binding the contract event 0xb53d448559d4a1f97e1fd48262b64223f5bdafd54d429a33599236f73e900b1f.
//
// Solidity: event CredentialRevoked(bytes32 indexed digest, address indexed subject, address indexed revoker, uint256 revokedBlock, bytes32 reason)
func (_Issuer *IssuerFilterer) ParseCredentialRevoked(log types.Log) (*IssuerCredentialRevoked, error) {
	event := new(IssuerCredentialRevoked)
	if err := _Issuer.contract.UnpackLog(event, "CredentialRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IssuerCredentialSignedIterator is returned from FilterCredentialSigned and is used to iterate over the raw logs and unpacked data for CredentialSigned events raised by the Issuer contract.
type IssuerCredentialSignedIterator struct {
	Event *IssuerCredentialSigned // Event containing the contract specifics and raw log

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
func (it *IssuerCredentialSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IssuerCredentialSigned)
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
		it.Event = new(IssuerCredentialSigned)
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
func (it *IssuerCredentialSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IssuerCredentialSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IssuerCredentialSigned represents a CredentialSigned event raised by the Issuer contract.
type IssuerCredentialSigned struct {
	Signer      common.Address
	Digest      [32]byte
	SignedBlock *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCredentialSigned is a free log retrieval operation binding the contract event 0x3c4c49076c09b6a1214e1ba437ac70fd0cff1c6ea5baf6df4c5f6ba937c4bf12.
//
// Solidity: event CredentialSigned(address indexed signer, bytes32 indexed digest, uint256 signedBlock)
func (_Issuer *IssuerFilterer) FilterCredentialSigned(opts *bind.FilterOpts, signer []common.Address, digest [][32]byte) (*IssuerCredentialSignedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _Issuer.contract.FilterLogs(opts, "CredentialSigned", signerRule, digestRule)
	if err != nil {
		return nil, err
	}
	return &IssuerCredentialSignedIterator{contract: _Issuer.contract, event: "CredentialSigned", logs: logs, sub: sub}, nil
}

// WatchCredentialSigned is a free log subscription operation binding the contract event 0x3c4c49076c09b6a1214e1ba437ac70fd0cff1c6ea5baf6df4c5f6ba937c4bf12.
//
// Solidity: event CredentialSigned(address indexed signer, bytes32 indexed digest, uint256 signedBlock)
func (_Issuer *IssuerFilterer) WatchCredentialSigned(opts *bind.WatchOpts, sink chan<- *IssuerCredentialSigned, signer []common.Address, digest [][32]byte) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _Issuer.contract.WatchLogs(opts, "CredentialSigned", signerRule, digestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IssuerCredentialSigned)
				if err := _Issuer.contract.UnpackLog(event, "CredentialSigned", log); err != nil {
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

// ParseCredentialSigned is a log parse operation binding the contract event 0x3c4c49076c09b6a1214e1ba437ac70fd0cff1c6ea5baf6df4c5f6ba937c4bf12.
//
// Solidity: event CredentialSigned(address indexed signer, bytes32 indexed digest, uint256 signedBlock)
func (_Issuer *IssuerFilterer) ParseCredentialSigned(log types.Log) (*IssuerCredentialSigned, error) {
	event := new(IssuerCredentialSigned)
	if err := _Issuer.contract.UnpackLog(event, "CredentialSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IssuerOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Issuer contract.
type IssuerOwnerChangedIterator struct {
	Event *IssuerOwnerChanged // Event containing the contract specifics and raw log

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
func (it *IssuerOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IssuerOwnerChanged)
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
		it.Event = new(IssuerOwnerChanged)
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
func (it *IssuerOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IssuerOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IssuerOwnerChanged represents a OwnerChanged event raised by the Issuer contract.
type IssuerOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_Issuer *IssuerFilterer) FilterOwnerChanged(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*IssuerOwnerChangedIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Issuer.contract.FilterLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IssuerOwnerChangedIterator{contract: _Issuer.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_Issuer *IssuerFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *IssuerOwnerChanged, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Issuer.contract.WatchLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IssuerOwnerChanged)
				if err := _Issuer.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_Issuer *IssuerFilterer) ParseOwnerChanged(log types.Log) (*IssuerOwnerChanged, error) {
	event := new(IssuerOwnerChanged)
	if err := _Issuer.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeABI is the input ABI used to generate the binding from.
const NodeABI = "[{\"inputs\":[{\"internalType\":\"enumRole\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"registrars\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"quorum\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registrar\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"insertedBlock\",\"type\":\"uint256\"}],\"name\":\"CredentialIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"revoker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revokedBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"reason\",\"type\":\"bytes32\"}],\"name\":\"CredentialRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"signedBlock\",\"type\":\"uint256\"}],\"name\":\"CredentialSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"createdBy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"NodeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"addChild\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"digests\",\"type\":\"bytes32[]\"}],\"name\":\"aggregateCredentials\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"approveCredential\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChildren\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"getCredentialProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"signed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"insertedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"registrar\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"witnesses\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"evidenceRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structNotary.CredentialProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"getCredentialSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"getDigests\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"getEvidenceRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"getProof\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"proof\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"insertedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structCredentialSum.Root\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"getRevoked\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"getRevokedProof\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"registrar\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"revokedBlock\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"reason\",\"type\":\"bytes32\"}],\"internalType\":\"structNotary.RevocationProof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRole\",\"outputs\":[{\"internalType\":\"enumRole\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"getWitnesses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"hasRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"isApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isChild\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLeaf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"isQuorumSigned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"isRevoked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isSigned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"myParent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownersCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"recordExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"witnesses\",\"type\":\"address[]\"}],\"name\":\"registerCredential\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"reason\",\"type\":\"bytes32\"}],\"name\":\"revokeCredential\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"revokedCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"verifyCredential\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyCredentialRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"verifyCredentialTree\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"verifyIssuedCredentials\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"digests\",\"type\":\"bytes32[]\"}],\"name\":\"verifyRootOf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"witnessesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NodeBin is the compiled bytecode used for deploying new contracts.
var NodeBin = "0x60c0604052635041329d60e01b60a09081526200002190600b90600162000322565b503480156200002f57600080fd5b50604051620043a8380380620043a8833981016040819052620000529162000473565b81818181600082511180156200006a5750815160ff10155b620000bc5760405162461bcd60e51b815260206004820152601860248201527f4f776e6572732f6e6f7420656e6f756768206f776e657273000000000000000060448201526064015b60405180910390fd5b60008160ff16118015620000d4575081518160ff1611155b620001225760405162461bcd60e51b815260206004820152601a60248201527f4f776e6572732f71756f72756d206f7574206f662072616e67650000000000006044820152606401620000b3565b60005b82518160ff161015620002575760026000848360ff16815181106200015a57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff16158015620001cd575060006001600160a01b0316838260ff1681518110620001b957634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b031614155b620001e857634e487b7160e01b600052600160045260246000fd5b600160026000858460ff16815181106200021257634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff19169115159190911790556200024f8162000573565b905062000125565b5081516200026d906001906020850190620003d5565b5090516000805461ff00191661010060ff938416021760ff19169190921617905550339050620002e05760405162461bcd60e51b815260206004820152601760248201527f4e6f64652f73656e6465722063616e6e6f7420626520300000000000000000006044820152606401620000b3565b3360601b608052600c805484919060ff1916600183818111156200031457634e487b7160e01b600052602160045260246000fd5b0217905550505050620005b6565b82805482825590600052602060002090600701600890048101928215620003c35791602002820160005b838211156200038f57835183826101000a81548163ffffffff021916908360e01c021790555092602001926004016020816003010492830192600103026200034c565b8015620003c15782816101000a81549063ffffffff02191690556004016020816003010492830192600103026200038f565b505b50620003d19291506200042d565b5090565b828054828255906000526020600020908101928215620003c3579160200282015b82811115620003c357825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620003f6565b5b80821115620003d157600081556001016200042e565b80516001600160a01b03811681146200045c57600080fd5b919050565b805160ff811681146200045c57600080fd5b60008060006060848603121562000488578283fd5b83516002811062000497578384fd5b602085810151919450906001600160401b0380821115620004b6578485fd5b818701915087601f830112620004ca578485fd5b815181811115620004df57620004df620005a0565b838102604051601f19603f83011681018181108582111715620005065762000506620005a0565b604052828152858101935084860182860187018c101562000525578889fd5b8895505b8386101562000552576200053d8162000444565b85526001959095019493860193860162000529565b508097505050505050506200056a6040850162000461565b90509250925092565b600060ff821660ff8114156200059757634e487b7160e01b82526011600452602482fd5b60010192915050565b634e487b7160e01b600052604160045260246000fd5b60805160601c613dd3620005d560003960006105a00152613dd36000f3fe608060405234801561001057600080fd5b50600436106102325760003560e01c8063762b77bb11610130578063b9488546116100b8578063f537ff9d1161007c578063f537ff9d146105f0578063f79fe53814610613578063f936cc1514610626578063fc91a89714610639578063fde081741461065c57610232565b8063b948854614610562578063dac46a6214610573578063de07fc4514610593578063e2d02cea146105ca578063eedb973c146105dd57610232565b806399d1dbc0116100ff57806399d1dbc0146104d8578063a6f9dae1146104e0578063abbff984146104f3578063ae02c0cc14610547578063affe39c11461055a57610232565b8063762b77bb1461047f5780637f8d61e01461049f57806386509dd6146104b25780638ebd9623146104c557610232565b80633610d3ca116101be5780634892e8e8116101825780634892e8e81461042b57806348aefc32146104335780635f889e17146104465780636f28eca6146104595780637532ed261461046c57610232565b80633610d3ca146103495780633ac82e191461035c5780633eea79d11461036f5780634294857f146103f55780634302802d1461040857610232565b80631703a018116102055780631703a018146102b35780631eee993a146102cc5780632cf0695a146102e15780632f54bf6e146102f4578063342577fa1461032057610232565b806301ffc9a714610237578063038ec1b81461025f578063079cf76e1461027f5780630ae07028146102a0575b600080fd5b61024a610245366004613836565b61066f565b60405190151581526020015b60405180910390f35b61027261026d3660046137b6565b6106a8565b6040516102569190613a65565b61029261028d3660046135f9565b610811565b604051908152602001610256565b61024a6102ae3660046136ba565b61082f565b60005460ff165b60405160ff9091168152602001610256565b6102df6102da3660046135f9565b610844565b005b61024a6102ef3660046135f9565b610a20565b61024a6103023660046135f9565b6001600160a01b031660009081526002602052604090205460ff1690565b61029261032e3660046135f9565b6001600160a01b031660009081526008602052604090205490565b6102df610357366004613815565b610a69565b6102df61036a3660046136e5565b610aa6565b6103d361037d3660046135f9565b6040805160608082018352600080835260208084018290529284018190526001600160a01b03949094168452600a8252928290208251938401835280548452600181015491840191909152600201549082015290565b6040805182518152602080840151908201529181015190820152606001610256565b61024a6104033660046137b6565b610ed4565b6102926104163660046137b6565b60009081526006602052604090206008015490565b610272610f62565b61024a6104413660046137b6565b610fc4565b61024a6104543660046136ba565b611006565b6102726104673660046137b6565b611014565b61024a61047a3660046135f9565b611083565b61049261048d3660046135f9565b6110d4565b6040516102569190613a8b565b61024a6104ad3660046137e6565b61113f565b61024a6104c0366004613615565b6111dc565b6104926104d33660046135f9565b61122c565b61024a6113df565b6102df6104ee3660046135f9565b61140d565b6105066105013660046137b6565b61171b565b604051610256919081516001600160a01b03908116825260208084015190911690820152604080830151908201526060918201519181019190915260800190565b61024a6105553660046137b6565b6117c5565b610272611812565b6102ba600054610100900460ff1690565b6105866105813660046137b6565b611872565b6040516102569190613bb5565b6040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152602001610256565b6102926105d8366004613615565b61194c565b6102df6105eb3660046137b6565b611985565b6102926105fe3660046137b6565b60009081526006602052604090206007015490565b61024a6106213660046137b6565b611991565b61024a6106343660046135f9565b6119d3565b61024a6106473660046135f9565b600e6020526000908152604090205460ff1681565b600c5460ff166040516102569190613ae8565b60006001600160e01b031982166301ffc9a760e01b14806106a057506001600160e01b03198216635041329d60e01b145b90505b919050565b60015460609060009067ffffffffffffffff8111156106d757634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610700578160200160208202803683370190505b5090506000805b600154811015610808576000858152600960205260408120600180549192918490811061074457634e487b7160e01b600052603260045260246000fd5b60009182526020808320909101546001600160a01b0316835282019290925260400190205460ff16156107f6576001818154811061079257634e487b7160e01b600052603260045260246000fd5b9060005260206000200160009054906101000a90046001600160a01b03168383815181106107d057634e487b7160e01b600052603260045260246000fd5b6001600160a01b0390921660209283029190910190910152816107f281613d21565b9250505b8061080081613d21565b915050610707565b50909392505050565b6001600160a01b0381166000908152600a60205260408120546106a0565b600061083b8383611c0a565b90505b92915050565b3360009081526002602052604090205460ff1661087c5760405162461bcd60e51b815260040161087390613b47565b60405180910390fd5b6001600c5460ff1660018111156108a357634e487b7160e01b600052602160045260246000fd5b146108f05760405162461bcd60e51b815260206004820152601760248201527f4e6f64652f6e6f6465206d75737420626520496e6e65720000000000000000006044820152606401610873565b600061097930600b80548060200260200160405190810160405280929190818152602001828054801561096f57602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b031916815260200190600401906020826003010492830192600103820291508084116109315790505b5050505050611d17565b90508061099657634e487b7160e01b600052600160045260246000fd5b60008290506000816001600160a01b031663fde081746040518163ffffffff1660e01b815260040160206040518083038186803b1580156109d657600080fd5b505afa1580156109ea573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a0e919061385e565b9050610a1a8482611d8b565b50505050565b6001600160a01b0381166000908152600560205260408120548290610a575760405162461bcd60e51b815260040161087390613b10565b610a62600384611f82565b9392505050565b3360009081526002602052604090205460ff16610a985760405162461bcd60e51b815260040161087390613b47565b610aa28282611ff5565b5050565b3360009081526002602052604090205460ff16610ad55760405162461bcd60e51b815260040161087390613b47565b6000600c5460ff166001811115610afc57634e487b7160e01b600052602160045260246000fd5b1415610b7457805115610b515760405162461bcd60e51b815260206004820152601f60248201527f4e6f64652f4c6561662063616e6e6f742068617665207769746e6573736573006044820152606401610873565b60408051600080825260208201909252610b6f9185918591906120c5565b610ecf565b6001600c5460ff166001811115610b9b57634e487b7160e01b600052602160045260246000fd5b14610bb657634e487b7160e01b600052600160045260246000fd5b6000815111610c005760405162461bcd60e51b8152602060048201526016602482015275139bd9194bddda5d1b995cdcc81b9bdd08199bdd5b9960521b6044820152606401610873565b6000815167ffffffffffffffff811115610c2a57634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610c53578160200160208202803683370190505b50905060005b8251811015610e33576000838281518110610c8457634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b0381166000908152600e90925260409091205490915060ff16610cfc5760405162461bcd60e51b815260206004820152601b60248201527f4e6f64652f61646472657373206e6f7420617574686f72697a656400000000006044820152606401610873565b6000610d0f82635041329d60e01b612192565b905080610d2c57634e487b7160e01b600052600160045260246000fd5b6040516303ce7bb760e11b81526001600160a01b038881166004830152839160009183169063079cf76e9060240160206040518083038186803b158015610d7257600080fd5b505afa158015610d86573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610daa91906137ce565b905080610def5760405162461bcd60e51b8152602060048201526013602482015272139bd9194bdc9bdbdd081b9bdd08199bdd5b99606a1b6044820152606401610873565b80868681518110610e1057634e487b7160e01b600052603260045260246000fd5b602002602001018181525050505050508080610e2b90613d21565b915050610c59565b50604051637d71689560e11b815260009073__$9176519fc5fb44025266bd62b274a40525$__9063fae2d12a90610e6e908590600401613a78565b60206040518083038186803b158015610e8657600080fd5b505af4158015610e9a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ebe91906137ce565b9050610ecc858583866120c5565b50505b505050565b604051631999f00560e21b8152600360048201526024810182905260009073__$1d2b6a72c2c8582bd594e19b02f34b62b6$__90636667c014906044015b60206040518083038186803b158015610f2a57600080fd5b505af4158015610f3e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106a0919061379c565b6060600d805480602002602001604051908101604052809291908181526020018280548015610fba57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610f9c575b5050505050905090565b6040516312f4761160e31b8152600360048201526024810182905260009073__$1d2b6a72c2c8582bd594e19b02f34b62b6$__906397a3b08890604401610f12565b600061083b600384846121ae565b60008181526006602090815260409182902060070180548351818402810184019094528084526060939283018282801561107757602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611059575b50505050509050919050565b6001600160a01b0381166000908152600a60205260408082209051634650570f60e01b8152600481019190915273__$9176519fc5fb44025266bd62b274a40525$__90634650570f90602401610f12565b6001600160a01b03811660009081526005602090815260409182902080548351818402810184019094528084526060939283018282801561107757602002820191906000526020600020905b8154815260200190600101908083116111205750505050509050919050565b60405163a8b0bcfd60e01b815260036004820152602481018390526001600160a01b038216604482015260009073__$1d2b6a72c2c8582bd594e19b02f34b62b6$__9063a8b0bcfd906064015b60206040518083038186803b1580156111a457600080fd5b505af41580156111b8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061083b919061379c565b6001600160a01b0382166000908152600a6020526040808220905163070f6e8d60e01b815273__$9176519fc5fb44025266bd62b274a40525$__9163070f6e8d9161118c91908690600401613acf565b6001600160a01b03811660009081526005602052604090205460609082906112665760405162461bcd60e51b815260040161087390613b10565b6001600160a01b03831660009081526008602052604081205467ffffffffffffffff8111156112a557634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156112ce578160200160208202803683370190505b5090506000805b6001600160a01b0386166000908152600560205260409020548110156113d5576001600160a01b0386166000908152600560205260409020805461134291908390811061133257634e487b7160e01b600052603260045260246000fd5b9060005260206000200154610ed4565b156113c3576001600160a01b038616600090815260056020526040902080548290811061137f57634e487b7160e01b600052603260045260246000fd5b90600052602060002001548383815181106113aa57634e487b7160e01b600052603260045260246000fd5b6020908102919091010152816113bf81613d21565b9250505b806113cd81613d21565b9150506112d5565b5090949350505050565b600080600c5460ff16600181111561140757634e487b7160e01b600052602160045260246000fd5b14905090565b3360009081526002602052604090205460ff1661143c5760405162461bcd60e51b815260040161087390613b47565b6001600160a01b03811660009081526002602052604090205460ff1615801561146d57506001600160a01b03811615155b6114b95760405162461bcd60e51b815260206004820152601c60248201527f4f776e6572732f696e76616c6964206164647265737320676976656e000000006044820152606401610873565b600154158015906114cd575060015460ff10155b6114e757634e487b7160e01b600052600160045260246000fd5b60015460009067ffffffffffffffff81111561151357634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561153c578160200160208202803683370190505b50905060005b60015460ff8216101561167657336001600160a01b031660018260ff168154811061157d57634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b0316146116215760018160ff16815481106115bb57634e487b7160e01b600052603260045260246000fd5b9060005260206000200160009054906101000a90046001600160a01b0316828260ff16815181106115fc57634e487b7160e01b600052603260045260246000fd5b60200260200101906001600160a01b031690816001600160a01b031681525050611666565b82828260ff168151811061164557634e487b7160e01b600052603260045260246000fd5b60200260200101906001600160a01b031690816001600160a01b0316815250505b61166f81613d3c565b9050611542565b50600054815160ff9091161461169c57634e487b7160e01b600052600160045260246000fd5b6040516001600160a01b0383169033907fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c90600090a380516116e59060019060208401906134ef565b50506001600160a01b0316600090815260026020526040808220805460ff19908116600117909155338352912080549091169055565b6040805160808101825260008082526020820181905281830181905260608201529051636dbae62960e11b8152600360048201526024810183905273__$1d2b6a72c2c8582bd594e19b02f34b62b6$__9063db75cc529060440160806040518083038186803b15801561178d57600080fd5b505af41580156117a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106a09190613962565b60008054604051633e01fd5d60e21b8152600360048201526024810184905260ff909116604482015273__$1d2b6a72c2c8582bd594e19b02f34b62b6$__9063f807f57490606401610f12565b60606001805480602002602001604051908101604052809291908181526020018280548015610fba576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311610f9c575050505050905090565b604080516101408101825260008082526020820181905281830181905260608083018290526080830182905260a0830182905260c0830182905260e0830182905261010083015261012082015290516301c30fd760e71b8152600360048201526024810183905273__$1d2b6a72c2c8582bd594e19b02f34b62b6$__9063e187eb809060440160006040518083038186803b15801561191057600080fd5b505af4158015611924573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106a0919081019061387d565b3360009081526002602052604081205460ff1661197b5760405162461bcd60e51b815260040161087390613b47565b61083b83836122b6565b61198e81612417565b50565b60405163143ceab160e21b8152600360048201526024810182905260009073__$1d2b6a72c2c8582bd594e19b02f34b62b6$__906350f3aac490604401610f12565b6000806119df836110d4565b90506000815111611a325760405162461bcd60e51b815260206004820152601960248201527f4e6f64652f63726564656e7469616c206e6f7420666f756e64000000000000006044820152606401610873565b611a3b83611083565b15611a5857611a4a83826111dc565b611a585760009150506106a3565b60005b8151811015611c0057611a94828281518110611a8757634e487b7160e01b600052603260045260246000fd5b6020026020010151611991565b611aae57634e487b7160e01b600052600160045260246000fd5b611adf84838381518110611ad257634e487b7160e01b600052603260045260246000fd5b6020026020010151611006565b611aee576000925050506106a3565b6001600c5460ff166001811115611b1557634e487b7160e01b600052602160045260246000fd5b148015611b6157506000611b5f838381518110611b4257634e487b7160e01b600052603260045260246000fd5b602002602001015160009081526006602052604090206007015490565b115b15611bee57611bdf84611baa848481518110611b8d57634e487b7160e01b600052603260045260246000fd5b602002602001015160009081526006602052604090206008015490565b611bda858581518110611bcd57634e487b7160e01b600052603260045260246000fd5b6020026020010151611014565b6124ea565b611bee576000925050506106a3565b80611bf881613d21565b915050611a5b565b5060019392505050565b6001600160a01b0382166000908152600560205260408120548390611c415760405162461bcd60e51b815260040161087390613b10565b6001600160a01b0384166000908152600a60209081526040808320600590925291829020915163070f6e8d60e01b815273__$9176519fc5fb44025266bd62b274a40525$__9263070f6e8d92611c9c92909190600401613c89565b60206040518083038186803b158015611cb457600080fd5b505af4158015611cc8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cec919061379c565b8015611d0f57506001600160a01b0384166000908152600a602052604090205483145b949350505050565b6000611d2283612989565b611d2e5750600061083e565b60005b8251811015611c0057611d6b84848381518110611d5e57634e487b7160e01b600052603260045260246000fd5b60200260200101516129bc565b611d7957600091505061083e565b80611d8381613d21565b915050611d31565b306001600160a01b0383161415611ddd5760405162461bcd60e51b81526020600482015260166024820152752737b23297b1b0b73737ba1030b2321034ba39b2b63360511b6044820152606401610873565b6001600160a01b0382166000908152600e602052604090205460ff1615611e465760405162461bcd60e51b815260206004820152601760248201527f4e6f64652f6e6f646520616c72656164792061646465640000000000000000006044820152606401610873565b6000816001811115611e6857634e487b7160e01b600052602160045260246000fd5b1480611e9357506001816001811115611e9157634e487b7160e01b600052602160045260246000fd5b145b611edf5760405162461bcd60e51b815260206004820152601760248201527f4e6f64652f696e76616c6964206368696c6420726f6c650000000000000000006044820152606401610873565b6001600160a01b0382166000818152600e6020526040808220805460ff19166001908117909155600d805491820181559092527fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb590910180546001600160a01b031916831790555133907f202458425cc4a92ea08231128a0a57aa24cfa42968e8090aa7d151f3f0a4eb3f90611f76908590613ae8565b60405180910390a35050565b6001600160a01b03811660009081526002830160209081526040808320805482518185028101850190935280835261083b938793879390929091830182828015611feb57602002820191906000526020600020905b815481526020019060010190808311611fd7575b50505050506129df565b81611fff81610ed4565b1561201c5760405162461bcd60e51b815260040161087390613b7e565b600083815260066020819052604090912001546001600160a01b031661205a336001600160a01b031660009081526002602052604090205460ff1690565b8061206d57506001600160a01b03811633145b6120b95760405162461bcd60e51b815260206004820152601c60248201527f4973737565722f73656e646572206e6f7420617574686f72697a6564000000006044820152606401610873565b610a1a60038585612a49565b3360009081526002602052604090205460ff166120f45760405162461bcd60e51b815260040161087390613b47565b826120fe81610ed4565b1561211b5760405162461bcd60e51b815260040161087390613b7e565b6001600160a01b03851660009081526002602052604090205460ff16156121845760405162461bcd60e51b815260206004820152601a60248201527f4973737565722f666f7262696464656e207265676973747261720000000000006044820152606401610873565b610ecc600386868686612bf8565b600061219d83612989565b801561083b575061083b83836129bc565b600081815260038401602052604081206001015461220e5760405162461bcd60e51b815260206004820152601b60248201527f4e6f746172792f63726564656e7469616c206e6f7420666f756e6400000000006044820152606401610873565b60008281526003850160205260409020600601546001600160a01b0384811691161461227c5760405162461bcd60e51b815260206004820152601b60248201527f4e6f746172792f6e6f74206f776e6564206279207375626a65637400000000006044820152606401610873565b600082815260038501602052604090206005015460ff168015611d0f57505060009081526004929092016020525060409020600201541590565b3360009081526002602052604081205460ff166122e55760405162461bcd60e51b815260040161087390613b47565b6001600160a01b038316600090815260056020526040902054839061231c5760405162461bcd60e51b815260040161087390613b10565b612328600385856129df565b6123745760405162461bcd60e51b815260206004820152601e60248201527f4973737565722f68617320696e76616c69642063726564656e7469616c7300006044820152606401610873565b6001600160a01b0384166000908152600a6020526040908190209051631b33c8e760e31b815273__$9176519fc5fb44025266bd62b274a40525$__9163d99e4738916123c7919088908890600401613c5f565b60206040518083038186803b1580156123df57600080fd5b505af41580156123f3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d0f91906137ce565b8061242181610ed4565b1561243e5760405162461bcd60e51b815260040161087390613b7e565b60005460ff166124895760405162461bcd60e51b8152602060048201526016602482015275125cdcdd595c8bdb9bc81c5d5bdc9d5b48199bdd5b9960521b6044820152606401610873565b6124a58261249960005460ff1690565b6003919060ff1661324c565b610aa25760405162461bcd60e51b8152602060048201526016602482015275125cdcdd595c8bd85c1c1c9bdd985b0819985a5b195960521b6044820152606401610873565b6000826125395760405162461bcd60e51b815260206004820152601860248201527f4e6f64652f726f6f742063616e6e6f74206265206e756c6c00000000000000006044820152606401610873565b6000825167ffffffffffffffff81111561256357634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561258c578160200160208202803683370190505b50905060005b83518110156128f65760008482815181106125bd57634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b0381166000908152600e90925260409091205490915060ff166126355760405162461bcd60e51b815260206004820152601b60248201527f4e6f64652f61646472657373206e6f7420617574686f72697a656400000000006044820152606401610873565b600061264882635041329d60e01b612192565b90508061266557634e487b7160e01b600052600160045260246000fd5b6000829050806001600160a01b03166399d1dbc06040518163ffffffff1660e01b815260040160206040518083038186803b1580156126a357600080fd5b505afa1580156126b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126db919061379c565b15612855576040516303ce7bb760e11b81526001600160a01b038a8116600483015282169063079cf76e9060240160206040518083038186803b15801561272157600080fd5b505afa158015612735573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061275991906137ce565b85858151811061277957634e487b7160e01b600052603260045260246000fd5b602002602001018181525050806001600160a01b0316630ae070288a8787815181106127b557634e487b7160e01b600052603260045260246000fd5b60200260200101516040518363ffffffff1660e01b81526004016127ee9291906001600160a01b03929092168252602082015260400190565b60206040518083038186803b15801561280657600080fd5b505afa15801561281a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061283e919061379c565b61285057600095505050505050610a62565b6128e0565b60405163f936cc1560e01b81526001600160a01b038a8116600483015282169063f936cc159060240160206040518083038186803b15801561289657600080fd5b505afa1580156128aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128ce919061379c565b6128e057600095505050505050610a62565b50505080806128ee90613d21565b915050612592565b50604051630b6d99fb60e11b815273__$9176519fc5fb44025266bd62b274a40525$__906316db33f6906129309087908590600401613acf565b60206040518083038186803b15801561294857600080fd5b505af415801561295c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612980919061379c565b95945050505050565b600061299c826301ffc9a760e01b6129bc565b80156106a057506129b5826001600160e01b03196129bc565b1592915050565b60008060006129cb8585613404565b915091508180156129805750949350505050565b6000805b8251811015612a3e57612a1e8585858481518110612a1157634e487b7160e01b600052603260045260246000fd5b60200260200101516121ae565b612a2c576000915050610a62565b80612a3681613d21565b9150506129e3565b506001949350505050565b6000828152600384016020526040902060010154612aa95760405162461bcd60e51b815260206004820152601b60248201527f4e6f746172792f63726564656e7469616c206e6f7420666f756e6400000000006044820152606401610873565b60008281526003840160205260409020600601546001600160a01b031680612b135760405162461bcd60e51b815260206004820152601d60248201527f4e6f746172792f7375626a6563742063616e6e6f74206265207a65726f0000006044820152606401610873565b6001600160a01b038116600090815260058501602052604081208054909190612b3b90613d21565b9091555060408051608081018252338082526001600160a01b03848116602080850182815243868801818152606088018b815260008d815260048f0186528a9020985189549088166001600160a01b0319918216178a55935160018a0180549190981694169390931790955593516002870155516003909501949094558451918252928101869052909286917fb53d448559d4a1f97e1fd48262b64223f5bdafd54d429a33599236f73e900b1f910160405180910390a450505050565b6000838152600686016020908152604080832033845290915290205460ff1615612c645760405162461bcd60e51b815260206004820152601c60248201527f4e6f746172792f73656e64657220616c7265616479207369676e6564000000006044820152606401610873565b6000838152600386016020526040902060010154612fd4576001600160a01b038416600090815260018601602052604090205415612e2f576001600160a01b0384166000908152600180870160209081526040808420548452600389019091529091200154612ce357634e487b7160e01b600052600160045260246000fd5b6001600160a01b0380851660009081526001808801602090815260408084205484526003808b01835281852082516101408101845281548152948101548585015260028101548584015290810154606085015260048101546080850152600581015460ff8116151560a086015261010090819004871660c0860152600682015490961660e0850152600781018054835181860281018601909452808452959694959194860193909190830182828015612dc557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612da7575b50505050508152602001600882015481525050905043816020015110612e2d5760405162461bcd60e51b815260206004820152601d60248201527f4e6f746172792f626c6f636b206e756d6265722076696f6c6174696f6e0000006044820152606401610873565b505b60405180610140016040528060018152602001438152602001428152602001866000016000876001600160a01b03166001600160a01b0316815260200190815260200160002060008154612e8290613d21565b91829055508152602080820186905260006040808401829052336060808601919091526001600160a01b038a811660808088019190915260a08088018a905260c09788018b90528b865260038e81018852958590208951815589880151600182015594890151600286015592880151948401949094559286015160048301558501516005820180549587015160ff1990961691151591909117610100600160a81b031916610100958416860217905560e08501516006820180546001600160a01b0319169190931617909155918301518051612f6492600785019201906134ef565b5061012091909101516008909101556001600160a01b0384166000818152600187016020526040908190208590555133919085907fc6f6a1702b44006e35bd83045fd933510f3a924c1bb3877bff1061c14eb1c6be90612fc79043815260200190565b60405180910390a46131ea565b60008381526003808701602090815260408084208151610140810183528154815260018201548185015260028201548184015293810154606085015260048101546080850152600581015460ff8116151560a08601526001600160a01b0361010091829004811660c087015260068301541660e08601526007820180548451818702810187019095528085529294918601939290918301828280156130a257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613084575b505050505081526020016008820154815250509050846001600160a01b03168160e001516001600160a01b03161461311c5760405162461bcd60e51b815260206004820181905260248201527f4e6f746172792f64696765737420616c726561647920726567697374657265646044820152606401610873565b82816101200151146131705760405162461bcd60e51b815260206004820152601f60248201527f4e6f746172792f6d69736d6174636865642065766964656e636520726f6f74006044820152606401610873565b815181610100015151146131c65760405162461bcd60e51b815260206004820152601b60248201527f4e6f746172792f6d69736d617463686564207769746e657373657300000000006044820152606401610873565b6000848152600387016020526040812080549091906131e490613d21565b90915550505b600083815260068601602090815260408083203380855290835292819020805460ff19166001179055514381528592917f3c4c49076c09b6a1214e1ba437ac70fd0cff1c6ea5baf6df4c5f6ba937c4bf12910160405180910390a35050505050565b60008281526003840160205260408120600601546001600160a01b03163381146132af5760405162461bcd60e51b8152602060048201526014602482015273139bdd185c9e4bdddc9bdb99c81cdd589a9958dd60621b6044820152606401610873565b600084815260038601602052604090206005015460ff16156133135760405162461bcd60e51b815260206004820181905260248201527f4e6f746172792f63726564656e7469616c20616c7265616479207369676e65646044820152606401610873565b60008481526003860160205260409020548311156133735760405162461bcd60e51b815260206004820152601e60248201527f4e6f746172792f6e6f2071756f72756d206f66207369676e61747572657300006044820152606401610873565b60008481526003860160209081526040808320600501805460ff191660019081179091556001600160a01b0385168452600289018352818420805491820181558452919092200185905551849033907f3c4c49076c09b6a1214e1ba437ac70fd0cff1c6ea5baf6df4c5f6ba937c4bf12906133f19043815260200190565b60405180910390a3506001949350505050565b604080516001600160e01b0319831660248083019190915282518083039091018152604490910182526020810180516001600160e01b03166301ffc9a760e01b17905290516000918291829081906001600160a01b038816906175309061346c908690613a2c565b6000604051808303818686fa925050503d80600081146134a8576040519150601f19603f3d011682016040523d82523d6000602084013e6134ad565b606091505b50915091506020815110156134cb57600080945094505050506134e8565b81818060200190518101906134e0919061379c565b945094505050505b9250929050565b828054828255906000526020600020908101928215613544579160200282015b8281111561354457825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019061350f565b50613550929150613554565b5090565b5b808211156135505760008155600101613555565b80516106a381613d88565b600082601f830112613584578081fd5b8151602061359961359483613cfd565b613ccc565b82815281810190858301838502870184018810156135b5578586fd5b855b858110156135dc5781516135ca81613d88565b845292840192908401906001016135b7565b5090979650505050505050565b805180151581146106a357600080fd5b60006020828403121561360a578081fd5b8135610a6281613d88565b60008060408385031215613627578081fd5b823561363281613d88565b915060208381013567ffffffffffffffff81111561364e578283fd5b8401601f8101861361365e578283fd5b803561366c61359482613cfd565b81815283810190838501858402850186018a1015613688578687fd5b8694505b838510156136aa57803583526001949094019391850191850161368c565b5080955050505050509250929050565b600080604083850312156136cc578182fd5b82356136d781613d88565b946020939093013593505050565b6000806000606084860312156136f9578081fd5b833561370481613d88565b92506020848101359250604085013567ffffffffffffffff811115613727578283fd5b8501601f81018713613737578283fd5b803561374561359482613cfd565b81815283810190838501858402850186018b1015613761578687fd5b8694505b8385101561378c57803561377881613d88565b835260019490940193918501918501613765565b5080955050505050509250925092565b6000602082840312156137ad578081fd5b61083b826135e9565b6000602082840312156137c7578081fd5b5035919050565b6000602082840312156137df578081fd5b5051919050565b600080604083850312156137f8578182fd5b82359150602083013561380a81613d88565b809150509250929050565b60008060408385031215613827578182fd5b50508035926020909101359150565b600060208284031215613847578081fd5b81356001600160e01b031981168114610a62578182fd5b60006020828403121561386f578081fd5b815160028110610a62578182fd5b60006020828403121561388e578081fd5b815167ffffffffffffffff808211156138a5578283fd5b81840191506101408083870312156138bb578384fd5b6138c481613ccc565b905082518152602083015160208201526040830151604082015260608301516060820152608083015160808201526138fe60a084016135e9565b60a082015261390f60c08401613569565b60c082015261392060e08401613569565b60e08201526101008084015183811115613938578586fd5b61394488828701613574565b91830191909152506101209283015192810192909252509392505050565b600060808284031215613973578081fd5b61397d6080613ccc565b825161398881613d88565b8152602083015161399881613d88565b6020820152604083810151908201526060928301519281019290925250919050565b6000815180845260208085019450808401835b838110156139f25781516001600160a01b0316875295820195908201906001016139cd565b509495945050505050565b6000815180845260208085019450808401835b838110156139f257815187529582019590820190600101613a10565b60008251815b81811015613a4c5760208186018101518583015201613a32565b81811115613a5a5782828501525b509190910192915050565b60006020825261083b60208301846139ba565b60006020825261083b60208301846139fd565b6020808252825182820181905260009190848201906040850190845b81811015613ac357835183529284019291840191600101613aa7565b50909695505050505050565b600083825260406020830152611d0f60408301846139fd565b6020810160028310613b0a57634e487b7160e01b600052602160045260246000fd5b91905290565b6020808252601f908201527f4973737565722f746865726520617265206e6f2063726564656e7469616c7300604082015260600190565b6020808252601d908201527f4f776e6572732f73656e646572206973206e6f7420616e206f776e6572000000604082015260600190565b60208082526019908201527f4973737565722f63726564656e7469616c207265766f6b656400000000000000604082015260600190565b60006020825282516020830152602083015160408301526040830151606083015260608301516080830152608083015160a083015260a0830151613bfd60c084018215159052565b5060c08301516001600160a01b03811660e08401525060e0830151610100613c2f818501836001600160a01b03169052565b808501519150506101406101208181860152613c4f6101608601846139ba565b9501519301929092525090919050565b8381526001600160a01b0383166020820152606060408201819052600090612980908301846139fd565b60006040820184835260206040818501528185548084526060860191508685528285209350845b818110156135dc57845483526001948501949284019201613cb0565b604051601f8201601f1916810167ffffffffffffffff81118282101715613cf557613cf5613d72565b604052919050565b600067ffffffffffffffff821115613d1757613d17613d72565b5060209081020190565b6000600019821415613d3557613d35613d5c565b5060010190565b600060ff821660ff811415613d5357613d53613d5c565b60010192915050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461198e57600080fdfea26469706673582212206fd767caab50cdbb7ddbe1f282496bbbc5cd39518b962110ebaaf338c6c8b5be64736f6c63430008020033"

// DeployNode deploys a new Ethereum contract, binding an instance of Node to it.
func DeployNode(auth *bind.TransactOpts, backend bind.ContractBackend, role uint8, registrars []common.Address, quorum uint8) (common.Address, *types.Transaction, *Node, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	notaryAddr, _, _, _ := DeployNotary(auth, backend)
	NodeBin = strings.Replace(NodeBin, "__$1d2b6a72c2c8582bd594e19b02f34b62b6$__", notaryAddr.String()[2:], -1)

	credentialSumAddr, _, _, _ := DeployCredentialSum(auth, backend)
	NodeBin = strings.Replace(NodeBin, "__$9176519fc5fb44025266bd62b274a40525$__", credentialSumAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeBin), backend, role, registrars, quorum)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Node{NodeCaller: NodeCaller{contract: contract}, NodeTransactor: NodeTransactor{contract: contract}, NodeFilterer: NodeFilterer{contract: contract}}, nil
}

// Node is an auto generated Go binding around an Ethereum contract.
type Node struct {
	NodeCaller     // Read-only binding to the contract
	NodeTransactor // Write-only binding to the contract
	NodeFilterer   // Log filterer for contract events
}

// NodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeSession struct {
	Contract     *Node             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeCallerSession struct {
	Contract *NodeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeTransactorSession struct {
	Contract     *NodeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeRaw struct {
	Contract *Node // Generic contract binding to access the raw methods on
}

// NodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeCallerRaw struct {
	Contract *NodeCaller // Generic read-only contract binding to access the raw methods on
}

// NodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeTransactorRaw struct {
	Contract *NodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNode creates a new instance of Node, bound to a specific deployed contract.
func NewNode(address common.Address, backend bind.ContractBackend) (*Node, error) {
	contract, err := bindNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Node{NodeCaller: NodeCaller{contract: contract}, NodeTransactor: NodeTransactor{contract: contract}, NodeFilterer: NodeFilterer{contract: contract}}, nil
}

// NewNodeCaller creates a new read-only instance of Node, bound to a specific deployed contract.
func NewNodeCaller(address common.Address, caller bind.ContractCaller) (*NodeCaller, error) {
	contract, err := bindNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeCaller{contract: contract}, nil
}

// NewNodeTransactor creates a new write-only instance of Node, bound to a specific deployed contract.
func NewNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeTransactor, error) {
	contract, err := bindNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeTransactor{contract: contract}, nil
}

// NewNodeFilterer creates a new log filterer instance of Node, bound to a specific deployed contract.
func NewNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeFilterer, error) {
	contract, err := bindNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeFilterer{contract: contract}, nil
}

// bindNode binds a generic wrapper to an already deployed contract.
func bindNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.NodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.NodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Node *NodeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Node.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Node *NodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Node.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Node *NodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Node.Contract.contract.Transact(opts, method, params...)
}

// GetChildren is a free data retrieval call binding the contract method 0x4892e8e8.
//
// Solidity: function getChildren() view returns(address[])
func (_Node *NodeCaller) GetChildren(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getChildren")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetChildren is a free data retrieval call binding the contract method 0x4892e8e8.
//
// Solidity: function getChildren() view returns(address[])
func (_Node *NodeSession) GetChildren() ([]common.Address, error) {
	return _Node.Contract.GetChildren(&_Node.CallOpts)
}

// GetChildren is a free data retrieval call binding the contract method 0x4892e8e8.
//
// Solidity: function getChildren() view returns(address[])
func (_Node *NodeCallerSession) GetChildren() ([]common.Address, error) {
	return _Node.Contract.GetChildren(&_Node.CallOpts)
}

// GetCredentialProof is a free data retrieval call binding the contract method 0xdac46a62.
//
// Solidity: function getCredentialProof(bytes32 digest) view returns((uint256,uint256,uint256,uint256,bytes32,bool,address,address,address[],bytes32))
func (_Node *NodeCaller) GetCredentialProof(opts *bind.CallOpts, digest [32]byte) (NotaryCredentialProof, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getCredentialProof", digest)

	if err != nil {
		return *new(NotaryCredentialProof), err
	}

	out0 := *abi.ConvertType(out[0], new(NotaryCredentialProof)).(*NotaryCredentialProof)

	return out0, err

}

// GetCredentialProof is a free data retrieval call binding the contract method 0xdac46a62.
//
// Solidity: function getCredentialProof(bytes32 digest) view returns((uint256,uint256,uint256,uint256,bytes32,bool,address,address,address[],bytes32))
func (_Node *NodeSession) GetCredentialProof(digest [32]byte) (NotaryCredentialProof, error) {
	return _Node.Contract.GetCredentialProof(&_Node.CallOpts, digest)
}

// GetCredentialProof is a free data retrieval call binding the contract method 0xdac46a62.
//
// Solidity: function getCredentialProof(bytes32 digest) view returns((uint256,uint256,uint256,uint256,bytes32,bool,address,address,address[],bytes32))
func (_Node *NodeCallerSession) GetCredentialProof(digest [32]byte) (NotaryCredentialProof, error) {
	return _Node.Contract.GetCredentialProof(&_Node.CallOpts, digest)
}

// GetCredentialSigners is a free data retrieval call binding the contract method 0x038ec1b8.
//
// Solidity: function getCredentialSigners(bytes32 digest) view returns(address[])
func (_Node *NodeCaller) GetCredentialSigners(opts *bind.CallOpts, digest [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getCredentialSigners", digest)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCredentialSigners is a free data retrieval call binding the contract method 0x038ec1b8.
//
// Solidity: function getCredentialSigners(bytes32 digest) view returns(address[])
func (_Node *NodeSession) GetCredentialSigners(digest [32]byte) ([]common.Address, error) {
	return _Node.Contract.GetCredentialSigners(&_Node.CallOpts, digest)
}

// GetCredentialSigners is a free data retrieval call binding the contract method 0x038ec1b8.
//
// Solidity: function getCredentialSigners(bytes32 digest) view returns(address[])
func (_Node *NodeCallerSession) GetCredentialSigners(digest [32]byte) ([]common.Address, error) {
	return _Node.Contract.GetCredentialSigners(&_Node.CallOpts, digest)
}

// GetDigests is a free data retrieval call binding the contract method 0x762b77bb.
//
// Solidity: function getDigests(address subject) view returns(bytes32[])
func (_Node *NodeCaller) GetDigests(opts *bind.CallOpts, subject common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getDigests", subject)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetDigests is a free data retrieval call binding the contract method 0x762b77bb.
//
// Solidity: function getDigests(address subject) view returns(bytes32[])
func (_Node *NodeSession) GetDigests(subject common.Address) ([][32]byte, error) {
	return _Node.Contract.GetDigests(&_Node.CallOpts, subject)
}

// GetDigests is a free data retrieval call binding the contract method 0x762b77bb.
//
// Solidity: function getDigests(address subject) view returns(bytes32[])
func (_Node *NodeCallerSession) GetDigests(subject common.Address) ([][32]byte, error) {
	return _Node.Contract.GetDigests(&_Node.CallOpts, subject)
}

// GetEvidenceRoot is a free data retrieval call binding the contract method 0x4302802d.
//
// Solidity: function getEvidenceRoot(bytes32 digest) view returns(bytes32)
func (_Node *NodeCaller) GetEvidenceRoot(opts *bind.CallOpts, digest [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getEvidenceRoot", digest)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEvidenceRoot is a free data retrieval call binding the contract method 0x4302802d.
//
// Solidity: function getEvidenceRoot(bytes32 digest) view returns(bytes32)
func (_Node *NodeSession) GetEvidenceRoot(digest [32]byte) ([32]byte, error) {
	return _Node.Contract.GetEvidenceRoot(&_Node.CallOpts, digest)
}

// GetEvidenceRoot is a free data retrieval call binding the contract method 0x4302802d.
//
// Solidity: function getEvidenceRoot(bytes32 digest) view returns(bytes32)
func (_Node *NodeCallerSession) GetEvidenceRoot(digest [32]byte) ([32]byte, error) {
	return _Node.Contract.GetEvidenceRoot(&_Node.CallOpts, digest)
}

// GetProof is a free data retrieval call binding the contract method 0x3eea79d1.
//
// Solidity: function getProof(address subject) view returns((bytes32,uint256,uint256))
func (_Node *NodeCaller) GetProof(opts *bind.CallOpts, subject common.Address) (CredentialSumRoot, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getProof", subject)

	if err != nil {
		return *new(CredentialSumRoot), err
	}

	out0 := *abi.ConvertType(out[0], new(CredentialSumRoot)).(*CredentialSumRoot)

	return out0, err

}

// GetProof is a free data retrieval call binding the contract method 0x3eea79d1.
//
// Solidity: function getProof(address subject) view returns((bytes32,uint256,uint256))
func (_Node *NodeSession) GetProof(subject common.Address) (CredentialSumRoot, error) {
	return _Node.Contract.GetProof(&_Node.CallOpts, subject)
}

// GetProof is a free data retrieval call binding the contract method 0x3eea79d1.
//
// Solidity: function getProof(address subject) view returns((bytes32,uint256,uint256))
func (_Node *NodeCallerSession) GetProof(subject common.Address) (CredentialSumRoot, error) {
	return _Node.Contract.GetProof(&_Node.CallOpts, subject)
}

// GetRevoked is a free data retrieval call binding the contract method 0x8ebd9623.
//
// Solidity: function getRevoked(address subject) view returns(bytes32[])
func (_Node *NodeCaller) GetRevoked(opts *bind.CallOpts, subject common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getRevoked", subject)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetRevoked is a free data retrieval call binding the contract method 0x8ebd9623.
//
// Solidity: function getRevoked(address subject) view returns(bytes32[])
func (_Node *NodeSession) GetRevoked(subject common.Address) ([][32]byte, error) {
	return _Node.Contract.GetRevoked(&_Node.CallOpts, subject)
}

// GetRevoked is a free data retrieval call binding the contract method 0x8ebd9623.
//
// Solidity: function getRevoked(address subject) view returns(bytes32[])
func (_Node *NodeCallerSession) GetRevoked(subject common.Address) ([][32]byte, error) {
	return _Node.Contract.GetRevoked(&_Node.CallOpts, subject)
}

// GetRevokedProof is a free data retrieval call binding the contract method 0xabbff984.
//
// Solidity: function getRevokedProof(bytes32 digest) view returns((address,address,uint256,bytes32))
func (_Node *NodeCaller) GetRevokedProof(opts *bind.CallOpts, digest [32]byte) (NotaryRevocationProof, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getRevokedProof", digest)

	if err != nil {
		return *new(NotaryRevocationProof), err
	}

	out0 := *abi.ConvertType(out[0], new(NotaryRevocationProof)).(*NotaryRevocationProof)

	return out0, err

}

// GetRevokedProof is a free data retrieval call binding the contract method 0xabbff984.
//
// Solidity: function getRevokedProof(bytes32 digest) view returns((address,address,uint256,bytes32))
func (_Node *NodeSession) GetRevokedProof(digest [32]byte) (NotaryRevocationProof, error) {
	return _Node.Contract.GetRevokedProof(&_Node.CallOpts, digest)
}

// GetRevokedProof is a free data retrieval call binding the contract method 0xabbff984.
//
// Solidity: function getRevokedProof(bytes32 digest) view returns((address,address,uint256,bytes32))
func (_Node *NodeCallerSession) GetRevokedProof(digest [32]byte) (NotaryRevocationProof, error) {
	return _Node.Contract.GetRevokedProof(&_Node.CallOpts, digest)
}

// GetRole is a free data retrieval call binding the contract method 0xfde08174.
//
// Solidity: function getRole() view returns(uint8)
func (_Node *NodeCaller) GetRole(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getRole")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetRole is a free data retrieval call binding the contract method 0xfde08174.
//
// Solidity: function getRole() view returns(uint8)
func (_Node *NodeSession) GetRole() (uint8, error) {
	return _Node.Contract.GetRole(&_Node.CallOpts)
}

// GetRole is a free data retrieval call binding the contract method 0xfde08174.
//
// Solidity: function getRole() view returns(uint8)
func (_Node *NodeCallerSession) GetRole() (uint8, error) {
	return _Node.Contract.GetRole(&_Node.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x079cf76e.
//
// Solidity: function getRoot(address subject) view returns(bytes32)
func (_Node *NodeCaller) GetRoot(opts *bind.CallOpts, subject common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getRoot", subject)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x079cf76e.
//
// Solidity: function getRoot(address subject) view returns(bytes32)
func (_Node *NodeSession) GetRoot(subject common.Address) ([32]byte, error) {
	return _Node.Contract.GetRoot(&_Node.CallOpts, subject)
}

// GetRoot is a free data retrieval call binding the contract method 0x079cf76e.
//
// Solidity: function getRoot(address subject) view returns(bytes32)
func (_Node *NodeCallerSession) GetRoot(subject common.Address) ([32]byte, error) {
	return _Node.Contract.GetRoot(&_Node.CallOpts, subject)
}

// GetWitnesses is a free data retrieval call binding the contract method 0x6f28eca6.
//
// Solidity: function getWitnesses(bytes32 digest) view returns(address[])
func (_Node *NodeCaller) GetWitnesses(opts *bind.CallOpts, digest [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "getWitnesses", digest)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWitnesses is a free data retrieval call binding the contract method 0x6f28eca6.
//
// Solidity: function getWitnesses(bytes32 digest) view returns(address[])
func (_Node *NodeSession) GetWitnesses(digest [32]byte) ([]common.Address, error) {
	return _Node.Contract.GetWitnesses(&_Node.CallOpts, digest)
}

// GetWitnesses is a free data retrieval call binding the contract method 0x6f28eca6.
//
// Solidity: function getWitnesses(bytes32 digest) view returns(address[])
func (_Node *NodeCallerSession) GetWitnesses(digest [32]byte) ([]common.Address, error) {
	return _Node.Contract.GetWitnesses(&_Node.CallOpts, digest)
}

// HasRoot is a free data retrieval call binding the contract method 0x7532ed26.
//
// Solidity: function hasRoot(address subject) view returns(bool)
func (_Node *NodeCaller) HasRoot(opts *bind.CallOpts, subject common.Address) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "hasRoot", subject)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRoot is a free data retrieval call binding the contract method 0x7532ed26.
//
// Solidity: function hasRoot(address subject) view returns(bool)
func (_Node *NodeSession) HasRoot(subject common.Address) (bool, error) {
	return _Node.Contract.HasRoot(&_Node.CallOpts, subject)
}

// HasRoot is a free data retrieval call binding the contract method 0x7532ed26.
//
// Solidity: function hasRoot(address subject) view returns(bool)
func (_Node *NodeCallerSession) HasRoot(subject common.Address) (bool, error) {
	return _Node.Contract.HasRoot(&_Node.CallOpts, subject)
}

// IsApproved is a free data retrieval call binding the contract method 0x48aefc32.
//
// Solidity: function isApproved(bytes32 digest) view returns(bool)
func (_Node *NodeCaller) IsApproved(opts *bind.CallOpts, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "isApproved", digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0x48aefc32.
//
// Solidity: function isApproved(bytes32 digest) view returns(bool)
func (_Node *NodeSession) IsApproved(digest [32]byte) (bool, error) {
	return _Node.Contract.IsApproved(&_Node.CallOpts, digest)
}

// IsApproved is a free data retrieval call binding the contract method 0x48aefc32.
//
// Solidity: function isApproved(bytes32 digest) view returns(bool)
func (_Node *NodeCallerSession) IsApproved(digest [32]byte) (bool, error) {
	return _Node.Contract.IsApproved(&_Node.CallOpts, digest)
}

// IsChild is a free data retrieval call binding the contract method 0xfc91a897.
//
// Solidity: function isChild(address ) view returns(bool)
func (_Node *NodeCaller) IsChild(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "isChild", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChild is a free data retrieval call binding the contract method 0xfc91a897.
//
// Solidity: function isChild(address ) view returns(bool)
func (_Node *NodeSession) IsChild(arg0 common.Address) (bool, error) {
	return _Node.Contract.IsChild(&_Node.CallOpts, arg0)
}

// IsChild is a free data retrieval call binding the contract method 0xfc91a897.
//
// Solidity: function isChild(address ) view returns(bool)
func (_Node *NodeCallerSession) IsChild(arg0 common.Address) (bool, error) {
	return _Node.Contract.IsChild(&_Node.CallOpts, arg0)
}

// IsLeaf is a free data retrieval call binding the contract method 0x99d1dbc0.
//
// Solidity: function isLeaf() view returns(bool)
func (_Node *NodeCaller) IsLeaf(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "isLeaf")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLeaf is a free data retrieval call binding the contract method 0x99d1dbc0.
//
// Solidity: function isLeaf() view returns(bool)
func (_Node *NodeSession) IsLeaf() (bool, error) {
	return _Node.Contract.IsLeaf(&_Node.CallOpts)
}

// IsLeaf is a free data retrieval call binding the contract method 0x99d1dbc0.
//
// Solidity: function isLeaf() view returns(bool)
func (_Node *NodeCallerSession) IsLeaf() (bool, error) {
	return _Node.Contract.IsLeaf(&_Node.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_Node *NodeCaller) IsOwner(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "isOwner", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_Node *NodeSession) IsOwner(account common.Address) (bool, error) {
	return _Node.Contract.IsOwner(&_Node.CallOpts, account)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_Node *NodeCallerSession) IsOwner(account common.Address) (bool, error) {
	return _Node.Contract.IsOwner(&_Node.CallOpts, account)
}

// IsQuorumSigned is a free data retrieval call binding the contract method 0xae02c0cc.
//
// Solidity: function isQuorumSigned(bytes32 digest) view returns(bool)
func (_Node *NodeCaller) IsQuorumSigned(opts *bind.CallOpts, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "isQuorumSigned", digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsQuorumSigned is a free data retrieval call binding the contract method 0xae02c0cc.
//
// Solidity: function isQuorumSigned(bytes32 digest) view returns(bool)
func (_Node *NodeSession) IsQuorumSigned(digest [32]byte) (bool, error) {
	return _Node.Contract.IsQuorumSigned(&_Node.CallOpts, digest)
}

// IsQuorumSigned is a free data retrieval call binding the contract method 0xae02c0cc.
//
// Solidity: function isQuorumSigned(bytes32 digest) view returns(bool)
func (_Node *NodeCallerSession) IsQuorumSigned(digest [32]byte) (bool, error) {
	return _Node.Contract.IsQuorumSigned(&_Node.CallOpts, digest)
}

// IsRevoked is a free data retrieval call binding the contract method 0x4294857f.
//
// Solidity: function isRevoked(bytes32 digest) view returns(bool)
func (_Node *NodeCaller) IsRevoked(opts *bind.CallOpts, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "isRevoked", digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRevoked is a free data retrieval call binding the contract method 0x4294857f.
//
// Solidity: function isRevoked(bytes32 digest) view returns(bool)
func (_Node *NodeSession) IsRevoked(digest [32]byte) (bool, error) {
	return _Node.Contract.IsRevoked(&_Node.CallOpts, digest)
}

// IsRevoked is a free data retrieval call binding the contract method 0x4294857f.
//
// Solidity: function isRevoked(bytes32 digest) view returns(bool)
func (_Node *NodeCallerSession) IsRevoked(digest [32]byte) (bool, error) {
	return _Node.Contract.IsRevoked(&_Node.CallOpts, digest)
}

// IsSigned is a free data retrieval call binding the contract method 0x7f8d61e0.
//
// Solidity: function isSigned(bytes32 digest, address account) view returns(bool)
func (_Node *NodeCaller) IsSigned(opts *bind.CallOpts, digest [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "isSigned", digest, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigned is a free data retrieval call binding the contract method 0x7f8d61e0.
//
// Solidity: function isSigned(bytes32 digest, address account) view returns(bool)
func (_Node *NodeSession) IsSigned(digest [32]byte, account common.Address) (bool, error) {
	return _Node.Contract.IsSigned(&_Node.CallOpts, digest, account)
}

// IsSigned is a free data retrieval call binding the contract method 0x7f8d61e0.
//
// Solidity: function isSigned(bytes32 digest, address account) view returns(bool)
func (_Node *NodeCallerSession) IsSigned(digest [32]byte, account common.Address) (bool, error) {
	return _Node.Contract.IsSigned(&_Node.CallOpts, digest, account)
}

// MyParent is a free data retrieval call binding the contract method 0xde07fc45.
//
// Solidity: function myParent() view returns(address)
func (_Node *NodeCaller) MyParent(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "myParent")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MyParent is a free data retrieval call binding the contract method 0xde07fc45.
//
// Solidity: function myParent() view returns(address)
func (_Node *NodeSession) MyParent() (common.Address, error) {
	return _Node.Contract.MyParent(&_Node.CallOpts)
}

// MyParent is a free data retrieval call binding the contract method 0xde07fc45.
//
// Solidity: function myParent() view returns(address)
func (_Node *NodeCallerSession) MyParent() (common.Address, error) {
	return _Node.Contract.MyParent(&_Node.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address[])
func (_Node *NodeCaller) Owners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "owners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address[])
func (_Node *NodeSession) Owners() ([]common.Address, error) {
	return _Node.Contract.Owners(&_Node.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address[])
func (_Node *NodeCallerSession) Owners() ([]common.Address, error) {
	return _Node.Contract.Owners(&_Node.CallOpts)
}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() view returns(uint8)
func (_Node *NodeCaller) OwnersCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "ownersCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() view returns(uint8)
func (_Node *NodeSession) OwnersCount() (uint8, error) {
	return _Node.Contract.OwnersCount(&_Node.CallOpts)
}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() view returns(uint8)
func (_Node *NodeCallerSession) OwnersCount() (uint8, error) {
	return _Node.Contract.OwnersCount(&_Node.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint8)
func (_Node *NodeCaller) Quorum(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "quorum")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint8)
func (_Node *NodeSession) Quorum() (uint8, error) {
	return _Node.Contract.Quorum(&_Node.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint8)
func (_Node *NodeCallerSession) Quorum() (uint8, error) {
	return _Node.Contract.Quorum(&_Node.CallOpts)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 digest) view returns(bool)
func (_Node *NodeCaller) RecordExists(opts *bind.CallOpts, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "recordExists", digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 digest) view returns(bool)
func (_Node *NodeSession) RecordExists(digest [32]byte) (bool, error) {
	return _Node.Contract.RecordExists(&_Node.CallOpts, digest)
}

// RecordExists is a free data retrieval call binding the contract method 0xf79fe538.
//
// Solidity: function recordExists(bytes32 digest) view returns(bool)
func (_Node *NodeCallerSession) RecordExists(digest [32]byte) (bool, error) {
	return _Node.Contract.RecordExists(&_Node.CallOpts, digest)
}

// RevokedCounter is a free data retrieval call binding the contract method 0x342577fa.
//
// Solidity: function revokedCounter(address subject) view returns(uint256)
func (_Node *NodeCaller) RevokedCounter(opts *bind.CallOpts, subject common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "revokedCounter", subject)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RevokedCounter is a free data retrieval call binding the contract method 0x342577fa.
//
// Solidity: function revokedCounter(address subject) view returns(uint256)
func (_Node *NodeSession) RevokedCounter(subject common.Address) (*big.Int, error) {
	return _Node.Contract.RevokedCounter(&_Node.CallOpts, subject)
}

// RevokedCounter is a free data retrieval call binding the contract method 0x342577fa.
//
// Solidity: function revokedCounter(address subject) view returns(uint256)
func (_Node *NodeCallerSession) RevokedCounter(subject common.Address) (*big.Int, error) {
	return _Node.Contract.RevokedCounter(&_Node.CallOpts, subject)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Node *NodeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Node *NodeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Node.Contract.SupportsInterface(&_Node.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Node *NodeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Node.Contract.SupportsInterface(&_Node.CallOpts, interfaceId)
}

// VerifyCredential is a free data retrieval call binding the contract method 0x5f889e17.
//
// Solidity: function verifyCredential(address subject, bytes32 digest) view returns(bool)
func (_Node *NodeCaller) VerifyCredential(opts *bind.CallOpts, subject common.Address, digest [32]byte) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "verifyCredential", subject, digest)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyCredential is a free data retrieval call binding the contract method 0x5f889e17.
//
// Solidity: function verifyCredential(address subject, bytes32 digest) view returns(bool)
func (_Node *NodeSession) VerifyCredential(subject common.Address, digest [32]byte) (bool, error) {
	return _Node.Contract.VerifyCredential(&_Node.CallOpts, subject, digest)
}

// VerifyCredential is a free data retrieval call binding the contract method 0x5f889e17.
//
// Solidity: function verifyCredential(address subject, bytes32 digest) view returns(bool)
func (_Node *NodeCallerSession) VerifyCredential(subject common.Address, digest [32]byte) (bool, error) {
	return _Node.Contract.VerifyCredential(&_Node.CallOpts, subject, digest)
}

// VerifyCredentialRoot is a free data retrieval call binding the contract method 0x0ae07028.
//
// Solidity: function verifyCredentialRoot(address subject, bytes32 root) view returns(bool)
func (_Node *NodeCaller) VerifyCredentialRoot(opts *bind.CallOpts, subject common.Address, root [32]byte) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "verifyCredentialRoot", subject, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyCredentialRoot is a free data retrieval call binding the contract method 0x0ae07028.
//
// Solidity: function verifyCredentialRoot(address subject, bytes32 root) view returns(bool)
func (_Node *NodeSession) VerifyCredentialRoot(subject common.Address, root [32]byte) (bool, error) {
	return _Node.Contract.VerifyCredentialRoot(&_Node.CallOpts, subject, root)
}

// VerifyCredentialRoot is a free data retrieval call binding the contract method 0x0ae07028.
//
// Solidity: function verifyCredentialRoot(address subject, bytes32 root) view returns(bool)
func (_Node *NodeCallerSession) VerifyCredentialRoot(subject common.Address, root [32]byte) (bool, error) {
	return _Node.Contract.VerifyCredentialRoot(&_Node.CallOpts, subject, root)
}

// VerifyCredentialTree is a free data retrieval call binding the contract method 0xf936cc15.
//
// Solidity: function verifyCredentialTree(address subject) view returns(bool)
func (_Node *NodeCaller) VerifyCredentialTree(opts *bind.CallOpts, subject common.Address) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "verifyCredentialTree", subject)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyCredentialTree is a free data retrieval call binding the contract method 0xf936cc15.
//
// Solidity: function verifyCredentialTree(address subject) view returns(bool)
func (_Node *NodeSession) VerifyCredentialTree(subject common.Address) (bool, error) {
	return _Node.Contract.VerifyCredentialTree(&_Node.CallOpts, subject)
}

// VerifyCredentialTree is a free data retrieval call binding the contract method 0xf936cc15.
//
// Solidity: function verifyCredentialTree(address subject) view returns(bool)
func (_Node *NodeCallerSession) VerifyCredentialTree(subject common.Address) (bool, error) {
	return _Node.Contract.VerifyCredentialTree(&_Node.CallOpts, subject)
}

// VerifyIssuedCredentials is a free data retrieval call binding the contract method 0x2cf0695a.
//
// Solidity: function verifyIssuedCredentials(address subject) view returns(bool)
func (_Node *NodeCaller) VerifyIssuedCredentials(opts *bind.CallOpts, subject common.Address) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "verifyIssuedCredentials", subject)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyIssuedCredentials is a free data retrieval call binding the contract method 0x2cf0695a.
//
// Solidity: function verifyIssuedCredentials(address subject) view returns(bool)
func (_Node *NodeSession) VerifyIssuedCredentials(subject common.Address) (bool, error) {
	return _Node.Contract.VerifyIssuedCredentials(&_Node.CallOpts, subject)
}

// VerifyIssuedCredentials is a free data retrieval call binding the contract method 0x2cf0695a.
//
// Solidity: function verifyIssuedCredentials(address subject) view returns(bool)
func (_Node *NodeCallerSession) VerifyIssuedCredentials(subject common.Address) (bool, error) {
	return _Node.Contract.VerifyIssuedCredentials(&_Node.CallOpts, subject)
}

// VerifyRootOf is a free data retrieval call binding the contract method 0x86509dd6.
//
// Solidity: function verifyRootOf(address subject, bytes32[] digests) view returns(bool)
func (_Node *NodeCaller) VerifyRootOf(opts *bind.CallOpts, subject common.Address, digests [][32]byte) (bool, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "verifyRootOf", subject, digests)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyRootOf is a free data retrieval call binding the contract method 0x86509dd6.
//
// Solidity: function verifyRootOf(address subject, bytes32[] digests) view returns(bool)
func (_Node *NodeSession) VerifyRootOf(subject common.Address, digests [][32]byte) (bool, error) {
	return _Node.Contract.VerifyRootOf(&_Node.CallOpts, subject, digests)
}

// VerifyRootOf is a free data retrieval call binding the contract method 0x86509dd6.
//
// Solidity: function verifyRootOf(address subject, bytes32[] digests) view returns(bool)
func (_Node *NodeCallerSession) VerifyRootOf(subject common.Address, digests [][32]byte) (bool, error) {
	return _Node.Contract.VerifyRootOf(&_Node.CallOpts, subject, digests)
}

// WitnessesLength is a free data retrieval call binding the contract method 0xf537ff9d.
//
// Solidity: function witnessesLength(bytes32 digest) view returns(uint256)
func (_Node *NodeCaller) WitnessesLength(opts *bind.CallOpts, digest [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Node.contract.Call(opts, &out, "witnessesLength", digest)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WitnessesLength is a free data retrieval call binding the contract method 0xf537ff9d.
//
// Solidity: function witnessesLength(bytes32 digest) view returns(uint256)
func (_Node *NodeSession) WitnessesLength(digest [32]byte) (*big.Int, error) {
	return _Node.Contract.WitnessesLength(&_Node.CallOpts, digest)
}

// WitnessesLength is a free data retrieval call binding the contract method 0xf537ff9d.
//
// Solidity: function witnessesLength(bytes32 digest) view returns(uint256)
func (_Node *NodeCallerSession) WitnessesLength(digest [32]byte) (*big.Int, error) {
	return _Node.Contract.WitnessesLength(&_Node.CallOpts, digest)
}

// AddChild is a paid mutator transaction binding the contract method 0x1eee993a.
//
// Solidity: function addChild(address nodeAddress) returns()
func (_Node *NodeTransactor) AddChild(opts *bind.TransactOpts, nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "addChild", nodeAddress)
}

// AddChild is a paid mutator transaction binding the contract method 0x1eee993a.
//
// Solidity: function addChild(address nodeAddress) returns()
func (_Node *NodeSession) AddChild(nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.Contract.AddChild(&_Node.TransactOpts, nodeAddress)
}

// AddChild is a paid mutator transaction binding the contract method 0x1eee993a.
//
// Solidity: function addChild(address nodeAddress) returns()
func (_Node *NodeTransactorSession) AddChild(nodeAddress common.Address) (*types.Transaction, error) {
	return _Node.Contract.AddChild(&_Node.TransactOpts, nodeAddress)
}

// AggregateCredentials is a paid mutator transaction binding the contract method 0xe2d02cea.
//
// Solidity: function aggregateCredentials(address subject, bytes32[] digests) returns(bytes32)
func (_Node *NodeTransactor) AggregateCredentials(opts *bind.TransactOpts, subject common.Address, digests [][32]byte) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "aggregateCredentials", subject, digests)
}

// AggregateCredentials is a paid mutator transaction binding the contract method 0xe2d02cea.
//
// Solidity: function aggregateCredentials(address subject, bytes32[] digests) returns(bytes32)
func (_Node *NodeSession) AggregateCredentials(subject common.Address, digests [][32]byte) (*types.Transaction, error) {
	return _Node.Contract.AggregateCredentials(&_Node.TransactOpts, subject, digests)
}

// AggregateCredentials is a paid mutator transaction binding the contract method 0xe2d02cea.
//
// Solidity: function aggregateCredentials(address subject, bytes32[] digests) returns(bytes32)
func (_Node *NodeTransactorSession) AggregateCredentials(subject common.Address, digests [][32]byte) (*types.Transaction, error) {
	return _Node.Contract.AggregateCredentials(&_Node.TransactOpts, subject, digests)
}

// ApproveCredential is a paid mutator transaction binding the contract method 0xeedb973c.
//
// Solidity: function approveCredential(bytes32 digest) returns()
func (_Node *NodeTransactor) ApproveCredential(opts *bind.TransactOpts, digest [32]byte) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "approveCredential", digest)
}

// ApproveCredential is a paid mutator transaction binding the contract method 0xeedb973c.
//
// Solidity: function approveCredential(bytes32 digest) returns()
func (_Node *NodeSession) ApproveCredential(digest [32]byte) (*types.Transaction, error) {
	return _Node.Contract.ApproveCredential(&_Node.TransactOpts, digest)
}

// ApproveCredential is a paid mutator transaction binding the contract method 0xeedb973c.
//
// Solidity: function approveCredential(bytes32 digest) returns()
func (_Node *NodeTransactorSession) ApproveCredential(digest [32]byte) (*types.Transaction, error) {
	return _Node.Contract.ApproveCredential(&_Node.TransactOpts, digest)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Node *NodeTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Node *NodeSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Node.Contract.ChangeOwner(&_Node.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Node *NodeTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Node.Contract.ChangeOwner(&_Node.TransactOpts, newOwner)
}

// RegisterCredential is a paid mutator transaction binding the contract method 0x3ac82e19.
//
// Solidity: function registerCredential(address subject, bytes32 digest, address[] witnesses) returns()
func (_Node *NodeTransactor) RegisterCredential(opts *bind.TransactOpts, subject common.Address, digest [32]byte, witnesses []common.Address) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "registerCredential", subject, digest, witnesses)
}

// RegisterCredential is a paid mutator transaction binding the contract method 0x3ac82e19.
//
// Solidity: function registerCredential(address subject, bytes32 digest, address[] witnesses) returns()
func (_Node *NodeSession) RegisterCredential(subject common.Address, digest [32]byte, witnesses []common.Address) (*types.Transaction, error) {
	return _Node.Contract.RegisterCredential(&_Node.TransactOpts, subject, digest, witnesses)
}

// RegisterCredential is a paid mutator transaction binding the contract method 0x3ac82e19.
//
// Solidity: function registerCredential(address subject, bytes32 digest, address[] witnesses) returns()
func (_Node *NodeTransactorSession) RegisterCredential(subject common.Address, digest [32]byte, witnesses []common.Address) (*types.Transaction, error) {
	return _Node.Contract.RegisterCredential(&_Node.TransactOpts, subject, digest, witnesses)
}

// RevokeCredential is a paid mutator transaction binding the contract method 0x3610d3ca.
//
// Solidity: function revokeCredential(bytes32 digest, bytes32 reason) returns()
func (_Node *NodeTransactor) RevokeCredential(opts *bind.TransactOpts, digest [32]byte, reason [32]byte) (*types.Transaction, error) {
	return _Node.contract.Transact(opts, "revokeCredential", digest, reason)
}

// RevokeCredential is a paid mutator transaction binding the contract method 0x3610d3ca.
//
// Solidity: function revokeCredential(bytes32 digest, bytes32 reason) returns()
func (_Node *NodeSession) RevokeCredential(digest [32]byte, reason [32]byte) (*types.Transaction, error) {
	return _Node.Contract.RevokeCredential(&_Node.TransactOpts, digest, reason)
}

// RevokeCredential is a paid mutator transaction binding the contract method 0x3610d3ca.
//
// Solidity: function revokeCredential(bytes32 digest, bytes32 reason) returns()
func (_Node *NodeTransactorSession) RevokeCredential(digest [32]byte, reason [32]byte) (*types.Transaction, error) {
	return _Node.Contract.RevokeCredential(&_Node.TransactOpts, digest, reason)
}

// NodeCredentialIssuedIterator is returned from FilterCredentialIssued and is used to iterate over the raw logs and unpacked data for CredentialIssued events raised by the Node contract.
type NodeCredentialIssuedIterator struct {
	Event *NodeCredentialIssued // Event containing the contract specifics and raw log

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
func (it *NodeCredentialIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeCredentialIssued)
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
		it.Event = new(NodeCredentialIssued)
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
func (it *NodeCredentialIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeCredentialIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeCredentialIssued represents a CredentialIssued event raised by the Node contract.
type NodeCredentialIssued struct {
	Digest        [32]byte
	Subject       common.Address
	Registrar     common.Address
	InsertedBlock *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCredentialIssued is a free log retrieval operation binding the contract event 0xc6f6a1702b44006e35bd83045fd933510f3a924c1bb3877bff1061c14eb1c6be.
//
// Solidity: event CredentialIssued(bytes32 indexed digest, address indexed subject, address indexed registrar, uint256 insertedBlock)
func (_Node *NodeFilterer) FilterCredentialIssued(opts *bind.FilterOpts, digest [][32]byte, subject []common.Address, registrar []common.Address) (*NodeCredentialIssuedIterator, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "CredentialIssued", digestRule, subjectRule, registrarRule)
	if err != nil {
		return nil, err
	}
	return &NodeCredentialIssuedIterator{contract: _Node.contract, event: "CredentialIssued", logs: logs, sub: sub}, nil
}

// WatchCredentialIssued is a free log subscription operation binding the contract event 0xc6f6a1702b44006e35bd83045fd933510f3a924c1bb3877bff1061c14eb1c6be.
//
// Solidity: event CredentialIssued(bytes32 indexed digest, address indexed subject, address indexed registrar, uint256 insertedBlock)
func (_Node *NodeFilterer) WatchCredentialIssued(opts *bind.WatchOpts, sink chan<- *NodeCredentialIssued, digest [][32]byte, subject []common.Address, registrar []common.Address) (event.Subscription, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "CredentialIssued", digestRule, subjectRule, registrarRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeCredentialIssued)
				if err := _Node.contract.UnpackLog(event, "CredentialIssued", log); err != nil {
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

// ParseCredentialIssued is a log parse operation binding the contract event 0xc6f6a1702b44006e35bd83045fd933510f3a924c1bb3877bff1061c14eb1c6be.
//
// Solidity: event CredentialIssued(bytes32 indexed digest, address indexed subject, address indexed registrar, uint256 insertedBlock)
func (_Node *NodeFilterer) ParseCredentialIssued(log types.Log) (*NodeCredentialIssued, error) {
	event := new(NodeCredentialIssued)
	if err := _Node.contract.UnpackLog(event, "CredentialIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeCredentialRevokedIterator is returned from FilterCredentialRevoked and is used to iterate over the raw logs and unpacked data for CredentialRevoked events raised by the Node contract.
type NodeCredentialRevokedIterator struct {
	Event *NodeCredentialRevoked // Event containing the contract specifics and raw log

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
func (it *NodeCredentialRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeCredentialRevoked)
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
		it.Event = new(NodeCredentialRevoked)
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
func (it *NodeCredentialRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeCredentialRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeCredentialRevoked represents a CredentialRevoked event raised by the Node contract.
type NodeCredentialRevoked struct {
	Digest       [32]byte
	Subject      common.Address
	Revoker      common.Address
	RevokedBlock *big.Int
	Reason       [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCredentialRevoked is a free log retrieval operation binding the contract event 0xb53d448559d4a1f97e1fd48262b64223f5bdafd54d429a33599236f73e900b1f.
//
// Solidity: event CredentialRevoked(bytes32 indexed digest, address indexed subject, address indexed revoker, uint256 revokedBlock, bytes32 reason)
func (_Node *NodeFilterer) FilterCredentialRevoked(opts *bind.FilterOpts, digest [][32]byte, subject []common.Address, revoker []common.Address) (*NodeCredentialRevokedIterator, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var revokerRule []interface{}
	for _, revokerItem := range revoker {
		revokerRule = append(revokerRule, revokerItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "CredentialRevoked", digestRule, subjectRule, revokerRule)
	if err != nil {
		return nil, err
	}
	return &NodeCredentialRevokedIterator{contract: _Node.contract, event: "CredentialRevoked", logs: logs, sub: sub}, nil
}

// WatchCredentialRevoked is a free log subscription operation binding the contract event 0xb53d448559d4a1f97e1fd48262b64223f5bdafd54d429a33599236f73e900b1f.
//
// Solidity: event CredentialRevoked(bytes32 indexed digest, address indexed subject, address indexed revoker, uint256 revokedBlock, bytes32 reason)
func (_Node *NodeFilterer) WatchCredentialRevoked(opts *bind.WatchOpts, sink chan<- *NodeCredentialRevoked, digest [][32]byte, subject []common.Address, revoker []common.Address) (event.Subscription, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var revokerRule []interface{}
	for _, revokerItem := range revoker {
		revokerRule = append(revokerRule, revokerItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "CredentialRevoked", digestRule, subjectRule, revokerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeCredentialRevoked)
				if err := _Node.contract.UnpackLog(event, "CredentialRevoked", log); err != nil {
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

// ParseCredentialRevoked is a log parse operation binding the contract event 0xb53d448559d4a1f97e1fd48262b64223f5bdafd54d429a33599236f73e900b1f.
//
// Solidity: event CredentialRevoked(bytes32 indexed digest, address indexed subject, address indexed revoker, uint256 revokedBlock, bytes32 reason)
func (_Node *NodeFilterer) ParseCredentialRevoked(log types.Log) (*NodeCredentialRevoked, error) {
	event := new(NodeCredentialRevoked)
	if err := _Node.contract.UnpackLog(event, "CredentialRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeCredentialSignedIterator is returned from FilterCredentialSigned and is used to iterate over the raw logs and unpacked data for CredentialSigned events raised by the Node contract.
type NodeCredentialSignedIterator struct {
	Event *NodeCredentialSigned // Event containing the contract specifics and raw log

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
func (it *NodeCredentialSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeCredentialSigned)
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
		it.Event = new(NodeCredentialSigned)
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
func (it *NodeCredentialSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeCredentialSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeCredentialSigned represents a CredentialSigned event raised by the Node contract.
type NodeCredentialSigned struct {
	Signer      common.Address
	Digest      [32]byte
	SignedBlock *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCredentialSigned is a free log retrieval operation binding the contract event 0x3c4c49076c09b6a1214e1ba437ac70fd0cff1c6ea5baf6df4c5f6ba937c4bf12.
//
// Solidity: event CredentialSigned(address indexed signer, bytes32 indexed digest, uint256 signedBlock)
func (_Node *NodeFilterer) FilterCredentialSigned(opts *bind.FilterOpts, signer []common.Address, digest [][32]byte) (*NodeCredentialSignedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "CredentialSigned", signerRule, digestRule)
	if err != nil {
		return nil, err
	}
	return &NodeCredentialSignedIterator{contract: _Node.contract, event: "CredentialSigned", logs: logs, sub: sub}, nil
}

// WatchCredentialSigned is a free log subscription operation binding the contract event 0x3c4c49076c09b6a1214e1ba437ac70fd0cff1c6ea5baf6df4c5f6ba937c4bf12.
//
// Solidity: event CredentialSigned(address indexed signer, bytes32 indexed digest, uint256 signedBlock)
func (_Node *NodeFilterer) WatchCredentialSigned(opts *bind.WatchOpts, sink chan<- *NodeCredentialSigned, signer []common.Address, digest [][32]byte) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "CredentialSigned", signerRule, digestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeCredentialSigned)
				if err := _Node.contract.UnpackLog(event, "CredentialSigned", log); err != nil {
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

// ParseCredentialSigned is a log parse operation binding the contract event 0x3c4c49076c09b6a1214e1ba437ac70fd0cff1c6ea5baf6df4c5f6ba937c4bf12.
//
// Solidity: event CredentialSigned(address indexed signer, bytes32 indexed digest, uint256 signedBlock)
func (_Node *NodeFilterer) ParseCredentialSigned(log types.Log) (*NodeCredentialSigned, error) {
	event := new(NodeCredentialSigned)
	if err := _Node.contract.UnpackLog(event, "CredentialSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeNodeAddedIterator is returned from FilterNodeAdded and is used to iterate over the raw logs and unpacked data for NodeAdded events raised by the Node contract.
type NodeNodeAddedIterator struct {
	Event *NodeNodeAdded // Event containing the contract specifics and raw log

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
func (it *NodeNodeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeNodeAdded)
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
		it.Event = new(NodeNodeAdded)
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
func (it *NodeNodeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeNodeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeNodeAdded represents a NodeAdded event raised by the Node contract.
type NodeNodeAdded struct {
	CreatedBy   common.Address
	NodeAddress common.Address
	Role        uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeAdded is a free log retrieval operation binding the contract event 0x202458425cc4a92ea08231128a0a57aa24cfa42968e8090aa7d151f3f0a4eb3f.
//
// Solidity: event NodeAdded(address indexed createdBy, address indexed nodeAddress, uint8 role)
func (_Node *NodeFilterer) FilterNodeAdded(opts *bind.FilterOpts, createdBy []common.Address, nodeAddress []common.Address) (*NodeNodeAddedIterator, error) {

	var createdByRule []interface{}
	for _, createdByItem := range createdBy {
		createdByRule = append(createdByRule, createdByItem)
	}
	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "NodeAdded", createdByRule, nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &NodeNodeAddedIterator{contract: _Node.contract, event: "NodeAdded", logs: logs, sub: sub}, nil
}

// WatchNodeAdded is a free log subscription operation binding the contract event 0x202458425cc4a92ea08231128a0a57aa24cfa42968e8090aa7d151f3f0a4eb3f.
//
// Solidity: event NodeAdded(address indexed createdBy, address indexed nodeAddress, uint8 role)
func (_Node *NodeFilterer) WatchNodeAdded(opts *bind.WatchOpts, sink chan<- *NodeNodeAdded, createdBy []common.Address, nodeAddress []common.Address) (event.Subscription, error) {

	var createdByRule []interface{}
	for _, createdByItem := range createdBy {
		createdByRule = append(createdByRule, createdByItem)
	}
	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "NodeAdded", createdByRule, nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeNodeAdded)
				if err := _Node.contract.UnpackLog(event, "NodeAdded", log); err != nil {
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

// ParseNodeAdded is a log parse operation binding the contract event 0x202458425cc4a92ea08231128a0a57aa24cfa42968e8090aa7d151f3f0a4eb3f.
//
// Solidity: event NodeAdded(address indexed createdBy, address indexed nodeAddress, uint8 role)
func (_Node *NodeFilterer) ParseNodeAdded(log types.Log) (*NodeNodeAdded, error) {
	event := new(NodeNodeAdded)
	if err := _Node.contract.UnpackLog(event, "NodeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Node contract.
type NodeOwnerChangedIterator struct {
	Event *NodeOwnerChanged // Event containing the contract specifics and raw log

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
func (it *NodeOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeOwnerChanged)
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
		it.Event = new(NodeOwnerChanged)
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
func (it *NodeOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeOwnerChanged represents a OwnerChanged event raised by the Node contract.
type NodeOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_Node *NodeFilterer) FilterOwnerChanged(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*NodeOwnerChangedIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Node.contract.FilterLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NodeOwnerChangedIterator{contract: _Node.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_Node *NodeFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *NodeOwnerChanged, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Node.contract.WatchLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeOwnerChanged)
				if err := _Node.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_Node *NodeFilterer) ParseOwnerChanged(log types.Log) (*NodeOwnerChanged, error) {
	event := new(NodeOwnerChanged)
	if err := _Node.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeInterfaceABI is the input ABI used to generate the binding from.
const NodeInterfaceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"createdBy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumRole\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"NodeAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"}],\"name\":\"addChild\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"digests\",\"type\":\"bytes32[]\"}],\"name\":\"aggregateCredentials\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"}],\"name\":\"approveCredential\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRole\",\"outputs\":[{\"internalType\":\"enumRole\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLeaf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"myParent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"witnesses\",\"type\":\"address[]\"}],\"name\":\"registerCredential\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"reason\",\"type\":\"bytes32\"}],\"name\":\"revokeCredential\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"verifyCredentialRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"verifyCredentialTree\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NodeInterface is an auto generated Go binding around an Ethereum contract.
type NodeInterface struct {
	NodeInterfaceCaller     // Read-only binding to the contract
	NodeInterfaceTransactor // Write-only binding to the contract
	NodeInterfaceFilterer   // Log filterer for contract events
}

// NodeInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeInterfaceSession struct {
	Contract     *NodeInterface    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeInterfaceCallerSession struct {
	Contract *NodeInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// NodeInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeInterfaceTransactorSession struct {
	Contract     *NodeInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// NodeInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeInterfaceRaw struct {
	Contract *NodeInterface // Generic contract binding to access the raw methods on
}

// NodeInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeInterfaceCallerRaw struct {
	Contract *NodeInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// NodeInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeInterfaceTransactorRaw struct {
	Contract *NodeInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeInterface creates a new instance of NodeInterface, bound to a specific deployed contract.
func NewNodeInterface(address common.Address, backend bind.ContractBackend) (*NodeInterface, error) {
	contract, err := bindNodeInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeInterface{NodeInterfaceCaller: NodeInterfaceCaller{contract: contract}, NodeInterfaceTransactor: NodeInterfaceTransactor{contract: contract}, NodeInterfaceFilterer: NodeInterfaceFilterer{contract: contract}}, nil
}

// NewNodeInterfaceCaller creates a new read-only instance of NodeInterface, bound to a specific deployed contract.
func NewNodeInterfaceCaller(address common.Address, caller bind.ContractCaller) (*NodeInterfaceCaller, error) {
	contract, err := bindNodeInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeInterfaceCaller{contract: contract}, nil
}

// NewNodeInterfaceTransactor creates a new write-only instance of NodeInterface, bound to a specific deployed contract.
func NewNodeInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeInterfaceTransactor, error) {
	contract, err := bindNodeInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeInterfaceTransactor{contract: contract}, nil
}

// NewNodeInterfaceFilterer creates a new log filterer instance of NodeInterface, bound to a specific deployed contract.
func NewNodeInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeInterfaceFilterer, error) {
	contract, err := bindNodeInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeInterfaceFilterer{contract: contract}, nil
}

// bindNodeInterface binds a generic wrapper to an already deployed contract.
func bindNodeInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeInterface *NodeInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeInterface.Contract.NodeInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeInterface *NodeInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeInterface.Contract.NodeInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeInterface *NodeInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeInterface.Contract.NodeInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeInterface *NodeInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeInterface *NodeInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeInterface *NodeInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeInterface.Contract.contract.Transact(opts, method, params...)
}

// GetRole is a free data retrieval call binding the contract method 0xfde08174.
//
// Solidity: function getRole() view returns(uint8)
func (_NodeInterface *NodeInterfaceCaller) GetRole(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NodeInterface.contract.Call(opts, &out, "getRole")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetRole is a free data retrieval call binding the contract method 0xfde08174.
//
// Solidity: function getRole() view returns(uint8)
func (_NodeInterface *NodeInterfaceSession) GetRole() (uint8, error) {
	return _NodeInterface.Contract.GetRole(&_NodeInterface.CallOpts)
}

// GetRole is a free data retrieval call binding the contract method 0xfde08174.
//
// Solidity: function getRole() view returns(uint8)
func (_NodeInterface *NodeInterfaceCallerSession) GetRole() (uint8, error) {
	return _NodeInterface.Contract.GetRole(&_NodeInterface.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x079cf76e.
//
// Solidity: function getRoot(address subject) view returns(bytes32)
func (_NodeInterface *NodeInterfaceCaller) GetRoot(opts *bind.CallOpts, subject common.Address) ([32]byte, error) {
	var out []interface{}
	err := _NodeInterface.contract.Call(opts, &out, "getRoot", subject)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x079cf76e.
//
// Solidity: function getRoot(address subject) view returns(bytes32)
func (_NodeInterface *NodeInterfaceSession) GetRoot(subject common.Address) ([32]byte, error) {
	return _NodeInterface.Contract.GetRoot(&_NodeInterface.CallOpts, subject)
}

// GetRoot is a free data retrieval call binding the contract method 0x079cf76e.
//
// Solidity: function getRoot(address subject) view returns(bytes32)
func (_NodeInterface *NodeInterfaceCallerSession) GetRoot(subject common.Address) ([32]byte, error) {
	return _NodeInterface.Contract.GetRoot(&_NodeInterface.CallOpts, subject)
}

// IsLeaf is a free data retrieval call binding the contract method 0x99d1dbc0.
//
// Solidity: function isLeaf() view returns(bool)
func (_NodeInterface *NodeInterfaceCaller) IsLeaf(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NodeInterface.contract.Call(opts, &out, "isLeaf")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLeaf is a free data retrieval call binding the contract method 0x99d1dbc0.
//
// Solidity: function isLeaf() view returns(bool)
func (_NodeInterface *NodeInterfaceSession) IsLeaf() (bool, error) {
	return _NodeInterface.Contract.IsLeaf(&_NodeInterface.CallOpts)
}

// IsLeaf is a free data retrieval call binding the contract method 0x99d1dbc0.
//
// Solidity: function isLeaf() view returns(bool)
func (_NodeInterface *NodeInterfaceCallerSession) IsLeaf() (bool, error) {
	return _NodeInterface.Contract.IsLeaf(&_NodeInterface.CallOpts)
}

// MyParent is a free data retrieval call binding the contract method 0xde07fc45.
//
// Solidity: function myParent() view returns(address)
func (_NodeInterface *NodeInterfaceCaller) MyParent(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NodeInterface.contract.Call(opts, &out, "myParent")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MyParent is a free data retrieval call binding the contract method 0xde07fc45.
//
// Solidity: function myParent() view returns(address)
func (_NodeInterface *NodeInterfaceSession) MyParent() (common.Address, error) {
	return _NodeInterface.Contract.MyParent(&_NodeInterface.CallOpts)
}

// MyParent is a free data retrieval call binding the contract method 0xde07fc45.
//
// Solidity: function myParent() view returns(address)
func (_NodeInterface *NodeInterfaceCallerSession) MyParent() (common.Address, error) {
	return _NodeInterface.Contract.MyParent(&_NodeInterface.CallOpts)
}

// VerifyCredentialRoot is a free data retrieval call binding the contract method 0x0ae07028.
//
// Solidity: function verifyCredentialRoot(address subject, bytes32 root) view returns(bool)
func (_NodeInterface *NodeInterfaceCaller) VerifyCredentialRoot(opts *bind.CallOpts, subject common.Address, root [32]byte) (bool, error) {
	var out []interface{}
	err := _NodeInterface.contract.Call(opts, &out, "verifyCredentialRoot", subject, root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyCredentialRoot is a free data retrieval call binding the contract method 0x0ae07028.
//
// Solidity: function verifyCredentialRoot(address subject, bytes32 root) view returns(bool)
func (_NodeInterface *NodeInterfaceSession) VerifyCredentialRoot(subject common.Address, root [32]byte) (bool, error) {
	return _NodeInterface.Contract.VerifyCredentialRoot(&_NodeInterface.CallOpts, subject, root)
}

// VerifyCredentialRoot is a free data retrieval call binding the contract method 0x0ae07028.
//
// Solidity: function verifyCredentialRoot(address subject, bytes32 root) view returns(bool)
func (_NodeInterface *NodeInterfaceCallerSession) VerifyCredentialRoot(subject common.Address, root [32]byte) (bool, error) {
	return _NodeInterface.Contract.VerifyCredentialRoot(&_NodeInterface.CallOpts, subject, root)
}

// VerifyCredentialTree is a free data retrieval call binding the contract method 0xf936cc15.
//
// Solidity: function verifyCredentialTree(address subject) view returns(bool)
func (_NodeInterface *NodeInterfaceCaller) VerifyCredentialTree(opts *bind.CallOpts, subject common.Address) (bool, error) {
	var out []interface{}
	err := _NodeInterface.contract.Call(opts, &out, "verifyCredentialTree", subject)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyCredentialTree is a free data retrieval call binding the contract method 0xf936cc15.
//
// Solidity: function verifyCredentialTree(address subject) view returns(bool)
func (_NodeInterface *NodeInterfaceSession) VerifyCredentialTree(subject common.Address) (bool, error) {
	return _NodeInterface.Contract.VerifyCredentialTree(&_NodeInterface.CallOpts, subject)
}

// VerifyCredentialTree is a free data retrieval call binding the contract method 0xf936cc15.
//
// Solidity: function verifyCredentialTree(address subject) view returns(bool)
func (_NodeInterface *NodeInterfaceCallerSession) VerifyCredentialTree(subject common.Address) (bool, error) {
	return _NodeInterface.Contract.VerifyCredentialTree(&_NodeInterface.CallOpts, subject)
}

// AddChild is a paid mutator transaction binding the contract method 0x1eee993a.
//
// Solidity: function addChild(address nodeAddress) returns()
func (_NodeInterface *NodeInterfaceTransactor) AddChild(opts *bind.TransactOpts, nodeAddress common.Address) (*types.Transaction, error) {
	return _NodeInterface.contract.Transact(opts, "addChild", nodeAddress)
}

// AddChild is a paid mutator transaction binding the contract method 0x1eee993a.
//
// Solidity: function addChild(address nodeAddress) returns()
func (_NodeInterface *NodeInterfaceSession) AddChild(nodeAddress common.Address) (*types.Transaction, error) {
	return _NodeInterface.Contract.AddChild(&_NodeInterface.TransactOpts, nodeAddress)
}

// AddChild is a paid mutator transaction binding the contract method 0x1eee993a.
//
// Solidity: function addChild(address nodeAddress) returns()
func (_NodeInterface *NodeInterfaceTransactorSession) AddChild(nodeAddress common.Address) (*types.Transaction, error) {
	return _NodeInterface.Contract.AddChild(&_NodeInterface.TransactOpts, nodeAddress)
}

// AggregateCredentials is a paid mutator transaction binding the contract method 0xe2d02cea.
//
// Solidity: function aggregateCredentials(address subject, bytes32[] digests) returns(bytes32)
func (_NodeInterface *NodeInterfaceTransactor) AggregateCredentials(opts *bind.TransactOpts, subject common.Address, digests [][32]byte) (*types.Transaction, error) {
	return _NodeInterface.contract.Transact(opts, "aggregateCredentials", subject, digests)
}

// AggregateCredentials is a paid mutator transaction binding the contract method 0xe2d02cea.
//
// Solidity: function aggregateCredentials(address subject, bytes32[] digests) returns(bytes32)
func (_NodeInterface *NodeInterfaceSession) AggregateCredentials(subject common.Address, digests [][32]byte) (*types.Transaction, error) {
	return _NodeInterface.Contract.AggregateCredentials(&_NodeInterface.TransactOpts, subject, digests)
}

// AggregateCredentials is a paid mutator transaction binding the contract method 0xe2d02cea.
//
// Solidity: function aggregateCredentials(address subject, bytes32[] digests) returns(bytes32)
func (_NodeInterface *NodeInterfaceTransactorSession) AggregateCredentials(subject common.Address, digests [][32]byte) (*types.Transaction, error) {
	return _NodeInterface.Contract.AggregateCredentials(&_NodeInterface.TransactOpts, subject, digests)
}

// ApproveCredential is a paid mutator transaction binding the contract method 0xeedb973c.
//
// Solidity: function approveCredential(bytes32 digest) returns()
func (_NodeInterface *NodeInterfaceTransactor) ApproveCredential(opts *bind.TransactOpts, digest [32]byte) (*types.Transaction, error) {
	return _NodeInterface.contract.Transact(opts, "approveCredential", digest)
}

// ApproveCredential is a paid mutator transaction binding the contract method 0xeedb973c.
//
// Solidity: function approveCredential(bytes32 digest) returns()
func (_NodeInterface *NodeInterfaceSession) ApproveCredential(digest [32]byte) (*types.Transaction, error) {
	return _NodeInterface.Contract.ApproveCredential(&_NodeInterface.TransactOpts, digest)
}

// ApproveCredential is a paid mutator transaction binding the contract method 0xeedb973c.
//
// Solidity: function approveCredential(bytes32 digest) returns()
func (_NodeInterface *NodeInterfaceTransactorSession) ApproveCredential(digest [32]byte) (*types.Transaction, error) {
	return _NodeInterface.Contract.ApproveCredential(&_NodeInterface.TransactOpts, digest)
}

// RegisterCredential is a paid mutator transaction binding the contract method 0x3ac82e19.
//
// Solidity: function registerCredential(address subject, bytes32 digest, address[] witnesses) returns()
func (_NodeInterface *NodeInterfaceTransactor) RegisterCredential(opts *bind.TransactOpts, subject common.Address, digest [32]byte, witnesses []common.Address) (*types.Transaction, error) {
	return _NodeInterface.contract.Transact(opts, "registerCredential", subject, digest, witnesses)
}

// RegisterCredential is a paid mutator transaction binding the contract method 0x3ac82e19.
//
// Solidity: function registerCredential(address subject, bytes32 digest, address[] witnesses) returns()
func (_NodeInterface *NodeInterfaceSession) RegisterCredential(subject common.Address, digest [32]byte, witnesses []common.Address) (*types.Transaction, error) {
	return _NodeInterface.Contract.RegisterCredential(&_NodeInterface.TransactOpts, subject, digest, witnesses)
}

// RegisterCredential is a paid mutator transaction binding the contract method 0x3ac82e19.
//
// Solidity: function registerCredential(address subject, bytes32 digest, address[] witnesses) returns()
func (_NodeInterface *NodeInterfaceTransactorSession) RegisterCredential(subject common.Address, digest [32]byte, witnesses []common.Address) (*types.Transaction, error) {
	return _NodeInterface.Contract.RegisterCredential(&_NodeInterface.TransactOpts, subject, digest, witnesses)
}

// RevokeCredential is a paid mutator transaction binding the contract method 0x3610d3ca.
//
// Solidity: function revokeCredential(bytes32 digest, bytes32 reason) returns()
func (_NodeInterface *NodeInterfaceTransactor) RevokeCredential(opts *bind.TransactOpts, digest [32]byte, reason [32]byte) (*types.Transaction, error) {
	return _NodeInterface.contract.Transact(opts, "revokeCredential", digest, reason)
}

// RevokeCredential is a paid mutator transaction binding the contract method 0x3610d3ca.
//
// Solidity: function revokeCredential(bytes32 digest, bytes32 reason) returns()
func (_NodeInterface *NodeInterfaceSession) RevokeCredential(digest [32]byte, reason [32]byte) (*types.Transaction, error) {
	return _NodeInterface.Contract.RevokeCredential(&_NodeInterface.TransactOpts, digest, reason)
}

// RevokeCredential is a paid mutator transaction binding the contract method 0x3610d3ca.
//
// Solidity: function revokeCredential(bytes32 digest, bytes32 reason) returns()
func (_NodeInterface *NodeInterfaceTransactorSession) RevokeCredential(digest [32]byte, reason [32]byte) (*types.Transaction, error) {
	return _NodeInterface.Contract.RevokeCredential(&_NodeInterface.TransactOpts, digest, reason)
}

// NodeInterfaceNodeAddedIterator is returned from FilterNodeAdded and is used to iterate over the raw logs and unpacked data for NodeAdded events raised by the NodeInterface contract.
type NodeInterfaceNodeAddedIterator struct {
	Event *NodeInterfaceNodeAdded // Event containing the contract specifics and raw log

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
func (it *NodeInterfaceNodeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeInterfaceNodeAdded)
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
		it.Event = new(NodeInterfaceNodeAdded)
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
func (it *NodeInterfaceNodeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeInterfaceNodeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeInterfaceNodeAdded represents a NodeAdded event raised by the NodeInterface contract.
type NodeInterfaceNodeAdded struct {
	CreatedBy   common.Address
	NodeAddress common.Address
	Role        uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeAdded is a free log retrieval operation binding the contract event 0x202458425cc4a92ea08231128a0a57aa24cfa42968e8090aa7d151f3f0a4eb3f.
//
// Solidity: event NodeAdded(address indexed createdBy, address indexed nodeAddress, uint8 role)
func (_NodeInterface *NodeInterfaceFilterer) FilterNodeAdded(opts *bind.FilterOpts, createdBy []common.Address, nodeAddress []common.Address) (*NodeInterfaceNodeAddedIterator, error) {

	var createdByRule []interface{}
	for _, createdByItem := range createdBy {
		createdByRule = append(createdByRule, createdByItem)
	}
	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _NodeInterface.contract.FilterLogs(opts, "NodeAdded", createdByRule, nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return &NodeInterfaceNodeAddedIterator{contract: _NodeInterface.contract, event: "NodeAdded", logs: logs, sub: sub}, nil
}

// WatchNodeAdded is a free log subscription operation binding the contract event 0x202458425cc4a92ea08231128a0a57aa24cfa42968e8090aa7d151f3f0a4eb3f.
//
// Solidity: event NodeAdded(address indexed createdBy, address indexed nodeAddress, uint8 role)
func (_NodeInterface *NodeInterfaceFilterer) WatchNodeAdded(opts *bind.WatchOpts, sink chan<- *NodeInterfaceNodeAdded, createdBy []common.Address, nodeAddress []common.Address) (event.Subscription, error) {

	var createdByRule []interface{}
	for _, createdByItem := range createdBy {
		createdByRule = append(createdByRule, createdByItem)
	}
	var nodeAddressRule []interface{}
	for _, nodeAddressItem := range nodeAddress {
		nodeAddressRule = append(nodeAddressRule, nodeAddressItem)
	}

	logs, sub, err := _NodeInterface.contract.WatchLogs(opts, "NodeAdded", createdByRule, nodeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeInterfaceNodeAdded)
				if err := _NodeInterface.contract.UnpackLog(event, "NodeAdded", log); err != nil {
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

// ParseNodeAdded is a log parse operation binding the contract event 0x202458425cc4a92ea08231128a0a57aa24cfa42968e8090aa7d151f3f0a4eb3f.
//
// Solidity: event NodeAdded(address indexed createdBy, address indexed nodeAddress, uint8 role)
func (_NodeInterface *NodeInterfaceFilterer) ParseNodeAdded(log types.Log) (*NodeInterfaceNodeAdded, error) {
	event := new(NodeInterfaceNodeAdded)
	if err := _NodeInterface.contract.UnpackLog(event, "NodeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NotaryABI is the input ABI used to generate the binding from.
const NotaryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registrar\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"insertedBlock\",\"type\":\"uint256\"}],\"name\":\"CredentialIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"revoker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revokedBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"reason\",\"type\":\"bytes32\"}],\"name\":\"CredentialRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"digest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"signedBlock\",\"type\":\"uint256\"}],\"name\":\"CredentialSigned\",\"type\":\"event\"}]"

// NotaryBin is the compiled bytecode used for deploying new contracts.
var NotaryBin = "0x61055961003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100875760003560e01c8063a8b0bcfd11610065578063a8b0bcfd1461011d578063db75cc5214610158578063e187eb8014610218578063f807f5741461023857610087565b806350f3aac41461008c5780636667c014146100ca57806397a3b088146100f3575b600080fd5b6100b561009a3660046103a0565b60009081526003919091016020526040902060010154151590565b60405190151581526020015b60405180910390f35b6100b56100d83660046103a0565b60009081526004919091016020526040902060020154151590565b6100b56101013660046103a0565b6000908152600391909101602052604090206005015460ff1690565b6100b561012b3660046103c1565b6000918252600692909201602090815260408083206001600160a01b039094168352929052205460ff1690565b6101d76101663660046103a0565b60408051608080820183526000808352602080840182905283850182905260609384018290529481526004959095018452938290208251948501835280546001600160a01b039081168652600182015416938501939093526002830154918401919091526003909101549082015290565b6040516100c1919081516001600160a01b03908116825260208084015190911690820152604080830151908201526060918201519181019190915260800190565b61022b6102263660046103a0565b610264565b6040516100c19190610479565b6100b5610246366004610404565b600082815260038401602052604090205460ff821611159392505050565b604080516101408101825260008082526020820181905291810182905260608082018390526080820183905260a0820183905260c0820183905260e082018390526101008201526101208101919091526000828152600380850160209081526040928390208351610140810185528154815260018201548184015260028201548186015292810154606084015260048101546080840152600581015460ff8116151560a08501526001600160a01b0361010091829004811660c086015260068301541660e0850152600782018054865181860281018601909752808752949592949186019390919083018282801561038557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610367575b50505050508152602001600882015481525050905092915050565b600080604083850312156103b2578182fd5b50508035926020909101359150565b6000806000606084860312156103d5578081fd5b833592506020840135915060408401356001600160a01b03811681146103f9578182fd5b809150509250925092565b600080600060608486031215610418578283fd5b8335925060208401359150604084013560ff811681146103f9578182fd5b6000815180845260208085019450808401835b8381101561046e5781516001600160a01b031687529582019590820190600101610449565b509495945050505050565b60006020825282516020830152602083015160408301526040830151606083015260608301516080830152608083015160a083015260a08301516104c160c084018215159052565b5060c08301516001600160a01b03811660e08401525060e08301516101006104f3818501836001600160a01b03169052565b808501519150506101406101208181860152610513610160860184610436565b950151930192909252509091905056fea264697066735822122076d579de4af2c19fd1bcae67940fb39805306c07b16878303d0525576e74256164736f6c63430008020033"

// DeployNotary deploys a new Ethereum contract, binding an instance of Notary to it.
func DeployNotary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Notary, error) {
	parsed, err := abi.JSON(strings.NewReader(NotaryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NotaryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Notary{NotaryCaller: NotaryCaller{contract: contract}, NotaryTransactor: NotaryTransactor{contract: contract}, NotaryFilterer: NotaryFilterer{contract: contract}}, nil
}

// Notary is an auto generated Go binding around an Ethereum contract.
type Notary struct {
	NotaryCaller     // Read-only binding to the contract
	NotaryTransactor // Write-only binding to the contract
	NotaryFilterer   // Log filterer for contract events
}

// NotaryCaller is an auto generated read-only Go binding around an Ethereum contract.
type NotaryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotaryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NotaryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotaryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NotaryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NotarySession struct {
	Contract     *Notary           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NotaryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NotaryCallerSession struct {
	Contract *NotaryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NotaryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NotaryTransactorSession struct {
	Contract     *NotaryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NotaryRaw is an auto generated low-level Go binding around an Ethereum contract.
type NotaryRaw struct {
	Contract *Notary // Generic contract binding to access the raw methods on
}

// NotaryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NotaryCallerRaw struct {
	Contract *NotaryCaller // Generic read-only contract binding to access the raw methods on
}

// NotaryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NotaryTransactorRaw struct {
	Contract *NotaryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNotary creates a new instance of Notary, bound to a specific deployed contract.
func NewNotary(address common.Address, backend bind.ContractBackend) (*Notary, error) {
	contract, err := bindNotary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Notary{NotaryCaller: NotaryCaller{contract: contract}, NotaryTransactor: NotaryTransactor{contract: contract}, NotaryFilterer: NotaryFilterer{contract: contract}}, nil
}

// NewNotaryCaller creates a new read-only instance of Notary, bound to a specific deployed contract.
func NewNotaryCaller(address common.Address, caller bind.ContractCaller) (*NotaryCaller, error) {
	contract, err := bindNotary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NotaryCaller{contract: contract}, nil
}

// NewNotaryTransactor creates a new write-only instance of Notary, bound to a specific deployed contract.
func NewNotaryTransactor(address common.Address, transactor bind.ContractTransactor) (*NotaryTransactor, error) {
	contract, err := bindNotary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NotaryTransactor{contract: contract}, nil
}

// NewNotaryFilterer creates a new log filterer instance of Notary, bound to a specific deployed contract.
func NewNotaryFilterer(address common.Address, filterer bind.ContractFilterer) (*NotaryFilterer, error) {
	contract, err := bindNotary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NotaryFilterer{contract: contract}, nil
}

// bindNotary binds a generic wrapper to an already deployed contract.
func bindNotary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NotaryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Notary *NotaryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Notary.Contract.NotaryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Notary *NotaryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Notary.Contract.NotaryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Notary *NotaryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Notary.Contract.NotaryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Notary *NotaryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Notary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Notary *NotaryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Notary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Notary *NotaryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Notary.Contract.contract.Transact(opts, method, params...)
}

// NotaryCredentialIssuedIterator is returned from FilterCredentialIssued and is used to iterate over the raw logs and unpacked data for CredentialIssued events raised by the Notary contract.
type NotaryCredentialIssuedIterator struct {
	Event *NotaryCredentialIssued // Event containing the contract specifics and raw log

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
func (it *NotaryCredentialIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotaryCredentialIssued)
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
		it.Event = new(NotaryCredentialIssued)
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
func (it *NotaryCredentialIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotaryCredentialIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotaryCredentialIssued represents a CredentialIssued event raised by the Notary contract.
type NotaryCredentialIssued struct {
	Digest        [32]byte
	Subject       common.Address
	Registrar     common.Address
	InsertedBlock *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCredentialIssued is a free log retrieval operation binding the contract event 0xc6f6a1702b44006e35bd83045fd933510f3a924c1bb3877bff1061c14eb1c6be.
//
// Solidity: event CredentialIssued(bytes32 indexed digest, address indexed subject, address indexed registrar, uint256 insertedBlock)
func (_Notary *NotaryFilterer) FilterCredentialIssued(opts *bind.FilterOpts, digest [][32]byte, subject []common.Address, registrar []common.Address) (*NotaryCredentialIssuedIterator, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}

	logs, sub, err := _Notary.contract.FilterLogs(opts, "CredentialIssued", digestRule, subjectRule, registrarRule)
	if err != nil {
		return nil, err
	}
	return &NotaryCredentialIssuedIterator{contract: _Notary.contract, event: "CredentialIssued", logs: logs, sub: sub}, nil
}

// WatchCredentialIssued is a free log subscription operation binding the contract event 0xc6f6a1702b44006e35bd83045fd933510f3a924c1bb3877bff1061c14eb1c6be.
//
// Solidity: event CredentialIssued(bytes32 indexed digest, address indexed subject, address indexed registrar, uint256 insertedBlock)
func (_Notary *NotaryFilterer) WatchCredentialIssued(opts *bind.WatchOpts, sink chan<- *NotaryCredentialIssued, digest [][32]byte, subject []common.Address, registrar []common.Address) (event.Subscription, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}

	logs, sub, err := _Notary.contract.WatchLogs(opts, "CredentialIssued", digestRule, subjectRule, registrarRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotaryCredentialIssued)
				if err := _Notary.contract.UnpackLog(event, "CredentialIssued", log); err != nil {
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

// ParseCredentialIssued is a log parse operation binding the contract event 0xc6f6a1702b44006e35bd83045fd933510f3a924c1bb3877bff1061c14eb1c6be.
//
// Solidity: event CredentialIssued(bytes32 indexed digest, address indexed subject, address indexed registrar, uint256 insertedBlock)
func (_Notary *NotaryFilterer) ParseCredentialIssued(log types.Log) (*NotaryCredentialIssued, error) {
	event := new(NotaryCredentialIssued)
	if err := _Notary.contract.UnpackLog(event, "CredentialIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NotaryCredentialRevokedIterator is returned from FilterCredentialRevoked and is used to iterate over the raw logs and unpacked data for CredentialRevoked events raised by the Notary contract.
type NotaryCredentialRevokedIterator struct {
	Event *NotaryCredentialRevoked // Event containing the contract specifics and raw log

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
func (it *NotaryCredentialRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotaryCredentialRevoked)
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
		it.Event = new(NotaryCredentialRevoked)
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
func (it *NotaryCredentialRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotaryCredentialRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotaryCredentialRevoked represents a CredentialRevoked event raised by the Notary contract.
type NotaryCredentialRevoked struct {
	Digest       [32]byte
	Subject      common.Address
	Revoker      common.Address
	RevokedBlock *big.Int
	Reason       [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCredentialRevoked is a free log retrieval operation binding the contract event 0xb53d448559d4a1f97e1fd48262b64223f5bdafd54d429a33599236f73e900b1f.
//
// Solidity: event CredentialRevoked(bytes32 indexed digest, address indexed subject, address indexed revoker, uint256 revokedBlock, bytes32 reason)
func (_Notary *NotaryFilterer) FilterCredentialRevoked(opts *bind.FilterOpts, digest [][32]byte, subject []common.Address, revoker []common.Address) (*NotaryCredentialRevokedIterator, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var revokerRule []interface{}
	for _, revokerItem := range revoker {
		revokerRule = append(revokerRule, revokerItem)
	}

	logs, sub, err := _Notary.contract.FilterLogs(opts, "CredentialRevoked", digestRule, subjectRule, revokerRule)
	if err != nil {
		return nil, err
	}
	return &NotaryCredentialRevokedIterator{contract: _Notary.contract, event: "CredentialRevoked", logs: logs, sub: sub}, nil
}

// WatchCredentialRevoked is a free log subscription operation binding the contract event 0xb53d448559d4a1f97e1fd48262b64223f5bdafd54d429a33599236f73e900b1f.
//
// Solidity: event CredentialRevoked(bytes32 indexed digest, address indexed subject, address indexed revoker, uint256 revokedBlock, bytes32 reason)
func (_Notary *NotaryFilterer) WatchCredentialRevoked(opts *bind.WatchOpts, sink chan<- *NotaryCredentialRevoked, digest [][32]byte, subject []common.Address, revoker []common.Address) (event.Subscription, error) {

	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}
	var subjectRule []interface{}
	for _, subjectItem := range subject {
		subjectRule = append(subjectRule, subjectItem)
	}
	var revokerRule []interface{}
	for _, revokerItem := range revoker {
		revokerRule = append(revokerRule, revokerItem)
	}

	logs, sub, err := _Notary.contract.WatchLogs(opts, "CredentialRevoked", digestRule, subjectRule, revokerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotaryCredentialRevoked)
				if err := _Notary.contract.UnpackLog(event, "CredentialRevoked", log); err != nil {
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

// ParseCredentialRevoked is a log parse operation binding the contract event 0xb53d448559d4a1f97e1fd48262b64223f5bdafd54d429a33599236f73e900b1f.
//
// Solidity: event CredentialRevoked(bytes32 indexed digest, address indexed subject, address indexed revoker, uint256 revokedBlock, bytes32 reason)
func (_Notary *NotaryFilterer) ParseCredentialRevoked(log types.Log) (*NotaryCredentialRevoked, error) {
	event := new(NotaryCredentialRevoked)
	if err := _Notary.contract.UnpackLog(event, "CredentialRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NotaryCredentialSignedIterator is returned from FilterCredentialSigned and is used to iterate over the raw logs and unpacked data for CredentialSigned events raised by the Notary contract.
type NotaryCredentialSignedIterator struct {
	Event *NotaryCredentialSigned // Event containing the contract specifics and raw log

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
func (it *NotaryCredentialSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotaryCredentialSigned)
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
		it.Event = new(NotaryCredentialSigned)
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
func (it *NotaryCredentialSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotaryCredentialSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotaryCredentialSigned represents a CredentialSigned event raised by the Notary contract.
type NotaryCredentialSigned struct {
	Signer      common.Address
	Digest      [32]byte
	SignedBlock *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCredentialSigned is a free log retrieval operation binding the contract event 0x3c4c49076c09b6a1214e1ba437ac70fd0cff1c6ea5baf6df4c5f6ba937c4bf12.
//
// Solidity: event CredentialSigned(address indexed signer, bytes32 indexed digest, uint256 signedBlock)
func (_Notary *NotaryFilterer) FilterCredentialSigned(opts *bind.FilterOpts, signer []common.Address, digest [][32]byte) (*NotaryCredentialSignedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _Notary.contract.FilterLogs(opts, "CredentialSigned", signerRule, digestRule)
	if err != nil {
		return nil, err
	}
	return &NotaryCredentialSignedIterator{contract: _Notary.contract, event: "CredentialSigned", logs: logs, sub: sub}, nil
}

// WatchCredentialSigned is a free log subscription operation binding the contract event 0x3c4c49076c09b6a1214e1ba437ac70fd0cff1c6ea5baf6df4c5f6ba937c4bf12.
//
// Solidity: event CredentialSigned(address indexed signer, bytes32 indexed digest, uint256 signedBlock)
func (_Notary *NotaryFilterer) WatchCredentialSigned(opts *bind.WatchOpts, sink chan<- *NotaryCredentialSigned, signer []common.Address, digest [][32]byte) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _Notary.contract.WatchLogs(opts, "CredentialSigned", signerRule, digestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotaryCredentialSigned)
				if err := _Notary.contract.UnpackLog(event, "CredentialSigned", log); err != nil {
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

// ParseCredentialSigned is a log parse operation binding the contract event 0x3c4c49076c09b6a1214e1ba437ac70fd0cff1c6ea5baf6df4c5f6ba937c4bf12.
//
// Solidity: event CredentialSigned(address indexed signer, bytes32 indexed digest, uint256 signedBlock)
func (_Notary *NotaryFilterer) ParseCredentialSigned(log types.Log) (*NotaryCredentialSigned, error) {
	event := new(NotaryCredentialSigned)
	if err := _Notary.contract.UnpackLog(event, "CredentialSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnersABI is the input ABI used to generate the binding from.
const OwnersABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"ownersList\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"quorumSize\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownersCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OwnersBin is the compiled bytecode used for deploying new contracts.
var OwnersBin = "0x60806040523480156200001157600080fd5b5060405162000a4938038062000a49833981016040819052620000349162000322565b60008251118015620000485750815160ff10155b6200009a5760405162461bcd60e51b815260206004820152601860248201527f4f776e6572732f6e6f7420656e6f756768206f776e657273000000000000000060448201526064015b60405180910390fd5b60008160ff16118015620000b2575081518160ff1611155b620001005760405162461bcd60e51b815260206004820152601a60248201527f4f776e6572732f71756f72756d206f7574206f662072616e6765000000000000604482015260640162000091565b60005b82518160ff161015620002355760026000848360ff16815181106200013857634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff16158015620001ab575060006001600160a01b0316838260ff16815181106200019757634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b031614155b620001c657634e487b7160e01b600052600160045260246000fd5b600160026000858460ff1681518110620001f057634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff19169115159190911790556200022d8162000409565b905062000103565b5081516200024b90600190602085019062000272565b5090516000805461ff00191661010060ff938416021760ff1916919092161790556200044c565b828054828255906000526020600020908101928215620002ca579160200282015b82811115620002ca57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000293565b50620002d8929150620002dc565b5090565b5b80821115620002d85760008155600101620002dd565b80516001600160a01b03811681146200030b57600080fd5b919050565b805160ff811681146200030b57600080fd5b6000806040838503121562000335578182fd5b82516001600160401b03808211156200034c578384fd5b818501915085601f83011262000360578384fd5b815160208282111562000377576200037762000436565b808202604051601f19603f830116810181811086821117156200039e576200039e62000436565b604052838152828101945085830182870184018b1015620003bd578889fd5b8896505b84871015620003ea57620003d581620002f3565b865260019690960195948301948301620003c1565b509650620003fc905087820162000310565b9450505050509250929050565b600060ff821660ff8114156200042d57634e487b7160e01b82526011600452602482fd5b60010192915050565b634e487b7160e01b600052604160045260246000fd5b6105ed806200045c6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80631703a0181461005c5780632f54bf6e1461007a578063a6f9dae1146100b6578063affe39c1146100cb578063b9488546146100e0575b600080fd5b60005460ff165b60405160ff90911681526020015b60405180910390f35b6100a6610088366004610510565b6001600160a01b031660009081526002602052604090205460ff1690565b6040519015158152602001610071565b6100c96100c4366004610510565b6100f1565b005b6100d3610434565b604051610071919061053e565b610063600054610100900460ff1690565b3360009081526002602052604090205460ff166101555760405162461bcd60e51b815260206004820152601d60248201527f4f776e6572732f73656e646572206973206e6f7420616e206f776e657200000060448201526064015b60405180910390fd5b6001600160a01b03811660009081526002602052604090205460ff1615801561018657506001600160a01b03811615155b6101d25760405162461bcd60e51b815260206004820152601c60248201527f4f776e6572732f696e76616c6964206164647265737320676976656e00000000604482015260640161014c565b600154158015906101e6575060015460ff10155b61020057634e487b7160e01b600052600160045260246000fd5b60015460009067ffffffffffffffff81111561022c57634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610255578160200160208202803683370190505b50905060005b60015460ff8216101561038f57336001600160a01b031660018260ff168154811061029657634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b03161461033a5760018160ff16815481106102d457634e487b7160e01b600052603260045260246000fd5b9060005260206000200160009054906101000a90046001600160a01b0316828260ff168151811061031557634e487b7160e01b600052603260045260246000fd5b60200260200101906001600160a01b031690816001600160a01b03168152505061037f565b82828260ff168151811061035e57634e487b7160e01b600052603260045260246000fd5b60200260200101906001600160a01b031690816001600160a01b0316815250505b6103888161058b565b905061025b565b50600054815160ff909116146103b557634e487b7160e01b600052600160045260246000fd5b6040516001600160a01b0383169033907fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c90600090a380516103fe906001906020840190610496565b50506001600160a01b0316600090815260026020526040808220805460ff19908116600117909155338352912080549091169055565b6060600180548060200260200160405190810160405280929190818152602001828054801561048c57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161046e575b5050505050905090565b8280548282559060005260206000209081019282156104eb579160200282015b828111156104eb57825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906104b6565b506104f79291506104fb565b5090565b5b808211156104f757600081556001016104fc565b600060208284031215610521578081fd5b81356001600160a01b0381168114610537578182fd5b9392505050565b6020808252825182820181905260009190848201906040850190845b8181101561057f5783516001600160a01b03168352928401929184019160010161055a565b50909695505050505050565b600060ff821660ff8114156105ae57634e487b7160e01b82526011600452602482fd5b6001019291505056fea264697066735822122067b038cb3060718bf05b180f202a64a128f0bf686651bbce38819770633c7ea564736f6c63430008020033"

// DeployOwners deploys a new Ethereum contract, binding an instance of Owners to it.
func DeployOwners(auth *bind.TransactOpts, backend bind.ContractBackend, ownersList []common.Address, quorumSize uint8) (common.Address, *types.Transaction, *Owners, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnersBin), backend, ownersList, quorumSize)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owners{OwnersCaller: OwnersCaller{contract: contract}, OwnersTransactor: OwnersTransactor{contract: contract}, OwnersFilterer: OwnersFilterer{contract: contract}}, nil
}

// Owners is an auto generated Go binding around an Ethereum contract.
type Owners struct {
	OwnersCaller     // Read-only binding to the contract
	OwnersTransactor // Write-only binding to the contract
	OwnersFilterer   // Log filterer for contract events
}

// OwnersCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnersSession struct {
	Contract     *Owners           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnersCallerSession struct {
	Contract *OwnersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnersTransactorSession struct {
	Contract     *OwnersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnersRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnersRaw struct {
	Contract *Owners // Generic contract binding to access the raw methods on
}

// OwnersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnersCallerRaw struct {
	Contract *OwnersCaller // Generic read-only contract binding to access the raw methods on
}

// OwnersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnersTransactorRaw struct {
	Contract *OwnersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwners creates a new instance of Owners, bound to a specific deployed contract.
func NewOwners(address common.Address, backend bind.ContractBackend) (*Owners, error) {
	contract, err := bindOwners(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owners{OwnersCaller: OwnersCaller{contract: contract}, OwnersTransactor: OwnersTransactor{contract: contract}, OwnersFilterer: OwnersFilterer{contract: contract}}, nil
}

// NewOwnersCaller creates a new read-only instance of Owners, bound to a specific deployed contract.
func NewOwnersCaller(address common.Address, caller bind.ContractCaller) (*OwnersCaller, error) {
	contract, err := bindOwners(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnersCaller{contract: contract}, nil
}

// NewOwnersTransactor creates a new write-only instance of Owners, bound to a specific deployed contract.
func NewOwnersTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnersTransactor, error) {
	contract, err := bindOwners(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnersTransactor{contract: contract}, nil
}

// NewOwnersFilterer creates a new log filterer instance of Owners, bound to a specific deployed contract.
func NewOwnersFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnersFilterer, error) {
	contract, err := bindOwners(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnersFilterer{contract: contract}, nil
}

// bindOwners binds a generic wrapper to an already deployed contract.
func bindOwners(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owners *OwnersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owners.Contract.OwnersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owners *OwnersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owners.Contract.OwnersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owners *OwnersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owners.Contract.OwnersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owners *OwnersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Owners.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owners *OwnersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owners.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owners *OwnersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owners.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_Owners *OwnersCaller) IsOwner(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Owners.contract.Call(opts, &out, "isOwner", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_Owners *OwnersSession) IsOwner(account common.Address) (bool, error) {
	return _Owners.Contract.IsOwner(&_Owners.CallOpts, account)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address account) view returns(bool)
func (_Owners *OwnersCallerSession) IsOwner(account common.Address) (bool, error) {
	return _Owners.Contract.IsOwner(&_Owners.CallOpts, account)
}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address[])
func (_Owners *OwnersCaller) Owners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Owners.contract.Call(opts, &out, "owners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address[])
func (_Owners *OwnersSession) Owners() ([]common.Address, error) {
	return _Owners.Contract.Owners(&_Owners.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0xaffe39c1.
//
// Solidity: function owners() view returns(address[])
func (_Owners *OwnersCallerSession) Owners() ([]common.Address, error) {
	return _Owners.Contract.Owners(&_Owners.CallOpts)
}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() view returns(uint8)
func (_Owners *OwnersCaller) OwnersCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Owners.contract.Call(opts, &out, "ownersCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() view returns(uint8)
func (_Owners *OwnersSession) OwnersCount() (uint8, error) {
	return _Owners.Contract.OwnersCount(&_Owners.CallOpts)
}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() view returns(uint8)
func (_Owners *OwnersCallerSession) OwnersCount() (uint8, error) {
	return _Owners.Contract.OwnersCount(&_Owners.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint8)
func (_Owners *OwnersCaller) Quorum(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Owners.contract.Call(opts, &out, "quorum")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint8)
func (_Owners *OwnersSession) Quorum() (uint8, error) {
	return _Owners.Contract.Quorum(&_Owners.CallOpts)
}

// Quorum is a free data retrieval call binding the contract method 0x1703a018.
//
// Solidity: function quorum() view returns(uint8)
func (_Owners *OwnersCallerSession) Quorum() (uint8, error) {
	return _Owners.Contract.Quorum(&_Owners.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Owners *OwnersTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Owners.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Owners *OwnersSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Owners.Contract.ChangeOwner(&_Owners.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Owners *OwnersTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Owners.Contract.ChangeOwner(&_Owners.TransactOpts, newOwner)
}

// OwnersOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Owners contract.
type OwnersOwnerChangedIterator struct {
	Event *OwnersOwnerChanged // Event containing the contract specifics and raw log

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
func (it *OwnersOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnersOwnerChanged)
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
		it.Event = new(OwnersOwnerChanged)
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
func (it *OwnersOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnersOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnersOwnerChanged represents a OwnerChanged event raised by the Owners contract.
type OwnersOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_Owners *OwnersFilterer) FilterOwnerChanged(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*OwnersOwnerChangedIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Owners.contract.FilterLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnersOwnerChangedIterator{contract: _Owners.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_Owners *OwnersFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *OwnersOwnerChanged, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Owners.contract.WatchLogs(opts, "OwnerChanged", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnersOwnerChanged)
				if err := _Owners.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed oldOwner, address indexed newOwner)
func (_Owners *OwnersFilterer) ParseOwnerChanged(log types.Log) (*OwnersOwnerChanged, error) {
	event := new(OwnersOwnerChanged)
	if err := _Owners.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
