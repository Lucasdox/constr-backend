package query

import "github.com/google/uuid"

type ListConstructionsFromCompanyQueryProjection []ConstructionByCompanyId

type ConstructionByCompanyId struct {
	Id   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
}

type ListCompanyQueryProjection []Company

type Company struct {
	Id   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
}