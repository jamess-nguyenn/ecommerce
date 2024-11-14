package routes

import (
	"ecommerce/http/handlers"
	"ecommerce/http/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func BuildApiRouter() *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()

	// user management
	apiUser := api.PathPrefix("/users").Subrouter()

	apiUser.HandleFunc("", handlers.CreateUser).Methods(http.MethodPost).Name("CreateUser")
	apiUser.HandleFunc("", handlers.GetListUser).Methods(http.MethodGet).Name("GetListUser")
	apiUser.HandleFunc("/{id:[0-9]+}", handlers.ShowUser).Methods(http.MethodGet).Name("ShowUser")
	apiUser.HandleFunc("/{id:[0-9]+}", handlers.UpdateUserById).Methods(http.MethodPut).Name("UpdateUserById")
	apiUser.HandleFunc("/{id:[0-9]+}", handlers.DeleteUserById).Methods(http.MethodDelete).Name("DeleteUserById")

	// product management
	apiProduct := api.PathPrefix("/products").Subrouter()

	apiProduct.HandleFunc("", handlers.CreateProduct).Methods(http.MethodPost).Name("CreateProduct")
	apiProduct.HandleFunc("", handlers.GetListProduct).Methods(http.MethodGet).Name("GetListProduct")
	apiProduct.HandleFunc("/{id:[0-9]+}", handlers.ShowProduct).Methods(http.MethodGet).Name("ShowProduct")
	apiProduct.HandleFunc("/{id:[0-9]+}", handlers.UpdateProductById).Methods(http.MethodPut).Name("UpdateProductById")
	apiProduct.HandleFunc("/{id:[0-9]+}", handlers.DeleteProductById).Methods(http.MethodDelete).Name("DeleteProductById")

	// use middleware
	router.Use(middleware.LogRequest, middleware.Auth, mux.CORSMethodMiddleware(router))

	return router
}
