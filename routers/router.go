package routers

//This file joins all the routes that we have on separate files

import (
	"github.com/gorilla/mux"
)

//InitRoutes create de router and set all routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetDatabaseRoutes(router)
	router = SetAuthenticationRoutes(router)
	return router
}
