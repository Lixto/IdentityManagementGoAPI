package core

//This file generate de token

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"time"

	"../parameters"

	jwt "github.com/dgrijalva/jwt-go"
)

//JWTAuthenticationBackend struct to save keys
type JWTAuthenticationBackend struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

var authBackendInstance *JWTAuthenticationBackend

//InitJWTAuthenticationBackend set keys if they aren't
func InitJWTAuthenticationBackend() *JWTAuthenticationBackend {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{
			privateKey: getPrivateKey(),
			PublicKey:  getPublicKey(),
		}
	}
	return authBackendInstance
}

//GenerateToken funtion generates one json token, argument is the user email, better change it to UUID?
func (backend *JWTAuthenticationBackend) GenerateToken(userEmail string) (string, int64, error) {
	token := jwt.New(jwt.SigningMethodRS512) //SigningMethodES256

	var expir = time.Now().Add(time.Minute * 60).Unix()

	/*claims := models.Claims{
		userEmail,
		jwt.StandardClaims{
			ExpiresAt: expir,
		},
	}

	token.Claims = claims*/

	token.Claims = jwt.MapClaims{
		"exp": expir,
		"nbf": time.Now().Unix(),
		"foo": userEmail,
	}

	tokenString, err := token.SignedString(backend.privateKey)
	if err != nil {
		return "", time.Now().Unix(), err
	}

	//Return token and the expiration time
	//old, now we are going to sentd unix time and tranform in LoginService
	//return tokenString, fmt.Sprintf("%q", time.Unix(expir, 0)), nil
	return tokenString, expir, nil
}

//getPrivateKey get the rsa private key
func getPrivateKey() *rsa.PrivateKey {
	privateKeyFile, err := os.Open(parameters.PrivateKeyPath)

	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	pembytes := make([]byte, pemfileinfo.Size())

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	if err != nil {
		panic(err)
	}

	data, _ := pem.Decode([]byte(pembytes))
	privateKeyFile.Close()

	privateKey, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	return privateKey
}

//getPublicKey get the rsa public key
func getPublicKey() *rsa.PublicKey {
	publicKeyFile, err := os.Open(parameters.PublicKeyPath)

	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	pembytes := make([]byte, pemfileinfo.Size())

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)

	if err != nil {
		panic(err)
	}

	data, _ := pem.Decode([]byte(pembytes))

	publicKeyFile.Close()

	publicKey, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	rsaPublicKey, errBool := publicKey.(*rsa.PublicKey)

	if !errBool {
		panic(err)
	}

	return rsaPublicKey
}
