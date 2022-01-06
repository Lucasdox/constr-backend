package application

import (
	"context"
	"github.com/Lucasdox/constr-backend/internal/application/query"
	"github.com/google/uuid"
)

type ConstructionService interface {
	ListConstructions(ctx context.Context, companyId uuid.UUID) (cntrs []query.ListConstructionsFromCompanyQueryResponse, err error)
}

type ConstructionServiceImpl struct {

}

func (c ConstructionServiceImpl) ListConstructions(ctx context.Context, companyId uuid.UUID) (cntrs []query.ListConstructionsFromCompanyQueryResponse, err error) {
	panic("implement me")
}

func NewConstructionService() ConstructionService {
	return &ConstructionServiceImpl{}
}
