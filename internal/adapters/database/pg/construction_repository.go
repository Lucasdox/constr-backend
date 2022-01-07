package pg

import (
	"github.com/Lucasdox/constr-backend/internal/adapters/database"
	"github.com/Lucasdox/constr-backend/internal/domain"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	SELECT_CONSTRUCTION_BY_COMPANY_ID = `SELECT id, name, initial_date, due_date, company_id, created_at, updated_at FROM construction WHERE company_id = $1;`
	INSERT_CONSTRUCTION = `INSERT INTO construction (id, name, initial_date, due_date, company_id) VALUES ($1, $2, $3, $4, $5);`
)

type ConstructionRepositoryImpl struct {
	db *database.DBImpl
	log *zap.Logger
}

func (r ConstructionRepositoryImpl) ListByCompanyId(companyId uuid.UUID) ([]domain.Construction, error) {
	r.log.Info("Fetching constructions from repository")

	var slc []domain.Construction
	rows, err := r.db.Query(SELECT_CONSTRUCTION_BY_COMPANY_ID, companyId)
	defer rows.Close()
	if err != nil {
		r.log.Warn("Couldn't retrieve constructions from db.", zap.Error(err), zap.String("company_id", companyId.String()))
		return nil, err
	}
	for rows.Next() {
		c := &domain.Construction{}
		err = rows.Scan(c.Id, c.Name, c.InitialDate, c.DueDate, c.CompanyId, c.CreatedAt, c.UpdatedAt)
		if err != nil {
			r.log.Warn("Error parsing rows", zap.Error(err))
			return nil, err
		}
		slc = append(slc, *c)
	}

	return slc, nil
}

func (r ConstructionRepositoryImpl) Save(c domain.Construction) error {
	_, err := r.db.Exec(INSERT_CONSTRUCTION, c.Id, c.Name, c.InitialDate, c.DueDate, c.CompanyId)
	if err != nil {
		r.log.Warn("Error inserting Company", zap.Error(err))
		return err
	}
	r.log.Info("Inserted company successfully", zap.String("name", c.Name))
	return nil
}

func NewConstructionRepository(db *database.DBImpl) domain.ConstructionRepository{
	log := zap.L().Named("construction_repository")
	return &ConstructionRepositoryImpl{db: db, log: log}
}
