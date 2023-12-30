package routes

import (
	"github.com/gorilla/mux"
	"github.com/nikenhpsr/task-5-pbi-btpns-niken/controllers/photocontroller"
	"github.com/nikenhpsr/task-5-pbi-btpns-niken/middleware"
)

func PhotoRoutes(router *mux.Router) {
	photo := router.PathPrefix("/photos").Subrouter()

	// Middleware
	photo.Use(middleware.Auth)

	photo.HandleFunc("", photocontroller.ListPhoto).Methods("GET")
	photo.HandleFunc("", photocontroller.CreatePhoto).Methods("POST")
	photo.HandleFunc("/{photoId}", photocontroller.ShowDetailPhoto).Methods("GET")
	photo.HandleFunc("/{photoId}", photocontroller.UpdatePhoto).Methods("PUT")
	photo.HandleFunc("/{photoId}", photocontroller.DeletePhoto).Methods("DELETE")
}
