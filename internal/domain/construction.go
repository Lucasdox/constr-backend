package domain

import (
	"github.com/google/uuid"
	"time"
)

type Construction struct {
	Id uuid.UUID
	CompanyId uuid.UUID
	Name string
	CreatedAt time.Time
	InitialDate time.Time
	DueDate time.Time
}

type ConstructionRepository interface {
	List(companyId uuid.UUID) ([]Construction, error)
	Save(con Construction) error
}
