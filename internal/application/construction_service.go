package application

import (
	"context"
	"github.com/Lucasdox/constr-backend/internal/domain"
	"github.com/google/uuid"
)

type ConstructionService interface {
	ListConstructions(ctx context.Context, companyId uuid.UUID) (cntrs []domain.Construction, err error)
}

type ConstructionServiceImpl struct {

}

func (c ConstructionServiceImpl) ListConstructions(ctx context.Context, companyId uuid.UUID) (cntrs []domain.Construction, err error) {
	panic("implement me")
}

func NewConstructionService() ConstructionService {
	return &ConstructionServiceImpl{}
}
