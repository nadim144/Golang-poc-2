package Routers

import (
	"wipro-toyota-poc/Handlers"

	"github.com/gorilla/mux"
)

func InitializeRouter(r *mux.Router) {
	r.HandleFunc("/product", Handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/product", Handlers.GetAllProducts).Methods("GET")
	r.HandleFunc("/product/{id}", Handlers.GetProduct).Methods("GET")
	r.HandleFunc("/product/{id}", Handlers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/product/{id}", Handlers.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/product/", Handlers.DeleteAllProducts).Methods("DELETE")
}
