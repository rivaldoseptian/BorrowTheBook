package main

import (
	"fmt"
	"server/config"
	"server/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	config.LoadConfig()
	config.ConnectDB()

	r := mux.NewRouter()
	routes.IndexRouter(r)

	log.Println("Server Running on port", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)
}
