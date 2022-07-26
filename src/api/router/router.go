package router

import (
	"github.com/gorilla/mux"
	"gitlab.com/raspberry.tech/wireguard-manager-and-api/src/api/middleware"
)

func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()                   //Router for routes
	router.Use(middleware.EnableCORSMiddleware) //need to allow CORS and OPTIONS
	router.Use(middleware.AuthMiddleware)

	manager := router.PathPrefix("/manager").Subrouter() //main subrouter

	keys := manager.PathPrefix("/key").Subrouter() //specific subrouter
	keys.HandleFunc("", getKeys).Methods("GET")
	keys.HandleFunc("", keyCreate).Methods("POST")          //post route for adding keys
	keys.HandleFunc("", keyRemove).Methods("DELETE")        //delete route for removing keys
	keys.HandleFunc("/enable", keyEnable).Methods("POST")   //post route for enabling key
	keys.HandleFunc("/disable", keyDisable).Methods("POST") //post route for disabling key

	subscriptions := manager.PathPrefix("/subscription").Subrouter() //specific subrouter
	subscriptions.HandleFunc("/all", getSubscriptions).Methods("GET")
	subscriptions.HandleFunc("/edit", keySetSubscription).Methods("POST") //for editing subscription
	subscriptions.HandleFunc("", getKeySub).Methods("POST")

	// router.MethodNotAllowedHandler = http.HandlerFunc(middleware.SetCORSHeaders) //if method is not found allow OPTIONS
	return router
}
