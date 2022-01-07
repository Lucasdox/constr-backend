package domain

import (
	"github.com/google/uuid"
)

type Company struct {
	Id uuid.UUID
	Name string
	Constructions []Construction
}

type CompanyRepository interface {
	Save(company Company) error
}

func NewCompany(name string) Company{
	return Company{
		Id:            uuid.New(),
		Name:          name,
	}
}
