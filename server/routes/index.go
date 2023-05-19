package routes

import "github.com/gorilla/mux"

func IndexRouter(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()
	AuthRoutes(api)
	BorrowRoutes(api)
	BookRoutes(api)
}
