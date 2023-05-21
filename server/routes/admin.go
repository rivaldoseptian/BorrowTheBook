package routes

import (
	"server/controllers"

	"github.com/gorilla/mux"
)

func AdminRoutes(r *mux.Router) {
	router := r.PathPrefix("/admin").Subrouter()
	router.HandleFunc("/registeradmin", controllers.RegisterAdmin).Methods("POST")
}
