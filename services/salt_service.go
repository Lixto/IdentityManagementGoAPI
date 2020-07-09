package services

import (
	"crypto/rand"

	"../parameters"
)

//GenerateSaltService one service that generate one rand salt
func GenerateSaltService() []byte {
	salt := make([]byte, parameters.SaltBytes)
	if _, err := rand.Read(salt); err != nil {
		return nil
	}
	return salt
}
