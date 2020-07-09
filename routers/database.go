package routers

import (
	authentication "../middlewares"

	"../controllers"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//SetDatabaseRoutes set hadles and routes for database access
func SetDatabaseRoutes(router *mux.Router) *mux.Router {

	router.PathPrefix("/users").Handler(negroni.New(
		negroni.HandlerFunc(authentication.TokenAuthentication),
		negroni.HandlerFunc(authentication.RefreshToken),
		negroni.HandlerFunc(controllers.GetUsersController),
	)).Methods("GET")

	router.HandleFunc("/register", controllers.RegisterUserController).Methods("POST")
	return router
}
