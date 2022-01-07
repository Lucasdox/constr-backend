package pg

import (
	"github.com/Lucasdox/constr-backend/internal/adapters/database"
	"github.com/Lucasdox/constr-backend/internal/domain"
	"go.uber.org/zap"
)

const (
	SELECT = `SELECT id, name from company;`
	INSERT = `INSERT INTO company (name) VALUES ($1);`
)

type CompanyRepositoryImpl struct {
	db database.DBImpl
}

func (r *CompanyRepositoryImpl) Insert(name string) error {
	l := zap.L()
	_, err := r.db.Exec(INSERT, name)
	if err != nil {
		l.Warn("Error inserting Company", zap.Error(err))
		return err
	}
	l.Info("Inserted company successfully", zap.String("name", name))
	return nil
}

func (r *CompanyRepositoryImpl) List() {
	var companySlc []domain.Company
	rows, err := r.db.Query(SELECT)
}

func NewCompanyRepository(db database.DBImpl) domain.CompanyRepository{
	return &CompanyRepositoryImpl{db: db}
}
