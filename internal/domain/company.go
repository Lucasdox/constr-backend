package domain

import "github.com/google/uuid"

type Company struct {
	Id uuid.UUID
	Name string
	Constructions []Construction
}