package pg

import (
	"github.com/Lucasdox/constr-backend/internal/adapters/database"
	"github.com/Lucasdox/constr-backend/internal/domain"
	"github.com/google/uuid"
)

const (
	SELECT_FIELDS = `id, name, initial_date, due_date, created_at, updated_at`
)

type ConstructionRepositoryImpl struct {
	db *database.DBImpl
}

func (c ConstructionRepositoryImpl) List(companyId uuid.UUID) ([]domain.Construction, error) {
	panic("implement me")
}

func (c ConstructionRepositoryImpl) Save(con domain.Construction) error {
	panic("implement me")
}

func NewConstructionRepository(db *database.DBImpl) domain.ConstructionRepository{
	return &ConstructionRepositoryImpl{db: db}
}
