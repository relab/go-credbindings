// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package notary

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
