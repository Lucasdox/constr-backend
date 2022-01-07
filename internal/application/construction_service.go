package application

import (
	"context"
	"github.com/Lucasdox/constr-backend/internal/application/command"
	"github.com/Lucasdox/constr-backend/internal/application/query"
	"github.com/Lucasdox/constr-backend/internal/domain"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ConstructionService interface {
	ListByCompanyId(context.Context, uuid.UUID) ([]query.ListConstructionsFromCompanyQueryProjection, error)
	Create(context.Context, command.CreateConstructionCommand) error
}

type ConstructionServiceImpl struct {
	log *zap.Logger
	repository domain.ConstructionRepository
}

func (s *ConstructionServiceImpl) ListByCompanyId(ctx context.Context, companyId uuid.UUID) (query.ListConstructionsFromCompanyQueryProjection, error) {
	var queryProj query.ListConstructionsFromCompanyQueryProjection
	slc, err := s.repository.ListByCompanyId(companyId)

	if err != nil {
		s.log.Warn("Failed to fetch constructions.")
		return nil, err
	}

	for _, c := range slc {
		var con query.ConstructionByCompanyId
		con.Id = c.Id
		queryProj = append(queryProj, )
	}
}

func (s *ConstructionServiceImpl) Create(ctx context.Context, cmd command.CreateConstructionCommand) error {
	c := domain.NewConstruction(cmd.CompanyId, cmd.ConstructionName, cmd.InitialDate, cmd.DueDate)

	err := s.repository.Save(c)

	if err != nil {
		s.log.Warn("Failed to create construction.", zap.Error(err), zap.String("company_id", cmd.CompanyId.String()), zap.String("name", cmd.ConstructionName))
		return err
	}

	return nil
}

func NewConstructionService(r domain.ConstructionRepository) ConstructionService {
	log := zap.L().Named("construction_service")
	return &ConstructionServiceImpl{repository: r, log: log}
}
