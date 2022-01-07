package pg

import (
	"github.com/Lucasdox/constr-backend/internal/adapters/database"
	"github.com/Lucasdox/constr-backend/internal/domain"
	"go.uber.org/zap"
)

const (
	SELECT_COMPANY = `SELECT id, name from company;`
	INSERT_COMPANY = `INSERT INTO company (id, name) VALUES ($1, $2);`
)

type CompanyRepositoryImpl struct {
	log *zap.Logger
	db *database.DBImpl
}

func (r *CompanyRepositoryImpl) Save(c domain.Company) error {
	_, err := r.db.Exec(INSERT_COMPANY, c.Id, c.Name)
	if err != nil {
		r.log.Warn("Error inserting Company", zap.Error(err))
		return err
	}
	r.log.Info("Inserted company successfully", zap.String("name", c.Name))
	return nil
}

func (r *CompanyRepositoryImpl) List() ([]domain.Company, error) {
	var companySlc []domain.Company
	rows, err := r.db.Query(SELECT_COMPANY)
	defer rows.Close()
	if err != nil {
		r.log.Warn("Couldn't retrieve companies from db.", zap.Error(err))
		return nil, err
	}
	for rows.Next() {
		company := domain.Company{}
		err = rows.Scan(&company.Id, &company.Name)
		if err != nil {
			r.log.Warn("Error parsing rows", zap.Error(err))
			return nil, err
		}
		companySlc = append(companySlc, company)
	}
	return companySlc, nil
}

func NewCompanyRepository(db *database.DBImpl) domain.CompanyRepository{
	log := zap.L().Named("company_repository")
	return &CompanyRepositoryImpl{db: db, log: log}
}
