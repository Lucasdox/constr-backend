package router

import (
	"github.com/Lucasdox/constr-backend/internal/adapters/database"
	"github.com/Lucasdox/constr-backend/internal/adapters/database/pg"
	"github.com/Lucasdox/constr-backend/internal/adapters/http/router/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func Router(db *database.DBImpl) *mux.Router {
	constructionRepository := pg.NewConstructionRepository(db)
	companyRepository := pg.NewCompanyRepository(db)

	r := mux.NewRouter().StrictSlash(true)

	conHandler := handlers.NewConstructionHandler(constructionRepository)
	cHandler := handlers.NewCompanyHandler(companyRepository)

	c := r.PathPrefix("/companies").Subrouter()
	c.Path("").Methods(http.MethodGet).HandlerFunc(cHandler.ListCompanies)
	c.Path("").Methods(http.MethodPost).HandlerFunc(cHandler.CreateCompany)

	// /companies/{id}/constructions
	c.Path("/{id}/constructions").Methods(http.MethodGet).HandlerFunc(conHandler.ListConstruction)

	return r
}
