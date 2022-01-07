package application

import (
	"context"
	"github.com/Lucasdox/constr-backend/internal/application/command"
	"github.com/Lucasdox/constr-backend/internal/application/query"
	"github.com/Lucasdox/constr-backend/internal/domain"
	"github.com/google/uuid"
)

type ConstructionService interface {
	List(context.Context, uuid.UUID) ([]query.ListConstructionsFromCompanyQueryResponse, error)
	Insert(context.Context, command.CreateConstructionCommand) error
}

type ConstructionServiceImpl struct {
	repository domain.ConstructionRepository
}

func (c *ConstructionServiceImpl) List(ctx context.Context, companyId uuid.UUID) ([]query.ListConstructionsFromCompanyQueryResponse, error) {
	panic("implement me")
}

func (c *ConstructionServiceImpl) Insert(ctx context.Context, command command.CreateConstructionCommand) error {
	panic("implement me")
}

func NewConstructionService(r domain.ConstructionRepository) ConstructionService {
	return &ConstructionServiceImpl{repository: r}
}
