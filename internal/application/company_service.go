package application

import (
	"context"
	"github.com/Lucasdox/constr-backend/internal/application/command"
	"github.com/Lucasdox/constr-backend/internal/application/query"
	"github.com/Lucasdox/constr-backend/internal/domain"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type CompanyService interface {
	ListCompanies(context.Context) (*query.ListCompanyQueryProjection, error)
	CreateCompanyAndPersist(context.Context, command.CreateCompanyCommand) (*uuid.UUID, error)
}

type CompanyServiceImpl struct {
	log *zap.Logger
	repository domain.CompanyRepository
}

func (s *CompanyServiceImpl) ListCompanies(ctx context.Context) (*query.ListCompanyQueryProjection, error) {
	var queryProj query.ListCompanyQueryProjection
	slc, err := s.repository.List()

	if err != nil {
		return nil, err
	}

	for _, c := range slc {
		var cProj query.Company
		cProj.Name = c.Name
		cProj.Id = c.Id
		queryProj = append(queryProj, cProj)
	}

	return &queryProj, nil
}


func (s *CompanyServiceImpl) CreateCompanyAndPersist(ctx context.Context, command command.CreateCompanyCommand) (*uuid.UUID, error) {
	company := domain.NewCompany(command.CompanyName)

	err := s.repository.Save(company)
	if err != nil {
		s.log.Warn("Failed to create company.", zap.String("name", command.CompanyName), zap.Error(err))
		return nil, err
	}
	s.log.Info("Created company successfully.", zap.String("name", command.CompanyName))

	id := company.Id

	return &id, nil
}

func NewCompanyService(repository domain.CompanyRepository) CompanyService {
	log := zap.L().Named("company_service")
	log.Info("Successfully instantiated Company Service.")

	return &CompanyServiceImpl{
		log:        log,
		repository: repository,
	}
}
