package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nikenhpsr/task-5-pbi-btpns-niken/config"
	"github.com/nikenhpsr/task-5-pbi-btpns-niken/routes"
)

func main() {
	config.LoadConfig()
	config.ConnectDatabase()

	router := mux.NewRouter()
	routes.RoutesIndex(router)

	log.Println("[APP] Server is listening to port", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), router)
}
