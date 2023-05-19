package routes

import (
	"server/controllers"

	"github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router) {
	router := r.PathPrefix("/book").Subrouter()
	router.HandleFunc("", controllers.GetBook).Methods("GET")
	router.HandleFunc("/{id}", controllers.ReturnBook).Methods("DELETE")
}
