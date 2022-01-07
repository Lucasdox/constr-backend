package handlers

import (
	"github.com/Lucasdox/constr-backend/internal/application"
	"github.com/Lucasdox/constr-backend/internal/domain"
	"net/http"
)

type ConstructionHandler struct {
	Service application.ConstructionService
}

func (h *ConstructionHandler) ListConstruction(w http.ResponseWriter, r *http.Request) {

}

func NewConstructionHandler(r domain.ConstructionRepository) *ConstructionHandler {
	s := application.NewConstructionService(r)
	return &ConstructionHandler{Service: s}
}
