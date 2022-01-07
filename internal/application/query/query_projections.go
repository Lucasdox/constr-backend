package query

import (
	"github.com/google/uuid"
	"time"
)

type ListConstructionsFromCompanyQueryProjection []ConstructionByCompanyId

type ConstructionByCompanyId struct {
	Id   uuid.UUID
	Name string
	InitialDate time.Time
	DueDate time.Time
}

type ListCompanyQueryProjection []Company

type Company struct {
	Id   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
}