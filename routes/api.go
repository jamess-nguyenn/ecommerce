package routes

import (
	"ecommerce/database/connection"
	"ecommerce/http/handlers"
	"ecommerce/http/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func BuildApiRouter(db *connection.MysqlDatabase) *mux.Router {
	router := mux.NewRouter()

	api := router.PathPrefix("/api/v1").Subrouter()

	// storefront management
	apiStorefront := api.PathPrefix("/storefronts").Subrouter()

	apiStorefront.HandleFunc("", handlers.CreateStorefront(db)).Methods(http.MethodPost).Name("CreateStorefront")
	apiStorefront.HandleFunc("", handlers.GetListStorefront).Methods(http.MethodGet).Name("GetListStorefront")
	apiStorefront.HandleFunc("/{id:[0-9]+}", handlers.ShowStorefront(db)).Methods(http.MethodGet).Name("ShowStorefront")
	apiStorefront.HandleFunc("/{id:[0-9]+}", handlers.UpdateStorefrontById(db)).Methods(http.MethodPut).Name("UpdateStorefrontById")
	apiStorefront.HandleFunc("/{id:[0-9]+}", handlers.DeleteStorefrontById).Methods(http.MethodDelete).Name("DeleteStorefrontById")

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
