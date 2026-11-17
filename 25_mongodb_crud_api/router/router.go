// Package router configures the REST routing table.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Hellthefox808/GoLang-Practice-code/25_mongodb_crud_api/controller"
)

// SetupRouter registers endpoints with their respective HTTP verbs.
func SetupRouter(ctrl *controller.MovieController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/movies", ctrl.GetAllMovies).Methods(http.MethodGet)
	r.HandleFunc("/api/movies", ctrl.CreateMovie).Methods(http.MethodPost)
	r.HandleFunc("/api/movies/{id}", ctrl.MarkAsWatched).Methods(http.MethodPut)
	r.HandleFunc("/api/movies/{id}", ctrl.DeleteMovie).Methods(http.MethodDelete)
	r.HandleFunc("/api/movies", ctrl.DeleteAllMovies).Methods(http.MethodDelete)

	return r
}
