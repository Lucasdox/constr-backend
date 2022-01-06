package router

import (
	"github.com/Lucasdox/constr-backend/internal/adapters/http/router/handlers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	ch := handlers.NewConstructionHandler()

	c := r.PathPrefix("/constructions").Subrouter()
	c.Path("/").Methods("GET").HandlerFunc(ch.ListConstruction)

	return r
}
