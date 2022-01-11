package handlers

import (
	"encoding/json"
	"github.com/Lucasdox/constr-backend/internal/application"
	"github.com/Lucasdox/constr-backend/internal/application/command"
	"github.com/Lucasdox/constr-backend/internal/domain"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type ConstructionHandler struct {
	service application.ConstructionService
}

func (h *ConstructionHandler) ListConstruction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	companyId, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	q, err := h.service.ListByCompanyId(r.Context(), companyId)
	if err != nil {
		http.Error(w, "error searching constructions", http.StatusInternalServerError)
	}

	res, _ := json.Marshal(q)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *ConstructionHandler) CreateConstruction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	companyId, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var cmd command.CreateConstructionCommand
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err = dec.Decode(&cmd)
	if err != nil {
		parseIncomingJsonError(err, w)
		return
	}
	cmd.CompanyId = companyId

	idP, err := h.service.Create(r.Context(), cmd)

	ctrId := *idP

	if err != nil {
		http.Error(w, "Error creating construction", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", ctrId.String())
}

func NewConstructionHandler(r domain.ConstructionRepository) *ConstructionHandler {
	s := application.NewConstructionService(r)
	return &ConstructionHandler{service: s}
}
