package pg

import (
	"github.com/Lucasdox/constr-backend/internal/adapters/database"
	"github.com/Lucasdox/constr-backend/internal/domain"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	SELECT = `SELECT id, name from company;`
	INSERT = `INSERT INTO company (id, name) VALUES ($1, $2);`
)

type CompanyRepositoryImpl struct {
	db database.DBImpl
}

func (r *CompanyRepositoryImpl) Insert(id uuid.UUID, name string) error {
	l := zap.L()
	_, err := r.db.Exec(INSERT, id, name)
	if err != nil {
		l.Warn("Error inserting Company", zap.Error(err))
		return err
	}
	l.Info("Inserted company successfully", zap.String("name", name))
	return nil
}

func (r *CompanyRepositoryImpl) List() ([]domain.Company, error) {
	l := zap.L()
	var companySlc []domain.Company
	rows, err := r.db.Query(SELECT)
	defer rows.Close()
	if err != nil {
		l.Named("company_repository").Warn("Couldn't retrieve companies from db.", zap.Error(err))
		return nil, err
	}
	for rows.Next() {
		company := domain.Company{}
		err = rows.Scan(&company.Id, &company.Name)
		companySlc = append(companySlc, company)
	}
	return companySlc, nil
}

func NewCompanyRepository(db database.DBImpl) domain.CompanyRepository{
	return &CompanyRepositoryImpl{db: db}
}
