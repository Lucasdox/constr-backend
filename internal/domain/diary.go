package domain

import "github.com/google/uuid"

type ConstructionDiary struct {
	Id uuid.UUID
	ConstructionId uuid.UUID
}
