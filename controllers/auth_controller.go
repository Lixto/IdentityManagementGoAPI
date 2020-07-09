package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "../models"
	"../services"
)

//LoginController is the function that will be call by the API when the user's send one login request
//This function uses the Login service from this API
func LoginController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	//Get de user you want to test
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	responseStatus, token := services.LoginService(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
	//json.NewEncoder(w).Encode(res)
}

//LoginMultiController is the same as normal login, but with out the token
func LoginMultiController(w http.ResponseWriter, r *http.Request) {
	//Get de user you want to test
	var user models.User
	var s string
	//_ = json.NewDecoder(r.Body).Decode(&user)
	_ = json.Unmarshal([]byte(s), &user)
	fmt.Println(s)
	//decoder := json.NewDecoder(r.Body)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	responseStatus := services.LoginMultiService(user)
	w.WriteHeader(responseStatus)
}

//InsertMultiController is the same as normal login, but with out the token
func InsertMultiController(w http.ResponseWriter, r *http.Request) {
	//Get de user you want to test
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	responseStatus := services.InsertMultiService(user)
	w.WriteHeader(responseStatus)
}
