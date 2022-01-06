package router

import "github.com/gorilla/mux"

func Router() *mux.Router {
	r := mux.NewRouter()

	c := r.PathPrefix("/constructions").Subrouter()
	c.Path("/").Methods("GET")
}
