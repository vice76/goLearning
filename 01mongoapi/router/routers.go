package router

import (
	"github.com/gorilla/mux"
	"github.com/vice76/mongoapi/controller"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkedAsWatch).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteAllMovie", controller.DeleteAllMovie).Methods("DELETE")

	return router
}
