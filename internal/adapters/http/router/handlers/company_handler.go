package handlers

import (
	"encoding/json"
	"github.com/Lucasdox/constr-backend/internal/application"
	"github.com/Lucasdox/constr-backend/internal/application/command"
	"github.com/Lucasdox/constr-backend/internal/domain"
	"go.uber.org/zap"
	"net/http"
)

type CompanyHandler struct {
	log *zap.Logger
	service application.CompanyService
}

func (h *CompanyHandler) ListCompanies(w http.ResponseWriter, r *http.Request) {

}

func (h *CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	defer r.Body.Close()
	var cmd command.CreateCompanyCommand

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&cmd)
	if err != nil {
		parseIncomingJsonError(err, w)
		return
	}

	idP, err := h.service.CreateCompanyAndPersist(r.Context(), cmd)

	id := *idP

	if err != nil {
		http.Error(w, "Error creating company", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", id.String())
}

func NewCompanyHandler(r domain.CompanyRepository) *CompanyHandler {
	s := application.NewCompanyService(r)
	l := zap.L().Named("company_handler")
	return &CompanyHandler{service: s, log: l}
}