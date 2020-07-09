package contracts

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

// Contrato3ABI is the input ABI used to generate the binding from.
const Contrato3ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"customerAdress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ExpirationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCustomer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IsAuthenticated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isAuthenticated\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"}],\"name\":\"SetAuthentication\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Contrato3FuncSigs maps the 4-byte function signature to its string representation.
var Contrato3FuncSigs = map[string]string{
	"b521fae5": "ExpirationTime()",
	"e38f39e7": "GetCustomer()",
	"0ae50a39": "GetOwner()",
	"7b0bfc27": "IsAuthenticated()",
	"f5728984": "SetAuthentication(bool,uint256)",
}

// Contrato3Bin is the compiled bytecode used for deploying new contracts.
var Contrato3Bin = "0x60806040526000805460ff1916905534801561001a57600080fd5b506040516102ac3803806102ac8339818101604052602081101561003d57600080fd5b505160028054336001600160a01b031991821617909155600380549091166001600160a01b039092169190911790556102318061007b6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630ae50a391461005c5780637b0bfc2714610080578063b521fae51461009c578063e38f39e7146100b6578063f5728984146100be575b600080fd5b6100646100e5565b604080516001600160a01b039092168252519081900360200190f35b6100886100f4565b604080519115158252519081900360200190f35b6100a46100fd565b60408051918252519081900360200190f35b610064610103565b6100e3600480360360408110156100d457600080fd5b5080351515906020013561016b565b005b6002546001600160a01b031690565b60005460ff1681565b60015481565b6002546000906001600160a01b0316331461015b576040805162461bcd60e51b815260206004820152601360248201527221b0b63632b91034b9903737ba1037bbb732b960691b604482015290519081900360640190fd5b506003546001600160a01b031690565b6002546001600160a01b031633146101c0576040805162461bcd60e51b815260206004820152601360248201527221b0b63632b91034b9903737ba1037bbb732b960691b604482015290519081900360640190fd5b6000805460ff19169215159290921790915560015556fea264697066735822122083472d8c9eef464c45a8bfd55e5d14e7f9ee08b78aa7ac51a4279e7e25bb9ea064736f6c637826302e362e382d646576656c6f702e323032302e352e352b636f6d6d69742e31646537336131360057"

// DeployContrato3 deploys a new Ethereum contract, binding an instance of Contrato3 to it.
func DeployContrato3(auth *bind.TransactOpts, backend bind.ContractBackend, customerAdress common.Address) (common.Address, *types.Transaction, *Contrato3, error) {
	parsed, err := abi.JSON(strings.NewReader(Contrato3ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Contrato3Bin), backend, customerAdress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contrato3{Contrato3Caller: Contrato3Caller{contract: contract}, Contrato3Transactor: Contrato3Transactor{contract: contract}, Contrato3Filterer: Contrato3Filterer{contract: contract}}, nil
}

// Contrato3 is an auto generated Go binding around an Ethereum contract.
type Contrato3 struct {
	Contrato3Caller     // Read-only binding to the contract
	Contrato3Transactor // Write-only binding to the contract
	Contrato3Filterer   // Log filterer for contract events
}

// Contrato3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Contrato3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Contrato3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Contrato3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Contrato3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Contrato3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Contrato3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Contrato3Session struct {
	Contract     *Contrato3        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Contrato3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Contrato3CallerSession struct {
	Contract *Contrato3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Contrato3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Contrato3TransactorSession struct {
	Contract     *Contrato3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Contrato3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Contrato3Raw struct {
	Contract *Contrato3 // Generic contract binding to access the raw methods on
}

// Contrato3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Contrato3CallerRaw struct {
	Contract *Contrato3Caller // Generic read-only contract binding to access the raw methods on
}

// Contrato3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Contrato3TransactorRaw struct {
	Contract *Contrato3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewContrato3 creates a new instance of Contrato3, bound to a specific deployed contract.
func NewContrato3(address common.Address, backend bind.ContractBackend) (*Contrato3, error) {
	contract, err := bindContrato3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contrato3{Contrato3Caller: Contrato3Caller{contract: contract}, Contrato3Transactor: Contrato3Transactor{contract: contract}, Contrato3Filterer: Contrato3Filterer{contract: contract}}, nil
}

// NewContrato3Caller creates a new read-only instance of Contrato3, bound to a specific deployed contract.
func NewContrato3Caller(address common.Address, caller bind.ContractCaller) (*Contrato3Caller, error) {
	contract, err := bindContrato3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Contrato3Caller{contract: contract}, nil
}

// NewContrato3Transactor creates a new write-only instance of Contrato3, bound to a specific deployed contract.
func NewContrato3Transactor(address common.Address, transactor bind.ContractTransactor) (*Contrato3Transactor, error) {
	contract, err := bindContrato3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Contrato3Transactor{contract: contract}, nil
}

// NewContrato3Filterer creates a new log filterer instance of Contrato3, bound to a specific deployed contract.
func NewContrato3Filterer(address common.Address, filterer bind.ContractFilterer) (*Contrato3Filterer, error) {
	contract, err := bindContrato3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Contrato3Filterer{contract: contract}, nil
}

// bindContrato3 binds a generic wrapper to an already deployed contract.
func bindContrato3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Contrato3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contrato3 *Contrato3Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contrato3.Contract.Contrato3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contrato3 *Contrato3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contrato3.Contract.Contrato3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contrato3 *Contrato3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contrato3.Contract.Contrato3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contrato3 *Contrato3CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contrato3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contrato3 *Contrato3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contrato3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contrato3 *Contrato3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contrato3.Contract.contract.Transact(opts, method, params...)
}

// ExpirationTime is a free data retrieval call binding the contract method 0xb521fae5.
//
// Solidity: function ExpirationTime() constant returns(uint256)
func (_Contrato3 *Contrato3Caller) ExpirationTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contrato3.contract.Call(opts, out, "ExpirationTime")
	return *ret0, err
}

// ExpirationTime is a free data retrieval call binding the contract method 0xb521fae5.
//
// Solidity: function ExpirationTime() constant returns(uint256)
func (_Contrato3 *Contrato3Session) ExpirationTime() (*big.Int, error) {
	return _Contrato3.Contract.ExpirationTime(&_Contrato3.CallOpts)
}

// ExpirationTime is a free data retrieval call binding the contract method 0xb521fae5.
//
// Solidity: function ExpirationTime() constant returns(uint256)
func (_Contrato3 *Contrato3CallerSession) ExpirationTime() (*big.Int, error) {
	return _Contrato3.Contract.ExpirationTime(&_Contrato3.CallOpts)
}

// GetCustomer is a free data retrieval call binding the contract method 0xe38f39e7.
//
// Solidity: function GetCustomer() constant returns(address)
func (_Contrato3 *Contrato3Caller) GetCustomer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contrato3.contract.Call(opts, out, "GetCustomer")
	return *ret0, err
}

// GetCustomer is a free data retrieval call binding the contract method 0xe38f39e7.
//
// Solidity: function GetCustomer() constant returns(address)
func (_Contrato3 *Contrato3Session) GetCustomer() (common.Address, error) {
	return _Contrato3.Contract.GetCustomer(&_Contrato3.CallOpts)
}

// GetCustomer is a free data retrieval call binding the contract method 0xe38f39e7.
//
// Solidity: function GetCustomer() constant returns(address)
func (_Contrato3 *Contrato3CallerSession) GetCustomer() (common.Address, error) {
	return _Contrato3.Contract.GetCustomer(&_Contrato3.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x0ae50a39.
//
// Solidity: function GetOwner() constant returns(address)
func (_Contrato3 *Contrato3Caller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contrato3.contract.Call(opts, out, "GetOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x0ae50a39.
//
// Solidity: function GetOwner() constant returns(address)
func (_Contrato3 *Contrato3Session) GetOwner() (common.Address, error) {
	return _Contrato3.Contract.GetOwner(&_Contrato3.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x0ae50a39.
//
// Solidity: function GetOwner() constant returns(address)
func (_Contrato3 *Contrato3CallerSession) GetOwner() (common.Address, error) {
	return _Contrato3.Contract.GetOwner(&_Contrato3.CallOpts)
}

// IsAuthenticated is a free data retrieval call binding the contract method 0x7b0bfc27.
//
// Solidity: function IsAuthenticated() constant returns(bool)
func (_Contrato3 *Contrato3Caller) IsAuthenticated(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contrato3.contract.Call(opts, out, "IsAuthenticated")
	return *ret0, err
}

// IsAuthenticated is a free data retrieval call binding the contract method 0x7b0bfc27.
//
// Solidity: function IsAuthenticated() constant returns(bool)
func (_Contrato3 *Contrato3Session) IsAuthenticated() (bool, error) {
	return _Contrato3.Contract.IsAuthenticated(&_Contrato3.CallOpts)
}

// IsAuthenticated is a free data retrieval call binding the contract method 0x7b0bfc27.
//
// Solidity: function IsAuthenticated() constant returns(bool)
func (_Contrato3 *Contrato3CallerSession) IsAuthenticated() (bool, error) {
	return _Contrato3.Contract.IsAuthenticated(&_Contrato3.CallOpts)
}

// SetAuthentication is a paid mutator transaction binding the contract method 0xf5728984.
//
// Solidity: function SetAuthentication(bool isAuthenticated, uint256 expirationTime) returns()
func (_Contrato3 *Contrato3Transactor) SetAuthentication(opts *bind.TransactOpts, isAuthenticated bool, expirationTime *big.Int) (*types.Transaction, error) {
	return _Contrato3.contract.Transact(opts, "SetAuthentication", isAuthenticated, expirationTime)
}

// SetAuthentication is a paid mutator transaction binding the contract method 0xf5728984.
//
// Solidity: function SetAuthentication(bool isAuthenticated, uint256 expirationTime) returns()
func (_Contrato3 *Contrato3Session) SetAuthentication(isAuthenticated bool, expirationTime *big.Int) (*types.Transaction, error) {
	return _Contrato3.Contract.SetAuthentication(&_Contrato3.TransactOpts, isAuthenticated, expirationTime)
}

// SetAuthentication is a paid mutator transaction binding the contract method 0xf5728984.
//
// Solidity: function SetAuthentication(bool isAuthenticated, uint256 expirationTime) returns()
func (_Contrato3 *Contrato3TransactorSession) SetAuthentication(isAuthenticated bool, expirationTime *big.Int) (*types.Transaction, error) {
	return _Contrato3.Contract.SetAuthentication(&_Contrato3.TransactOpts, isAuthenticated, expirationTime)
}
