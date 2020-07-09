package contracts

//package contract

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

// Contrato1ABI is the input ABI used to generate the binding from.
const Contrato1ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"customerAdress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"GetCustomer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IsAuthenticated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isAuthenticated\",\"type\":\"bool\"}],\"name\":\"SetIsAuthentication\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Contrato1FuncSigs maps the 4-byte function signature to its string representation.
var Contrato1FuncSigs = map[string]string{
	"e38f39e7": "GetCustomer()",
	"7b0bfc27": "IsAuthenticated()",
	"ca8474a1": "SetIsAuthentication(bool)",
}

// Contrato1Bin is the compiled bytecode used for deploying new contracts.
var Contrato1Bin = "0x60806040526000805460ff1916905534801561001a57600080fd5b506040516102683803806102688339818101604052602081101561003d57600080fd5b505160008054610100600160a81b0319163361010002178155600180546001600160a01b0319166001600160a01b03909316929092179091556101e290819061008690396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80637b0bfc2714610046578063ca8474a114610062578063e38f39e714610083575b600080fd5b61004e6100a7565b604080519115158252519081900360200190f35b6100816004803603602081101561007857600080fd5b503515156100b0565b005b61008b61011d565b604080516001600160a01b039092168252519081900360200190f35b60005460ff1681565b60005461010090046001600160a01b0316331461010a576040805162461bcd60e51b815260206004820152601360248201527221b0b63632b91034b9903737ba1037bbb732b960691b604482015290519081900360640190fd5b6000805460ff1916911515919091179055565b6000805461010090046001600160a01b03163314610178576040805162461bcd60e51b815260206004820152601360248201527221b0b63632b91034b9903737ba1037bbb732b960691b604482015290519081900360640190fd5b506001546001600160a01b03169056fea26469706673582212201a47246228661c727b601d97f49d7846177f596e1bc713f0b304f3e057ecf0f864736f6c637826302e362e362d646576656c6f702e323032302e342e382b636f6d6d69742e37613965303234650057"

// DeployContrato1 deploys a new Ethereum contract, binding an instance of Contrato1 to it.
func DeployContrato1(auth *bind.TransactOpts, backend bind.ContractBackend, customerAdress common.Address) (common.Address, *types.Transaction, *Contrato1, error) {
	parsed, err := abi.JSON(strings.NewReader(Contrato1ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Contrato1Bin), backend, customerAdress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contrato1{Contrato1Caller: Contrato1Caller{contract: contract}, Contrato1Transactor: Contrato1Transactor{contract: contract}, Contrato1Filterer: Contrato1Filterer{contract: contract}}, nil
}

// Contrato1 is an auto generated Go binding around an Ethereum contract.
type Contrato1 struct {
	Contrato1Caller     // Read-only binding to the contract
	Contrato1Transactor // Write-only binding to the contract
	Contrato1Filterer   // Log filterer for contract events
}

// Contrato1Caller is an auto generated read-only Go binding around an Ethereum contract.
type Contrato1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Contrato1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Contrato1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Contrato1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Contrato1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Contrato1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Contrato1Session struct {
	Contract     *Contrato1        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Contrato1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Contrato1CallerSession struct {
	Contract *Contrato1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Contrato1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Contrato1TransactorSession struct {
	Contract     *Contrato1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Contrato1Raw is an auto generated low-level Go binding around an Ethereum contract.
type Contrato1Raw struct {
	Contract *Contrato1 // Generic contract binding to access the raw methods on
}

// Contrato1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Contrato1CallerRaw struct {
	Contract *Contrato1Caller // Generic read-only contract binding to access the raw methods on
}

// Contrato1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Contrato1TransactorRaw struct {
	Contract *Contrato1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewContrato1 creates a new instance of Contrato1, bound to a specific deployed contract.
func NewContrato1(address common.Address, backend bind.ContractBackend) (*Contrato1, error) {
	contract, err := bindContrato1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contrato1{Contrato1Caller: Contrato1Caller{contract: contract}, Contrato1Transactor: Contrato1Transactor{contract: contract}, Contrato1Filterer: Contrato1Filterer{contract: contract}}, nil
}

// NewContrato1Caller creates a new read-only instance of Contrato1, bound to a specific deployed contract.
func NewContrato1Caller(address common.Address, caller bind.ContractCaller) (*Contrato1Caller, error) {
	contract, err := bindContrato1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Contrato1Caller{contract: contract}, nil
}

// NewContrato1Transactor creates a new write-only instance of Contrato1, bound to a specific deployed contract.
func NewContrato1Transactor(address common.Address, transactor bind.ContractTransactor) (*Contrato1Transactor, error) {
	contract, err := bindContrato1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Contrato1Transactor{contract: contract}, nil
}

// NewContrato1Filterer creates a new log filterer instance of Contrato1, bound to a specific deployed contract.
func NewContrato1Filterer(address common.Address, filterer bind.ContractFilterer) (*Contrato1Filterer, error) {
	contract, err := bindContrato1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Contrato1Filterer{contract: contract}, nil
}

// bindContrato1 binds a generic wrapper to an already deployed contract.
func bindContrato1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Contrato1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contrato1 *Contrato1Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contrato1.Contract.Contrato1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contrato1 *Contrato1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contrato1.Contract.Contrato1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contrato1 *Contrato1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contrato1.Contract.Contrato1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contrato1 *Contrato1CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contrato1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contrato1 *Contrato1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contrato1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contrato1 *Contrato1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contrato1.Contract.contract.Transact(opts, method, params...)
}

// GetCustomer is a free data retrieval call binding the contract method 0xe38f39e7.
//
// Solidity: function GetCustomer() constant returns(address)
func (_Contrato1 *Contrato1Caller) GetCustomer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contrato1.contract.Call(opts, out, "GetCustomer")
	return *ret0, err
}

// GetCustomer is a free data retrieval call binding the contract method 0xe38f39e7.
//
// Solidity: function GetCustomer() constant returns(address)
func (_Contrato1 *Contrato1Session) GetCustomer() (common.Address, error) {
	return _Contrato1.Contract.GetCustomer(&_Contrato1.CallOpts)
}

// GetCustomer is a free data retrieval call binding the contract method 0xe38f39e7.
//
// Solidity: function GetCustomer() constant returns(address)
func (_Contrato1 *Contrato1CallerSession) GetCustomer() (common.Address, error) {
	return _Contrato1.Contract.GetCustomer(&_Contrato1.CallOpts)
}

// IsAuthenticated is a free data retrieval call binding the contract method 0x7b0bfc27.
//
// Solidity: function IsAuthenticated() constant returns(bool)
func (_Contrato1 *Contrato1Caller) IsAuthenticated(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contrato1.contract.Call(opts, out, "IsAuthenticated")
	return *ret0, err
}

// IsAuthenticated is a free data retrieval call binding the contract method 0x7b0bfc27.
//
// Solidity: function IsAuthenticated() constant returns(bool)
func (_Contrato1 *Contrato1Session) IsAuthenticated() (bool, error) {
	return _Contrato1.Contract.IsAuthenticated(&_Contrato1.CallOpts)
}

// IsAuthenticated is a free data retrieval call binding the contract method 0x7b0bfc27.
//
// Solidity: function IsAuthenticated() constant returns(bool)
func (_Contrato1 *Contrato1CallerSession) IsAuthenticated() (bool, error) {
	return _Contrato1.Contract.IsAuthenticated(&_Contrato1.CallOpts)
}

// SetIsAuthentication is a paid mutator transaction binding the contract method 0xca8474a1.
//
// Solidity: function SetIsAuthentication(bool isAuthenticated) returns()
func (_Contrato1 *Contrato1Transactor) SetIsAuthentication(opts *bind.TransactOpts, isAuthenticated bool) (*types.Transaction, error) {
	return _Contrato1.contract.Transact(opts, "SetIsAuthentication", isAuthenticated)
}

// SetIsAuthentication is a paid mutator transaction binding the contract method 0xca8474a1.
//
// Solidity: function SetIsAuthentication(bool isAuthenticated) returns()
func (_Contrato1 *Contrato1Session) SetIsAuthentication(isAuthenticated bool) (*types.Transaction, error) {
	return _Contrato1.Contract.SetIsAuthentication(&_Contrato1.TransactOpts, isAuthenticated)
}

// SetIsAuthentication is a paid mutator transaction binding the contract method 0xca8474a1.
//
// Solidity: function SetIsAuthentication(bool isAuthenticated) returns()
func (_Contrato1 *Contrato1TransactorSession) SetIsAuthentication(isAuthenticated bool) (*types.Transaction, error) {
	return _Contrato1.Contract.SetIsAuthentication(&_Contrato1.TransactOpts, isAuthenticated)
}
