package middlewares

//This file is one middelware to check the token on every single request

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"

	core "../core"
	"../models"
	services "../services"

	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
)

//BlockchainAuthentication check is user is autehnticate inside the  blockchain
func BlockchainAuthentication(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	w.Header().Set("Content-Type", "application/json")
	if services.CheckContract(user.Contract) {
		responseStatus, token := services.TokenResponse(user)
		w.WriteHeader(responseStatus)
		w.Write(token)
	} else {
		responseStatus, token := services.LoginService(user)
		w.WriteHeader(responseStatus)
		w.Write(token)
	}
}

// TokenAuthentication middleware for vaidation token
func TokenAuthentication(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	authBackend := core.InitJWTAuthenticationBackend()

	token, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"]) // For the moment we don't have this
		} else {
			return authBackend.PublicKey, nil
		}
	})

	if err == nil && token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		if claims.VerifyExpiresAt(time.Now().Unix(), false) != false {
			//If the token haven't been expired we continue with the request, this case we refresh teh token
			next(rw, req)
		} else {
			//Token is expired
			ResetMultiFactor(rw, req, next)
			rw.WriteHeader(http.StatusUnauthorized)
		}
	} else {
		ResetMultiFactor(rw, req, next)
		rw.WriteHeader(http.StatusUnauthorized)
	}
}

//RefreshToken will see if the token is over expiration time
func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	authBackend := core.InitJWTAuthenticationBackend()

	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		} else {
			return authBackend.PublicKey, nil
		}
	})

	if err == nil && token.Valid {
		/*claims := token.Claims.(jwt.MapClaims)
		aux := claims["ExpiresAt"]
		aux2 := claims["nbf"]*/
		var user models.User
		_ = json.NewDecoder(r.Body).Decode(&user)
		w.Header().Set("Content-Type", "application/json")
		// If the token is valid, we use the service to refresh it
		w.Write(services.RefreshTokenService(user))
		next(w, r)
	} else {
		ResetMultiFactor(w, r, next)
	}
}

//ResetMultiFactor will set all the multifactor value to false
func ResetMultiFactor(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.Fingerprint = false
	res := services.InsertMultiService(user)
	if res == 200 {
		instance := services.LoadContract(user.Contract)
		services.SetAuthentication(instance, false, big.NewInt(int64(time.Now().Unix())))
	}
}
