// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package owners

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

// OwnersABI is the input ABI used to generate the binding from.
const OwnersABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"ownersList\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"quorumSize\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownersCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OwnersBin is the compiled bytecode used for deploying new contracts.
var OwnersBin = "0x60806040523480156200001157600080fd5b5060405162000a4938038062000a49833981016040819052620000349162000322565b60008251118015620000485750815160ff10155b6200009a5760405162461bcd60e51b815260206004820152601860248201527f4f776e6572732f6e6f7420656e6f756768206f776e657273000000000000000060448201526064015b60405180910390fd5b60008160ff16118015620000b2575081518160ff1611155b620001005760405162461bcd60e51b815260206004820152601a60248201527f4f776e6572732f71756f72756d206f7574206f662072616e6765000000000000604482015260640162000091565b60005b82518160ff161015620002355760026000848360ff16815181106200013857634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff16158015620001ab575060006001600160a01b0316838260ff16815181106200019757634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b031614155b620001c657634e487b7160e01b600052600160045260246000fd5b600160026000858460ff1681518110620001f057634e487b7160e01b600052603260045260246000fd5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff19169115159190911790556200022d816200040a565b905062000103565b5081516200024b90600190602085019062000272565b5090516000805461ffff191661010060ff9384160260ff191617919092161790556200044d565b828054828255906000526020600020908101928215620002ca579160200282015b82811115620002ca57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000293565b50620002d8929150620002dc565b5090565b5b80821115620002d85760008155600101620002dd565b80516001600160a01b03811681146200030b57600080fd5b919050565b805160ff811681146200030b57600080fd5b6000806040838503121562000335578182fd5b82516001600160401b03808211156200034c578384fd5b818501915085601f83011262000360578384fd5b815160208282111562000377576200037762000437565b8160051b604051601f19603f830116810181811086821117156200039f576200039f62000437565b604052838152828101945085830182870184018b1015620003be578889fd5b8896505b84871015620003eb57620003d681620002f3565b865260019690960195948301948301620003c2565b509650620003fd905087820162000310565b9450505050509250929050565b600060ff821660ff8114156200042e57634e487b7160e01b82526011600452602482fd5b60010192915050565b634e487b7160e01b600052604160045260246000fd5b6105ec806200045d6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80631703a0181461005c5780632f54bf6e1461007a578063a6f9dae1146100b6578063affe39c1146100cb578063b9488546146100e0575b600080fd5b60005460ff165b60405160ff90911681526020015b60405180910390f35b6100a661008836600461050f565b6001600160a01b031660009081526002602052604090205460ff1690565b6040519015158152602001610071565b6100c96100c436600461050f565b6100f0565b005b6100d3610433565b604051610071919061053d565b600054610100900460ff16610063565b3360009081526002602052604090205460ff166101545760405162461bcd60e51b815260206004820152601d60248201527f4f776e6572732f73656e646572206973206e6f7420616e206f776e657200000060448201526064015b60405180910390fd5b6001600160a01b03811660009081526002602052604090205460ff1615801561018557506001600160a01b03811615155b6101d15760405162461bcd60e51b815260206004820152601c60248201527f4f776e6572732f696e76616c6964206164647265737320676976656e00000000604482015260640161014b565b600154158015906101e5575060015460ff10155b6101ff57634e487b7160e01b600052600160045260246000fd5b60015460009067ffffffffffffffff81111561022b57634e487b7160e01b600052604160045260246000fd5b604051908082528060200260200182016040528015610254578160200160208202803683370190505b50905060005b60015460ff8216101561038e57336001600160a01b031660018260ff168154811061029557634e487b7160e01b600052603260045260246000fd5b6000918252602090912001546001600160a01b0316146103395760018160ff16815481106102d357634e487b7160e01b600052603260045260246000fd5b9060005260206000200160009054906101000a90046001600160a01b0316828260ff168151811061031457634e487b7160e01b600052603260045260246000fd5b60200260200101906001600160a01b031690816001600160a01b03168152505061037e565b82828260ff168151811061035d57634e487b7160e01b600052603260045260246000fd5b60200260200101906001600160a01b031690816001600160a01b0316815250505b6103878161058a565b905061025a565b50600054815160ff909116146103b457634e487b7160e01b600052600160045260246000fd5b6040516001600160a01b0383169033907fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c90600090a380516103fd906001906020840190610495565b50506001600160a01b0316600090815260026020526040808220805460ff19908116600117909155338352912080549091169055565b6060600180548060200260200160405190810160405280929190818152602001828054801561048b57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161046d575b5050505050905090565b8280548282559060005260206000209081019282156104ea579160200282015b828111156104ea57825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906104b5565b506104f69291506104fa565b5090565b5b808211156104f657600081556001016104fb565b600060208284031215610520578081fd5b81356001600160a01b0381168114610536578182fd5b9392505050565b6020808252825182820181905260009190848201906040850190845b8181101561057e5783516001600160a01b031683529284019291840191600101610559565b50909695505050505050565b600060ff821660ff8114156105ad57634e487b7160e01b82526011600452602482fd5b6001019291505056fea2646970667358221220ffa961b5735597882bd22579bf9351979c02a7426ea816ffcc3b3349f80dfd8464736f6c63430008040033"

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
