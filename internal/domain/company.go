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
	Insert(id uuid.UUID, name string) error
}

func NewCompany(name string) Company{
	return Company{
		Id:            uuid.New(),
		Name:          name,
	}
}
