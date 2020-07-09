package services

import (
	"context"
	cont "context"
	"crypto/subtle"
	sql "database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	core "../core"
	models "../models"
	"../parameters"

	_ "github.com/denisenkom/go-mssqldb" //To access to db
	"golang.org/x/crypto/argon2"
)

//LoginService gets if login is true o false
func LoginService(user models.User) (int, []byte) {

	conn, err := sql.Open("sqlserver", parameters.ConnString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	ctx := context.Background()
	err = conn.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	unencodedEmail := argon2.IDKey([]byte(user.Email), parameters.SaltEmails, parameters.Time, parameters.Memory, parameters.Threads, parameters.KeyLen)
	encodedEmail := base64.StdEncoding.EncodeToString(unencodedEmail)

	//We will check for the biometric factor to be set on the BD
	var auth bool
	auth, err = CheckMultiFactor(encodedEmail, conn, ctx)
	if err != nil {
		return http.StatusInternalServerError, []byte("")
	}

	//Mis dudas, de momento se queda pero no
	if !auth {
		timeOld := time.Now().Unix()

		//120 are 2 minutes in seconds, so you will try to get the value while auth = falte for 2 minutes, esle you will get and error
		for !auth || timeOld < timeOld+120 {
			auth, err = CheckMultiFactor(encodedEmail, conn, ctx)
			if err != nil {
				return http.StatusInternalServerError, []byte("")
			}

			if auth {
				break
			}
		}
	}

	if !auth {
		return http.StatusInternalServerError, []byte("You must use auth multi factor")
	}

	// Execute query
	tsql := fmt.Sprintf("SELECT Password, Salt, Address FROM [dbo].[Users] WHERE Email = @Email;")

	rows, err := conn.QueryContext(ctx, tsql, sql.Named("Email", encodedEmail))
	if err != nil {
		//log.Fatal(err.Error())
		return http.StatusInternalServerError, []byte("")
	}
	defer rows.Close()

	var password, salt, address string
	for rows.Next() {
		err := rows.Scan(&password, &salt, &address)
		if err != nil {
			//log.Fatal(err.Error())
			return http.StatusInternalServerError, []byte("")
		}
	}

	decodePassword, _ := base64.StdEncoding.DecodeString(password)
	decodeSalt, _ := base64.StdEncoding.DecodeString(salt)
	var auxPassword = argon2.IDKey([]byte(user.Password), decodeSalt, parameters.Time, parameters.Memory, parameters.Threads, parameters.KeyLen)

	//Compare if both password have the same bytes
	if subtle.ConstantTimeCompare(decodePassword, auxPassword) == 1 {
		decodeAddress, _ := base64.StdEncoding.DecodeString(address)
		var auxAddress = argon2.IDKey([]byte(user.Address), decodeSalt, parameters.Time, parameters.Memory, parameters.Threads, parameters.KeyLen)
		if subtle.ConstantTimeCompare(decodeAddress, auxAddress) == 1 {
			return TokenResponse(user)
		}
	}

	return http.StatusInternalServerError, []byte("")
}

//TokenResponse function to generate token response when you are log in
func TokenResponse(user models.User) (int, []byte) {
	//Aquí debe generar el tocken
	authBackend := core.InitJWTAuthenticationBackend()
	token, expir, err := authBackend.GenerateToken(user.Email)
	if err != nil {
		return http.StatusInternalServerError, []byte("")
	}

	// Old method, now we are going to make de transformation here
	//response, _ := json.Marshal(models.TokenAuthentication{token, expir})
	response, _ := json.Marshal(models.TokenAuthentication{
		Token:                 token,
		ExpirationTimeForUser: fmt.Sprintf("%q", time.Unix(expir, 0)),
		ExpirationTime:        expir,
	})

	instance := LoadContract(user.Contract)
	tran, err := SetAuthentication(instance, true, big.NewInt(int64(expir)))
	if err != nil {
		return http.StatusInternalServerError, []byte("")
	}
	fmt.Println(tran)

	//Only for debug
	fmt.Println(GetExpirationTime(instance))
	fmt.Println(GetAuthentication(instance))

	if GetAuthentication(instance) {
		return http.StatusOK, response
	} else {
		return http.StatusInternalServerError, []byte("")
	}
}

//LoginMultiService gets if login is true o false
func LoginMultiService(user models.User) int {

	conn, err := sql.Open("sqlserver", parameters.ConnString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	ctx := context.Background()
	err = conn.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	// Execute query
	tsql := fmt.Sprintf("SELECT Password, Salt FROM [dbo].Users WHERE Email = @Email;")

	unencodedEmail := argon2.IDKey([]byte(user.Email), parameters.SaltEmails, parameters.Time, parameters.Memory, parameters.Threads, parameters.KeyLen)
	encodedEmail := base64.StdEncoding.EncodeToString(unencodedEmail)

	rows, err := conn.QueryContext(ctx, tsql, sql.Named("Email", encodedEmail))
	if err != nil {
		//log.Fatal(err.Error())
		return http.StatusInternalServerError
	}
	defer rows.Close()

	var password, salt string
	for rows.Next() {
		err := rows.Scan(&password, &salt)
		if err != nil {
			//log.Fatal(err.Error())
			return http.StatusInternalServerError
		}
	}

	decodePassword, _ := base64.StdEncoding.DecodeString(password)
	decodeSalt, _ := base64.StdEncoding.DecodeString(salt)
	var auxPassword = argon2.IDKey([]byte(user.Password), decodeSalt, parameters.Time, parameters.Memory, parameters.Threads, parameters.KeyLen)

	//Compare if both password have the same bytes
	if subtle.ConstantTimeCompare(decodePassword, auxPassword) == 1 {
		return http.StatusOK
	}

	return http.StatusInternalServerError
}

//RefreshTokenService one service that uses the Login service to refresh the token one hour
func RefreshTokenService(user models.User) []byte {
	status, response := LoginService(user)
	if status == http.StatusOK {
		return response
	}
	return []byte("")
}

//CheckMultiFactor check if the user has one second auth factor and it's value
func CheckMultiFactor(email string, conn *sql.DB, ctx cont.Context) (bool, error) {

	// Execute query
	tsql := fmt.Sprintf("SELECT Fingerprint FROM [dbo].[UserAuthentication] WHERE Email = @Email;")

	rows, err := conn.QueryContext(ctx, tsql, sql.Named("Email", email))

	if err != nil {
		defer rows.Close()
		return false, err
	}
	defer rows.Close()

	var fingerprint bool
	for rows.Next() {
		err := rows.Scan(&fingerprint)
		if err != nil {
			return false, err
		}
	}
	return fingerprint, nil
}

//InsertMultiService one service that insert one multifactor on the DB
func InsertMultiService(auth models.User) int {

	conn, err := sql.Open("sqlserver", parameters.ConnString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	ctx := context.Background()
	err = conn.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	// Execute query
	tsql := fmt.Sprintf("SELECT Password, Salt FROM [dbo].[Users] WHERE Email = @Email;")

	unencodedEmail := argon2.IDKey([]byte(auth.Email), parameters.SaltEmails, parameters.Time, parameters.Memory, parameters.Threads, parameters.KeyLen)
	encodedEmail := base64.StdEncoding.EncodeToString(unencodedEmail)

	rows, err := conn.QueryContext(ctx, tsql, sql.Named("Email", encodedEmail))
	if err != nil {
		//log.Fatal(err.Error())
		return http.StatusInternalServerError
	}
	defer rows.Close()

	var password, salt string
	for rows.Next() {
		err := rows.Scan(&password, &salt)
		if err != nil {
			//log.Fatal(err.Error())
			return http.StatusInternalServerError
		}
	}

	decodePassword, _ := base64.StdEncoding.DecodeString(password)
	decodeSalt, _ := base64.StdEncoding.DecodeString(salt)
	var auxPassword = argon2.IDKey([]byte(auth.Password), decodeSalt, parameters.Time, parameters.Memory, parameters.Threads, parameters.KeyLen)

	//Compare if both password have the same bytes
	if subtle.ConstantTimeCompare(decodePassword, auxPassword) == 1 {

		// Prepare query
		tsql := fmt.Sprintf("UPDATE [dbo].[UserAuthentication] SET Fingerprint = @Fingerprint WHERE Email = @Email")

		_, err = conn.ExecContext(
			ctx,
			tsql,
			sql.Named("Email", encodedEmail),
			sql.Named("Fingerprint", auth.Fingerprint))

		if err != nil {
			return http.StatusInternalServerError
		}
		return http.StatusOK
	}
	return http.StatusInternalServerError
}
