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
	ListByCompanyId(context.Context, uuid.UUID) (query.ListConstructionsFromCompanyQueryProjection, error)
	Create(context.Context, command.CreateConstructionCommand) (*uuid.UUID, error)
}

type ConstructionServiceImpl struct {
	log *zap.Logger
	repository domain.ConstructionRepository
}

func (s *ConstructionServiceImpl) ListByCompanyId(ctx context.Context, companyId uuid.UUID) (query.ListConstructionsFromCompanyQueryProjection, error) {
	queryProj := query.ListConstructionsFromCompanyQueryProjection{}
	slc, err := s.repository.ListByCompanyId(companyId)

	if err != nil {
		s.log.Warn("Failed to fetch constructions.")
		return nil, err
	}

	for _, c := range slc {
		var con query.ConstructionByCompanyId
		con.Id = c.Id
		con.Name = c.Name
		con.InitialDate = c.InitialDate
		con.DueDate = c.DueDate
		queryProj = append(queryProj, con)
	}

	s.log.Info("ListByCompanyId success")

	return queryProj, nil
}

func (s *ConstructionServiceImpl) Create(ctx context.Context, cmd command.CreateConstructionCommand) (*uuid.UUID, error) {
	c := domain.NewConstruction(cmd.CompanyId, cmd.ConstructionName, cmd.InitialDate, cmd.DueDate)

	err := s.repository.Save(c)

	if err != nil {
		s.log.Warn("Failed to create construction.", zap.Error(err), zap.String("company_id", cmd.CompanyId.String()), zap.String("name", cmd.ConstructionName))
		return nil, err
	}

	return &c.Id, nil
}

func NewConstructionService(r domain.ConstructionRepository) ConstructionService {
	log := zap.L().Named("construction_service")
	return &ConstructionServiceImpl{repository: r, log: log}
}
