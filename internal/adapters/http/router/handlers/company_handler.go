package handlers

import (
	"github.com/Lucasdox/constr-backend/internal/application"
	"github.com/Lucasdox/constr-backend/internal/domain"
	"net/http"
)

type CompanyHandler struct {
	service application.CompanyService
}

func (h *CompanyHandler) List(w http.ResponseWriter, r *http.Request) {

}

func (h *CompanyHandler) Insert(w http.ResponseWriter, r *http.Request) {

}

func NewCompanyHandler(r domain.CompanyRepository) *CompanyHandler {
	s := application.NewCompanyService(r)
	return &CompanyHandler{service: s}
}