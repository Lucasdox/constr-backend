package command

import (
	"github.com/google/uuid"
	"time"
)

type CreateConstructionCommand struct {
	CompanyId uuid.UUID
	ConstructionName string
	InitialDate time.Time
	DueDate time.Time
}
