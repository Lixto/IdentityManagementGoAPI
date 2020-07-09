package controllers

import (
	"encoding/json"
	"net/http"

	"../services"
)

// GetUsersController controller that will get all users of one table in my DB with the GetUserService
func GetUsersController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	services.GetUsersService(w)
}

//RegisterUserController controller that will introduce one user from the request on the DB with DoRegister service
func RegisterUserController(w http.ResponseWriter, r *http.Request) {
	result := services.DoRegisterService(w, r)
	if result.Ok {
		address, _ := services.DeployContract(result.Address)
		result.Contract = address
	}
	//msg, _ := json.MarshalIndent(result, "", "  ")
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
