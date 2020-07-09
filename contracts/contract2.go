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

// Contrato2ABI is the input ABI used to generate the binding from.
const Contrato2ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"customerAdress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ExpirationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCustomer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IsAuthenticated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isAuthenticated\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"}],\"name\":\"SetAuthentication\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Contrato2FuncSigs maps the 4-byte function signature to its string representation.
var Contrato2FuncSigs = map[string]string{
	"b521fae5": "ExpirationTime()",
	"e38f39e7": "GetCustomer()",
	"7b0bfc27": "IsAuthenticated()",
	"f5728984": "SetAuthentication(bool,uint256)",
}

// Contrato2Bin is the compiled bytecode used for deploying new contracts.
var Contrato2Bin = "0x60806040526000805460ff1916905534801561001a57600080fd5b5060405161028b38038061028b8339818101604052602081101561003d57600080fd5b505160028054336001600160a01b031991821617909155600380549091166001600160a01b039092169190911790556102108061007b6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80637b0bfc2714610051578063b521fae51461006d578063e38f39e714610087578063f5728984146100ab575b600080fd5b6100596100d2565b604080519115158252519081900360200190f35b6100756100db565b60408051918252519081900360200190f35b61008f6100e1565b604080516001600160a01b039092168252519081900360200190f35b6100d0600480360360408110156100c157600080fd5b50803515159060200135610149565b005b60005460ff1681565b60015481565b6002546000906001600160a01b03163314610139576040805162461bcd60e51b815260206004820152601360248201527221b0b63632b91034b9903737ba1037bbb732b960691b604482015290519081900360640190fd5b506003546001600160a01b031690565b6002546001600160a01b0316331461019e576040805162461bcd60e51b815260206004820152601360248201527221b0b63632b91034b9903737ba1037bbb732b960691b604482015290519081900360640190fd5b6000805460ff19169215159290921790915560015556fea26469706673582212205dc5048c2dbd6c13e3903aab6bf1efdde925ca7e53c7b8dd7a6d1879d774fa2b64736f6c637827302e362e372d646576656c6f702e323032302e342e32372b636f6d6d69742e39396161383231340058"

// DeployContrato2 deploys a new Ethereum contract, binding an instance of Contrato2 to it.
func DeployContrato2(auth *bind.TransactOpts, backend bind.ContractBackend, customerAdress common.Address) (common.Address, *types.Transaction, *Contrato2, error) {
	parsed, err := abi.JSON(strings.NewReader(Contrato2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Contrato2Bin), backend, customerAdress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contrato2{Contrato2Caller: Contrato2Caller{contract: contract}, Contrato2Transactor: Contrato2Transactor{contract: contract}, Contrato2Filterer: Contrato2Filterer{contract: contract}}, nil
}

// Contrato2 is an auto generated Go binding around an Ethereum contract.
type Contrato2 struct {
	Contrato2Caller     // Read-only binding to the contract
	Contrato2Transactor // Write-only binding to the contract
	Contrato2Filterer   // Log filterer for contract events
}

// Contrato2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Contrato2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Contrato2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Contrato2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Contrato2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Contrato2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Contrato2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Contrato2Session struct {
	Contract     *Contrato2        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Contrato2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Contrato2CallerSession struct {
	Contract *Contrato2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Contrato2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Contrato2TransactorSession struct {
	Contract     *Contrato2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Contrato2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Contrato2Raw struct {
	Contract *Contrato2 // Generic contract binding to access the raw methods on
}

// Contrato2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Contrato2CallerRaw struct {
	Contract *Contrato2Caller // Generic read-only contract binding to access the raw methods on
}

// Contrato2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Contrato2TransactorRaw struct {
	Contract *Contrato2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewContrato2 creates a new instance of Contrato2, bound to a specific deployed contract.
func NewContrato2(address common.Address, backend bind.ContractBackend) (*Contrato2, error) {
	contract, err := bindContrato2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contrato2{Contrato2Caller: Contrato2Caller{contract: contract}, Contrato2Transactor: Contrato2Transactor{contract: contract}, Contrato2Filterer: Contrato2Filterer{contract: contract}}, nil
}

// NewContrato2Caller creates a new read-only instance of Contrato2, bound to a specific deployed contract.
func NewContrato2Caller(address common.Address, caller bind.ContractCaller) (*Contrato2Caller, error) {
	contract, err := bindContrato2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Contrato2Caller{contract: contract}, nil
}

// NewContrato2Transactor creates a new write-only instance of Contrato2, bound to a specific deployed contract.
func NewContrato2Transactor(address common.Address, transactor bind.ContractTransactor) (*Contrato2Transactor, error) {
	contract, err := bindContrato2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Contrato2Transactor{contract: contract}, nil
}

// NewContrato2Filterer creates a new log filterer instance of Contrato2, bound to a specific deployed contract.
func NewContrato2Filterer(address common.Address, filterer bind.ContractFilterer) (*Contrato2Filterer, error) {
	contract, err := bindContrato2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Contrato2Filterer{contract: contract}, nil
}

// bindContrato2 binds a generic wrapper to an already deployed contract.
func bindContrato2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Contrato2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contrato2 *Contrato2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contrato2.Contract.Contrato2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contrato2 *Contrato2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contrato2.Contract.Contrato2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contrato2 *Contrato2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contrato2.Contract.Contrato2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contrato2 *Contrato2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contrato2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contrato2 *Contrato2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contrato2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contrato2 *Contrato2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contrato2.Contract.contract.Transact(opts, method, params...)
}

// ExpirationTime is a free data retrieval call binding the contract method 0xb521fae5.
//
// Solidity: function ExpirationTime() constant returns(uint256)
func (_Contrato2 *Contrato2Caller) ExpirationTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contrato2.contract.Call(opts, out, "ExpirationTime")
	return *ret0, err
}

// ExpirationTime is a free data retrieval call binding the contract method 0xb521fae5.
//
// Solidity: function ExpirationTime() constant returns(uint256)
func (_Contrato2 *Contrato2Session) ExpirationTime() (*big.Int, error) {
	return _Contrato2.Contract.ExpirationTime(&_Contrato2.CallOpts)
}

// ExpirationTime is a free data retrieval call binding the contract method 0xb521fae5.
//
// Solidity: function ExpirationTime() constant returns(uint256)
func (_Contrato2 *Contrato2CallerSession) ExpirationTime() (*big.Int, error) {
	return _Contrato2.Contract.ExpirationTime(&_Contrato2.CallOpts)
}

// GetCustomer is a free data retrieval call binding the contract method 0xe38f39e7.
//
// Solidity: function GetCustomer() constant returns(address)
func (_Contrato2 *Contrato2Caller) GetCustomer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contrato2.contract.Call(opts, out, "GetCustomer")
	return *ret0, err
}

// GetCustomer is a free data retrieval call binding the contract method 0xe38f39e7.
//
// Solidity: function GetCustomer() constant returns(address)
func (_Contrato2 *Contrato2Session) GetCustomer() (common.Address, error) {
	return _Contrato2.Contract.GetCustomer(&_Contrato2.CallOpts)
}

// GetCustomer is a free data retrieval call binding the contract method 0xe38f39e7.
//
// Solidity: function GetCustomer() constant returns(address)
func (_Contrato2 *Contrato2CallerSession) GetCustomer() (common.Address, error) {
	return _Contrato2.Contract.GetCustomer(&_Contrato2.CallOpts)
}

// IsAuthenticated is a free data retrieval call binding the contract method 0x7b0bfc27.
//
// Solidity: function IsAuthenticated() constant returns(bool)
func (_Contrato2 *Contrato2Caller) IsAuthenticated(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contrato2.contract.Call(opts, out, "IsAuthenticated")
	return *ret0, err
}

// IsAuthenticated is a free data retrieval call binding the contract method 0x7b0bfc27.
//
// Solidity: function IsAuthenticated() constant returns(bool)
func (_Contrato2 *Contrato2Session) IsAuthenticated() (bool, error) {
	return _Contrato2.Contract.IsAuthenticated(&_Contrato2.CallOpts)
}

// IsAuthenticated is a free data retrieval call binding the contract method 0x7b0bfc27.
//
// Solidity: function IsAuthenticated() constant returns(bool)
func (_Contrato2 *Contrato2CallerSession) IsAuthenticated() (bool, error) {
	return _Contrato2.Contract.IsAuthenticated(&_Contrato2.CallOpts)
}

// SetAuthentication is a paid mutator transaction binding the contract method 0xf5728984.
//
// Solidity: function SetAuthentication(bool isAuthenticated, uint256 expirationTime) returns()
func (_Contrato2 *Contrato2Transactor) SetAuthentication(opts *bind.TransactOpts, isAuthenticated bool, expirationTime *big.Int) (*types.Transaction, error) {
	return _Contrato2.contract.Transact(opts, "SetAuthentication", isAuthenticated, expirationTime)
}

// SetAuthentication is a paid mutator transaction binding the contract method 0xf5728984.
//
// Solidity: function SetAuthentication(bool isAuthenticated, uint256 expirationTime) returns()
func (_Contrato2 *Contrato2Session) SetAuthentication(isAuthenticated bool, expirationTime *big.Int) (*types.Transaction, error) {
	return _Contrato2.Contract.SetAuthentication(&_Contrato2.TransactOpts, isAuthenticated, expirationTime)
}

// SetAuthentication is a paid mutator transaction binding the contract method 0xf5728984.
//
// Solidity: function SetAuthentication(bool isAuthenticated, uint256 expirationTime) returns()
func (_Contrato2 *Contrato2TransactorSession) SetAuthentication(isAuthenticated bool, expirationTime *big.Int) (*types.Transaction, error) {
	return _Contrato2.Contract.SetAuthentication(&_Contrato2.TransactOpts, isAuthenticated, expirationTime)
}
