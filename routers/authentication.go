package routers

import (
	"../controllers"
	authentication "../middlewares"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//SetAuthenticationRoutes set hadles and routes for authentication
func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.PathPrefix("/login").Handler(negroni.New(
		negroni.HandlerFunc(authentication.BlockchainAuthentication),
		negroni.HandlerFunc(controllers.LoginController),
	)).Methods("POST")
	router.PathPrefix("/logout").Handler(negroni.New(negroni.HandlerFunc(authentication.ResetMultiFactor))).Methods("POST")
	router.HandleFunc("/loginMultiFactor", controllers.LoginMultiController).Methods("POST")
	router.HandleFunc("/insertMultiFactor", controllers.InsertMultiController).Methods("POST")
	router.HandleFunc("/checkToken", controllers.InsertMultiController).Handler(negroni.New(negroni.HandlerFunc(authentication.TokenAuthentication))).Methods("GET")
	return router
}
