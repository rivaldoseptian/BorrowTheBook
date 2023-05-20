package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gorilla/mux"
)

func BorrowRoutes(r *mux.Router) {
	router := r.PathPrefix("/borrow").Subrouter()
	router.Use(middleware.Auth)
	router.HandleFunc("/{id}", controllers.BorrowBook).Methods("POST")
	router.HandleFunc("/me", controllers.ListBorrow).Methods("GET")
}
