package routes

import (
	"github.com/gorilla/mux"
	"github.com/nikenhpsr/task-5-pbi-btpns-niken/controllers/usercontroller"
	"github.com/nikenhpsr/task-5-pbi-btpns-niken/middleware"
)

func UserRoutes(router *mux.Router) {
	user := router.PathPrefix("/users").Subrouter()

	// Middleware
	user.Use(middleware.Auth)

	user.HandleFunc("", usercontroller.ListUser).Methods("GET")
	user.HandleFunc("/{userId}", usercontroller.UpdateUser).Methods("PUT")
	user.HandleFunc("/{userId}", usercontroller.DeleteUser).Methods("DELETE")
}
