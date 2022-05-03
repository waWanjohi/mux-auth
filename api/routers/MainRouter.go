package routers

import (
	"github.com/gorilla/mux"
	"github.com/waWanjohi/mux-auth/api/handlers"
	"github.com/waWanjohi/mux-auth/api/middlewares"
)

func MainRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	router.HandleFunc("/login", middlewares.SetMiddlewareJSON(handlers.LoginHandler)).Methods("POST")

	protected := router.PathPrefix("/").Subrouter()
	// protected.Use(middlewares.AuthenticationMiddleware)
	protected.HandleFunc("/resource", middlewares.SetMiddlewareAuthentication(handlers.ProtectedHandler)).Methods("GET", "POST")

	return router
}
