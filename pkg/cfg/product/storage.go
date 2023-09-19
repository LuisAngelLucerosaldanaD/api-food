package product

import (
	"github.com/jmoiron/sqlx"

	"api-food/internal/logger"
)

const (
	Postgresql = "postgres"
)

type ServicesProductRepository interface {
	create(m *Product) error
	update(m *Product) error
	delete(id int64) error
	getByID(id int64) (*Product, error)
	getAll() ([]*Product, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesProductRepository {
	var s ServicesProductRepository
	engine := db.DriverName()
	switch engine {
	case Postgresql:
		return newProductPsqlRepository(db, txID)
	default:
		logger.Error.Println("el motor de base de datos no está implementado.", engine)
	}
	return s
}
