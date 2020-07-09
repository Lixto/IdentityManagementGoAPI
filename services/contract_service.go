package services

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"../../API-Go-Blockchain/contracts"
	"../../API-Go-Blockchain/parameters"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getClient() *ethclient.Client {
	//Connect to the client
	conn, err := ethclient.Dial(parameters.ConnURL)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func getTransactOptions() *bind.TransactOpts {
	conn := getClient()
	//Private key from ganache only for dev proposes
	privateKey, err := crypto.HexToECDSA(parameters.OwnerPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := conn.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	return auth
}

func getCustomerAddress() common.Address {
	customerAddress := common.HexToAddress(parameters.OwnerAddress)
	return customerAddress
}

//DeployContract function to deploy one contract into the blockchain
func DeployContract(userAddress string) (string, *contracts.Contrato3) {
	address, tx, instance, err := contracts.DeployContrato3(getTransactOptions(), getClient(), common.HexToAddress(userAddress))

	if err != nil {
		log.Fatal(err)
	}

	//Address of the contract
	fmt.Println(address.Hex())
	//Transaction object
	fmt.Println(tx.Hash().Hex())

	return address.Hex(), instance
}

//LoadContract to load one contract from the blockchain
func LoadContract(rawAddress string) *contracts.Contrato3 {
	address := common.HexToAddress(rawAddress)
	instance, err := contracts.NewContrato3(address, getClient())
	if err != nil {
		log.Fatal(err)
	}

	return instance
}

//GetAuthentication get the autherithation value for the customer contract
func GetAuthentication(instance *contracts.Contrato3) bool {
	isAuthenticated, err := instance.IsAuthenticated(nil)
	if err != nil {
		//log.Fatal("Error with contract")
		return false
	}
	return isAuthenticated
}

//GetExpirationTime get the expiration time value for the customer contract
func GetExpirationTime(instance *contracts.Contrato3) *big.Int {
	expirationTime, err := instance.ExpirationTime(nil)
	if err != nil {
		//log.Fatal("Error with contract")
		return big.NewInt(int64(time.Now().Unix()))
	}
	return expirationTime
}

//SetAuthentication set autherithation value for the customer contract
func SetAuthentication(instance *contracts.Contrato3, value bool, expirationTime *big.Int) (*types.Transaction, error) {
	return instance.SetAuthentication(getTransactOptions(), value, expirationTime)
}

func checkExpirationTime(expTime *big.Int) bool {
	//Cmp to compare pointers of big int x < y == -1; x = y == 0; x > y == +1
	if big.NewInt(int64(time.Now().Unix())).Cmp(expTime) == -1 {
		return true
	}
	return false
}

func checkTrustClients() bool {
	return true
}

//CheckContract to check the authentication from the contract
func CheckContract(rawAddress string) bool {
	instance := LoadContract(rawAddress)
	if checkTrustClients() && checkExpirationTime(GetExpirationTime(instance)) && GetAuthentication(instance) {
		return true
	}
	return false
}
