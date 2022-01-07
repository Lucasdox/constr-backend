package domain

import (
	"github.com/google/uuid"
	"time"
)

type Construction struct {
	Id uuid.UUID
	CompanyId uuid.UUID
	Name string
	InitialDate time.Time
	DueDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ConstructionRepository interface {
	ListByCompanyId(companyId uuid.UUID) ([]Construction, error)
	Save(con Construction) error
}

func NewConstruction(companyId uuid.UUID, name string, initDate time.Time, dueDate time.Time) Construction {
	return Construction{
		Id:          uuid.New(),
		CompanyId:   companyId,
		Name:        name,
		InitialDate: initDate,
		DueDate:     dueDate,
	}
}
