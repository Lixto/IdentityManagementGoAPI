package models

//This file contains one approach for the Cliams, you can use this or the ones ont he jwt-go.
//On this API we are using the ones on the jwt-go

import jwt "github.com/dgrijalva/jwt-go"

//Claims for the token
type Claims struct {
	user string `json:"user"`
	jwt.StandardClaims
}
