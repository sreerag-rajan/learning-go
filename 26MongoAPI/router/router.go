package router

import (
	"github.com/gorilla/mux"
	"github.coms/sreerag-rajan/mongoapi/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controller.DeleteOneMovies).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
