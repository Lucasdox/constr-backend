package pg

import (
	"github.com/Lucasdox/constr-backend/internal/adapters/database"
	"github.com/Lucasdox/constr-backend/internal/domain"
	"github.com/google/uuid"
	"time"
)

type ConstructionRepositoryImpl struct {
	db database.DBImpl
}

func (c ConstructionRepositoryImpl) List(companyId uuid.UUID) ([]domain.Construction, error) {
	panic("implement me")
}

func (c ConstructionRepositoryImpl) Insert(companyId uuid.UUID, name string, initialDate time.Time, dueDate time.Time) error {
	panic("implement me")
}

func NewConstructionRepository(db database.DBImpl) domain.ConstructionRepository{
	return &ConstructionRepositoryImpl{db: db}
}
