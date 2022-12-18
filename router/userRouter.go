package router

import (
	"github.com/gorilla/mux"
	"crud/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// create
	router.HandleFunc("/api/user/add", controller.CreateUsers).Methods("POST")

	// retrieve
	router.HandleFunc("/api/user/getAll", controller.GetAllUsers).Methods("GET")

	// update
	router.HandleFunc("/api/user/update/{id}", controller.UpdateDescription).Methods("PUT")
	
	// delete
	router.HandleFunc("/api/user/delete/{id}", controller.DeleteAUser).Methods("DELETE")

	return router
}