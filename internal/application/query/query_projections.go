package query

import (
	"github.com/google/uuid"
	"time"
)

type ListConstructionsFromCompanyQueryProjection []ConstructionByCompanyId

type ConstructionByCompanyId struct {
	Id          uuid.UUID `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	InitialDate time.Time `json:"initial_date"`
	DueDate     time.Time `json:"due_date"`
}

type ListCompanyQueryProjection []Company

type Company struct {
	Id   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
}