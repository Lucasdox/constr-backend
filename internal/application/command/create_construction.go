package command

import (
	"github.com/google/uuid"
	"time"
)

type CreateConstructionCommand struct {
	CompanyId        uuid.UUID `json:"company_id,omitempty"`
	ConstructionName string    `json:"construction_name,omitempty"`
	InitialDate      time.Time `json:"initial_date"`
	DueDate          time.Time `json:"due_date"`
}
