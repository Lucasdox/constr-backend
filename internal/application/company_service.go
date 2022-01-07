package application

import (
	"context"
	"github.com/Lucasdox/constr-backend/internal/application/command"
	"github.com/Lucasdox/constr-backend/internal/application/query"
)

type CompanyService interface {
	List(context.Context) ([]query.ListCompanyQueryResponse, error)
	Insert(context.Context, command.CreateCompanyCommand) error
}
