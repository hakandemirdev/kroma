// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
	_ = abi.ConvertType
)

// ProposerFeeVaultMetaData contains all meta data concerning the ProposerFeeVault contract.
var ProposerFeeVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_WITHDRAWAL_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECIPIENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalProcessed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x61012060405234801561001157600080fd5b506040516108e53803806108e58339810160408190526100309161005d565b678ac7230489e800006080526001600160a01b031660a052600060c0819052600160e0526101005261008d565b60006020828403121561006f57600080fd5b81516001600160a01b038116811461008657600080fd5b9392505050565b60805160a05160c05160e051610100516108006100e560003960006103d3015260006103aa01526000610381015260008181607c015281816102570152610319015260008181610137015261015b01526108006000f3fe60806040526004361061005e5760003560e01c806354fd4d501161004357806354fd4d50146100df57806384411d6514610101578063d3e5792b1461012557600080fd5b80630d9019e11461006a5780633ccfd60b146100c857600080fd5b3661006557005b600080fd5b34801561007657600080fd5b5061009e7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156100d457600080fd5b506100dd610159565b005b3480156100eb57600080fd5b506100f461037a565b6040516100bf91906105d4565b34801561010d57600080fd5b5061011760005481565b6040519081526020016100bf565b34801561013157600080fd5b506101177f000000000000000000000000000000000000000000000000000000000000000081565b7f0000000000000000000000000000000000000000000000000000000000000000471015610233576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f4665655661756c743a207769746864726177616c20616d6f756e74206d75737460448201527f2062652067726561746572207468616e206d696e696d756d207769746864726160648201527f77616c20616d6f756e7400000000000000000000000000000000000000000000608482015260a40160405180910390fd5b600047905080600080828254610249919061061d565b9091555050604080518281527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166020820152338183015290517fc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba9181900360600190a1604080516020810182526000815290517fe11013dd0000000000000000000000000000000000000000000000000000000081527342000000000000000000000000000000000000099163e11013dd918491610345917f0000000000000000000000000000000000000000000000000000000000000000916188b891600401610635565b6000604051808303818588803b15801561035e57600080fd5b505af1158015610372573d6000803e3d6000fd5b505050505050565b60606103a57f000000000000000000000000000000000000000000000000000000000000000061041d565b6103ce7f000000000000000000000000000000000000000000000000000000000000000061041d565b6103f77f000000000000000000000000000000000000000000000000000000000000000061041d565b60405160200161040993929190610679565b604051602081830303815290604052905090565b60608160000361046057505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b811561048a5780610474816106ef565b91506104839050600a83610756565b9150610464565b60008167ffffffffffffffff8111156104a5576104a561076a565b6040519080825280601f01601f1916602001820160405280156104cf576020820181803683370190505b5090505b8415610552576104e4600183610799565b91506104f1600a866107b0565b6104fc90603061061d565b60f81b818381518110610511576105116107c4565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535061054b600a86610756565b94506104d3565b949350505050565b60005b8381101561057557818101518382015260200161055d565b83811115610584576000848401525b50505050565b600081518084526105a281602086016020860161055a565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006105e7602083018461058a565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115610630576106306105ee565b500190565b73ffffffffffffffffffffffffffffffffffffffff8416815263ffffffff83166020820152606060408201526000610670606083018461058a565b95945050505050565b6000845161068b81846020890161055a565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516106c7816001850160208a0161055a565b600192019182015283516106e281600284016020880161055a565b0160020195945050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610720576107206105ee565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261076557610765610727565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000828210156107ab576107ab6105ee565b500390565b6000826107bf576107bf610727565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a",
}

// ProposerFeeVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use ProposerFeeVaultMetaData.ABI instead.
var ProposerFeeVaultABI = ProposerFeeVaultMetaData.ABI

// ProposerFeeVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProposerFeeVaultMetaData.Bin instead.
var ProposerFeeVaultBin = ProposerFeeVaultMetaData.Bin

// DeployProposerFeeVault deploys a new Ethereum contract, binding an instance of ProposerFeeVault to it.
func DeployProposerFeeVault(auth *bind.TransactOpts, backend bind.ContractBackend, _recipient common.Address) (common.Address, *types.Transaction, *ProposerFeeVault, error) {
	parsed, err := ProposerFeeVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProposerFeeVaultBin), backend, _recipient)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProposerFeeVault{ProposerFeeVaultCaller: ProposerFeeVaultCaller{contract: contract}, ProposerFeeVaultTransactor: ProposerFeeVaultTransactor{contract: contract}, ProposerFeeVaultFilterer: ProposerFeeVaultFilterer{contract: contract}}, nil
}

// ProposerFeeVault is an auto generated Go binding around an Ethereum contract.
type ProposerFeeVault struct {
	ProposerFeeVaultCaller     // Read-only binding to the contract
	ProposerFeeVaultTransactor // Write-only binding to the contract
	ProposerFeeVaultFilterer   // Log filterer for contract events
}

// ProposerFeeVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProposerFeeVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposerFeeVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProposerFeeVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposerFeeVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProposerFeeVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposerFeeVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProposerFeeVaultSession struct {
	Contract     *ProposerFeeVault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProposerFeeVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProposerFeeVaultCallerSession struct {
	Contract *ProposerFeeVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ProposerFeeVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProposerFeeVaultTransactorSession struct {
	Contract     *ProposerFeeVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ProposerFeeVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProposerFeeVaultRaw struct {
	Contract *ProposerFeeVault // Generic contract binding to access the raw methods on
}

// ProposerFeeVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProposerFeeVaultCallerRaw struct {
	Contract *ProposerFeeVaultCaller // Generic read-only contract binding to access the raw methods on
}

// ProposerFeeVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProposerFeeVaultTransactorRaw struct {
	Contract *ProposerFeeVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProposerFeeVault creates a new instance of ProposerFeeVault, bound to a specific deployed contract.
func NewProposerFeeVault(address common.Address, backend bind.ContractBackend) (*ProposerFeeVault, error) {
	contract, err := bindProposerFeeVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProposerFeeVault{ProposerFeeVaultCaller: ProposerFeeVaultCaller{contract: contract}, ProposerFeeVaultTransactor: ProposerFeeVaultTransactor{contract: contract}, ProposerFeeVaultFilterer: ProposerFeeVaultFilterer{contract: contract}}, nil
}

// NewProposerFeeVaultCaller creates a new read-only instance of ProposerFeeVault, bound to a specific deployed contract.
func NewProposerFeeVaultCaller(address common.Address, caller bind.ContractCaller) (*ProposerFeeVaultCaller, error) {
	contract, err := bindProposerFeeVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProposerFeeVaultCaller{contract: contract}, nil
}

// NewProposerFeeVaultTransactor creates a new write-only instance of ProposerFeeVault, bound to a specific deployed contract.
func NewProposerFeeVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*ProposerFeeVaultTransactor, error) {
	contract, err := bindProposerFeeVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProposerFeeVaultTransactor{contract: contract}, nil
}

// NewProposerFeeVaultFilterer creates a new log filterer instance of ProposerFeeVault, bound to a specific deployed contract.
func NewProposerFeeVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*ProposerFeeVaultFilterer, error) {
	contract, err := bindProposerFeeVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProposerFeeVaultFilterer{contract: contract}, nil
}

// bindProposerFeeVault binds a generic wrapper to an already deployed contract.
func bindProposerFeeVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProposerFeeVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProposerFeeVault *ProposerFeeVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProposerFeeVault.Contract.ProposerFeeVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProposerFeeVault *ProposerFeeVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposerFeeVault.Contract.ProposerFeeVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProposerFeeVault *ProposerFeeVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProposerFeeVault.Contract.ProposerFeeVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProposerFeeVault *ProposerFeeVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProposerFeeVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProposerFeeVault *ProposerFeeVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposerFeeVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProposerFeeVault *ProposerFeeVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProposerFeeVault.Contract.contract.Transact(opts, method, params...)
}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_ProposerFeeVault *ProposerFeeVaultCaller) MINWITHDRAWALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProposerFeeVault.contract.Call(opts, &out, "MIN_WITHDRAWAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_ProposerFeeVault *ProposerFeeVaultSession) MINWITHDRAWALAMOUNT() (*big.Int, error) {
	return _ProposerFeeVault.Contract.MINWITHDRAWALAMOUNT(&_ProposerFeeVault.CallOpts)
}

// MINWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xd3e5792b.
//
// Solidity: function MIN_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_ProposerFeeVault *ProposerFeeVaultCallerSession) MINWITHDRAWALAMOUNT() (*big.Int, error) {
	return _ProposerFeeVault.Contract.MINWITHDRAWALAMOUNT(&_ProposerFeeVault.CallOpts)
}

// RECIPIENT is a free data retrieval call binding the contract method 0x0d9019e1.
//
// Solidity: function RECIPIENT() view returns(address)
func (_ProposerFeeVault *ProposerFeeVaultCaller) RECIPIENT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProposerFeeVault.contract.Call(opts, &out, "RECIPIENT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RECIPIENT is a free data retrieval call binding the contract method 0x0d9019e1.
//
// Solidity: function RECIPIENT() view returns(address)
func (_ProposerFeeVault *ProposerFeeVaultSession) RECIPIENT() (common.Address, error) {
	return _ProposerFeeVault.Contract.RECIPIENT(&_ProposerFeeVault.CallOpts)
}

// RECIPIENT is a free data retrieval call binding the contract method 0x0d9019e1.
//
// Solidity: function RECIPIENT() view returns(address)
func (_ProposerFeeVault *ProposerFeeVaultCallerSession) RECIPIENT() (common.Address, error) {
	return _ProposerFeeVault.Contract.RECIPIENT(&_ProposerFeeVault.CallOpts)
}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_ProposerFeeVault *ProposerFeeVaultCaller) TotalProcessed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProposerFeeVault.contract.Call(opts, &out, "totalProcessed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_ProposerFeeVault *ProposerFeeVaultSession) TotalProcessed() (*big.Int, error) {
	return _ProposerFeeVault.Contract.TotalProcessed(&_ProposerFeeVault.CallOpts)
}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_ProposerFeeVault *ProposerFeeVaultCallerSession) TotalProcessed() (*big.Int, error) {
	return _ProposerFeeVault.Contract.TotalProcessed(&_ProposerFeeVault.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ProposerFeeVault *ProposerFeeVaultCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProposerFeeVault.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ProposerFeeVault *ProposerFeeVaultSession) Version() (string, error) {
	return _ProposerFeeVault.Contract.Version(&_ProposerFeeVault.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_ProposerFeeVault *ProposerFeeVaultCallerSession) Version() (string, error) {
	return _ProposerFeeVault.Contract.Version(&_ProposerFeeVault.CallOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ProposerFeeVault *ProposerFeeVaultTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposerFeeVault.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ProposerFeeVault *ProposerFeeVaultSession) Withdraw() (*types.Transaction, error) {
	return _ProposerFeeVault.Contract.Withdraw(&_ProposerFeeVault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ProposerFeeVault *ProposerFeeVaultTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ProposerFeeVault.Contract.Withdraw(&_ProposerFeeVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ProposerFeeVault *ProposerFeeVaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposerFeeVault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ProposerFeeVault *ProposerFeeVaultSession) Receive() (*types.Transaction, error) {
	return _ProposerFeeVault.Contract.Receive(&_ProposerFeeVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ProposerFeeVault *ProposerFeeVaultTransactorSession) Receive() (*types.Transaction, error) {
	return _ProposerFeeVault.Contract.Receive(&_ProposerFeeVault.TransactOpts)
}

// ProposerFeeVaultWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the ProposerFeeVault contract.
type ProposerFeeVaultWithdrawalIterator struct {
	Event *ProposerFeeVaultWithdrawal // Event containing the contract specifics and raw log

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
func (it *ProposerFeeVaultWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposerFeeVaultWithdrawal)
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
		it.Event = new(ProposerFeeVaultWithdrawal)
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
func (it *ProposerFeeVaultWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposerFeeVaultWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposerFeeVaultWithdrawal represents a Withdrawal event raised by the ProposerFeeVault contract.
type ProposerFeeVaultWithdrawal struct {
	Value *big.Int
	To    common.Address
	From  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_ProposerFeeVault *ProposerFeeVaultFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*ProposerFeeVaultWithdrawalIterator, error) {

	logs, sub, err := _ProposerFeeVault.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &ProposerFeeVaultWithdrawalIterator{contract: _ProposerFeeVault.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_ProposerFeeVault *ProposerFeeVaultFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ProposerFeeVaultWithdrawal) (event.Subscription, error) {

	logs, sub, err := _ProposerFeeVault.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposerFeeVaultWithdrawal)
				if err := _ProposerFeeVault.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_ProposerFeeVault *ProposerFeeVaultFilterer) ParseWithdrawal(log types.Log) (*ProposerFeeVaultWithdrawal, error) {
	event := new(ProposerFeeVaultWithdrawal)
	if err := _ProposerFeeVault.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
