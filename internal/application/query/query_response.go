package query

import "github.com/Lucasdox/constr-backend/internal/domain"

type ListConstructionsFromCompanyQueryResponse struct {
	Constructions []domain.Construction `json:"constructions"`
}
