package category

import (
	"github.com/jmoiron/sqlx"

	"api-food/internal/logger"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesCategoryRepository interface {
	create(m *Category) error
	update(m *Category) error
	delete(id int64) error
	getByID(id int64) (*Category, error)
	getAll() ([]*Category, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesCategoryRepository {
	var s ServicesCategoryRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newCategoryPsqlRepository(db, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
