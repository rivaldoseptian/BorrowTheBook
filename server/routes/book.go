package routes

import (
	"net/http"
	"server/controllers"
	"server/middleware"

	"github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router) {
	router := r.PathPrefix("/book").Subrouter()
	router.Use(middleware.Auth)
	router.HandleFunc("", controllers.GetBook).Methods("GET")
	router.Handle("/{id}", middleware.AdminAuth(http.HandlerFunc(controllers.ReturnBook))).Methods("DELETE")

}
