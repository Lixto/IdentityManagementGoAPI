package services

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	core "../core"
	models "../models"
	"../parameters"

	_ "github.com/denisenkom/go-mssqldb"
	"golang.org/x/crypto/argon2"
)

// GetUsersService one service that gets all users from the DB
func GetUsersService(w http.ResponseWriter) {

	// Create connection pool
	//Open is used to create a database handle
	conn, err := sql.Open("sqlserver", parameters.ConnString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	// Check if database is alive and conection
	ctx := context.Background()
	err = conn.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	// Execute query
	tsql := fmt.Sprintf("SELECT * FROM [dbo].[Users];")

	rows, err := conn.QueryContext(ctx, tsql)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer rows.Close()

	var users []models.User
	// Iterate through the result set.
	for rows.Next() {
		var email, password, salt string

		// Get values from row.
		err := rows.Scan(&email, &password, &salt)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		var user models.User
		user.Email = email
		user.Password = password
		user.Salt = salt
		users = append(users, user)

		fmt.Printf("Email: %s, Password: %s\n", email, password)
	}

	json.NewEncoder(w).Encode(users)
}

//DoRegisterService one service that insert one user on the DB
func DoRegisterService(w http.ResponseWriter, r *http.Request) models.ResponseRegister {
	var user models.User //_ = json.NewDecoder(r.Body).Decode(&user)
	body := json.NewDecoder(r.Body)
	body.DisallowUnknownFields()
	err := body.Decode(&user)
	customPassword := ""
	var res models.ResponseRegister

	if err != nil || user.Email == "" {
		res = models.ResponseRegister{Ok: false, Msg: "Error on the body"}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		var pass string

		if user.Password != "" {
			pass = user.Password
		} else {
			pass = core.GeneratePassRand(user.LongPassword, user.MayusPassword, user.SpecialPassword, user.NumbersPassword)
			if pass == "" {
				res = models.ResponseRegister{Ok: false, Msg: "Password must be at least length 8"}
				w.WriteHeader(http.StatusBadRequest)
			} else {
				customPassword = pass
			}
		}

		// Create connection pool
		//Open is used to create a database handle
		conn, err := sql.Open("sqlserver", parameters.ConnString)
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
		defer conn.Close()

		// Check if database is alive and conection
		ctx := context.Background()
		err = conn.PingContext(ctx)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("Connected!\n")

		// Prepare query
		tsql := fmt.Sprintf("INSERT INTO [dbo].[Users] VALUES (@Email, @Password, @Salt, @Address)")

		//Check User Address
		var privateKey, publicKey, address string
		if user.Address == "" {
			privateKey, publicKey, address = core.GeneratePublicAddress()
			user.Address = address
		}

		//Generate salt
		salt := GenerateSaltService()
		unencodedEmail := argon2.IDKey([]byte(user.Email), parameters.SaltEmails, parameters.Time, parameters.Memory, parameters.Threads, parameters.KeyLen)
		unencodedPassword := argon2.IDKey([]byte(pass), salt, parameters.Time, parameters.Memory, parameters.Threads, parameters.KeyLen)
		unencodedAddress := argon2.IDKey([]byte(user.Address), salt, parameters.Time, parameters.Memory, parameters.Threads, parameters.KeyLen)
		encodedEmail := base64.StdEncoding.EncodeToString(unencodedEmail)
		encodedPassword := base64.StdEncoding.EncodeToString(unencodedPassword)
		encodedAddress := base64.StdEncoding.EncodeToString(unencodedAddress)
		encodedSalt := base64.StdEncoding.EncodeToString(salt)

		_, err = conn.ExecContext(
			ctx,
			tsql,
			sql.Named("Email", encodedEmail),
			sql.Named("Password", encodedPassword),
			sql.Named("Salt", encodedSalt),
			sql.Named("Address", encodedAddress))

		if err != nil {
			res = models.ResponseRegister{Ok: false, Msg: "Error registering user in the DB, maybe the email is duplicated"}
		} else {

			tsql := fmt.Sprintf("INSERT INTO [dbo].[UserAuthentication] VALUES (@Email, 0)")

			_, err = conn.ExecContext(
				ctx,
				tsql,
				sql.Named("Email", encodedEmail))

			if err != nil {
				//log.Fatal(err.Error())
				res = models.ResponseRegister{Ok: false, Msg: "Error registering user in the DB try it again"}
			} else {
				if privateKey != "" && publicKey != "" && address != "" {
					res.PrivateKey = privateKey
					res.PublicKey = publicKey
					res.Address = address
				} else {
					res.Address = user.Address
				}

				if customPassword != "" {
					res.Password = customPassword
				}
				res.Ok = true
				res.Msg = "User has been registed correctly"
			}
		}
	}
	return res
}

// GetTrustClientsService one service that gets all clients from the DB
func GetTrustClientsService() []models.TrustClients {

	// Create connection pool
	//Open is used to create a database handle
	conn, err := sql.Open("sqlserver", parameters.ConnString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	// Check if database is alive and conection
	ctx := context.Background()
	err = conn.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	// Execute query
	tsql := fmt.Sprintf("SELECT * FROM [dbo].[TrustClients];")

	rows, err := conn.QueryContext(ctx, tsql)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	defer rows.Close()

	var clients []models.TrustClients
	// Iterate through the result set.
	for rows.Next() {
		var clientAddress string
		var trust int

		// Get values from row.
		err := rows.Scan(&clientAddress, &trust)
		if err != nil {
			log.Fatal(err.Error())
			return nil
		}

		var client models.TrustClients
		client.Client = clientAddress
		client.Trust = trust
		clients = append(clients, client)
	}
	return clients
}
