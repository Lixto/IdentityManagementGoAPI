package core

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

const kPath = "./keystores"

func generatePrivateKey() *ecdsa.PrivateKey {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	return privateKey
}

func generatePublicKey(privateKey *ecdsa.PrivateKey) *ecdsa.PublicKey {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	return publicKeyECDSA
}

//GeneratePublicAddress generate the public address which is what you're used to seeing
func GeneratePublicAddress() (string, string, string) {
	privateKey := generatePrivateKey()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyString := hexutil.Encode(privateKeyBytes)[2:]

	publicKey := generatePublicKey(privateKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	publicKeyString := hexutil.Encode(publicKeyBytes)[4:]

	address := crypto.PubkeyToAddress(*publicKey).Hex()

	return privateKeyString, publicKeyString, address
}

//CreateKeyStore function creates a keystore file
func CreateKeyStore(password string) string {
	ks := keystore.NewKeyStore(kPath, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	return account.Address.Hex()
}

//ImportKeyStore function to import jeystore file
func ImportKeyStore(file []byte, password string) string {
	ks := keystore.NewKeyStore(kPath, keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := ks.Import(file, password, password)
	if err != nil {
		log.Fatal(err)
	}
	return account.Address.Hex()
}
